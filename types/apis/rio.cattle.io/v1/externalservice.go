package v1

import (
	"github.com/rancher/norman/condition"
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ExternalService struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ExternalServiceSpec   `json:"spec,omitempty"`
	Status            ExternalServiceStatus `json:"status,omitempty"`
}

type ExternalServiceSpec struct {
	// External service located outside mesh, represented by IPs
	IPAddresses []string `json:"ipAddresses,omitempty"`
	// External service located outside mesh, represented by DNS
	FQDN string `json:"fqdn,omitempty"`
	// In-Mesh service name in another stack in current project
	Service string `json:"service,omitempty"`
	StackScoped
}

type ExternalServiceStatus struct {
	Conditions []condition.GenericCondition `json:"conditions,omitempty"`
}
