// Code generated from Pkl module `o47monad.sercon.MongodbConfig`. DO NOT EDIT.
package mongodbconf

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type MongodbConfig struct {
	Enabled bool `pkl:"enabled"`

	Uri *string `pkl:"uri"`

	DbName string `pkl:"dbName"`

	UriBuilder *URIBuilder `pkl:"uriBuilder"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a MongodbConfig
func LoadFromPath(ctx context.Context, path string) (ret *MongodbConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a MongodbConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*MongodbConfig, error) {
	var ret MongodbConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
