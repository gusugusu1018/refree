package main

import (
	"net/http"
	"time"

	"./data"
	log "github.com/cihub/seelog"
)

func main() {
	initialize()
	log.Debug("Start main")
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Server.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	// GET
	mux.HandleFunc("/", index)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/info", serverInfo)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	// POST
	mux.HandleFunc("/authenticate", authenticate)
	mux.HandleFunc("/signup-account", signupAccount)
	server := &http.Server{
		Addr:           config.Server.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.Server.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.Server.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	log.Info("Start Server" + config.Server.Address)
	server.ListenAndServe()
}

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : err")
	vals := request.URL.Query()
	t := parseTemplateFiles("layout", "error")
	err := t.ExecuteTemplate(writer, "layout", vals.Get("msg"))
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Error(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : index")
	var u data.User
	sess, err := session(writer, request)
	var IsAliveSession string
	if err == nil {
		IsAliveSession = "true"
		u, _ = sess.User()
	}
	m := map[string]string{
		"IsAliveSession": IsAliveSession,
		"UserName":       u.Name,
	}
	t := parseTemplateFiles("layout", "index")
	err = t.ExecuteTemplate(writer, "layout", m)
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Error(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}

func about(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : about")
	var u data.User
	sess, err := session(writer, request)
	var IsAliveSession string
	if err == nil {
		IsAliveSession = "true"
		u, _ = sess.User()
	}
	m := map[string]string{
		"IsAliveSession": IsAliveSession,
		"UserName":       u.Name,
	}
	t := parseTemplateFiles("layout", "about")
	err = t.ExecuteTemplate(writer, "layout", m)
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Error(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}

func contact(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : contact")
	var u data.User
	sess, err := session(writer, request)
	var IsAliveSession string
	if err == nil {
		IsAliveSession = "true"
		u, _ = sess.User()
	}
	m := map[string]string{
		"IsAliveSession": IsAliveSession,
		"UserName":       u.Name,
	}
	t := parseTemplateFiles("layout", "contact")
	err = t.ExecuteTemplate(writer, "layout", m)
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Error(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}

func serverInfo(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : info")
	_, err := session(writer, request)
	var infoMsg string
	if err != nil {
		infoMsg = "Can't get session"
	} else {
		infoMsg = "Session alive"
	}
	m := map[string]string{
		"InfoMsg": infoMsg,
	}
	t := parseTemplateFiles("layout", "info")
	err = t.ExecuteTemplate(writer, "layout", m)
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Error(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}

func signup(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : signup")
	t := parseTemplateFiles("layout", "signup")
	err := t.ExecuteTemplate(writer, "layout", nil)
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Error(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}

func login(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : login")
	vals := request.URL.Query()
	if vals != nil {
		_, err := session(writer, request)
		if err == nil {
			http.Redirect(writer, request, "/", 302)
		}
	}
	t := parseTemplateFiles("layout", "login")
	err := t.ExecuteTemplate(writer, "layout", vals.Get("msg"))
	if err != nil {
		errMsg := "Error : Can't execute template"
		log.Error(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
}

// POST /signup
// Create the user account
func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		errMsg := "Error :Cannot parse form"
		log.Error(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		errMsg := "Error :Cannot create user"
		log.Error(errMsg)
		error_message(writer, request, errMsg)
		panic(err)
	}
	http.Redirect(writer, request, "/login?msg=please login", 302)
}

// POST /authenticate
// Authenticate the user given the email and password
func authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		log.Error("Cannot find user")
		http.Redirect(writer, request, "/login?msg=not found", 302)
	}
	if user.Password == data.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			errMsg := "Error :Cannot create session"
			log.Error(errMsg)
			error_message(writer, request, errMsg)
			panic(err)
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login?msg=miss password", 302)
	}
}

// GET /logout
// Logs the user out
func logout(writer http.ResponseWriter, request *http.Request) {
	log.Debug("Access : logout")
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		log.Debug(err, "Failed to get cookie")
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}
