package design

import . "goa.design/goa/v3/dsl"

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("Adding number")
	Server("calc", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

var _ = Service("calc", func() {
	Description("The calc service")
	Method("add", func() {
		Payload(func() {
			Attribute("a", Int, "Left operand")
			Attribute("b", Int, "Right operand")
			Required("a", "b")
		})
		Result(Int)
		HTTP(func() {
			GET("add/{a}/{b}")
			Response(StatusOK)
		})
	})
})
