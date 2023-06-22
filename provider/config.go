package provider

import (
	"fmt"
)

// provider types
type (
	IntProvider interface {
		IntGetter() func() (int64, error)
	}
	StringProvider interface {
		StringGetter() func() (string, error)
	}
	FloatProvider interface {
		FloatGetter() func() (float64, error)
	}
	BoolProvider interface {
		BoolGetter() func() (bool, error)
	}
	SetIntProvider interface {
		IntSetter(param string) func(int64) error
	}
	SetStringProvider interface {
		StringSetter(param string) func(string) error
	}
	SetBoolProvider interface {
		BoolSetter(param string) func(bool) error
	}
)

type providerRegistry map[string]func(map[string]interface{}) (IntProvider, error)

func (r providerRegistry) Add(name string, factory func(map[string]interface{}) (IntProvider, error)) {
	if _, exists := r[name]; exists {
		panic(fmt.Sprintf("cannot register duplicate plugin type: %s", name))
	}
	r[name] = factory
}

func (r providerRegistry) Get(name string) (func(map[string]interface{}) (IntProvider, error), error) {
	factory, exists := r[name]
	if !exists {
		return nil, fmt.Errorf("invalid plugin source: %s", name)
	}
	return factory, nil
}

var registry providerRegistry = make(map[string]func(map[string]interface{}) (IntProvider, error))

// Config is the general provider config
type Config struct {
	Source string
	Other  map[string]interface{} `mapstructure:",remain"`
}

// NewIntGetterFromConfig creates a IntGetter from config
func NewIntGetterFromConfig(config Config) (res func() (int64, error), err error) {
	factory, err := registry.Get(config.Source)
	if err == nil {
		var provider IntProvider
		provider, err = factory(config.Other)

		if err == nil {
			res = provider.IntGetter()
		}
	}

	if err == nil && res == nil {
		err = fmt.Errorf("invalid plugin source: %s", config.Source)
	}

	return
}

// NewFloatGetterFromConfig creates a FloatGetter from config
func NewFloatGetterFromConfig(config Config) (res func() (float64, error), err error) {
	factory, err := registry.Get(config.Source)
	if err == nil {
		var provider IntProvider
		provider, err = factory(config.Other)

		if prov, ok := provider.(FloatProvider); ok {
			res = prov.FloatGetter()
		}
	}

	if err == nil && res == nil {
		err = fmt.Errorf("invalid plugin source: %s", config.Source)
	}

	return
}

// NewStringGetterFromConfig creates a StringGetter from config
func NewStringGetterFromConfig(config Config) (res func() (string, error), err error) {
	switch typ := config.Source; typ {
	case "combined", "openwb":
		res, err = NewOpenWBStatusProviderFromConfig(config.Other)

	default:
		var factory func(map[string]interface{}) (IntProvider, error)
		factory, err = registry.Get(typ)
		if err == nil {
			var provider IntProvider
			provider, err = factory(config.Other)

			if prov, ok := provider.(StringProvider); ok {
				res = prov.StringGetter()
			}
		}

		if err == nil && res == nil {
			err = fmt.Errorf("invalid plugin source: %s", config.Source)
		}
	}

	return
}

// NewBoolGetterFromConfig creates a BoolGetter from config
func NewBoolGetterFromConfig(config Config) (res func() (bool, error), err error) {
	factory, err := registry.Get(config.Source)
	if err == nil {
		var provider IntProvider
		provider, err = factory(config.Other)

		if prov, ok := provider.(BoolProvider); ok {
			res = prov.BoolGetter()
		}
	}

	if err == nil && res == nil {
		err = fmt.Errorf("invalid plugin source: %s", config.Source)
	}

	return
}

// NewIntSetterFromConfig creates a IntSetter from config
func NewIntSetterFromConfig(param string, config Config) (res func(int64) error, err error) {
	factory, err := registry.Get(config.Source)
	if err == nil {
		var provider IntProvider
		provider, err = factory(config.Other)

		if prov, ok := provider.(SetIntProvider); ok {
			res = prov.IntSetter(param)
		}
	}

	if err == nil && res == nil {
		err = fmt.Errorf("invalid plugin source: %s", config.Source)
	}

	return
}

// NewBoolSetterFromConfig creates a BoolSetter from config
func NewBoolSetterFromConfig(param string, config Config) (res func(bool) error, err error) {
	factory, err := registry.Get(config.Source)
	if err == nil {
		var provider IntProvider
		provider, err = factory(config.Other)

		if prov, ok := provider.(SetBoolProvider); ok {
			res = prov.BoolSetter(param)
		}
	}

	if err == nil && res == nil {
		err = fmt.Errorf("invalid plugin source: %s", config.Source)
	}

	return
}
