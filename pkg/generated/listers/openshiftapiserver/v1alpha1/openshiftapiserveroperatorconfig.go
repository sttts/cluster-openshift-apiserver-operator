// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/cluster-openshift-apiserver-operator/pkg/apis/openshiftapiserver/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OpenShiftAPIServerOperatorConfigLister helps list OpenShiftAPIServerOperatorConfigs.
type OpenShiftAPIServerOperatorConfigLister interface {
	// List lists all OpenShiftAPIServerOperatorConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.OpenShiftAPIServerOperatorConfig, err error)
	// Get retrieves the OpenShiftAPIServerOperatorConfig from the index for a given name.
	Get(name string) (*v1alpha1.OpenShiftAPIServerOperatorConfig, error)
	OpenShiftAPIServerOperatorConfigListerExpansion
}

// openShiftAPIServerOperatorConfigLister implements the OpenShiftAPIServerOperatorConfigLister interface.
type openShiftAPIServerOperatorConfigLister struct {
	indexer cache.Indexer
}

// NewOpenShiftAPIServerOperatorConfigLister returns a new OpenShiftAPIServerOperatorConfigLister.
func NewOpenShiftAPIServerOperatorConfigLister(indexer cache.Indexer) OpenShiftAPIServerOperatorConfigLister {
	return &openShiftAPIServerOperatorConfigLister{indexer: indexer}
}

// List lists all OpenShiftAPIServerOperatorConfigs in the indexer.
func (s *openShiftAPIServerOperatorConfigLister) List(selector labels.Selector) (ret []*v1alpha1.OpenShiftAPIServerOperatorConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpenShiftAPIServerOperatorConfig))
	})
	return ret, err
}

// Get retrieves the OpenShiftAPIServerOperatorConfig from the index for a given name.
func (s *openShiftAPIServerOperatorConfigLister) Get(name string) (*v1alpha1.OpenShiftAPIServerOperatorConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("openshiftapiserveroperatorconfig"), name)
	}
	return obj.(*v1alpha1.OpenShiftAPIServerOperatorConfig), nil
}
