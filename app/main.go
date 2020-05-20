package main

import (
	"net/http"

	"github.com/aachi/goprojects/mysqlapp/app/user"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var svc user.UserService
	svc = user.NewService()
	svc = user.NewLoggingMiddleware(svc)
	svc = user.NewInstrumentingMiddleware(svc)
	http.Handle("/get-user", user.MakeGetUserHttpHandler(svc))
	http.Handle("/create-user", user.MakeCreateUserHttpHandler(svc))
	http.Handle("/matric", promhttp.Handler())
	http.ListenAndServe(":5000", nil)
}
