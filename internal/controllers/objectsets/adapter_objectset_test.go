package objectsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1alpha1 "package-operator.run/apis/core/v1alpha1"
)

func TestGenericObjectSet(t *testing.T) {
	t.Parallel()

	objectSet := newGenericObjectSet(testScheme).(*GenericObjectSet)

	co := objectSet.ClientObject()
	assert.IsType(t, &corev1alpha1.ObjectSet{}, co)

	objectSet.Status.Conditions = []metav1.Condition{{}}
	assert.Equal(t, objectSet.Status.Conditions, *objectSet.GetConditions())

	assert.False(t, objectSet.IsPaused())
	assert.False(t, objectSet.IsArchived())
	objectSet.Spec.LifecycleState = corev1alpha1.ObjectSetLifecycleStatePaused
	assert.True(t, objectSet.IsPaused())
	objectSet.Spec.LifecycleState = corev1alpha1.ObjectSetLifecycleStateArchived
	assert.True(t, objectSet.IsArchived())

	phases := []corev1alpha1.ObjectSetTemplatePhase{{}}
	objectSet.SetPhases(phases)
	assert.Equal(t, phases, objectSet.GetPhases())

	objectSet.Spec.AvailabilityProbes = []corev1alpha1.ObjectSetProbe{{}}
	assert.Equal(t, objectSet.Spec.AvailabilityProbes,
		objectSet.GetAvailabilityProbes())

	var revision int64 = 34
	objectSet.SetRevision(revision)
	assert.Equal(t, revision, objectSet.GetRevision())

	objectSet.Spec.Previous = []corev1alpha1.PreviousRevisionReference{
		{},
	}
	assert.Equal(t, objectSet.Spec.Previous, objectSet.GetPrevious())

	remotes := []corev1alpha1.RemotePhaseReference{{}}
	objectSet.SetRemotePhases(remotes)
	assert.Equal(t, remotes, objectSet.GetRemotePhases())

	controllerOf := []corev1alpha1.ControlledObjectReference{{}}
	objectSet.SetStatusControllerOf(controllerOf)
	assert.Equal(t, controllerOf, objectSet.Status.ControllerOf)
}

func TestGenericClusterObjectSet(t *testing.T) {
	t.Parallel()

	objectSet := newGenericClusterObjectSet(testScheme).(*GenericClusterObjectSet)

	co := objectSet.ClientObject()
	assert.IsType(t, &corev1alpha1.ClusterObjectSet{}, co)

	objectSet.Status.Conditions = []metav1.Condition{{}}
	assert.Equal(t, objectSet.Status.Conditions, *objectSet.GetConditions())

	assert.False(t, objectSet.IsPaused())
	assert.False(t, objectSet.IsArchived())
	objectSet.Spec.LifecycleState = corev1alpha1.ObjectSetLifecycleStatePaused
	assert.True(t, objectSet.IsPaused())
	objectSet.Spec.LifecycleState = corev1alpha1.ObjectSetLifecycleStateArchived
	assert.True(t, objectSet.IsArchived())

	phases := []corev1alpha1.ObjectSetTemplatePhase{{}}
	objectSet.SetPhases(phases)
	assert.Equal(t, phases, objectSet.GetPhases())

	objectSet.Spec.AvailabilityProbes = []corev1alpha1.ObjectSetProbe{{}}
	assert.Equal(t, objectSet.Spec.AvailabilityProbes,
		objectSet.GetAvailabilityProbes())

	var revision int64 = 34
	objectSet.SetRevision(revision)
	assert.Equal(t, revision, objectSet.GetRevision())

	objectSet.Spec.Previous = []corev1alpha1.PreviousRevisionReference{
		{},
	}
	assert.Equal(t, objectSet.Spec.Previous, objectSet.GetPrevious())

	remotes := []corev1alpha1.RemotePhaseReference{{}}
	objectSet.SetRemotePhases(remotes)
	assert.Equal(t, remotes, objectSet.GetRemotePhases())

	controllerOf := []corev1alpha1.ControlledObjectReference{{}}
	objectSet.SetStatusControllerOf(controllerOf)
	assert.Equal(t, controllerOf, objectSet.Status.ControllerOf)
}
