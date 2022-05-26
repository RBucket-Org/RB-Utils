package get_env

import (
	"fmt"
	"os"

	"github.com/RBucket-Org/RB-Utils/utils/env"
	"github.com/RBucket-Org/RB-Utils/utils/env/get_keys"
)

func (ef *envfetch) GetMysqlPassword() (string, error) {
	key, err := env.InitEnv.Read(get_keys.MysqlPasswordFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.MysqlPassword(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetMysqlDatabase() (string, error) {
	key, err := env.InitEnv.Read(get_keys.MysqlDatabaseFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.MysqlDatabase(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetMysqlUsername() (string, error) {
	fmt.Println(os.Getenv("MYSQL_PASSWORD"))
	key, err := env.InitEnv.Read(get_keys.MysqlUsernameFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.MysqlUsername(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetMysqlHost() (string, error) {
	key, err := env.InitEnv.Read(get_keys.MysqlHostFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.MysqlHost(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetMysqlPort() (string, error) {
	key, err := env.InitEnv.Read(get_keys.MysqlPortFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return get_keys.MysqlPort(), nil
	} else {
		return key, nil
	}
}
