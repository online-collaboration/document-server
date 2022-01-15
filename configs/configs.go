package configs

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type Config struct {
	Server server `toml:"server"`
	Mysql  mysql  `toml:"mysql"`
}

type server struct {
	Port    string `toml:"port"`
	IDC     string `toml:"idc"`
	Name    string `toml:"name"`
	MaxCpus int    `toml:"max_cups"`
}

type mysql struct {
	Address  string `toml:"address"`
	Password string `toml:"password"`
}

var (
	Conf *Config
	once sync.Once
)

func getEnv() string {
	env := os.Getenv("APP_ENV")
	if len(env) == 0 {
		env = "dev"
	}
	return env
}

func Initialize() *Config {
	once.Do(func() {
		env := getEnv()
		log.Println("env=", env)
		if env == "" {
			env = "dev"
		}
		filePath, err := filepath.Abs("configs/" + env + ".conf.toml")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("parse toml file once. filePath: %s\n", filePath)
		if _, err := toml.DecodeFile(filePath, &Conf); err != nil {
			fmt.Println(err)
		}

	})
	return Conf
}
