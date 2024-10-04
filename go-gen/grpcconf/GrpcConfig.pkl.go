// Code generated from Pkl module `o47monad.sercon.GrpcConfig`. DO NOT EDIT.
package grpcconf

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type GrpcConfig struct {
	ReflectionEnabled bool `pkl:"reflectionEnabled"`

	HealthCheckEnabled bool `pkl:"healthCheckEnabled"`

	LoggingEnabled bool `pkl:"loggingEnabled"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a GrpcConfig
func LoadFromPath(ctx context.Context, path string) (ret *GrpcConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a GrpcConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*GrpcConfig, error) {
	var ret GrpcConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
