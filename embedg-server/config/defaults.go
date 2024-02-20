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

	v.SetDefault("app.public_url", "https://ggffgfhayaya.onrender.com/app")

	// API defaults
	v.SetDefault("api.host", "https://ggffgfhayaya.onrender.com")
	v.SetDefault("api.port", 10000)
	v.SetDefault("api.public_url", "https://ggffgfhayaya.onrender.com/api")

	// CDN defaults
	v.SetDefault("cdn.public_url", "https://ggffgfhayaya.onrender.com/cdn")
// this token and client secret just used for tutorial in youtube and after the tutorial its was reset the token and secret id
	v.SetDefault("token", "MTIwODkxNTkxNjc4MDMzOTIzMA.GG_JKf.RkxiMVN5x6rGAPloiQ6DvVH-ckVkRSu41_78ug")
	v.SetDefault("client_id", "1208915916780339230")
	v.SetDefault("client_secret", "qGzvhAP17PyQazH_eggalvNXztpuLyUt")
}
