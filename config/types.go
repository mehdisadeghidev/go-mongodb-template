package config

var (
	App      app
	Cors     cors
	Database database
)

type app struct {
	Env string `mapstructure:"APP_ENV"`
	Url string `mapstructure:"APP_URL"`
}

type cors struct {
	AllowedOrigins string `mapstructure:"ALLOWED_ORIGINS"`
}

type database struct {
	Uri  string `mapstructure:"DB_URI"`
	Name string `mapstructure:"DB_NAME"`
}
