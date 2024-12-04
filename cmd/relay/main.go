package main

import (
	"flag"

	"github.com/47monad/sercon"
)

func main() {
	var (
		pklPath    string
		envPath    string
		outputPath string
	)

	flag.StringVar(&pklPath, "pkl", "", "pkl file name")
	flag.StringVar(&envPath, "env", "", "env file name")
	flag.StringVar(&outputPath, "output", "", "output file name")

	flag.Parse()

	err := sercon.Build(pklPath, outputPath, envPath)
	if err != nil {
		panic(err)
	}
}
