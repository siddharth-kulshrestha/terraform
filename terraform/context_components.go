package terraform

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/addrs"
	"github.com/hashicorp/terraform/providers"
	"github.com/hashicorp/terraform/provisioners"
)

// contextComponentFactory is the interface that Context uses
// to initialize various components such as providers and provisioners.
// This factory gets more information than the raw maps using to initialize
// a Context. This information is used for debugging.
type contextComponentFactory interface {
	// ResourceProvider creates a new ResourceProvider with the given type.
	ResourceProvider(typ addrs.Provider) (providers.Interface, error)
	ResourceProviders() []string

	// ResourceProvisioner creates a new ResourceProvisioner with the given
	// type.
	ResourceProvisioner(typ string) (provisioners.Interface, error)
	ResourceProvisioners() []string
}

// basicComponentFactory just calls a factory from a map directly.
type basicComponentFactory struct {
	providers    map[addrs.Provider]providers.Factory
	provisioners map[string]provisioners.Factory
}

func (c *basicComponentFactory) ResourceProviders() []string {
	var result []string
	log.Println("CUSTOM_LOG_SID: calling basicComponentFactory::ResourceProviders!")
	for k := range c.providers {
		result = append(result, k.String())
	}
	return result
}

func (c *basicComponentFactory) ResourceProvisioners() []string {
	var result []string
	for k := range c.provisioners {
		result = append(result, k)
	}

	return result
}

func (c *basicComponentFactory) ResourceProvider(typ addrs.Provider) (providers.Interface, error) {
	log.Println("CUSTOM_LOG_SID: in ResourceProvider:")
	log.Println(c.providers)
	log.Println(typ)
	f, ok := c.providers[typ]
	if !ok {
		return nil, fmt.Errorf("unknown provider %q", typ.String())
	}
	log.Println("Here plugin is starting up!!!")
	p, err := f()
	log.Println("CUSTOM_LOG_SID: here is the response from provider factory method: ")
	log.Println(p)
	log.Printf("%T\n", p)
	return p, err
}

func (c *basicComponentFactory) ResourceProvisioner(typ string) (provisioners.Interface, error) {
	f, ok := c.provisioners[typ]
	if !ok {
		return nil, fmt.Errorf("unknown provisioner %q", typ)
	}

	return f()
}
