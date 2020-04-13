package k8scrdclient

import (
	"context"
	"fmt"

	"github.com/giantswarm/backoff"
	"github.com/giantswarm/microerror"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/giantswarm/k8sclient/k8sversion"
)

// EnsureCreated ensures the given CRD exists, is active (aka. established) and
// does not have conflicting names.
func (c *CRDClient) EnsureCreatedV1(ctx context.Context, crd *apiextensionsv1.CustomResourceDefinition, b backoff.Interface) error {
	var err error

	err = c.ensureCreatedV1(ctx, crd)
	if err != nil {
		return microerror.Mask(err)
	}

	err = c.ensureUpdatedV1(ctx, crd, b)
	if err != nil {
		return microerror.Mask(err)
	}

	err = c.validateStatusV1(crd, b)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// EnsureDeleted ensures the given CRD does not exist.
func (c *CRDClient) EnsureDeletedV1(ctx context.Context, crd *apiextensionsv1.CustomResourceDefinition, b backoff.Interface) error {
	o := func() error {
		err := c.k8sExtClient.ApiextensionsV1().CustomResourceDefinitions().Delete(crd.Name, nil)
		if errors.IsNotFound(err) {
			// Fall trough. We reached our goal.
		} else if err != nil {
			return microerror.Mask(err)
		}

		return nil
	}

	err := backoff.Retry(o, b)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (c *CRDClient) ensureCreatedV1(ctx context.Context, crd *apiextensionsv1.CustomResourceDefinition) error {
	// We explicitly check for the existence of the CRD by fetching it to
	// workaround nasty side effects with error handling of multiple creation
	// attempts. It may happen that a CRD creation is forbidden while it exists
	// and actually has to be updated.
	{
		c.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("finding CRD %#q", crd.Name))

		_, err := c.k8sExtClient.ApiextensionsV1().CustomResourceDefinitions().Get(crd.Name, metav1.GetOptions{})
		if errors.IsNotFound(err) {
			c.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("did not find CRD %#q", crd.Name))
		} else if err != nil {
			return microerror.Mask(err)
		} else {
			c.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("found CRD %#q", crd.Name))
			return nil
		}
	}

	{
		c.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("creating CRD %#q", crd.Name))

		_, err := c.k8sExtClient.ApiextensionsV1().CustomResourceDefinitions().Create(crd)
		if err != nil {
			return microerror.Mask(err)
		}

		c.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("created CRD %#q", crd.Name))
	}

	return nil
}

// ensureUpdated ensures if the CRD changed it is updated accordingly. This is
// needed if e.g. a previous version of the CRD without the status subresource
// is present where it should actually be set. Another example would be the CRD
// apiversion changing, which tends to happen every now and then over the
// runtime object lifecycle and community adoption.
func (c *CRDClient) ensureUpdatedV1(ctx context.Context, desired *apiextensionsv1.CustomResourceDefinition, b backoff.Interface) error {
	o := func() error {
		current, err := c.k8sExtClient.ApiextensionsV1().CustomResourceDefinitions().Get(desired.Name, metav1.GetOptions{})
		if err != nil {
			return microerror.Mask(err)
		}

		equal := desired.Spec.String() == current.Spec.String()
		latest, err := crdVersionLatestV1(desired, current)
		if err != nil {
			return microerror.Mask(err)
		}

		if latest && !equal {
			c.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("updating CRD %#q", desired.Name))

			// Since we get a pointer of the desired CRD we do not want to mess around
			// with it. Thus we take a copy of desired and use that instead for the
			// update.
			copy := desired.DeepCopy()
			copy.SetResourceVersion(current.ResourceVersion)

			_, err = c.k8sExtClient.ApiextensionsV1().CustomResourceDefinitions().Update(copy)
			if err != nil {
				return microerror.Mask(err)
			}

			c.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("updated CRD %#q", desired.Name))
		} else {
			c.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("not updating CRD %#q", desired.Name))
			c.logger.LogCtx(ctx, "level", "debug", "message", fmt.Sprintf("CRD %#q already updated", desired.Name))
		}

		return nil
	}

	err := backoff.Retry(o, b)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func (c *CRDClient) validateStatusV1(crd *apiextensionsv1.CustomResourceDefinition, b backoff.Interface) error {
	var err error

	o := func() error {
		manifest, err := c.k8sExtClient.ApiextensionsV1().CustomResourceDefinitions().Get(crd.Name, metav1.GetOptions{})
		if err != nil {
			return microerror.Mask(err)
		}

		// In case the CRDs names are not accepted we have to stop processing here
		// and return the reason of the failing condition. Therefore we stop retries
		// permanently.
		{
			con, ok := statusConditionV1(manifest.Status.Conditions, apiextensionsv1.NamesAccepted)
			if ok && statusConditionFalseV1(con) {
				return backoff.Permanent(microerror.Maskf(nameConflictError, con.Message))
			}
		}
		// In case the CRD is non-structural we have to stop processing here and
		// return the reason of the failing condition. Therefore we stop retries
		// permanently.
		{
			con, ok := statusConditionV1(manifest.Status.Conditions, apiextensionsv1.NonStructuralSchema)
			if ok && statusConditionTrueV1(con) {
				return backoff.Permanent(microerror.Maskf(notEstablishedError, con.Message))
			}
		}
		// In case the CRD is not yet established we have to retry and only return a
		// normal error so that the backoff can do its job.
		{
			con, ok := statusConditionV1(manifest.Status.Conditions, apiextensionsv1.Established)
			if ok && statusConditionFalseV1(con) {
				return microerror.Maskf(notEstablishedError, con.Message)
			}
		}

		return nil
	}

	err = backoff.Retry(o, b)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// crdVersionLatest returns true when the desired version of the given CRD
// represents the latest version of them all. In case the current version of the
// given CRD which we just read from the system appears to be newer, this means
// another version of the same operator updated the CRD to its own version,
// because it was in fact the latest version. Consider the following
// relationships between operators and CRD versions they are aware of. Below o1
// is operator-1 being aware of apiversion v1. In turn operator-2 knows about
// apiversion v1 and v2, where v1 must be the exact same version as operator-1
// defined already in order to keep the CRD schema version lifecycle in tact.
// Note that in the example below o1 must not run anymore in cae o3 drops its
// support. There is some orchestration involved the system's maintainers have
// to get right.
//
//         o1    |     o2     |    o3
//     ---------------------------------
//         v1    |   v1, v2   |   v2, v3
//
// Each operator compares the latest version it finds in desired and current. If
// the latest version of desired is below the latest version of current, the
// desired version of the given CRD is not considered to be the latest known
// version and thus not allowed to update the system's CRD.
//
// Note that the given versions must be in the format of usual Kubernetes
// APIVersions, e.g. v1alpha1, v2beta5, v2.
//
//     https://kubernetes.io/docs/concepts/overview/kubernetes-api/#api-versioning
//
func crdVersionLatestV1(desired *apiextensionsv1.CustomResourceDefinition, current *apiextensionsv1.CustomResourceDefinition) (bool, error) {
	desiredVersions := crdVersionsV1(desired)
	currentVersions := crdVersionsV1(current)

	// In case there are no versions given at all, we do not want to do anything.
	if len(desiredVersions) == 0 && len(currentVersions) == 0 {
		return false, nil
	}
	// In case there are only current versions given, we do not want to overwrite
	// them.
	if len(desiredVersions) == 0 && len(currentVersions) != 0 {
		return false, nil
	}
	// In case there are only desired versions given, we want to update to them.
	if len(desiredVersions) != 0 && len(currentVersions) == 0 {
		return true, nil
	}

	// All code below handles the situation in which both desired and current
	// versions are given. In this case we need to figure out if desired or
	// current contains the latest version.

	desiredLatest, err := k8sversion.Latest(desiredVersions)
	if err != nil {
		return false, microerror.Mask(err)
	}
	currentLatest, err := k8sversion.Latest(crdVersionsV1(current))
	if err != nil {
		return false, microerror.Mask(err)
	}

	less, err := k8sversion.Less(desiredLatest, currentLatest)
	if err != nil {
		return false, microerror.Mask(err)
	}
	if less {
		return false, nil
	}

	return true, nil
}

func crdVersionsV1(crd *apiextensionsv1.CustomResourceDefinition) []string {
	var versions []string

	if crd == nil {
		return versions
	}

	for _, v := range crd.Spec.Versions {
		versions = append(versions, v.Name)
	}

	return versions
}

func statusConditionV1(conditions []apiextensionsv1.CustomResourceDefinitionCondition, t apiextensionsv1.CustomResourceDefinitionConditionType) (apiextensionsv1.CustomResourceDefinitionCondition, bool) {
	for _, con := range conditions {
		if con.Type == t {
			return con, true
		}
	}

	return apiextensionsv1.CustomResourceDefinitionCondition{}, false
}

func statusConditionFalseV1(con apiextensionsv1.CustomResourceDefinitionCondition) bool {
	return con.Status == apiextensionsv1.ConditionFalse
}

func statusConditionTrueV1(con apiextensionsv1.CustomResourceDefinitionCondition) bool {
	return con.Status == apiextensionsv1.ConditionTrue
}
