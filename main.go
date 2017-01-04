package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port              int
	Connection_String string
}

func main() {

	// mock env variable
	os.Setenv("Connection_String", "test")
	os.Setenv("Port", "8081")

	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		os.Exit(500)
	}

	fmt.Println(configuration.Port)
	fmt.Println(configuration.Connection_String)

}

func getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := []string{"config/", "config.", env, ".json"}
	dirname, _ := os.Getwd()
	filePath := path.Join(dirname, strings.Join(filename, ""))

	return filePath
}
