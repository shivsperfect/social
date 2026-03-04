package main

import "net/http"

func (app *app) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}
