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
	// get file name
	env := os.Getenv("GIN_MODE")
	fileName := "conf.yml"
	if env != "" {
		fileName = fmt.Sprintf("conf.%s.yml", env)
	}

	// load config
	file := fmt.Sprintf("./config/%s", fileName)

	filebytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(filebytes, &Cfg)
	if err != nil {
		panic(err)
	}

	log.Printf("Loaded configs from %s in %s mode .", file, env)
}
