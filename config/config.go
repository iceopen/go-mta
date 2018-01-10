package config

var (
	APP_ID     string
	SECRET_KEY string
)

func SetConfig(app_id string, secret_key string) error {
	APP_ID = app_id
	SECRET_KEY = secret_key
	return nil
}
