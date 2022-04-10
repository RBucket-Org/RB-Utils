package env


func (ef *envfetch) GetProtectedTokenExtractionKey() (string, error) {
	key, err := read(protectedTokenExtractionKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return protectedTokenExtractionKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetPublicTokenExtractionKey() (string, error) {
	key, err := read(publicTokenExtractionKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return publicTokenExtractionKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetPublicAuth() (string, error) {
	key, err := read(publicAuthFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return publicAuth(), nil
	} else {
		return key, nil
	}
}