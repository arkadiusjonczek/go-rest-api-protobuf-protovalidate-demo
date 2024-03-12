package app

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
)

func NewCrudRouter[T Entity](handler *GenericCrudHandler[T]) *mux.Router {
	mux := mux.NewRouter()

	name := reflect.TypeFor[T]().Name()

	path := fmt.Sprintf("/%s", strings.ToLower(name))
	pathWithId := fmt.Sprintf("%s/{id}", path)

	log.Printf("Using path '%s' and '%s' for '%s'", path, pathWithId, name)

	mux.HandleFunc(path, handler.ReadAll).Methods(http.MethodGet)
	mux.HandleFunc(path, handler.Create).Methods(http.MethodPost)
	mux.HandleFunc(pathWithId, handler.Read).Methods(http.MethodGet)
	mux.HandleFunc(pathWithId, handler.Update).Methods(http.MethodPut)
	mux.HandleFunc(pathWithId, handler.Delete).Methods(http.MethodDelete)

	return mux
}
