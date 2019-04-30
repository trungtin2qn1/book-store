package SetUp

import "net/http"

var ErrorMap map[string]int //global variables

//InitStuffs ...
func InitStuffs() {
	ErrorMap = map[string]int{
		"": http.StatusServiceUnavailable,
	}
}
