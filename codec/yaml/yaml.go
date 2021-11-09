package yaml

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

var (
	Marshal = yaml.Marshal
	// Unmarshal = yaml.Unmarshal

	NewEncoder = yaml.NewEncoder
	NewDecoder = yaml.NewDecoder
)

func MustMarshal(v interface{}) []byte {
	bs, _ := Marshal(v)
	return bs
}

func MustMarshalToString(v interface{}) string {
	return string(MustMarshal(v))
}

type Config struct {
	cleanup bool
}

type ConfigFunc = func(config *Config)

func WithCleanup() ConfigFunc {
	return func(config *Config) {
		config.cleanup = true
	}
}

func Unmarshal(bs []byte, v interface{}, opts ...ConfigFunc) error {
	config := &Config{}
	for i := 0; i < len(opts); i++ {
		opts[i](config)
	}

	if config.cleanup {
		var x interface{}
		if err := yaml.Unmarshal(bs, &x); err != nil {
			return err
		}
		*v.(*interface{}) = value(x)
		return nil
	}

	return yaml.Unmarshal(bs, v)
}

func UnmarshalFromString(s string, v interface{}, opts ...ConfigFunc) error {
	return Unmarshal([]byte(s), v, opts...)
}

func value(in interface{}) interface{} {
	switch x := in.(type) {
	case []interface{}:
		y := make([]interface{}, len(x))
		for i, v := range x {
			y[i] = value(v)
		}
		return y
	case map[string]interface{}:
		y := make(map[string]interface{})
		for k, v := range x {
			y[k] = value(v)
		}
		return y
	case map[interface{}]interface{}:
		y := make(map[string]interface{})
		for k, v := range x {
			y[fmt.Sprintf("%v", k)] = value(v)
		}
		return y
	default:
		return x
	}
}
