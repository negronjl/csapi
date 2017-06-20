package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("health", func() {
	BasePath("/health")
	Action("health", func() {
		Routing(GET("/:dc/:hgroup/:hostname"))
		Description("Get the health from servers monitored by the requested Zabbix API.  " +
			"If provided, you can narrow the list of servers down by Host Group or hostname")
		Params(func() {
			Param("dc", String, "Data Center", func() {
				Example("DC1")
			})
			Param("hgroup", String, "Host Group name", func() {
				Example("HG1")
			})
			Param("hostname", String, "Host name", func() {
				Example("monitored-box.mydomain.tld")
			})
		})
		Response(OK, "application/json")
		Response(NotFound)
	})

})

