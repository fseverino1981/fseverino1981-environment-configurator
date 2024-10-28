package config

import (
	"io/ioutil"
	"sync"
    "fmt"
    "os"
    "gopkg.in/yaml.v2"
)

type Config struct{
    AppName string `yaml:"app_name"`
    Database struct{
        Host string `yaml:"host"`
        Port int `yaml:"port"`
        Username string `yaml:"username"`
        Password string `yam:"password"`
    } `yaml:"database"`
}

var instance *Config
var once sync.Once

func GetConfig(env string) *Config{
    once.Do(func(){
        instance = &Config{}
        err := instance.loadConfig(env)
        if err != nil{
            fmt.Printf("Erro ao carregar a configuração: %v\n", err)
            os.Exit(1)
        }
    })
    return instance
}

func (c *Config) loadConfig(env string) error{
    filePath := fmt.Sprintf("config/%s.yaml", env)
    file, err := ioutil.ReadFile(filePath)
    if err != nil{
        return err
    }
    return yaml.Unmarshal(file, c)
}