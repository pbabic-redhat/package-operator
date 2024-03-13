//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	corev1alpha1 "package-operator.run/apis/core/v1alpha1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageEnvironment) DeepCopyInto(out *PackageEnvironment) {
	*out = *in
	out.Kubernetes = in.Kubernetes
	if in.OpenShift != nil {
		in, out := &in.OpenShift, &out.OpenShift
		*out = new(PackageEnvironmentOpenShift)
		**out = **in
	}
	if in.Proxy != nil {
		in, out := &in.Proxy, &out.Proxy
		*out = new(PackageEnvironmentProxy)
		**out = **in
	}
	if in.HyperShift != nil {
		in, out := &in.HyperShift, &out.HyperShift
		*out = new(PackageEnvironmentHyperShift)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageEnvironment.
func (in *PackageEnvironment) DeepCopy() *PackageEnvironment {
	if in == nil {
		return nil
	}
	out := new(PackageEnvironment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageEnvironmentHyperShift) DeepCopyInto(out *PackageEnvironmentHyperShift) {
	*out = *in
	if in.HostedCluster != nil {
		in, out := &in.HostedCluster, &out.HostedCluster
		*out = new(PackageEnvironmentHyperShiftHostedCluster)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageEnvironmentHyperShift.
func (in *PackageEnvironmentHyperShift) DeepCopy() *PackageEnvironmentHyperShift {
	if in == nil {
		return nil
	}
	out := new(PackageEnvironmentHyperShift)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageEnvironmentHyperShiftHostedCluster) DeepCopyInto(out *PackageEnvironmentHyperShiftHostedCluster) {
	*out = *in
	in.TemplateContextObjectMeta.DeepCopyInto(&out.TemplateContextObjectMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageEnvironmentHyperShiftHostedCluster.
func (in *PackageEnvironmentHyperShiftHostedCluster) DeepCopy() *PackageEnvironmentHyperShiftHostedCluster {
	if in == nil {
		return nil
	}
	out := new(PackageEnvironmentHyperShiftHostedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageEnvironmentKubernetes) DeepCopyInto(out *PackageEnvironmentKubernetes) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageEnvironmentKubernetes.
func (in *PackageEnvironmentKubernetes) DeepCopy() *PackageEnvironmentKubernetes {
	if in == nil {
		return nil
	}
	out := new(PackageEnvironmentKubernetes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageEnvironmentOpenShift) DeepCopyInto(out *PackageEnvironmentOpenShift) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageEnvironmentOpenShift.
func (in *PackageEnvironmentOpenShift) DeepCopy() *PackageEnvironmentOpenShift {
	if in == nil {
		return nil
	}
	out := new(PackageEnvironmentOpenShift)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageEnvironmentProxy) DeepCopyInto(out *PackageEnvironmentProxy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageEnvironmentProxy.
func (in *PackageEnvironmentProxy) DeepCopy() *PackageEnvironmentProxy {
	if in == nil {
		return nil
	}
	out := new(PackageEnvironmentProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifest) DeepCopyInto(out *PackageManifest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Test.DeepCopyInto(&out.Test)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifest.
func (in *PackageManifest) DeepCopy() *PackageManifest {
	if in == nil {
		return nil
	}
	out := new(PackageManifest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageManifest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestCelMacro) DeepCopyInto(out *PackageManifestCelMacro) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestCelMacro.
func (in *PackageManifestCelMacro) DeepCopy() *PackageManifestCelMacro {
	if in == nil {
		return nil
	}
	out := new(PackageManifestCelMacro)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestComponentsConfig) DeepCopyInto(out *PackageManifestComponentsConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestComponentsConfig.
func (in *PackageManifestComponentsConfig) DeepCopy() *PackageManifestComponentsConfig {
	if in == nil {
		return nil
	}
	out := new(PackageManifestComponentsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestConditionalPath) DeepCopyInto(out *PackageManifestConditionalPath) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestConditionalPath.
func (in *PackageManifestConditionalPath) DeepCopy() *PackageManifestConditionalPath {
	if in == nil {
		return nil
	}
	out := new(PackageManifestConditionalPath)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestImage) DeepCopyInto(out *PackageManifestImage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestImage.
func (in *PackageManifestImage) DeepCopy() *PackageManifestImage {
	if in == nil {
		return nil
	}
	out := new(PackageManifestImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestLock) DeepCopyInto(out *PackageManifestLock) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestLock.
func (in *PackageManifestLock) DeepCopy() *PackageManifestLock {
	if in == nil {
		return nil
	}
	out := new(PackageManifestLock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageManifestLock) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestLockImage) DeepCopyInto(out *PackageManifestLockImage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestLockImage.
func (in *PackageManifestLockImage) DeepCopy() *PackageManifestLockImage {
	if in == nil {
		return nil
	}
	out := new(PackageManifestLockImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestLockSpec) DeepCopyInto(out *PackageManifestLockSpec) {
	*out = *in
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]PackageManifestLockImage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestLockSpec.
func (in *PackageManifestLockSpec) DeepCopy() *PackageManifestLockSpec {
	if in == nil {
		return nil
	}
	out := new(PackageManifestLockSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestPhase) DeepCopyInto(out *PackageManifestPhase) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestPhase.
func (in *PackageManifestPhase) DeepCopy() *PackageManifestPhase {
	if in == nil {
		return nil
	}
	out := new(PackageManifestPhase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestSpec) DeepCopyInto(out *PackageManifestSpec) {
	*out = *in
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]PackageManifestScope, len(*in))
		copy(*out, *in)
	}
	if in.Phases != nil {
		in, out := &in.Phases, &out.Phases
		*out = make([]PackageManifestPhase, len(*in))
		copy(*out, *in)
	}
	if in.AvailabilityProbes != nil {
		in, out := &in.AvailabilityProbes, &out.AvailabilityProbes
		*out = make([]corev1alpha1.ObjectSetProbe, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Config.DeepCopyInto(&out.Config)
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]PackageManifestImage, len(*in))
		copy(*out, *in)
	}
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = new(PackageManifestComponentsConfig)
		**out = **in
	}
	if in.CelMacros != nil {
		in, out := &in.CelMacros, &out.CelMacros
		*out = make([]PackageManifestCelMacro, len(*in))
		copy(*out, *in)
	}
	if in.ConditionalPaths != nil {
		in, out := &in.ConditionalPaths, &out.ConditionalPaths
		*out = make([]PackageManifestConditionalPath, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestSpec.
func (in *PackageManifestSpec) DeepCopy() *PackageManifestSpec {
	if in == nil {
		return nil
	}
	out := new(PackageManifestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestSpecConfig) DeepCopyInto(out *PackageManifestSpecConfig) {
	*out = *in
	if in.OpenAPIV3Schema != nil {
		in, out := &in.OpenAPIV3Schema, &out.OpenAPIV3Schema
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestSpecConfig.
func (in *PackageManifestSpecConfig) DeepCopy() *PackageManifestSpecConfig {
	if in == nil {
		return nil
	}
	out := new(PackageManifestSpecConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestTest) DeepCopyInto(out *PackageManifestTest) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = make([]PackageManifestTestCaseTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Kubeconform != nil {
		in, out := &in.Kubeconform, &out.Kubeconform
		*out = new(PackageManifestTestKubeconform)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestTest.
func (in *PackageManifestTest) DeepCopy() *PackageManifestTest {
	if in == nil {
		return nil
	}
	out := new(PackageManifestTest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestTestCaseTemplate) DeepCopyInto(out *PackageManifestTestCaseTemplate) {
	*out = *in
	in.Context.DeepCopyInto(&out.Context)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestTestCaseTemplate.
func (in *PackageManifestTestCaseTemplate) DeepCopy() *PackageManifestTestCaseTemplate {
	if in == nil {
		return nil
	}
	out := new(PackageManifestTestCaseTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageManifestTestKubeconform) DeepCopyInto(out *PackageManifestTestKubeconform) {
	*out = *in
	if in.SchemaLocations != nil {
		in, out := &in.SchemaLocations, &out.SchemaLocations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageManifestTestKubeconform.
func (in *PackageManifestTestKubeconform) DeepCopy() *PackageManifestTestKubeconform {
	if in == nil {
		return nil
	}
	out := new(PackageManifestTestKubeconform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateContext) DeepCopyInto(out *TemplateContext) {
	*out = *in
	in.Package.DeepCopyInto(&out.Package)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	in.Environment.DeepCopyInto(&out.Environment)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateContext.
func (in *TemplateContext) DeepCopy() *TemplateContext {
	if in == nil {
		return nil
	}
	out := new(TemplateContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateContextObjectMeta) DeepCopyInto(out *TemplateContextObjectMeta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateContextObjectMeta.
func (in *TemplateContextObjectMeta) DeepCopy() *TemplateContextObjectMeta {
	if in == nil {
		return nil
	}
	out := new(TemplateContextObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateContextPackage) DeepCopyInto(out *TemplateContextPackage) {
	*out = *in
	in.TemplateContextObjectMeta.DeepCopyInto(&out.TemplateContextObjectMeta)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateContextPackage.
func (in *TemplateContextPackage) DeepCopy() *TemplateContextPackage {
	if in == nil {
		return nil
	}
	out := new(TemplateContextPackage)
	in.DeepCopyInto(out)
	return out
}
