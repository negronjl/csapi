package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)
// API
var _ = API("CSAPI", func() {
	Description("This API exposes a REST endpoint to query a Zabbix monitoring server.")
	BasePath("/api/v1")
	Host("localhost:8080")
	Scheme("http")
	Consumes("application/json")
	Produces("application/json")
})
