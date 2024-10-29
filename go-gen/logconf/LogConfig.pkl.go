// Code generated from Pkl module `o47monad.sercon.LogConfig`. DO NOT EDIT.
package logconf

import (
	"context"

	"github.com/47monad/sercon/go-gen/logconf/level"
	"github.com/apple/pkl-go/pkl"
)

type LogConfig struct {
	Level level.Level `pkl:"level" json:"level,omitempty"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a LogConfig
func LoadFromPath(ctx context.Context, path string) (ret *LogConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a LogConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*LogConfig, error) {
	var ret LogConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
