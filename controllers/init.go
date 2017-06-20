package controllers

import (
	"strings"
	"os"
	"errors"
	"github.com/rday/zabbix"
	"fmt"
)

func Init_env() (ZAPIStructure, error) {
	zapi_env_list := os.Getenv("zapis")
	if zapi_env_list == "" {
		return nil, errors.New("Invalid or undefined zapis environment variable.")
	}
	zapi_env_entries := strings.Split(zapi_env_list, "],[")
	if len(zapi_env_entries) < 1 {
		return nil, errors.New("Unable to parse zapi list.")
	}
	//retval := make(map[string]ZStruct)
	retval := ZAPIStructure{}
	for _, zapi_env_entry := range zapi_env_entries {
		items := strings.TrimPrefix(zapi_env_entry, "[")
		items = strings.TrimSuffix(items, "]")
		item := strings.Split(items, ",")
		zapi_object, err := zabbix.NewAPI(item[1], item[2], item[3])
		fmt.Printf("Alias: [%s]\n", item[0])
		fmt.Printf("URL: [%s]\n", item[1])
		fmt.Printf("Username: [%s]\n", item[2])
		fmt.Printf("Password: [%s]\n", item[3])
		if err != nil {
			return nil, errors.New("Unable to instantiate Zabbix API internal object")
		}
		login, err := zapi_object.Login()
		if err != nil {
			fmt.Printf("Error login in: %v\n", err)
			fmt.Printf("login value: %v\n", login)
			return nil, errors.New("Unable to log in")
		}
		retval[item[0]] = ZStruct{
			zapi_object:   zapi_object,
			zapi_url:      item[1],
			zapi_username: item[2],
			zapi_password: item[3],
		}
	}
	return retval, nil
}
