package api

import (
	"cloudkarafka-mgmt/zookeeper"

	"github.com/gorilla/mux"

	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Name, Password string
}

func Whoami(w http.ResponseWriter, r *http.Request, p zookeeper.Permissions) {
	writeJson(w, p)
}

func Users(w http.ResponseWriter, r *http.Request, p zookeeper.Permissions) {
	switch r.Method {
	case "GET":
		users(w, p)
	case "POST":
		if !p.ClusterWrite() {
			http.NotFound(w, r)
			return
		}
		u, err := decodeUser(r)
		if err != nil {
			internalError(w, err.Error())
		} else {
			createUser(w, u)
		}
	}
}

func User(w http.ResponseWriter, r *http.Request, p zookeeper.Permissions) {
	vars := mux.Vars(r)
	switch r.Method {
	case "GET":
		if !p.ClusterRead() && vars["name"] != p.Username {
			http.NotFound(w, r)
			return
		}
		user := zookeeper.PermissionsFor(vars["name"])
		writeJson(w, user)
	case "DELETE":
		fmt.Println(p.ClusterWrite())
		if !p.ClusterWrite() {
			http.NotFound(w, r)
			return
		}
		zookeeper.DeleteUser(vars["name"])
		w.WriteHeader(http.StatusNoContent)
	}
}

func decodeUser(r *http.Request) (user, error) {
	var (
		u   user
		err error
	)
	switch r.Header.Get("content-type") {
	case "application/json":
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&u)
	default:
		err = r.ParseForm()
		u = user{Name: r.PostForm.Get("name"), Password: r.PostForm.Get("password")}
	}
	return u, err
}

func users(w http.ResponseWriter, p zookeeper.Permissions) {
	users, err := zookeeper.Users(p)
	if err != nil {
		internalError(w, err.Error())
	} else {
		writeJson(w, users)
	}
}

func createUser(w http.ResponseWriter, u user) {
	err := zookeeper.CreateUser(u.Name, u.Password)
	if err != nil {
		internalError(w, err.Error())
		return
	}
}
