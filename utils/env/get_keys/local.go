package get_keys

import (
	"os"

	"github.com/RBucket-Org/RB-Utils/utils/env/constants"
)

func EnvScope() string { return os.Getenv(constants.ENV_SCOPE) }

func MysqlPassword() string { return os.Getenv(constants.MYSQL_PASSWORD) }
func MysqlDatabase() string { return os.Getenv(constants.MYSQL_DATABASE) }
func MysqlUsername() string { return os.Getenv(constants.MYSQL_USERNAME) }
func MysqlHost() string     { return os.Getenv(constants.MYSQL_HOST) }
func MysqlPort() string     { return os.Getenv(constants.MYSQL_PORT) }

func AccessKey() string     { return os.Getenv(constants.ACCESS_TOKEN) }
func ActivationKey() string { return os.Getenv(constants.ACTIVATION_TOKEN) }
func PasswordKey() string   { return os.Getenv(constants.PASSWORD_TOKEN) }
func RefreshKey() string    { return os.Getenv(constants.REFRESH_TOKEN) }
func EmailAPIKey() string      { return os.Getenv(constants.EMAIL_API_KEY) }

func ProtectedTokenExtractionKey() string {
	return os.Getenv(constants.PROTECTED_TOKEN_EXTRACTION_KEY)
}
func PublicTokenExtractionKey() string { return os.Getenv(constants.PUBLIC_TOKEN_EXTRACTION_KEY) }
func PublicAuth() string               { return os.Getenv(constants.PUBLIC_AUTH) }
func Size() string                     { return os.Getenv(constants.SIZE) }

func Port() string            { return os.Getenv(constants.PORT) }
func LocalHTMLPath() string   { return os.Getenv(constants.LOCAL_HTML) }
func LocalAssetsPath() string { return os.Getenv(constants.LOCAL_ASSETS) }
func BaseURI() string         { return os.Getenv(constants.URI) }
func BaseUserURI() string     { return os.Getenv(constants.USER_URI) }
func NotificationAPIKey() string     { return os.Getenv(constants.NOTIFICATION_API_KEY) }
