package relay

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/47monad/sercon/go-gen/svcconf"
	"github.com/apple/pkl-go/pkl"
	"github.com/joho/godotenv"
)

type Options struct {
	BasePath       string
	PKLPath        string
	EnvPath        string
	Output         string
	fullPKLPath    string
	fullEnvPath    string
	fullOutputPath string
}

type Option func(*Options)

func WithBasePath(path string) Option {
	return func(opts *Options) {
		opts.BasePath = path
	}
}

func WithPKLPath(path string) Option {
	return func(opts *Options) {
		opts.BasePath = path
	}
}

func WithEnvPath(path string) Option {
	return func(opts *Options) {
		opts.BasePath = path
	}
}

func WithOutputPath(path string) Option {
	return func(opts *Options) {
		opts.BasePath = path
	}
}

func Create(opts ...Option) error {
	_opts := &Options{
		BasePath: "./config",
		PKLPath:  "app.pkl",
		EnvPath:  ".env",
		Output:   "build/app.json",
	}

	for _, opt := range opts {
		opt(_opts)
	}

	_opts.fullEnvPath = fmt.Sprintf("%s/%s", _opts.BasePath, _opts.EnvPath)
	_opts.fullPKLPath = fmt.Sprintf("%s/%s", _opts.BasePath, _opts.PKLPath)
	_opts.fullOutputPath = fmt.Sprintf("%s/%s", _opts.BasePath, _opts.Output)

	if err := LoadEnv(_opts.fullEnvPath); err != nil {
		return err
	}

	conf, err := LoadConfig(_opts.fullPKLPath)
	if err != nil {
		return err
	}

	if err := WriteToJson(conf, _opts.fullOutputPath); err != nil {
		return err
	}

	return nil
}

func LoadEnv(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

func WriteToJson(c *svcconf.ServiceConfig, outputPath string) error {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig(pklPath string) (*svcconf.ServiceConfig, error) {
	ctx := context.Background()
	options := func(opts *pkl.EvaluatorOptions) {
		pkl.WithDefaultAllowedResources(opts)
		pkl.WithOsEnv(opts)
		pkl.WithDefaultAllowedModules(opts)
		opts.Logger = pkl.NoopLogger
		opts.OutputFormat = "json"
	}
	evaluator, err := pkl.NewEvaluator(ctx, options)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()

	conf, err := load(ctx, evaluator, pkl.FileSource(pklPath))
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*svcconf.ServiceConfig, error) {
	var ret svcconf.ServiceConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}