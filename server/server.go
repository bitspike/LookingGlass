package server

import (
	"cloudkarafka-mgmt/server/api"
	"cloudkarafka-mgmt/zookeeper"

	"github.com/gorilla/mux"

	"net/http"
	"time"
)

func ah(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, _ := r.BasicAuth()
		if !zookeeper.ValidateScramLogin(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fn(w, r)
	}
}

func apiRoutes(r *mux.Router) {
	a := r.PathPrefix("/api").Subrouter()
	a.HandleFunc("/acl/{topic}", ah(api.Acl))
	a.HandleFunc("/brokers", ah(api.Brokers))
	a.HandleFunc("/brokers/{id}", ah(api.Broker))
	a.HandleFunc("/topics", ah(api.Topics))
	a.HandleFunc("/topics/{topic}", ah(api.Topic))
	a.HandleFunc("/topics/{topic}/config", ah(api.Config))
	a.HandleFunc("/topics/{topic}/{partition}", ah(api.Partition))
	a.HandleFunc("/consumers", ah(api.Consumers))
	a.HandleFunc("/consumers/{name}", ah(api.Consumer))
	a.HandleFunc("/users", ah(api.Users))
	a.HandleFunc("/users/{name}", ah(api.User))
}

func Start(port string) {
	r := mux.NewRouter()
	apiRoutes(r)

	r.HandleFunc("/", ah(func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "server/views/home.html")
	}))

	r.HandleFunc("/topics/{topic}", ah(func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "server/views/topic.html")
	}))

	r.HandleFunc("/brokers/{id}", ah(func(w http.ResponseWriter, req *http.Request) {
		http.ServeFile(w, req, "server/views/broker.html")
	}))

	http.Handle("/js/", http.FileServer(http.Dir("server/public/")))
	http.Handle("/css/", http.FileServer(http.Dir("server/public/")))
	http.Handle("/", r)
	s := &http.Server{
		Addr:         port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}