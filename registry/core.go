package registry

import (
	"context"
	kreg "github.com/go-kratos/kratos/v2/registry"
)

// Registrar is service registrar.
type Registrar interface {
	kreg.Registrar
	// Update the registration.
	Update(ctx context.Context, service *kreg.ServiceInstance) error
}

type Discovery = kreg.Discovery

type ServiceInstance = kreg.ServiceInstance
