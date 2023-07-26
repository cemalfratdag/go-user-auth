package config

type Config struct {
	Db     DbConfig     `mapstructure:"db"`
	Server ServerConfig `mapstructure:"server"`
	Jwt    JwtConfig    `mapstructure:"jwt"`
}

type DbConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"dbname"`
	Port     string `mapstructure:"port"`
	Sslmode  string `mapstructure:"sslmode"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type JwtConfig struct {
	AccessSecret    string `mapstructure:"accessSecret"`
	RefreshSecret   string `mapstructure:"refreshSecret"`
	AccessInterval  int    `mapstructure:"accessInterval"`
	RefreshInterval int    `mapstructure:"refreshInterval"`
}
