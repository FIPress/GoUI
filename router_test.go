package goui

import "testing"

func TestParse(t *testing.T) {
	url := "book/get/:id/:name"
	route := new(route)
	parseRoute(url, route)

	_, params := dispatch("book/get/12/my book")
	t.Log(params)
}
