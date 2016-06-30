package design

import (
    . "github.com/goadesign/goa/design"
    . "github.com/goadesign/goa/design/apidsl"
)

var _ = API("playgrounds", func() {
    Title("Playgrounds Directory")
    Description("A simple goa service")
    Scheme("http")
    Host("localhost:9876")
})

var _ = Resource("field", func() {
    BasePath("/fields")
    DefaultMedia(FieldMedia)
    Action("show", func() {
        Description("Get field by ID")
        Routing(GET("/:fieldID"))
        Params(func() {
            Param("fieldID", String, "Field ID")
        })
        Response(OK)
        Response(NotFound)
    })
})

var FieldMedia = MediaType("application/vnd.sdqali.field+json", func() {
    Description("A plying field")
    Attributes(func() {
        Attribute("id", String, "Unique field ID (UUID)")
        Attribute("href", String, "API href for the field")
        Attribute("name", String, "Name of field")
        Attribute("address", String, "Address for location of field")
        Required("id", "href", "name", "address")
    })
    View("default", func() {
        Attribute("id")
        Attribute("name")
        Attribute("href")
        Attribute("address")
    })
})
