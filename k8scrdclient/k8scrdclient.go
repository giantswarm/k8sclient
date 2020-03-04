package k8scrdclient

import (
	"context"

	"github.com/giantswarm/backoff"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Config struct {
	K8sExtClient apiextensionsclient.Interface
	Logger       micrologger.Logger
}

type CRDClient struct {
	k8sExtClient apiextensionsclient.Interface
	logger       micrologger.Logger
}

func New(config Config) (*CRDClient, error) {
	if config.K8sExtClient == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.K8sExtClient must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	crdClient := &CRDClient{
		k8sExtClient: config.K8sExtClient,
		logger:       config.Logger,
	}

	return crdClient, nil
}

// EnsureCreated ensures the given CRD exists, is active (aka. established) and
// does not have conflicting names.
func (c *CRDClient) EnsureCreated(ctx context.Context, crd *apiextensionsv1beta1.CustomResourceDefinition, b backoff.Interface) error {
	var err error

	err = c.ensureCreated(ctx, crd, b)
	if err != nil {
		return microerror.Mask(err)
	}

	err = c.ensureStatusSubresourceCreated(ctx, crd, b)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

// EnsureDeleted ensures the given CRD does not exist.
func (c *CRDClient) EnsureDeleted(ctx context.Context, crd *apiextensionsv1beta1.CustomResourceDefinition, b backoff.Interface) error {
	o := func() error {
		err := c.k8sExtClient.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(crd.Name, nil)
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

func (c *CRDClient) ensureCreated(ctx context.Context, crd *apiextensionsv1beta1.CustomResourceDefinition, b backoff.Interface) error {
	_, err := c.k8sExtClient.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
	if errors.IsAlreadyExists(err) {
		// Fall trough. We need to check CRD status.
	} else if err != nil {
		return microerror.Mask(err)
	}

	o := func() error {
		manifest, err := c.k8sExtClient.ApiextensionsV1beta1().CustomResourceDefinitions().Get(crd.Name, metav1.GetOptions{})
		if err != nil {
			return microerror.Mask(err)
		}

		// In case the CRDs names are not accepted we have to stop processing here
		// and return the reason of the failing condition. Therefore we stop retries
		// permanently.
		{
			con, ok := statusCondition(manifest.Status.Conditions, apiextensionsv1beta1.NamesAccepted)
			if ok && statusConditionFalse(con) {
				return backoff.Permanent(microerror.Maskf(nameConflictError, con.Reason))
			}
		}
		// In case the CRD is non-structural we have to stop processing here and
		// return the reason of the failing condition. Therefore we stop retries
		// permanently.
		{
			con, ok := statusCondition(manifest.Status.Conditions, apiextensionsv1beta1.NonStructuralSchema)
			if ok && statusConditionTrue(con) {
				return backoff.Permanent(microerror.Maskf(notEstablishedError, con.Reason))
			}
		}
		// In case the CRD is not yet established we have to retry and only return a
		// normal error so that the backoff can do its job.
		{
			con, ok := statusCondition(manifest.Status.Conditions, apiextensionsv1beta1.Established)
			if ok && statusConditionFalse(con) {
				return microerror.Maskf(notEstablishedError, con.Reason)
			}
		}

		return nil
	}

	err = backoff.Retry(o, b)
	if err != nil {
		deleteErr := c.k8sExtClient.ApiextensionsV1beta1().CustomResourceDefinitions().Delete(crd.Name, nil)
		if deleteErr != nil {
			return microerror.Mask(deleteErr)
		}

		return microerror.Mask(err)
	}

	return nil
}

// ensureStatusSubresourceCreated ensures if the CRD has a status subresource
// it is created. This is needed if a previous version of the CRD without the
// status subresource is present.
func (c *CRDClient) ensureStatusSubresourceCreated(ctx context.Context, crd *apiextensionsv1beta1.CustomResourceDefinition, b backoff.Interface) error {
	if crd.Spec.Subresources == nil || crd.Spec.Subresources.Status == nil {
		// Nothing to do.
		return nil
	}

	o := func() error {
		manifest, err := c.k8sExtClient.ApiextensionsV1beta1().CustomResourceDefinitions().Get(crd.Name, metav1.GetOptions{})
		if err != nil {
			return microerror.Mask(err)
		}

		if manifest.Spec.Subresources == nil || manifest.Spec.Subresources.Status == nil {
			crd.SetResourceVersion(manifest.ResourceVersion)
			_, err = c.k8sExtClient.ApiextensionsV1beta1().CustomResourceDefinitions().Update(crd)
			if err != nil {
				return microerror.Mask(err)
			}
		}

		return nil
	}

	err := backoff.Retry(o, b)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func statusCondition(conditions []apiextensionsv1beta1.CustomResourceDefinitionCondition, t apiextensionsv1beta1.CustomResourceDefinitionConditionType) (apiextensionsv1beta1.CustomResourceDefinitionCondition, bool) {
	for _, con := range conditions {
		if con.Type == t {
			return con, true
		}
	}

	return apiextensionsv1beta1.CustomResourceDefinitionCondition{}, false
}

func statusConditionFalse(con apiextensionsv1beta1.CustomResourceDefinitionCondition) bool {
	return con.Status == apiextensionsv1beta1.ConditionFalse
}

func statusConditionTrue(con apiextensionsv1beta1.CustomResourceDefinitionCondition) bool {
	return con.Status == apiextensionsv1beta1.ConditionTrue
}
