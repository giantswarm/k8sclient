package fake

import (
	"testing"

	"github.com/giantswarm/micrologger"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"github.com/giantswarm/k8sclient/v8/pkg/k8sclient"
)

// TestInterface make sure Clients struct is compatible with k8sClient.Interface.
func TestInterface(t *testing.T) {
	var err error

	var logger micrologger.Logger
	{
		c := micrologger.Config{}

		logger, err = micrologger.New(c)
		if err != nil {
			t.Fatal(err)
		}
	}

	var k8sClient k8sclient.Interface
	{
		c := k8sclient.ClientsConfig{
			Logger:        logger,
			SchemeBuilder: k8sclient.SchemeBuilder(appsv1.SchemeBuilder),
		}

		k8sClient, err = NewClients(c, &corev1.Node{})
		if err != nil {
			t.Fatal(err)
		}
	}

	// Needed to avoid: k8sClient declared but not used error.
	k8sClient.RESTClient()
}
