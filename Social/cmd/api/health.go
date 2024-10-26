package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func (app *application) saySomething(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello say function"))
}