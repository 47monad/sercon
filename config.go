package sercon

import (
	"context"
	"github.com/47monad/sercon/go-gen/svcconf"
	"github.com/joho/godotenv"
)

func Load(ctx context.Context, configFileName string, envFileName string) (*svcconf.ServiceConfig, error) {
	if envFileName != "" {
		err := godotenv.Load(envFileName)
		if err != nil {
			return nil, err
		}
	}
	conf, err := svcconf.LoadFromPath(ctx, configFileName)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
