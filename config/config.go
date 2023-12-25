package config

type Config struct {
	Host     string `config:"host,required"`
	Port     int32  `config:"port,required"`
	SSL      bool   `config:"ssl"`
	Login    string `config:"login,required"`
	Password string `config:"password,required"`
	Api      string `config:"api,required"`
}
