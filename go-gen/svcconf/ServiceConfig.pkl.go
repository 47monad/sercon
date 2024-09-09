// Code generated from Pkl module `o47monad.sercon.ServiceConfig`. DO NOT EDIT.
package svcconf

import (
	"context"

	"github.com/47monad/sercon/go-gen/mongodbconf"
	"github.com/47monad/sercon/go-gen/prometheusconf"
	"github.com/47monad/sercon/go-gen/svcconf/loglevel"
	"github.com/apple/pkl-go/pkl"
)

type ServiceConfig struct {
	Name string `pkl:"name"`

	Version string `pkl:"version"`

	Ports map[string]uint16 `pkl:"ports"`

	Host string `pkl:"host"`

	Env string `pkl:"env"`

	DebugMode bool `pkl:"debugMode"`

	LogLevel loglevel.LogLevel `pkl:"logLevel"`

	Mongodb *mongodbconf.MongodbConfig `pkl:"mongodb"`

	Prometheus *prometheusconf.PrometheusConfig `pkl:"prometheus"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a ServiceConfig
func LoadFromPath(ctx context.Context, path string) (ret *ServiceConfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a ServiceConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*ServiceConfig, error) {
	var ret ServiceConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
