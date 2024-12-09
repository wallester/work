package devenv

import (
	"github.com/juju/errors"
	"github.com/wallester/monorepo/pkg/common/slices"
)

// ServicesToMap converts a slice of Service objects into a map with service short names and directories as keys.
func ServicesToMap(services []*Service) (map[string]*Service, error) {
	res := make(map[string]*Service, len(services)*2)
	for _, service := range services {
		if res[service.ShortName] != nil {
			return nil, errors.Errorf("duplicate service short name: %s", service.ShortName)
		}

		res[service.ShortName] = service

		// e.g. api
		if service.ShortName == service.Directory {
			continue
		}

		if res[service.Directory] != nil {
			return nil, errors.Errorf("duplicate service directory found: %s", service.Directory)
		}

		res[service.Directory] = service
	}

	return res, nil
}

// FilterServices returns a subset of services which match the given names. If no names are provided, all services are returned.
func FilterServices(services []*Service, names []string) []*Service {
	if len(names) == 0 {
		return services
	}

	nameSet := slices.ToBoolMap(names)

	return slices.Select(services, func(service *Service) bool { return nameSet[service.ShortName] || nameSet[service.Directory] })
}

// FilterManagedServices returns a subset of services that are managed by mu
func FilterManagedServices(services []*Service) []*Service {
	return slices.Reject(services, func(service *Service) bool { return service.SkipManagement })
}
