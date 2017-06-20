package controllers

import (
	"fmt"
	"time"
	"github.com/rday/zabbix"
)

func GetDCStatus(zapi_structure ZStruct) (zabbix.JsonRPCResponse, error) {
	params := make(map[string]string, 0)
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)

	params["output"] = "extend"
	params["select_acknowledges"] = "extend"
	params["time_from"] = fmt.Sprintf("%d", yesterday.Unix())
	params["time_till"] = fmt.Sprintf("%d", now.Unix())
	params["selectHosts"] = "extend"
	params["selectRelatedObject"] = "extend"
	params["select_alerts"] = "extend"

	retval, err := zapi_structure.zapi_object.ZabbixRequest("event.get", params)
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
