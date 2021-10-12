package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type cfg struct {
	App struct {
		Listen string `yaml:"listen"`
	}

	DB struct {
		Url string `yaml:"url"`
	}

	JWT struct {
		Secret string `yaml:"secret"`
	}
}

var Cfg cfg

func init() {

	env := getEnv("GIN_MODE", "dev")
	filePath := fmt.Sprintf("./config/conf.%s.yml", env)

	filebytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(filebytes, &Cfg)
	if err != nil {
		panic(err)
	}
	log.Printf("Loaded configs from %s in %s mode .", filePath, env)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
