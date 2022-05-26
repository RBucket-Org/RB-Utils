package get_env

import (
	"github.com/RBucket-Org/RB-Utils/utils/env"
	"github.com/RBucket-Org/RB-Utils/utils/env/get_keys"
)

var GetKey EnvInterface = &envfetch{}


type envfetch struct{}

//EnvInterface : holds the signature function
type EnvInterface interface {
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
	GetNotificationAPIKey() (string, error)

	//MIDDLEWARE ENV
	GetProtectedTokenExtractionKey() (string, error)
	GetPublicTokenExtractionKey() (string, error)
	GetPublicAuth() (string, error)
}

//get the env scope to define in which enviroment the build is present
func (ef *envfetch) GetEnvScope() (string, error) {
	key, err := env.InitEnv.Read(get_keys.EnvScopeFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.EnvScope(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetAccessKey() (string, error) {
	key, err := env.InitEnv.Read(get_keys.AccessKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.AccessKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetActivationKey() (string, error) {
	key, err := env.InitEnv.Read(get_keys.ActivationKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.ActivationKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetRefereshKey() (string, error) {
	key, err := env.InitEnv.Read(get_keys.RefreshKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.RefreshKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetPasswordKey() (string, error) {
	key, err := env.InitEnv.Read(get_keys.PasswordKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.PasswordKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetEmailAPIKey() (string, error) {
	key, err := env.InitEnv.Read(get_keys.EmailAPIKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.EmailAPIKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetSize() (string, error) {
	key, err := env.InitEnv.Read(get_keys.SizeFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.Size(), nil
	} else {
		return key, nil
	}
}