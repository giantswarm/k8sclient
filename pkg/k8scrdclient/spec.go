package k8scrdclient

import (
	"context"

	"github.com/giantswarm/backoff"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type Interface interface {
	EnsureCreated(ctx context.Context, customResource *apiextensionsv1.CustomResourceDefinition, backOff backoff.Interface) error
	EnsureDeleted(ctx context.Context, customResource *apiextensionsv1.CustomResourceDefinition, backOff backoff.Interface) error
}
