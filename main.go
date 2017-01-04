package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port              int
	Connection_String string
}

func main() {

	os.Setenv("Connection_String", "test")
	os.Setenv("Port", "8081")

	configuration := Configuration{}
	err := gonfig.GetJsonConfiguration(getFileName(), &configuration)
	if err != nil {
		os.Exit(500)
	}

	gonfig.GetEnvVariablesConfiguration(&configuration)

	fmt.Println(configuration.Port)
	fmt.Println(configuration.Connection_String)
}

func getFileName() string {

	filePath := path.Join(filepath.Dir(os.Args[0]), "config.json")

	return filePath
}
