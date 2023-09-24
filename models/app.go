package models

type AppConfig struct {
	Name           string `env:"NAME_SERVER"`
	Port           string `env:"PORT"`
	Host           string `env:"HOST_SERVER"`
	DatabaseConfig DatabaseConfig
	RabbitMQConfig RabbitMQ
	EmailConfig    EmailConfig
	AppEnv         string `env:"APP_ENV"`
	SecretKey      string `env:"SECRET_KEY"`
	DSN_MQ         string `env:"DSN_MQ"`
}

type DatabaseConfig struct {
	Name     string `env:"NAME_DB"`
	Host     string `env:"HOST_DB"`
	Port     string `env:"PORT_DB"`
	User     string `env:"USER_DB"`
	Password string `env:"PASSWORD_DB"`
}

type RabbitMQ struct {
	Host     string `env:"HOST_RABBIT"`
	Port     string `env:"PORT_RABBIT"`
	User     string `env:"USER_RABBIT"`
	Password string `env:"PASSWORD_RABBIT"`
}

type EmailConfig struct {
	ConfigSmtpHost     string `env:"CONFIG_SMTP_HOST"`
	ConfigSmtpPort     int    `env:"CONFIG_SMTP_PORT"`
	ConfigSenderEmail  string `env:"CONFIG_SENDER_EMAIL"`
	ConfigAuthEmail    string `env:"CONFIG_AUTH_EMAIL"`
	ConfigAuthPassword string `env:"CONFIG_AUTH_PASSWORD"`
}
