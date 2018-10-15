package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"./data"
	"github.com/BurntSushi/toml"
	log "github.com/cihub/seelog"
	"github.com/fsnotify/fsnotify"
)

type ServerConfig struct {
	Address        string
	ReadTimeout    int64
	WriteTimeout   int64
	Static         string
	Logsettingfile string
}

type Config struct {
	Server   ServerConfig
	DBconfig data.Config
}

const File = "./config.toml"

var config Config

func initialize() {
	// Configuration
	watchConfig(&config)
	// Log
	logger, err := log.LoggerFromConfigAsFile(config.Server.Logsettingfile)
	if err != nil {
		panic(err)
	}
	log.ReplaceLogger(logger)
	defer log.Flush()
	log.Debug("Success : Setup log")
	// Database
	data.SetupDB(config.DBconfig)
	log.Debug("Success : Setup db")
}

// Convenience function to redirect to the error message page
func error_message(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// Checks if the user is logged in and has a session, if not err is not nil
func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// parse HTML templates
// pass in a list of file names, and get a template
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func watchConfig(conf *Config) {
	reload(conf)
	go monitor(conf)
}

func reload(conf *Config) {
	_, err := toml.DecodeFile(File, &conf)
	if err != nil {
		panic(err)
	}
}

func monitor(conf *Config) {
	var err error
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()
	for {
		err = watcher.Add(File)
		if err != nil {
			panic(err)
		}
		select {
		case <-watcher.Events:
			reload(conf)
			if err != nil {
				panic(err)
			}
		}
	}
}
