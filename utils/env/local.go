package env

import "os"

func envScope() string { return os.Getenv(ENV_SCOPE) }

func mysqlPassword() string { return os.Getenv(MYSQL_PASSWORD) }
func mysqlDatabase() string { return os.Getenv(MYSQL_DATABASE) }
func mysqlUsername() string { return os.Getenv(MYSQL_USERNAME) }
func mysqlHost() string     { return os.Getenv(MYSQL_HOST) }
func mysqlPort() string     { return os.Getenv(MYSQL_PORT) }

func accessKey() string     { return os.Getenv(ACCESS_TOKEN) }
func activationKey() string { return os.Getenv(ACTIVATION_TOKEN) }
func passwordKey() string   { return os.Getenv(PASSWORD_TOKEN) }
func refreshKey() string    { return os.Getenv(REFRESH_TOKEN) }
func emailAPIKey() string      { return os.Getenv(EMAIL_API_KEY) }

func protectedTokenExtractionKey() string {
	return os.Getenv(PROTECTED_TOKEN_EXTRACTION_KEY)
}
func publicTokenExtractionKey() string { return os.Getenv(PUBLIC_TOKEN_EXTRACTION_KEY) }
func publicAuth() string               { return os.Getenv(PUBLIC_AUTH) }
func size() string                     { return os.Getenv(SIZE) }

func port() string            { return os.Getenv(PORT) }
func localHTMLPath() string   { return os.Getenv(LOCAL_HTML) }
func localAssetsPath() string { return os.Getenv(LOCAL_ASSETS) }
func baseURI() string         { return os.Getenv(URI) }
func baseUserURI() string     { return os.Getenv(USER_URI) }
