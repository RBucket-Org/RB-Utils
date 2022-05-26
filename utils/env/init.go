package env

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type initenv struct {}

//Abstract layer get enviroment variable keys and method to load the env files
var (
	InitEnv InitEnvInterface = &initenv{}
)

type InitEnvInterface interface {
	Read(fileadd string) (string, error)
	LoadEnvFile(providePath func() string)
}

// Reads a secret from files from the docker and other container based services
func (ie *initenv) Read(fileadd string) (string, error) {
	var secret string

	//check if the file is empty
	if len(fileadd) == 0 {
		return "", nil
	}

	//read the file
	buf, err := ioutil.ReadFile(fileadd)
	if err != nil {
		return secret, err
	}

	//convert and assign
	secret = strings.TrimSpace(string(buf))
	return secret, nil
}

//get the info from local env
func (ie *initenv) LoadEnvFile(providePath func() string) {
	//get the path
	env := providePath()

	fmt.Println(env)
	fmt.Println("docker-env:::", os.Getenv("DOCKER_ENV"))

	loadErr := godotenv.Load(env)
	if loadErr != nil {
		panic(loadErr)
	}
}