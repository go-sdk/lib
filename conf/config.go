package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/imdario/mergo"
	"github.com/thoas/go-funk"

	"github.com/go-sdk/lib/codec/json"
	"github.com/go-sdk/lib/codec/yaml"
	"github.com/go-sdk/lib/val"
)

type Config struct {
	debug     bool
	skipError bool
	overwrite bool
}

type ConfigFunc = func(config *Config)

func WithDebug() ConfigFunc {
	return func(config *Config) {
		config.debug = true
	}
}

func WithSkipError() ConfigFunc {
	return func(config *Config) {
		config.skipError = true
	}
}

func WithOverwrite() ConfigFunc {
	return func(config *Config) {
		config.overwrite = true
	}
}

type Conf struct {
	*Config

	data map[string]interface{}
	env  map[string]string
}

func New(opts ...ConfigFunc) *Conf {
	config := &Conf{
		Config: &Config{},
		data:   map[string]interface{}{},
		env:    map[string]string{},
	}

	for i := 0; i < len(opts); i++ {
		opts[i](config.Config)
	}

	return config
}

// Load from files and env, env > files[n-1] > ... > files[0]
func Load(files ...string) error {
	return New().Load(files...)
}

func (conf *Conf) Load(paths ...string) error {
	for i := 0; i < len(paths); i++ {
		path, err := filepath.Abs(paths[i])
		if err != nil {
			continue
		}

		bs, err := os.ReadFile(path)
		if err := conf.err(err, "read config file (%s) fail", path); err != nil {
			return err
		}

		var v interface{}

		switch strings.ToLower(filepath.Ext(path)) {
		case ".yaml", ".yml":
			err = yaml.Unmarshal(bs, &v, yaml.WithCleanup())
		case ".json":
			err = json.Unmarshal(bs, &v)
		}

		if err := conf.err(err, "parse config file (%s) fail", path); err != nil {
			return err
		}

		if err != nil {
			continue
		}

		err = mergo.Merge(&conf.data, v, mergo.WithOverride, mergo.WithSliceDeepCopy)

		if err := conf.err(err, "merge config file (%s) fail", path); err != nil {
			return err
		}
	}

	conf.loadEnv()

	if conf.overwrite {
		_conf = conf
	}

	return nil
}

var (
	defaultConfigPath = []string{"config.yaml", "config.yml", "config.json"}

	_conf *Conf
)

func init() {
	_conf = New(WithSkipError())
	_ = _conf.Load(defaultConfigPath...)
}

func Get(key string) val.Value {
	return _conf.Get(key)
}

func (conf *Conf) Get(key string) val.Value {
	if v, ok := conf.env[key]; ok {
		return val.New(v)
	}
	v := funk.Get(conf.data, key, funk.WithAllowZero())
	return val.New(v)
}

func (conf *Conf) loadEnv() {
	es := os.Environ()
	for i := 0; i < len(es); i++ {
		ss := strings.SplitN(es[i], "=", 2)
		if len(ss) < 2 {
			continue
		}
		k, v := ss[0], ss[1]
		if k == "" {
			continue
		}
		conf.env[k] = v
	}
}

func (conf *Conf) err(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	err = fmt.Errorf(format+", "+err.Error(), args...)

	if conf.debug {
		fmt.Println(time.Now().Format("2006-01-02T15:04:05.999Z07:00") + " [config] " + err.Error())
	}

	if conf.skipError {
		return nil
	}

	return err
}
