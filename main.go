package main

import (
	"net/http"
	"screening/apis"
)

func main() {
	apis.SetupJsonApi()
	http.ListenAndServe(":80", nil)
}
