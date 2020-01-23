package k8scrdclient

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Test_crdAPIVersionEqual(t *testing.T) {
	testCases := []struct {
		name   string
		crd    *apiextensionsv1beta1.CustomResourceDefinition
		latest *apiextensionsv1beta1.CustomResourceDefinition
		equal  bool
	}{
		{
			name:  "case 0 only one crd given yields false",
			crd:   &apiextensionsv1beta1.CustomResourceDefinition{},
			equal: false,
		},
		{
			name:   "case 1 both the same yields true",
			crd:    &apiextensionsv1beta1.CustomResourceDefinition{},
			latest: &apiextensionsv1beta1.CustomResourceDefinition{},
			equal:  true,
		},
		{
			name: "case 2 different apiversions yields false",
			crd: &apiextensionsv1beta1.CustomResourceDefinition{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "a",
				},
			},
			latest: &apiextensionsv1beta1.CustomResourceDefinition{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "b",
				},
			},
			equal: false,
		},
		{
			name: "case 3 same apiversions yields true",
			crd: &apiextensionsv1beta1.CustomResourceDefinition{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "a",
				},
			},
			latest: &apiextensionsv1beta1.CustomResourceDefinition{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "a",
				},
			},
			equal: true,
		},
		{
			name: "case 4 same apiversions but different names yields true",
			crd: &apiextensionsv1beta1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name: "a",
				},
				TypeMeta: metav1.TypeMeta{
					APIVersion: "a",
				},
			},
			latest: &apiextensionsv1beta1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name: "b",
				},
				TypeMeta: metav1.TypeMeta{
					APIVersion: "a",
				},
			},
			equal: true,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			equal := crdAPIVersionEqual(tc.crd, tc.latest)

			if !cmp.Equal(equal, tc.equal) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.equal, equal))
			}
		})
	}
}

func Test_crdStatusEqual(t *testing.T) {
	testCases := []struct {
		name   string
		crd    *apiextensionsv1beta1.CustomResourceDefinition
		latest *apiextensionsv1beta1.CustomResourceDefinition
		equal  bool
	}{
		{
			name:  "case 0 only one crd given yields false",
			crd:   &apiextensionsv1beta1.CustomResourceDefinition{},
			equal: false,
		},
		{
			name:   "case 1 both the same yields true",
			crd:    &apiextensionsv1beta1.CustomResourceDefinition{},
			latest: &apiextensionsv1beta1.CustomResourceDefinition{},
			equal:  true,
		},
		{
			name: "case 2 one nil the other not nil yields false",
			crd: &apiextensionsv1beta1.CustomResourceDefinition{
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
						Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
					},
				},
			},
			latest: &apiextensionsv1beta1.CustomResourceDefinition{
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
						Status: nil,
					},
				},
			},
			equal: false,
		},
		{
			name: "case 3 same statuses yields true",
			crd: &apiextensionsv1beta1.CustomResourceDefinition{
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
						Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
					},
				},
			},
			latest: &apiextensionsv1beta1.CustomResourceDefinition{
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
						Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
					},
				},
			},
			equal: true,
		},
		{
			name: "case 4 same statuses but different names yields true",
			crd: &apiextensionsv1beta1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name: "a",
				},
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
						Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
					},
				},
			},
			latest: &apiextensionsv1beta1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name: "b",
				},
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Subresources: &apiextensionsv1beta1.CustomResourceSubresources{
						Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
					},
				},
			},
			equal: true,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			equal := crdStatusEqual(tc.crd, tc.latest)

			if !cmp.Equal(equal, tc.equal) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.equal, equal))
			}
		})
	}
}

func Test_crdValidationEqual(t *testing.T) {
	testCases := []struct {
		name   string
		crd    *apiextensionsv1beta1.CustomResourceDefinition
		latest *apiextensionsv1beta1.CustomResourceDefinition
		equal  bool
	}{
		{
			name:  "case 0 only one crd given yields false",
			crd:   &apiextensionsv1beta1.CustomResourceDefinition{},
			equal: false,
		},
		{
			name:   "case 1 both the same yields true",
			crd:    &apiextensionsv1beta1.CustomResourceDefinition{},
			latest: &apiextensionsv1beta1.CustomResourceDefinition{},
			equal:  true,
		},
		{
			name: "case 2 one nil the other not nil yields false",
			crd: &apiextensionsv1beta1.CustomResourceDefinition{
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Validation: &apiextensionsv1beta1.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{},
					},
				},
			},
			latest: &apiextensionsv1beta1.CustomResourceDefinition{
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Validation: &apiextensionsv1beta1.CustomResourceValidation{
						OpenAPIV3Schema: nil,
					},
				},
			},
			equal: false,
		},
		{
			name: "case 3 same statuses yields true",
			crd: &apiextensionsv1beta1.CustomResourceDefinition{
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Validation: &apiextensionsv1beta1.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{},
					},
				},
			},
			latest: &apiextensionsv1beta1.CustomResourceDefinition{
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Validation: &apiextensionsv1beta1.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{},
					},
				},
			},
			equal: true,
		},
		{
			name: "case 4 same statuses but different names yields true",
			crd: &apiextensionsv1beta1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name: "a",
				},
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Validation: &apiextensionsv1beta1.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{},
					},
				},
			},
			latest: &apiextensionsv1beta1.CustomResourceDefinition{
				ObjectMeta: metav1.ObjectMeta{
					Name: "b",
				},
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Validation: &apiextensionsv1beta1.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{},
					},
				},
			},
			equal: true,
		},
		{
			name: "case 5 different statuses yields false",
			crd: &apiextensionsv1beta1.CustomResourceDefinition{
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Validation: &apiextensionsv1beta1.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{
							Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
								"a": apiextensionsv1beta1.JSONSchemaProps{
									Pattern: "a",
								},
							},
						},
					},
				},
			},
			latest: &apiextensionsv1beta1.CustomResourceDefinition{
				Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
					Validation: &apiextensionsv1beta1.CustomResourceValidation{
						OpenAPIV3Schema: &apiextensionsv1beta1.JSONSchemaProps{
							Properties: map[string]apiextensionsv1beta1.JSONSchemaProps{
								"b": apiextensionsv1beta1.JSONSchemaProps{
									Pattern: "b",
								},
							},
						},
					},
				},
			},
			equal: false,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			equal := crdValidationEqual(tc.crd, tc.latest)

			if !cmp.Equal(equal, tc.equal) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.equal, equal))
			}
		})
	}
}
