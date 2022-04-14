// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "elasticgpu.io/elastic-gpu/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ElasticGPUs returns a ElasticGPUInformer.
	ElasticGPUs() ElasticGPUInformer
	// ElasticGPUClaims returns a ElasticGPUClaimInformer.
	ElasticGPUClaims() ElasticGPUClaimInformer
	// ElasticGPUClasses returns a ElasticGPUClassInformer.
	ElasticGPUClasses() ElasticGPUClassInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ElasticGPUs returns a ElasticGPUInformer.
func (v *version) ElasticGPUs() ElasticGPUInformer {
	return &elasticGPUInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ElasticGPUClaims returns a ElasticGPUClaimInformer.
func (v *version) ElasticGPUClaims() ElasticGPUClaimInformer {
	return &elasticGPUClaimInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ElasticGPUClasses returns a ElasticGPUClassInformer.
func (v *version) ElasticGPUClasses() ElasticGPUClassInformer {
	return &elasticGPUClassInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
