module github.com/giantswarm/k8sclient/v5

go 1.14

require (
	github.com/giantswarm/apiextensions/v3 v3.13.0
	github.com/giantswarm/backoff v0.2.0
	github.com/giantswarm/microerror v0.3.0
	github.com/giantswarm/micrologger v0.4.0
	github.com/google/go-cmp v0.5.4
	k8s.io/api v0.19.2
	k8s.io/apiextensions-apiserver v0.19.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v0.19.2
	sigs.k8s.io/controller-runtime v0.7.0
)

replace sigs.k8s.io/cluster-api => github.com/giantswarm/cluster-api v0.3.10-gs
