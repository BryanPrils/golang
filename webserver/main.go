package main

import (
	"net/http"

	"example.com/webserver/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
