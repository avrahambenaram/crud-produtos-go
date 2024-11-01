package configuration

import "github.com/spf13/viper"

type ServerConfig struct {
	Port uint
}

var (
	MysqlDSN string
	Server   ServerConfig
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	Server.Port = viper.GetUint("server.port")
	MysqlDSN = viper.GetString("mysql_dsn")
}
