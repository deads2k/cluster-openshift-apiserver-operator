// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// OpenShiftControllerManagerStatusApplyConfiguration represents an declarative configuration of the OpenShiftControllerManagerStatus type for use
// with apply.
type OpenShiftControllerManagerStatusApplyConfiguration struct {
	OperatorStatusApplyConfiguration `json:",inline"`
}

// OpenShiftControllerManagerStatusApplyConfiguration constructs an declarative configuration of the OpenShiftControllerManagerStatus type for use with
// apply.
func OpenShiftControllerManagerStatus() *OpenShiftControllerManagerStatusApplyConfiguration {
	return &OpenShiftControllerManagerStatusApplyConfiguration{}
}

// WithObservedGeneration sets the ObservedGeneration field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ObservedGeneration field is set to the value of the last call.
func (b *OpenShiftControllerManagerStatusApplyConfiguration) WithObservedGeneration(value int64) *OpenShiftControllerManagerStatusApplyConfiguration {
	b.ObservedGeneration = &value
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *OpenShiftControllerManagerStatusApplyConfiguration) WithConditions(values ...*OperatorConditionApplyConfiguration) *OpenShiftControllerManagerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithVersion sets the Version field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Version field is set to the value of the last call.
func (b *OpenShiftControllerManagerStatusApplyConfiguration) WithVersion(value string) *OpenShiftControllerManagerStatusApplyConfiguration {
	b.Version = &value
	return b
}

// WithReadyReplicas sets the ReadyReplicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadyReplicas field is set to the value of the last call.
func (b *OpenShiftControllerManagerStatusApplyConfiguration) WithReadyReplicas(value int32) *OpenShiftControllerManagerStatusApplyConfiguration {
	b.ReadyReplicas = &value
	return b
}

// WithLatestAvailableRevision sets the LatestAvailableRevision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LatestAvailableRevision field is set to the value of the last call.
func (b *OpenShiftControllerManagerStatusApplyConfiguration) WithLatestAvailableRevision(value int32) *OpenShiftControllerManagerStatusApplyConfiguration {
	b.LatestAvailableRevision = &value
	return b
}

// WithGenerations adds the given value to the Generations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Generations field.
func (b *OpenShiftControllerManagerStatusApplyConfiguration) WithGenerations(values ...*GenerationStatusApplyConfiguration) *OpenShiftControllerManagerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGenerations")
		}
		b.Generations = append(b.Generations, *values[i])
	}
	return b
}
