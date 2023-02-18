package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

type IP struct {
	Ip      string `json:"ip"`
	Port    int    `json:"port"`
	Secured bool   `json:"secured"`
}

var configdir string

func config_read() []IP {
	os.Chdir(configdir)
	content, err := ioutil.ReadFile("ip.json")
	if err != nil {
		mError("Could not read config file", filepath.Join(configdir, "ip.json")+":", err.Error())
	}
	config := []IP{}
	err = json.Unmarshal(content, &config)
	if err != nil {
		mError("Could not get data from config file", filepath.Join(configdir, "ip.json")+":", err.Error())
	}
	return config
}

func config_save(config []IP) {
	os.Chdir(configdir)
	content, err := json.Marshal(config)
	if err != nil {
		mError("Could not generate config file:", err.Error())
	}
	err = ioutil.WriteFile("ip.json", content, 0644)
	if err != nil {
		mError("Could not save to config file", filepath.Join(configdir, "ip.json")+":", err.Error())
	}
}

func password_read() string {
	os.Chdir(configdir)
	_password, err := ioutil.ReadFile("password")
	if err != nil {
		mError("Could not read config file", filepath.Join(configdir, "password")+":", err.Error())
	}
	return string(_password)
}

func password_save(password string) {
	err := ioutil.WriteFile("password", []byte(password), 0644)
	if err != nil {
		mError("Could not save to config file", filepath.Join(configdir, "password")+":", err.Error())
	}
}

func prepare_configdir() {
	_configdir, err := os.UserConfigDir()
	if err != nil {
		mError("Could not find user directory:", err.Error())
	}
	configdir = filepath.Join(_configdir, "wpkg2-cli")
	os.MkdirAll(configdir, 0664)
	os.Chdir(configdir)
	if _, err := os.Stat(filepath.Join(configdir, "ip.json")); os.IsNotExist(err) {
		err = os.WriteFile("ip.json", []byte(`[{"ip":"droplet-s1.medzik.xyz","port":3219,"secured":true}]`), 0664)
		if err != nil {
			mError("Could not create config file", filepath.Join(configdir, "ip.json")+":", err.Error())
		}
	}
	if _, err := os.Stat(filepath.Join(configdir, "password")); os.IsNotExist(err) {
		err = os.WriteFile("password", []byte(""), 0664)
		if err != nil {
			mError("Could not create config file", filepath.Join(configdir, "password")+":", err.Error())
		}
	}
}
