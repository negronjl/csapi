package controllers

import (
	"github.com/rday/zabbix"
)

func GetDCStatus(zapi_structure ZStruct) (zabbix.JsonRPCResponse, error) {
	params := make(map[string]string, 0)
	params["output"] = "extend"
	retval, err := zapi_structure.zapi_object.ZabbixRequest("alert.get", params)
	if err != nil {
		return zabbix.JsonRPCResponse{}, err
	}
	return retval, nil
}
//
//func GetHGStatus(zapi_structure ZStruct, host_group string) (zabbix.JsonRPCResponse, error) {
//
//}
//
//func GetHostStatus(zapi_structure ZStruct, host_name string) (zabbix.JsonRPCResponse, error) {
//
//}

