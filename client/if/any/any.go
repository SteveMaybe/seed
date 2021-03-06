package any

import (
	"qlova.org/seed/client"
	"qlova.org/seed/client/if/the"
	"qlova.org/seed/use/js"
)

//AreTrue returns a client.Bool that is true if any of its arguments are true.
func AreTrue(a, b client.Bool, others ...client.Bool) client.Bool {
	var expression string = "(%v || %v"
	var values = []client.Value{a, b}

	for _, n := range others {
		expression += " || %v"
		values = append(values, n)
	}

	expression += ")"

	return the.Bool(js.Bool{Value: js.NewValue(expression, values...)}, values...)
}

//AreFalse returns a client.Bool that is true if any of its arguments are false.
func AreFalse(a, b client.Bool, others ...client.Bool) client.Bool {
	var expression string = "(!(%v && %v"
	var values = []client.Value{a, b}

	for _, n := range others {
		expression += " && %v"
		values = append(values, n)
	}

	expression += "))"

	return the.Bool(js.Bool{Value: js.NewValue(expression, values...)}, values...)
}
