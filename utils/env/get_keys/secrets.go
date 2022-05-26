package get_keys

import (
	"os"

	"github.com/RBucket-Org/RB-Utils/utils/env/constants"
)

func EnvScopeFile() string { return os.Getenv(constants.ENV_SCOPE_FILE) }

func MysqlPasswordFile() string { return os.Getenv(constants.MYSQL_PASSWORD_FILE) }
func MysqlDatabaseFile() string { return os.Getenv(constants.MYSQL_DATABASE_FILE) }
func MysqlUsernameFile() string { return os.Getenv(constants.MYSQL_USERNAME_FILE) }
func MysqlHostFile() string     { return os.Getenv(constants.MYSQL_HOST_FILE) }
func MysqlPortFile() string     { return os.Getenv(constants.MYSQL_PORT_FILE) }

func AccessKeyFile() string     { return os.Getenv(constants.ACCESS_TOKEN_FILE) }
func ActivationKeyFile() string { return os.Getenv(constants.ACTIVATION_TOKEN_FILE) }
func PasswordKeyFile() string   { return os.Getenv(constants.PASSWORD_TOKEN_FILE) }
func RefreshKeyFile() string    { return os.Getenv(constants.REFRESH_TOKEN_FILE) }
func EmailAPIKeyFile() string   { return os.Getenv(constants.EMAIL_API_KEY_FILE) }

func ProtectedTokenExtractionKeyFile() string {
	return os.Getenv(constants.PROTECTED_TOKEN_EXTRACTION_KEY_FILE)
}
func PublicTokenExtractionKeyFile() string { return os.Getenv(constants.PUBLIC_TOKEN_EXTRACTION_KEY_FILE) }
func PublicAuthFile() string               { return os.Getenv(constants.PUBLIC_AUTH_FILE) }
func SizeFile() string                     { return os.Getenv(constants.SIZE_FILE) }

func PortFile() string               { return os.Getenv(constants.PORT_FILE) }
func LocalHTMLPathFile() string      { return os.Getenv(constants.LOCAL_HTML_FILE) }
func LocalAssetsPathFile() string    { return os.Getenv(constants.LOCAL_ASSETS_FILE) }
func BaseURIFile() string            { return os.Getenv(constants.URI_FILE) }
func BaseUserURIFile() string        { return os.Getenv(constants.USER_URI_FILE) }
func NotificationAPIKeyFile() string { return os.Getenv(constants.NOTIFICATION_API_KEY_FILE) }
