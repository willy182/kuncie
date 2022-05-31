//

package kuncie

import (
	"github.com/golangid/candi/codebase/factory"
	"github.com/golangid/candi/codebase/factory/dependency"
	"github.com/golangid/candi/codebase/factory/types"
	"github.com/golangid/candi/config"

	"kuncie/configs"
	"kuncie/internal/modules/items"
	"kuncie/internal/modules/orders"
	"kuncie/internal/modules/products"
)

// Service model
type Service struct {
	cfg          *config.Config
	deps         dependency.Dependency
	applications []factory.AppServerFactory
	modules      []factory.ModuleFactory
	name         types.Service
}

// NewService in this service
func NewService(cfg *config.Config) factory.ServiceFactory {
	deps := configs.LoadServiceConfigs(cfg)

	modules := []factory.ModuleFactory{
		items.NewModule(deps),
		orders.NewModule(deps),
		products.NewModule(deps),
	}

	s := &Service{
		cfg:     cfg,
		deps:    deps,
		modules: modules,
		name:    types.Service(cfg.ServiceName),
	}

	s.applications = configs.InitAppFromEnvironmentConfig(s)

	// Add custom application runner, must implement `factory.AppServerFactory` methods
	s.applications = append(s.applications, []factory.AppServerFactory{
		// customApplication
	}...)

	return s
}

// GetConfig method
func (s *Service) GetConfig() *config.Config {
	return s.cfg
}

// GetDependency method
func (s *Service) GetDependency() dependency.Dependency {
	return s.deps
}

// GetApplications method
func (s *Service) GetApplications() []factory.AppServerFactory {
	return s.applications
}

// GetModules method
func (s *Service) GetModules() []factory.ModuleFactory {
	return s.modules
}

// Name method
func (s *Service) Name() types.Service {
	return s.name
}
