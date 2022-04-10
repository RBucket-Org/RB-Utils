package env

import "os"

func envScopeFile() string { return os.Getenv(ENV_SCOPE_FILE) }

func mysqlPasswordFile() string { return os.Getenv(MYSQL_PASSWORD_FILE) }
func mysqlDatabaseFile() string { return os.Getenv(MYSQL_DATABASE_FILE) }
func mysqlUsernameFile() string { return os.Getenv(MYSQL_USERNAME_FILE) }
func mysqlHostFile() string     { return os.Getenv(MYSQL_HOST_FILE) }
func mysqlPortFile() string     { return os.Getenv(MYSQL_PORT_FILE) }

func accessKeyFile() string     { return os.Getenv(ACCESS_TOKEN_FILE) }
func activationKeyFile() string { return os.Getenv(ACTIVATION_TOKEN_FILE) }
func passwordKeyFile() string   { return os.Getenv(PASSWORD_TOKEN_FILE) }
func refreshKeyFile() string    { return os.Getenv(REFRESH_TOKEN_FILE) }
func emailAPIKeyFile() string      { return os.Getenv(EMAIL_API_KEY_FILE) }

func protectedTokenExtractionKeyFile() string {
	return os.Getenv(PROTECTED_TOKEN_EXTRACTION_KEY_FILE)
}
func publicTokenExtractionKeyFile() string { return os.Getenv(PUBLIC_TOKEN_EXTRACTION_KEY_FILE) }
func publicAuthFile() string               { return os.Getenv(PUBLIC_AUTH_FILE) }
func sizeFile() string                     { return os.Getenv(SIZE_FILE) }

func portFile() string            { return os.Getenv(PORT_FILE) }
func localHTMLPathFile() string   { return os.Getenv(LOCAL_HTML_FILE) }
func localAssetsPathFile() string { return os.Getenv(LOCAL_ASSETS_FILE) }
func baseURIFile() string         { return os.Getenv(URI_FILE) }
func baseUserURIFile() string     { return os.Getenv(USER_URI_FILE) }