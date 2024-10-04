package sercon

import (
	"context"
	"fmt"
	"github.com/47monad/sercon/go-gen/svcconf"
	"github.com/joho/godotenv"
)

type Config svcconf.ServiceConfig

type LoadOptions struct {
	ctx        context.Context
	configPath string
	envPath    string
}

type LoadOption func(*LoadOptions)

func Load(opts ...LoadOption) (*Config, error) {
	_opts := prepareOptions(opts)
	if _opts.configPath == "" {
		return nil, fmt.Errorf("missing config path")
	}
	if _opts.envPath != "" {
		err := godotenv.Load(_opts.envPath)
		if err != nil {
			return nil, err
		}
	}
	conf, err := svcconf.LoadFromPath(_opts.ctx, _opts.configPath)
	if err != nil {
		return nil, err
	}

	return (*Config)(conf), nil
}

func prepareOptions(options []LoadOption) *LoadOptions {
	opts := &LoadOptions{
		ctx: context.Background(),
	}

	for _, option := range options {
		option(opts)
	}

	return opts
}

func WithContext(ctx context.Context) LoadOption {
	return func(o *LoadOptions) {
		o.ctx = ctx
	}
}

func WithConfigPath(path string) LoadOption {
	return func(o *LoadOptions) {
		o.configPath = path
	}
}

func WithEnvPath(path string) LoadOption {
	return func(o *LoadOptions) {
		o.envPath = path
	}
}
