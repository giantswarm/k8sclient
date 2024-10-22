package k8sclient

import (
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/giantswarm/k8sclient/v8/pkg/k8scrdclient"
)

type Interface interface {
	CRDClient() k8scrdclient.Interface
	CtrlClient() client.Client
	DynClient() dynamic.Interface
	ExtClient() apiextensionsclient.Interface
	K8sClient() kubernetes.Interface
	RESTClient() rest.Interface
	RESTConfig() *rest.Config
	Scheme() *runtime.Scheme
}
