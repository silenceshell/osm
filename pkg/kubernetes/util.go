package kubernetes

import (
	"fmt"

	corev1 "k8s.io/api/core/v1"
)

const (
	clusterDomain = "cluster.local"
)

// GetDomainsForService returns a list of domains over which the service
// can be accessed within the local cluster.
func GetDomainsForService(service *corev1.Service) []string {
	var domains []string
	if service == nil {
		return domains
	}

	serviceName := service.Name
	namespace := service.Namespace

	domains = append(domains, serviceName)                                                        // service
	domains = append(domains, fmt.Sprintf("%s.%s", serviceName, namespace))                       // service.namespace
	domains = append(domains, fmt.Sprintf("%s.%s.svc", serviceName, namespace))                   // service.namespace.svc
	domains = append(domains, fmt.Sprintf("%s.%s.svc.cluster", serviceName, namespace))           // service.namespace.svc.cluster
	domains = append(domains, fmt.Sprintf("%s.%s.svc.%s", serviceName, namespace, clusterDomain)) // service.namespace.svc.cluster.local
	for _, portSpec := range service.Spec.Ports {
		port := portSpec.Port
		domains = append(domains, fmt.Sprintf("%s:%d", serviceName, port))                                     // service:port
		domains = append(domains, fmt.Sprintf("%s.%s:%d", serviceName, namespace, port))                       // service.namespace:port
		domains = append(domains, fmt.Sprintf("%s.%s.svc:%d", serviceName, namespace, port))                   // service.namespace.svc:port
		domains = append(domains, fmt.Sprintf("%s.%s.svc.cluster:%d", serviceName, namespace, port))           // service.namespace.svc.cluster:port
		domains = append(domains, fmt.Sprintf("%s.%s.svc.%s:%d", serviceName, namespace, clusterDomain, port)) // service.namespace.svc.cluster.local:port
	}
	return domains
}
