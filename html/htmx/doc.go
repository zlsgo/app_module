// Package htmx provides typed attribute names and small helpers for using
// htmx 2.x with the app_module/html element DSL.
//
// The package mirrors the htmx attribute reference so callers can avoid
// stringly-typed attribute names while still composing regular el.Node values:
//
//	button := el.BUTTON(
//		htmx.Get("/todos"),
//		htmx.Target("#list"),
//		htmx.Swap("innerHTML"),
//		el.Text("Refresh"),
//	)
//
// JSON-like attributes such as hx-vals, hx-headers and hx-request accept the
// same values as el.Attr. Passing ztype.Map lets the renderer serialize and
// HTML-escape the JSON safely at render time.
package htmx
