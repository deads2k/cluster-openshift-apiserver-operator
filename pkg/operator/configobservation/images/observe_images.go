package images

import (
	"bytes"
	"encoding/json"

	"github.com/golang/glog"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/openshift/library-go/pkg/operator/configobserver"
	"github.com/openshift/library-go/pkg/operator/events"

	"github.com/openshift/cluster-openshift-apiserver-operator/pkg/operator/configobservation"
)

func ObserveInternalRegistryHostname(genericListers configobserver.Listers, eventRecorder events.Recorder, existingConfig map[string]interface{}) (map[string]interface{}, []error) {
	listers := genericListers.(configobservation.Listers)
	var errs []error
	prevObservedConfig := map[string]interface{}{}

	// first observe all the existing config values so that if we get any errors
	// we can at least return those.
	internalRegistryHostnamePath := []string{"imagePolicyConfig", "internalRegistryHostname"}
	currentInternalRegistryHostname, _, err := unstructured.NestedString(existingConfig, internalRegistryHostnamePath...)
	if err != nil {
		return prevObservedConfig, append(errs, err)
	}
	if len(currentInternalRegistryHostname) > 0 {
		err := unstructured.SetNestedField(prevObservedConfig, currentInternalRegistryHostname, internalRegistryHostnamePath...)
		if err != nil {
			return prevObservedConfig, append(errs, err)
		}
	}

	if !listers.ImageConfigSynced() {
		glog.Warning("images.config.openshift.io not synced")
		return prevObservedConfig, errs
	}

	// now gather the cluster config and turn it into the observed config
	observedConfig := map[string]interface{}{}
	configImage, err := listers.ImageConfigLister.Get("cluster")
	if errors.IsNotFound(err) {
		glog.Warningf("image.config.openshift.io/cluster: not found")
		return observedConfig, errs
	}
	if err != nil {
		return prevObservedConfig, append(errs, err)
	}

	internalRegistryHostName := configImage.Status.InternalRegistryHostname
	if len(internalRegistryHostName) > 0 {
		err = unstructured.SetNestedField(observedConfig, internalRegistryHostName, internalRegistryHostnamePath...)
		if err != nil {
			return prevObservedConfig, append(errs, err)
		}
	}

	return observedConfig, errs
}

func ObserveExternalRegistryHostnames(genericListers configobserver.Listers, eventRecorder events.Recorder, existingConfig map[string]interface{}) (map[string]interface{}, []error) {
	listers := genericListers.(configobservation.Listers)
	var errs []error
	prevObservedConfig := map[string]interface{}{}

	// first observe all the existing config values so that if we get any errors
	// we can at least return those.
	externalRegistryHostnamePath := []string{"imagePolicyConfig", "externalRegistryHostnames"}
	o, _, err := unstructured.NestedSlice(existingConfig, externalRegistryHostnamePath...)
	if err != nil {
		return prevObservedConfig, append(errs, err)
	}
	err = unstructured.SetNestedSlice(prevObservedConfig, o, externalRegistryHostnamePath...)
	if err != nil {
		return prevObservedConfig, append(errs, err)
	}

	if !listers.ImageConfigSynced() {
		glog.Warning("images.config.openshift.io not synced")
		return prevObservedConfig, errs
	}

	// now gather the cluster config and turn it into the observed config
	observedConfig := map[string]interface{}{}
	configImage, err := listers.ImageConfigLister.Get("cluster")
	if errors.IsNotFound(err) {
		glog.Warningf("image.config.openshift.io/cluster: not found")
		return observedConfig, errs
	}
	if err != nil {
		return prevObservedConfig, append(errs, err)
	}

	// User provided values take precedence, first entry in the array
	// has special significance.
	externalRegistryHostnames := configImage.Spec.ExternalRegistryHostnames
	externalRegistryHostnames = append(externalRegistryHostnames, configImage.Status.ExternalRegistryHostnames...)

	hostnames, err := Convert(externalRegistryHostnames)
	if err != nil {
		return prevObservedConfig, append(errs, err)
	}
	err = unstructured.SetNestedField(observedConfig, hostnames, externalRegistryHostnamePath...)
	if err != nil {
		return prevObservedConfig, append(errs, err)
	}

	return observedConfig, errs
}

func ObserveAllowedRegistriesForImport(genericListers configobserver.Listers, eventRecorder events.Recorder, existingConfig map[string]interface{}) (map[string]interface{}, []error) {
	listers := genericListers.(configobservation.Listers)
	var errs []error
	prevObservedConfig := map[string]interface{}{}

	// first observe all the existing config values so that if we get any errors
	// we can at least return those.
	allowedRegistriesForImportPath := []string{"imagePolicyConfig", "allowedRegistriesForImport"}
	o, _, err := unstructured.NestedSlice(existingConfig, allowedRegistriesForImportPath...)
	if err != nil {
		return prevObservedConfig, append(errs, err)
	}
	err = unstructured.SetNestedSlice(prevObservedConfig, o, allowedRegistriesForImportPath...)
	if err != nil {
		return prevObservedConfig, append(errs, err)
	}

	if !listers.ImageConfigSynced() {
		glog.Warning("images.config.openshift.io not synced")
		return prevObservedConfig, errs
	}

	// now gather the cluster config and turn it into the observed config
	observedConfig := map[string]interface{}{}
	configImage, err := listers.ImageConfigLister.Get("cluster")
	if errors.IsNotFound(err) {
		glog.Warningf("image.config.openshift.io/cluster: not found")
		return observedConfig, errs
	}
	if err != nil {
		return prevObservedConfig, append(errs, err)
	}

	allowed, err := Convert(configImage.Spec.AllowedRegistriesForImport)
	if err != nil {
		return prevObservedConfig, append(errs, err)
	}
	err = unstructured.SetNestedField(observedConfig, allowed, allowedRegistriesForImportPath...)
	if err != nil {
		return prevObservedConfig, append(errs, err)
	}

	return observedConfig, errs
}

func Convert(o interface{}) (interface{}, error) {
	if o == nil {
		return nil, nil
	}
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(o); err != nil {
		return nil, err
	}

	ret := []interface{}{}
	if err := json.NewDecoder(buf).Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}
