package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2"
)

const (
	PostOnly   = true
	GetAllowed = false
)

type Server struct {
	Port           int
	MgoDb          *mgo.Database
	CollectionName string
	IsGerman       bool
	Debug          bool
	Handlers       map[string]http.HandlerFunc
}

func MakeServer(port int, mgoDb *mgo.Database, collectionName string, isGerman bool, debug bool) Server {
	server := Server{}

	server.Port = port
	server.MgoDb = mgoDb
	server.CollectionName = collectionName
	server.IsGerman = isGerman
	server.Debug = debug
	server.Handlers = map[string]http.HandlerFunc{}

	return server
}

func (s *Server) Start() {
	for path, handler := range s.Handlers {
		http.HandleFunc(path, handler)
	}

	fmt.Printf("Ready to listen: %d\n", s.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil)
}

func (s *Server) AddHandler(
	path string,
	fn func(http.ResponseWriter, *http.Request, *mgo.Database, string, bool, bool) error,
	isPostOnly bool,
) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if isPostOnly && r.Method != "POST" {
			http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)

			return
		}

		err := fn(w, r, s.MgoDb, s.CollectionName, s.IsGerman, s.Debug)
		if err != nil {
			json.NewEncoder(w).Encode(fmt.Sprint(err))
			LogErr(err, true)
		}
	}

	s.Handlers[path] = handler
}
