package env

import (
	"fmt"
	"os"
)

func (ef *envfetch) GetMysqlPassword() (string, error) {
	key, err := read(mysqlPasswordFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return mysqlPassword(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetMysqlDatabase() (string, error) {
	key, err := read(mysqlDatabaseFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return mysqlDatabase(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetMysqlUsername() (string, error) {
	fmt.Println(os.Getenv("MYSQL_PASSWORD"))
	key, err := read(mysqlUsernameFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return mysqlUsername(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetMysqlHost() (string, error) {
	key, err := read(mysqlHostFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return mysqlHost(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetMysqlPort() (string, error) {
	key, err := read(mysqlPortFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return mysqlPort(), nil
	} else {
		return key, nil
	}
}
