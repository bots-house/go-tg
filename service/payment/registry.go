package payment

import (
	"sync"
)

// GatewayRegistry contains registries
type GatewayRegistry struct {
	gateways map[string]Gateway
	lock     sync.RWMutex
}

// NewGatewayRegistry creates a gateway registry.
func NewGatewayRegistry(gateways ...Gateway) *GatewayRegistry {
	registry := &GatewayRegistry{
		gateways: make(map[string]Gateway),
	}

	for _, gateway := range gateways {
		registry.Add(gateway)
	}

	return registry
}

// Add gateway to registry.
func (registry *GatewayRegistry) Add(gateway Gateway) {
	registry.lock.Lock()
	registry.gateways[gateway.Name()] = gateway
	registry.lock.Unlock()
}

// Find registry by name.
func (registry *GatewayRegistry) Get(name string) Gateway {
	registry.lock.RLock()
	gateway := registry.gateways[name]
	registry.lock.RUnlock()

	return gateway
}

// List of gateways.
func (registry *GatewayRegistry) List() []Gateway {
	registry.lock.RLock()
	defer registry.lock.RUnlock()

	gateways := make([]Gateway, 0, len(registry.gateways))

	for _, gateway := range registry.gateways {
		gateways = append(gateways, gateway)
	}

	return gateways
}
