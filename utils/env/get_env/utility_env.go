package get_env

import (
	"github.com/RBucket-Org/RB-Utils/utils/env"
	"github.com/RBucket-Org/RB-Utils/utils/env/get_keys"
)

//get the server port
func (ef *envfetch) GetServerPort() (string, error) {
	key, err := env.InitEnv.Read(get_keys.PortFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.Port(), nil
	} else {
		return key, nil
	}
}

//get the server local assets path
func (ef *envfetch) GetLocalAssets() (string, error) {
	key, err := env.InitEnv.Read(get_keys.LocalAssetsPathFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.LocalAssetsPath(), nil
	} else {
		return key, nil
	}
}

//get the server local html file path
func (ef *envfetch) GetLocalHTML() (string, error) {
	key, err := env.InitEnv.Read(get_keys.LocalHTMLPathFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.LocalHTMLPath(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetBaseURI() (string, error) {
	key, err := env.InitEnv.Read(get_keys.BaseURIFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.BaseURI(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetUSERBaseURI() (string, error) {
	key, err := env.InitEnv.Read(get_keys.BaseUserURIFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.BaseUserURI(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetNotificationAPIKey() (string, error) {
	key, err := env.InitEnv.Read(get_keys.NotificationAPIKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.NotificationAPIKey(), nil
	} else {
		return key, nil
	}
}
