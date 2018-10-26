package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/tkanos/gonfig"
)

type Configuration struct {
	Port              int `env:"APP_PORT"`
	Connection_String string
}

func main() {

	// mock env variable
	os.Setenv("Connection_String", "test")
	os.Setenv("APP_PORT", "8081")

	configuration := Configuration{}
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		fmt.Println(err)
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
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
