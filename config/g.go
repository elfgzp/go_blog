package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	projectName := "go_blog"
	getConfig(projectName)
}

func getConfig(projectName string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", projectName))
	viper.AddConfigPath(fmt.Sprintf("/data/docker/config/%s", projectName))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
}

//
func GetMysqlConnectingString() string {
	usr := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true", usr, pwd, host, port, db, charset)
}
