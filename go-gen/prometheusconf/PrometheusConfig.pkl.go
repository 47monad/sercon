// Code generated from Pkl module `o47monad.sercon.PrometheusConfig`. DO NOT EDIT.
package prometheusconf

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type PrometheusConfig struct {
	Enabled bool `pkl:"enabled" json:"enabled,omitempty"`

	UseGrpcMetrics bool `pkl:"useGrpcMetrics" json:"useGrpcMetrics,omitempty"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a PrometheusConfig
func LoadFromPath(ctx context.Context, path string) (ret *PrometheusConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a PrometheusConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*PrometheusConfig, error) {
	var ret PrometheusConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
