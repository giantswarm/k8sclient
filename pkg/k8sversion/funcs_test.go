package k8sversion

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_K8sVersion_Latest(t *testing.T) {
	testCases := []struct {
		name         string
		versions     []string
		latest       string
		errorMatcher func(err error) bool
	}{
		{
			name: "case 0",
			versions: []string{
				"v1",
				"v2",
				"v1alpha1",
			},
			latest:       "v2",
			errorMatcher: nil,
		},
		{
			name: "case 1",
			versions: []string{
				"v1alpha1",
				"v1beta1",
				"v1alpha2",
			},
			latest:       "v1beta1",
			errorMatcher: nil,
		},
		{
			name: "case 2",
			versions: []string{
				"v1alpha2",
				"v2alpha3",
				"v10alpha3",
				"v10alpha30",
				"v1alpha3",
			},
			latest:       "v10alpha30",
			errorMatcher: nil,
		},
		{
			name: "case 3",
			versions: []string{
				"v1alpha2",
				"foo",
				"v10alpha3",
				"v10alpha30",
				"v1alpha3",
			},
			latest:       "",
			errorMatcher: IsInvalidKubeVersion,
		},
		{
			name: "case 4",
			versions: []string{
				"v1",
				"v1alpha1",
			},
			latest:       "v1",
			errorMatcher: nil,
		},
		{
			name: "case 5",
			versions: []string{
				"v1alpha1",
				"v1",
			},
			latest:       "v1",
			errorMatcher: nil,
		},
		{
			name:         "case 6",
			versions:     []string{},
			latest:       "",
			errorMatcher: IsInvalidKubeVersion,
		},
		{
			name:         "case 7",
			versions:     nil,
			latest:       "",
			errorMatcher: IsInvalidKubeVersion,
		},
		{
			name: "case 8",
			versions: []string{
				"v2alpha1",
				"v1beta1",
				"v1alpha2",
			},
			latest:       "v2alpha1",
			errorMatcher: nil,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(tc.name)

			latest, err := Latest(tc.versions)

			switch {
			case err == nil && tc.errorMatcher == nil:
				// correct; carry on
			case err != nil && tc.errorMatcher == nil:
				t.Fatalf("error == %#v, want nil", err)
			case err == nil && tc.errorMatcher != nil:
				t.Fatalf("error == nil, want non-nil")
			case !tc.errorMatcher(err):
				t.Fatalf("error == %#v, want matching", err)
			}

			if !cmp.Equal(latest, tc.latest) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.latest, latest))
			}
		})
	}
}

func Test_K8sVersion_Less(t *testing.T) {
	testCases := []struct {
		name         string
		a            string
		b            string
		less         bool
		errorMatcher func(err error) bool
	}{
		{
			name:         "case 0",
			a:            "v1",
			b:            "v2",
			less:         true,
			errorMatcher: nil,
		},
		{
			name:         "case 1",
			a:            "v2",
			b:            "v1",
			less:         false,
			errorMatcher: nil,
		},
		{
			name:         "case 2",
			a:            "v1alpha1",
			b:            "v1",
			less:         true,
			errorMatcher: nil,
		},
		{
			name:         "case 3",
			a:            "v1",
			b:            "v1alpha1",
			less:         false,
			errorMatcher: nil,
		},
		{
			name:         "case 4",
			a:            "v1",
			b:            "foo",
			less:         false,
			errorMatcher: IsInvalidKubeVersion,
		},
		{
			name:         "case 5",
			a:            "v1",
			b:            "v1",
			less:         false,
			errorMatcher: nil,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Log(tc.name)

			less, err := Less(tc.a, tc.b)

			switch {
			case err == nil && tc.errorMatcher == nil:
				// correct; carry on
			case err != nil && tc.errorMatcher == nil:
				t.Fatalf("error == %#v, want nil", err)
			case err == nil && tc.errorMatcher != nil:
				t.Fatalf("error == nil, want non-nil")
			case !tc.errorMatcher(err):
				t.Fatalf("error == %#v, want matching", err)
			}

			if !cmp.Equal(less, tc.less) {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.less, less))
			}
		})
	}
}
