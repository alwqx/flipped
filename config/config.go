package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/adolphlwq/flipped/entity"
)

const (
	CONFIG_PATH  = "config.json"
	EXAMPLE_PATH = "config.example.json"
)

func GetConfig(configPath string) *entity.FlippedConfig {
	if len(configPath) != 0 {
		logrus.Infof("Get config info from %s", configPath)
		return readConfig(configPath)
	}

	if checkFileExist(CONFIG_PATH) {
		logrus.Infof("Get config info from %s", CONFIG_PATH)
		return readConfig(CONFIG_PATH)
	} else if checkFileExist(EXAMPLE_PATH) {
		logrus.Infof("Get config info from %s", EXAMPLE_PATH)
		return readConfig(EXAMPLE_PATH)
	} else {
		logrus.Fatal("please provide config.json or config.example.json")
	}

	return nil
}

func checkFileExist(fileName string) bool {
	fileList := getFileList()
	for _, file := range fileList {
		if file == fileName {
			return true
		}
	}

	return false
}

func getFileList() []string {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		logrus.Fatalf("read file list error: %s", err.Error())
	}

	var fileList []string
	for _, file := range files {
		// judge is file
		if !file.IsDir() {
			fileList = append(fileList, file.Name())
		}
	}

	return fileList
}

func readConfig(fileName string) *entity.FlippedConfig {
	config, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Fatal(err)
	}

	var configJSON entity.FlippedConfig
	err = json.Unmarshal(config, &configJSON)
	if err != nil {
		logrus.Fatalf("parse json file error: %v", err)
	}

	return fillConfig(configJSON)
}

// fillConfig prefer use variables from env
func fillConfig(flippedConfig entity.FlippedConfig) *entity.FlippedConfig {
	// config Mongo
	if address := os.Getenv("mysql_address"); address != "" {
		flippedConfig.Mysql.Address = address
	}
	if port := os.Getenv("mysql_port"); port != "" {
		flippedConfig.Mysql.Port = port
	}
	if database := os.Getenv("mysql_dbname"); database != "" {
		flippedConfig.Mysql.DBName = database
	}
	if user := os.Getenv("mysql_user"); user != "" {
		flippedConfig.Mysql.User = user
	}
	if password := os.Getenv("mysql_password"); password != "" {
		flippedConfig.Mysql.Password = password
	}

	// config DCOS
	if host := os.Getenv("flipped_host"); host != "" {
		flippedConfig.Flipped.Host = host
	}
	if port := os.Getenv("flipped_port"); port != "" {
		flippedConfig.Flipped.Port = port
	}

	return &flippedConfig
}
