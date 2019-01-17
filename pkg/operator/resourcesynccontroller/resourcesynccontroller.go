package resourcesynccontroller

import (
	"github.com/openshift/cluster-openshift-apiserver-operator/pkg/operator/operatorclient"
	"github.com/openshift/library-go/pkg/operator/events"
	"github.com/openshift/library-go/pkg/operator/resourcesynccontroller"
	"github.com/openshift/library-go/pkg/operator/v1helpers"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
)

func NewResourceSyncController(
	operatorConfigClient v1helpers.OperatorClient,
	kubeInformersForNamespaces map[string]informers.SharedInformerFactory,
	kubeClient kubernetes.Interface,
	eventRecorder events.Recorder) (*resourcesynccontroller.ResourceSyncController, error) {

	resourceSyncController := resourcesynccontroller.NewResourceSyncController(
		operatorConfigClient,
		kubeInformersForNamespaces,
		kubeClient,
		eventRecorder,
	)
	if err := resourceSyncController.SyncConfigMap(
		resourcesynccontroller.ResourceLocation{Namespace: operatorclient.TargetNamespaceName, Name: "etcd-serving-ca"},
		resourcesynccontroller.ResourceLocation{Namespace: operatorclient.EtcdNamespaceName, Name: "etcd-serving-ca"},
	); err != nil {
		return nil, err
	}
	if err := resourceSyncController.SyncSecret(
		resourcesynccontroller.ResourceLocation{Namespace: operatorclient.TargetNamespaceName, Name: "etcd-client"},
		resourcesynccontroller.ResourceLocation{Namespace: operatorclient.EtcdNamespaceName, Name: "etcd-client"},
	); err != nil {
		return nil, err
	}
	if err := resourceSyncController.SyncConfigMap(
		resourcesynccontroller.ResourceLocation{Namespace: operatorclient.TargetNamespaceName, Name: "aggregator-client-ca"},
		resourcesynccontroller.ResourceLocation{Namespace: operatorclient.MachineSpecifiedGlobalConfigNamespace, Name: "aggregator-client-ca"},
	); err != nil {
		return nil, err
	}
	if err := resourceSyncController.SyncConfigMap(
		resourcesynccontroller.ResourceLocation{Namespace: operatorclient.OperatorNamespace, Name: "initial-client-ca"},
		resourcesynccontroller.ResourceLocation{Namespace: operatorclient.UserSpecifiedGlobalConfigNamespace, Name: "initial-client-ca"},
	); err != nil {
		return nil, err
	}
	// this ca bundle contains certs used to sign CSRs (kubelet serving and client certificates)
	if err := resourceSyncController.SyncConfigMap(
		resourcesynccontroller.ResourceLocation{Namespace: operatorclient.OperatorNamespace, Name: "csr-controller-ca"},
		resourcesynccontroller.ResourceLocation{Namespace: operatorclient.MachineSpecifiedGlobalConfigNamespace, Name: "csr-controller-ca"},
	); err != nil {
		return nil, err
	}
	// this ca bundle contains certs used by the kube-apiserver to verify client certs
	if err := resourceSyncController.SyncConfigMap(
		resourcesynccontroller.ResourceLocation{Namespace: operatorclient.MachineSpecifiedGlobalConfigNamespace, Name: "kube-apiserver-client-ca"},
		resourcesynccontroller.ResourceLocation{Namespace: operatorclient.TargetNamespaceName, Name: "client-ca"},
	); err != nil {
		return nil, err
	}

	return resourceSyncController, nil
}
