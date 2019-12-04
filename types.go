package k8sclient

import "k8s.io/apimachinery/pkg/runtime"

// SchemeBuilder is an optional way to extend the known types to the global
// client-go scheme. Make use of it for custom CRs.
type SchemeBuilder []func(*runtime.Scheme) error
