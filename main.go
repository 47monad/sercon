package main

import (
	"flag"
	"fmt"

	sercon "github.com/47monad/sercon/lib"
)

func main() {
	var (
		configBasePath string
		pklPath        string
		envPath        string
		outputPath     string
	)

	flag.StringVar(&configBasePath, "base", "", "config base path")
	flag.StringVar(&pklPath, "pkl", "", "pkl file name")
	flag.StringVar(&envPath, "env", "", "env file name")
	flag.StringVar(&outputPath, "output", "", "output file name")

	flag.Parse()

	if outputPath == "" {
		outputPath = ".sercon/app.json"
	}

	if configBasePath == "" {
		configBasePath = "./config"
	}

	if pklPath == "" {
		pklPath = "app.pkl"
	}

	if envPath == "" {
		envPath = ".env"
	}

	fullPKLPath := fmt.Sprintf("%s/%s", configBasePath, pklPath)
	fullEnvPath := fmt.Sprintf("%s/%s", configBasePath, envPath)

	err := sercon.Build(fullPKLPath, outputPath, fullEnvPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Config was built inside %s", outputPath)
}
