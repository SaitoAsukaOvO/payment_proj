package config

import (
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

var AppConfigPath = []string{"/etc/config/app_conf.yaml", "config/app_conf.yaml", "../config/app_conf.yaml"}
var appConfig = &AppConfig{}
var appConfigOnce sync.Once


// make it singleton
func LoadAppConfig() *AppConfig {
	appConfigOnce.Do(func() {
		for _, path := range AppConfigPath {
			if yamlFile, err := ioutil.ReadFile(path); err == nil {
				if err = yaml.Unmarshal(yamlFile, appConfig); err == nil {
					return
				}
			}
		}
		if projectRoot, err := GetProjectRootPath(); err == nil {
			if yamlFile, err := ioutil.ReadFile(projectRoot + "/app_conf.yaml"); err == nil {
				if err = yaml.Unmarshal(yamlFile, appConfig); err == nil {
					return
				}
			}
		}
		log.Printf("Warnning - Failed to load the app config yaml file and used the default empty app config instance")
	})
	return appConfig
}


// For tests only
func GetProjectRootPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	items := strings.Split(cwd, "/")
	rootPath := make([]string, 0)
	for _, item := range items {
		rootPath = append(rootPath, item)
		if item == "payment_proj" {
			rootPath = append(rootPath, "config")
			break
		}
	}
	return strings.Join(rootPath, "/"), nil
}


type AppConfig struct {
	Mysql struct {
		Host            string `yaml:"host"`
		Port            string `yaml:"port"`
		User            string `yaml:"user"`
		Password        string `yaml:"password"`
		Database        string `yaml:"database"`
		MaxConn         int    `yaml:"maxConn"`
		MaxIdle         int    `yaml:"maxIdle"`
		MaxConnLifetime int64  `yaml:"maxConnLifetime"`
	}
	Security map[string]string `yaml:"publicKey"`
}
