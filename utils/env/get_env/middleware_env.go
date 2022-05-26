package get_env

import (
	"github.com/RBucket-Org/RB-Utils/utils/env"
	"github.com/RBucket-Org/RB-Utils/utils/env/get_keys"
)

func (ef *envfetch) GetProtectedTokenExtractionKey() (string, error) {
	key, err := env.InitEnv.Read(get_keys.ProtectedTokenExtractionKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.ProtectedTokenExtractionKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetPublicTokenExtractionKey() (string, error) {
	key, err := env.InitEnv.Read(get_keys.PublicTokenExtractionKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.PublicTokenExtractionKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetPublicAuth() (string, error) {
	key, err := env.InitEnv.Read(get_keys.PublicAuthFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.PublicAuth(), nil
	} else {
		return key, nil
	}
}