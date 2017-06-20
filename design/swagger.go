package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)


var _ = Resource("swagger", func() {
	BasePath("/swagger")
	Origin("*", func() { // CORS policy that applies to all actions and file servers
		Methods("GET") // of "public" resource
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger.yaml", "swagger/swagger.yaml")
})
