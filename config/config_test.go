package config

import "testing"
import "os"

func TestGetFileList(t *testing.T) {
	ret := getFileList()
	if len(ret) != 2 {
		t.Errorf("expect get 2 files, but got %d", len(ret))
	}
}

func TestCheckFileExist(t *testing.T) {
	b1 := checkFileExist("config.go")
	if !b1 {
		t.Fail()
	}

	b2 := checkFileExist("config_test.go")
	if !b2 {
		t.Fail()
	}

	b3 := checkFileExist("aa.go")
	if b3 {
		t.Fail()
	}
}

func TestGetConfig(t *testing.T) {
	flippedConfig := GetConfig("../config.example.json")
	if flippedConfig.Mysql.Address != "127.0.0.1" {
		t.Errorf("mysql address should be 127.0.0.1 but got %s", flippedConfig.Mysql.Address)
	}
}

func TestReadConfig(t *testing.T) {
	flippedConfig := GetConfig("../config.example.json")

	if flippedConfig.Flipped.Host != "127.0.0.1" {
		t.Errorf("flipped host should be 127.0.0.1 but got %s", flippedConfig.Flipped.Host)
	}

	if flippedConfig.Mysql.Password != "test" {
		t.Errorf("mysql password should be test but got %s", flippedConfig.Mysql.Password)
	}
}

func TestFillConfig(t *testing.T) {
	os.Setenv("mysql_address", "192.168.10.134")
	os.Setenv("mysql_port", "3008")
	os.Setenv("flipped_host", "192.168.10.134")
	os.Setenv("flipped_port", "9999")

	flippedConfig := GetConfig("../config.example.json")

	if flippedConfig.Mysql.Address != "192.168.10.134" {
		t.Errorf("mysql address should be %s but got %s", "192.168.10.134", flippedConfig.Mysql.Address)
	}

	if flippedConfig.Mysql.Port != "3008" {
		t.Errorf("mysql port should be 3008 but got %s", flippedConfig.Mysql.Port)
	}

	if flippedConfig.Flipped.Host != "192.168.10.134" {
		t.Errorf("flipped host should be %s but got %s", "192.168.10.134", flippedConfig.Flipped.Host)
	}

	if flippedConfig.Flipped.Port != "9999" {
		t.Errorf("flipped port should be 9999 but got %s", flippedConfig.Flipped.Port)
	}
}
