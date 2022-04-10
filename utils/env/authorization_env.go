package env

func (ef *envfetch) GetAccessKey() (string, error) {
	key, err := read(accessKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return accessKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetActivationKey() (string, error) {
	key, err := read(activationKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return activationKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetRefereshKey() (string, error) {
	key, err := read(refreshKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return refreshKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetPasswordKey() (string, error) {
	key, err := read(passwordKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return passwordKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetEmailAPIKey() (string, error) {
	key, err := read(emailAPIKeyFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return emailAPIKey(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetSize() (string, error) {
	key, err := read(sizeFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return size(), nil
	} else {
		return key, nil
	}
}