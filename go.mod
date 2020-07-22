module github.com/giantswarm/k8sclient/v3

go 1.14

require (
	github.com/coreos/bbolt v1.3.3 // indirect
	github.com/coreos/etcd v3.3.17+incompatible // indirect
	github.com/giantswarm/apiextensions v0.4.16
	github.com/giantswarm/backoff v0.2.0
	github.com/giantswarm/microerror v0.2.0
	github.com/giantswarm/micrologger v0.3.1
	github.com/google/go-cmp v0.5.1
	k8s.io/api v0.17.2
	k8s.io/apiextensions-apiserver v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v0.17.2
	sigs.k8s.io/controller-runtime v0.5.2
	sigs.k8s.io/structured-merge-diff v1.0.1 // indirect
	sigs.k8s.io/testing_frameworks v0.1.2 // indirect
)
