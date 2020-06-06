package main

import (
	"gopkg.in/yaml.v2"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)


type Config struct {
	Credentials struct {
		Nim string `yaml:"nim"`
		Password string `yaml:"password"`
	}
}
var configFileName string = "config.yml"

func readConfig()(*Config,error){
	config := &Config{}
	Dir,err:=os.Executable()
	log.Print(Dir)
	if err != nil {
		log.Fatal(err)
	}
	configPath:=getConfigDirectory()+"/"+configFileName
	file,err:=os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}
	return config, nil
}

func getConfigDirectory()string{
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}