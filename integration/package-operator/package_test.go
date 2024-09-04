//go:build integration

package packageoperator

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/go-logr/logr/testr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"testing"

	corev1alpha1 "package-operator.run/apis/core/v1alpha1"
	manifestsv1alpha1 "package-operator.run/apis/manifests/v1alpha1"
)

func newPackage(meta metav1.ObjectMeta, spec corev1alpha1.PackageSpec, namespaced bool) client.Object {
	if !namespaced {
		return &corev1alpha1.ClusterPackage{
			ObjectMeta: meta,
			Spec:       spec,
		}
	}

	pkg := &corev1alpha1.Package{
		ObjectMeta: meta,
		Spec:       spec,
	}
	pkg.SetNamespace("default")
	return pkg
}

// testNamespacedAndCluster constructs a (Cluster)Package from the 'meta' and 'spec' parameters
// adding a namespace, if needed. Then it ensures successful deployment of both versions of the package
// and optionally runs 'postCheck'.
func testNamespacedAndCluster(
	t *testing.T,
	meta metav1.ObjectMeta, spec corev1alpha1.PackageSpec,
	postCheck func(ctx context.Context, t *testing.T, namespace string),
) {
	t.Helper()

	for _, tc := range []struct {
		name             string
		namespace        string
		pkg              client.Object
		objectDeployment client.Object
	}{
		{"cluster", "", newPackage(meta, spec, false), &corev1alpha1.ClusterObjectDeployment{}},
		{"namespaced", "default", newPackage(meta, spec, true), &corev1alpha1.ObjectDeployment{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx := logr.NewContext(context.Background(), testr.New(t))

			requireDeployPackage(ctx, t, tc.pkg, tc.objectDeployment)

			if postCheck != nil {
				postCheck(ctx, t, tc.namespace)
			}
		})
	}
}

func TestPackage_simple(t *testing.T) {
	meta := metav1.ObjectMeta{
		Name: "success",
	}
	spec := corev1alpha1.PackageSpec{
		Image: SuccessTestPackageImage,
		Config: &runtime.RawExtension{
			Raw: []byte(fmt.Sprintf(`{"testStubImage": "%s"}`, TestStubImage)),
		},
	}
	postCheck := func(ctx context.Context, t *testing.T, namespace string) {
		t.Helper()

		if namespace == "" {
			return
		}

		// Test if environment information is injected successfully.
		deploy := &appsv1.Deployment{}
		err := Client.Get(ctx, client.ObjectKey{
			Name:      "test-stub-success",
			Namespace: namespace,
		}, deploy)
		require.NoError(t, err)

		var env manifestsv1alpha1.PackageEnvironment
		te := deploy.Annotations["test-environment"]
		err = json.Unmarshal([]byte(te), &env)
		require.NoError(t, err)
		assert.NotEmpty(t, env.Kubernetes.Version)
	}

	testNamespacedAndCluster(t, meta, spec, postCheck)
}

func TestPackage_simpleWithSlices(t *testing.T) {
	meta := metav1.ObjectMeta{
		Name: "success-slices",
		Annotations: map[string]string{
			// Manually force one slice per object.
			"packages.package-operator.run/chunking-strategy": "EachObject",
		},
	}
	spec := corev1alpha1.PackageSpec{
		Image: SuccessTestPackageImage,
		Config: &runtime.RawExtension{
			Raw: []byte(fmt.Sprintf(`{"testStubImage": "%s"}`, TestStubImage)),
		},
	}

	postCheck := func(ctx context.Context, t *testing.T, namespace string) {
		t.Helper()

		// Reminder: Every rendered object get's wrapped into its own ObjectSlice.
		// Or in go terms: `len(renderedObjects) == len(resultingObjectSlices)`.
		// The test package renders a deployment plus an additional namespace object
		// when the package is installed as a ClusterPackage.
		// When `namespace` is not empty, the test package has been installed as a Package.
		// When `namespace` is empty, the test package has been installed as a ClusterPackage.
		assertedAmount := 2 // Deployment and Namespace
		if namespace != "" {
			assertedAmount = 1 // only a Deployment
		}

		assertAmountOfSliceObjectsControlledPerObjectSet(ctx, t, types.NamespacedName{
			Namespace: namespace,
			Name:      "success-slices",
		}, assertedAmount)
	}

	testNamespacedAndCluster(t, meta, spec, postCheck)
}

func TestPackage_simpleWithoutSlices(t *testing.T) {
	meta := metav1.ObjectMeta{
		Name: "success-no-slices",
		Annotations: map[string]string{
			// Manually disable slicing.
			"packages.package-operator.run/chunking-strategy": "NoOp",
		},
	}
	spec := corev1alpha1.PackageSpec{
		Image: SuccessTestPackageImage,
		Config: &runtime.RawExtension{
			Raw: []byte(fmt.Sprintf(`{"testStubImage": "%s"}`, TestStubImage)),
		},
	}

	postCheck := func(ctx context.Context, t *testing.T, namespace string) {
		t.Helper()
		assertAmountOfSliceObjectsControlledPerObjectSet(ctx, t, types.NamespacedName{
			Namespace: namespace,
			Name:      "success-no-slices",
		}, 0)
	}

	testNamespacedAndCluster(t, meta, spec, postCheck)
}

func TestPackage_multi(t *testing.T) {
	meta := metav1.ObjectMeta{
		Name:      "success-multi",
		Namespace: "default",
	}
	spec := corev1alpha1.PackageSpec{
		Image: SuccessTestMultiPackageImage,
		Config: &runtime.RawExtension{
			Raw: []byte(fmt.Sprintf(`{"testStubMultiPackageImage": "%s","testStubImage": "%s"}`,
				SuccessTestMultiPackageImage, TestStubImage,
			)),
		},
	}

	postCheck := func(ctx context.Context, t *testing.T, namespace string) {
		t.Helper()

		var pkgBE, pkgFE client.Object
		var ns string
		if namespace != "" {
			pkgBE = &corev1alpha1.Package{}
			pkgFE = &corev1alpha1.Package{}
			ns = namespace
		} else {
			pkgBE = &corev1alpha1.ClusterPackage{}
			pkgFE = &corev1alpha1.ClusterPackage{}
			ns = "success-multi"
		}

		require.NoError(t, Client.Get(ctx, client.ObjectKey{Name: "success-multi-backend", Namespace: ns}, pkgBE))
		require.NoError(t, Client.Get(ctx, client.ObjectKey{Name: "success-multi-frontend", Namespace: ns}, pkgFE))
	}

	testNamespacedAndCluster(t, meta, spec, postCheck)
}

func TestPackage_cel(t *testing.T) {
	meta := metav1.ObjectMeta{
		Name:      "success-cel",
		Namespace: "default",
	}
	spec := corev1alpha1.PackageSpec{
		Image: SuccessTestCelPackageImage,
		Config: &runtime.RawExtension{
			Raw: []byte(fmt.Sprintf(`{"testStubCelPackageImage": "%s","testStubImage": "%s"}`,
				SuccessTestCelPackageImage, TestStubImage,
			)),
		},
	}

	postCheck := func(ctx context.Context, t *testing.T, namespace string) {
		t.Helper()

		var ns string
		if namespace != "" {
			ns = namespace
		} else {
			ns = "success-cel"
		}

		// deployment should be there
		deploy := &appsv1.Deployment{}
		err := Client.Get(ctx, client.ObjectKey{
			Name:      "test-deployment",
			Namespace: ns,
		}, deploy)
		require.NoError(t, err)

		// configMap should not be there
		cm := &v1.ConfigMap{}
		err = Client.Get(ctx, client.ObjectKey{
			Name:      "test-cm",
			Namespace: ns,
		}, cm)
		require.EqualError(t, err, "configmaps \"test-cm\" not found")

		// ignored configMap should not be there
		err = Client.Get(ctx, client.ObjectKey{
			Name:      "ignored-cm",
			Namespace: ns,
		}, cm)
		require.EqualError(t, err, "configmaps \"ignored-cm\" not found")
	}

	testNamespacedAndCluster(t, meta, spec, postCheck)
}

func TestPackage_nonExistent(t *testing.T) {
	pkg := &corev1alpha1.Package{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "non-existent",
			Namespace: "default",
		},
		Spec: corev1alpha1.PackageSpec{
			Image: "quay.io/package-operator/non-existent:v123",
		},
	}

	ctx := logr.NewContext(context.Background(), testr.New(t))
	require.NoError(t, Client.Create(ctx, pkg))
	cleanupOnSuccess(ctx, t, pkg)

	require.NoError(t,
		Waiter.WaitForCondition(ctx, pkg, corev1alpha1.PackageUnpacked, metav1.ConditionFalse))

	existingPackage := &corev1alpha1.Package{}
	err := Client.Get(ctx, client.ObjectKey{Name: "non-existent", Namespace: "default"}, existingPackage)
	require.NoError(t, err)
	require.Equal(t, "ImagePullBackOff", existingPackage.Status.Conditions[0].Reason)
}
