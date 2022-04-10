package env

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	GetKey = &envfetch{}
}

//EnvInterface : holds the signature function
type EnvInterface interface {
	LoadEnvFile(providePath func() string)

	//DATABASE ENV
	GetMysqlPassword() (string, error)
	GetMysqlDatabase() (string, error)
	GetMysqlUsername() (string, error)
	GetMysqlHost() (string, error)
	GetMysqlPort() (string, error)

	//AUTHORIZATION ENV
	GetAccessKey() (string, error)
	GetActivationKey() (string, error)
	GetRefereshKey() (string, error)
	GetPasswordKey() (string, error)
	GetEmailAPIKey() (string, error)
	GetSize() (string, error)

	//UTILITY ENV
	GetServerPort() (string, error)
	GetLocalAssets() (string, error)
	GetLocalHTML() (string, error)
	GetBaseURI() (string, error)
	GetUSERBaseURI() (string, error)

	//MIDDLEWARE ENV
	GetProtectedTokenExtractionKey() (string, error)
	GetPublicTokenExtractionKey() (string, error)
	GetPublicAuth() (string, error)
}

type envfetch struct{}

//Abstract layer get enviroment variable keys and method to load the env files
var (
	GetKey EnvInterface
)

// Reads a secret from files from the docker and other container based services
func read(fileadd string) (string, error) {
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
func (ef *envfetch) LoadEnvFile(providePath func() string) {
	//get the path
	env := providePath()

	fmt.Println(env)
	fmt.Println("docker-env:::", os.Getenv("DOCKER_ENV"))

	loadErr := godotenv.Load(env)
	if loadErr != nil {
		panic(loadErr)
	}
}

//get the env scope to define in which enviroment the build is present
func (ef *envfetch) GetEnvScope() (string, error) {
	key, err := read(envScopeFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return envScope(), nil
	} else {
		return key, nil
	}
}
