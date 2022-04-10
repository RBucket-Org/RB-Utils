package env

//get the server port
func (ef *envfetch) GetServerPort() (string, error) {
	key, err := read(portFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return port(), nil
	} else {
		return key, nil
	}
}

//get the server local assets path
func (ef *envfetch) GetLocalAssets() (string, error) {
	key, err := read(localAssetsPathFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return localAssetsPath(), nil
	} else {
		return key, nil
	}
}

//get the server local html file path
func (ef *envfetch) GetLocalHTML() (string, error) {
	key, err := read(localHTMLPathFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return localHTMLPath(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetBaseURI() (string, error) {
	key, err := read(baseURIFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return baseURI(), nil
	} else {
		return key, nil
	}
}

func (ef *envfetch) GetUSERBaseURI() (string, error) {
	key, err := read(baseUserURIFile())
	if err != nil {
		return "", err
	}

	if len(key) == 0 {
		return baseUserURI(), nil
	} else {
		return key, nil
	}
}
