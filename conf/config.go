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

type Option struct {
	debug     bool
	skipError bool
	overwrite bool
}

type OptionFunc = func(config *Option)

func WithDebug(t bool) OptionFunc {
	return func(config *Option) {
		config.debug = t
	}
}

func WithSkipError(t bool) OptionFunc {
	return func(config *Option) {
		config.skipError = t
	}
}

func WithOverwrite(t bool) OptionFunc {
	return func(config *Option) {
		config.overwrite = t
	}
}

type Config struct {
	*Option

	files []string

	data map[string]interface{}
	env  map[string]string
	args map[string]string
}

func New(opts ...OptionFunc) *Config {
	conf := &Config{
		Option: &Option{},
		files:  []string{},
		data:   map[string]interface{}{},
		env:    map[string]string{},
		args:   map[string]string{},
	}

	for i := 0; i < len(opts); i++ {
		opts[i](conf.Option)
	}

	return conf
}

// Load from files and env, env > files[n-1] > ... > files[0]
func Load(files ...string) error {
	return New().Load(files...)
}

func (conf *Config) Load(paths ...string) error {
	for i := 0; i < len(paths); i++ {
		path, err := filepath.Abs(paths[i])
		if err != nil {
			continue
		}

		bs, err := os.ReadFile(path)
		if err != nil {
			if x := conf.log(err, "read config file (%s) fail", path); x != nil {
				return x
			}
			continue
		}

		var v interface{}

		switch strings.ToLower(filepath.Ext(path)) {
		case ".yaml", ".yml":
			err = yaml.Unmarshal(bs, &v, yaml.WithCleanup(true))
		case ".json":
			err = json.Unmarshal(bs, &v)
		}

		if err != nil {
			if x := conf.log(err, "parse config file (%s) fail", path); x != nil {
				return x
			}
			continue
		}

		err = mergo.Merge(&conf.data, v, mergo.WithOverride, mergo.WithSliceDeepCopy)

		if err != nil {
			if x := conf.log(err, "merge config file (%s) fail", path); x != nil {
				return x
			}
			continue
		}

		conf.files = append(conf.files, path)
	}

	conf.loadEnv()
	conf.loadArg(os.Args[1:])

	if conf.overwrite {
		config = conf
	}

	return nil
}

var (
	defaultConfigPaths = []string{
		"config.yaml", "config.yml", "config.json",
	}

	depth = 2

	config *Config

	envKeyFunc = func(key string) string {
		return strings.ReplaceAll(strings.ToUpper(key), ".", "_")
	}
)

func cps() (p []string) {
	p = make([]string, len(defaultConfigPaths))
	copy(p, defaultConfigPaths)
	for i := 1; i <= depth; i++ {
		prefix := ""
		for j := 1; j <= i; j++ {
			prefix += "../"
		}
		for k := 0; k < len(defaultConfigPaths); k++ {
			p = append(p, prefix+defaultConfigPaths[k])
		}
	}
	exec, _ := os.Executable()
	execDir := filepath.Dir(exec) + string(os.PathSeparator)
	for i, l := 0, len(p); i < l; i++ {
		p = append(p, execDir+p[i])
	}
	return
}

func init() {
	config = New(WithSkipError(true))
	_ = config.Load(cps()...)
}

func Get(key string) val.Value {
	return config.Get(key)
}

func Files() []string {
	return config.Files()
}

func (conf *Config) Get(key string) val.Value {
	if v, ok := conf.args[key]; ok {
		return val.New(v)
	}
	if v, ok := conf.env[envKeyFunc(key)]; ok {
		return val.New(v)
	}
	v := funk.Get(conf.data, key, funk.WithAllowZero())
	return val.New(v)
}

func (conf *Config) Files() []string {
	return conf.files
}

func (conf *Config) loadEnv() {
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

func (conf *Config) loadArg(args []string) {
	for i := 0; i < len(args); i++ {
		a := args[i]
		if len(a) < 2 || a[0] != '-' || a[1] == '=' {
			continue
		}
		name, value := a[1:], ""
		if ss := strings.Split(name, "="); len(ss) == 2 {
			name = ss[0]
			value = ss[1]
		} else if i < len(args)-1 {
			i++
			value = args[i]
		}
		conf.args[name] = value
	}
}

func (conf *Config) log(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	err = fmt.Errorf(format+", "+err.Error(), args...)

	if conf.debug {
		fmt.Println(time.Now().Format("2006-01-02T15:04:05.000Z07:00") + " [config] " + err.Error())
	}

	if conf.skipError {
		return nil
	}

	return err
}
