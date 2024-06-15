package core

import (
	"errors"
	service_registry "service_registry/core/models"
)

func registerService(service service_registry.Service) (service_registry.Id, error) {
	return "", errors.New("Not implemented")
}

func deregisterService(serviceId service_registry.Id) (service_registry.Id, error) {
	return "", errors.New("Not implemented")
}
