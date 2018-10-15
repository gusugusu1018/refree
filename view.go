package main

import (
	"net/http"

	log "github.com/cihub/seelog"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	/*
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
		} else {
			generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
		}
	*/
	log.Debug("Access : err")
	vals := request.URL.Query()
	t := parseTemplateFiles("layout", "error")
	err := t.ExecuteTemplate(writer, "layout", vals.Get("msg"))
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Debug(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : index")
	t := parseTemplateFiles("layout", "index")
	err := t.ExecuteTemplate(writer, "layout", nil)
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Debug(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}

func about(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : about")
	t := parseTemplateFiles("layout", "about")
	err := t.ExecuteTemplate(writer, "layout", nil)
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Debug(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}

func contact(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : contact")
	t := parseTemplateFiles("layout", "contact")
	err := t.ExecuteTemplate(writer, "layout", nil)
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Debug(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}

func serverInfo(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : info")
	m := map[string]string{
		"InfoMsg": "Hello World!",
	}
	t := parseTemplateFiles("layout", "info")
	err := t.ExecuteTemplate(writer, "layout", m)
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Debug(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}

func login(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : info")
	m := map[string]string{
		"InfoMsg": "Hello World!",
	}
	t := parseTemplateFiles("layout", "info")
	err := t.ExecuteTemplate(writer, "layout", m)
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Debug(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}

func signup(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : info")
	m := map[string]string{
		"InfoMsg": "Hello World!",
	}
	t := parseTemplateFiles("layout", "info")
	err := t.ExecuteTemplate(writer, "layout", m)
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Debug(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}
