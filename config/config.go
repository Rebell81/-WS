package config

type Config struct {
	Host     string `mapstructure:"host,required"`
	Port     int32  `mapstructure:"port,required"`
	SSL      bool   `mapstructure:"ssl"`
	Login    string `mapstructure:"login,required"`
	Password string `mapstructure:"password,required"`
	Api      string `mapstructure:"rutrackerApiToken,required"`
	Token    string `mapstructure:"telegramToken,required"`
	ChatId   int64  `mapstructure:"telegramChatId"`
}
