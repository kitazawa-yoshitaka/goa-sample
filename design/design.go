package design

import (
	. "goa.design/goa/v3/dsl"
)

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
			Response("DivByZero", StatusBadRequest)
			Response("Timeout", StatusRequestTimeout)
			Response("NotFound", StatusNotFound)
		})
	})

	Error("DivByZero", func() {
		Description("DivByZero is the error returned by the service methods when the right operand is 0.")
	})

	Error("NotFound", ErrorResult, "NotFound is the error returned when there is no bottle with the given ID.")

	Error("Timeout", ErrorResult, "operation timed out, retry later.", func() {
		// Timeout indicates an error due to a timeout.
		Timeout()
		// Temporary indicates that the request may be retried.
		Temporary()

	})
})
