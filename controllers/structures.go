package controllers

import (
	"github.com/rday/zabbix"
)

type ZStruct struct {
	zapi_object *zabbix.API
	zapi_url string
	zapi_username string
	zapi_password string
	zapi_version string
}

type ZAPIStructure map[string]ZStruct