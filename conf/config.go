package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Pool   pool `toml:"pool"`
	RedisMaster masterLink `toml:"redis-master"`
	RedisSlave slaveLink `toml:"redisslave"`
}

type pool struct {
	InitCap int
	MaxCap int
	IdleTimeout int
}
type masterLink struct{
	Server string `toml:"server"`
	Port  int `toml:"port"`
	Auth  string `toml:"auth"`
}
type slaveLink struct{
	Server []string `toml:"server"`
	Port  int `toml:"port"`
	Auth  string `toml:"auth"`
}


var Config tomlConfig

func init(){
	fmt.Println("start conf")


	toml.DecodeFile("conf/config.toml", &Config)

	fmt.Println(Config)
}
