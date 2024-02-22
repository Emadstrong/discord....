require('dotenv').config
package config

import "github.com/spf13/viper"

func setupDefaults() {
	v := viper.GetViper()

	v.SetDefault("discord.activity_name", "message.style")

	// Postgres defaults
	v.SetDefault("postgres.host", "dpg-cn99c56d3nmc73dh8ma0-a")
	v.SetDefault("postgres.port", 5432)
	v.SetDefault("postgres.dbname", "hellosjeejeww")
	v.SetDefault("postgres.user", "hellosjeejeww_user")
	v.SetDefault("postgres.password", "0jNqHD2FQmVbMKvhRkenJtYgig3TiVRo")

	// S3 defaults
	v.SetDefault("s3.endpoint", "localhost:9000")
	v.SetDefault("s3.access_key_id", "embedg")
	v.SetDefault("s3.secret_access_key", "1234567890")

	v.SetDefault("app.public_url", "https://discord-jsjsjwjaaa.onrender.com/app")

	// API defaults
	v.SetDefault("api.host", "https://discord-jsjsjwjaaa.onrender.com")
	v.SetDefault("api.port", 8080)
	v.SetDefault("api.public_url", "https://discord-jsjsjwjaaa.onrender.com/api")

	// CDN defaults
	v.SetDefault("cdn.public_url", "https://discord-jsjsjwjaaa.onrender.com/cdn")

	v.SetDefault("token", process.env.TOKEN || "BOT_TOKEN",)
	v.SetDefault("client_secret", "iJ2xN4w9pp64dJnTMoeSOqRnGovIIUby")
	v.SetDefault("client_id", "1210201359777800202")
}
