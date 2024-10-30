package main

import (
	"flag"

	"github.com/47monad/sercon/relay"
)

func main() {
	var (
		basePath string
		pklName  string
		envName  string
		output   string
	)

	flag.StringVar(&basePath, "base-path", "", "base path")
	flag.StringVar(&pklName, "pkl-file", "", "pkl file name")
	flag.StringVar(&envName, "env-file", "", "env file name")
	flag.StringVar(&output, "output", "", "output file name")

	flag.Parse()

	err := relay.Create(&relay.Options{
		BasePath: basePath,
	})

	if err != nil {
		panic(err)
	}
}
