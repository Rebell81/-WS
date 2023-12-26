package config

type Config struct {
	Host              string `config:"qb_host,required" json:"qb_host"`
	Port              int32  `config:"qb_port,required" json:"qb_port"`
	SSL               bool   `config:"ssl" json:"ssl"`
	Login             string `config:"qb_login,required" json:"qb_login"`
	Password          string `config:"qb_password,required" json:"qb_password"`
	RutrackerApiToken string `config:"rutracker_api_token,required" json:"rutracker_api_token"`
	TelegramToken     string `config:"telegram_token,required" json:"telegram_token"`
	ChatId            int64  `config:"telegram_chat_id,required" json:"telegram_chat_id"`
	DurationSeconds   int    `config:"duration_seconds,required" json:"duration_seconds"`
	ManualCheckOnly   bool   `config:"manual_check" json:"manual_check"`
}
