package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"

	v1 "github.com/arkadiusjonczek/go-rest-api-protobuf-protovalidate-demo.git/pkg/proto/demo/v1"
)

type CrudHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	Read(w http.ResponseWriter, r *http.Request)
	ReadAll(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type Entity interface {
	v1.Customer | v1.Customer2
}

var _ CrudHandler = (*GenericCrudHandler[v1.Customer])(nil)

type GenericCrudHandler[E Entity] struct {
	decoder Decoder[E]
	store   Store[E]
}

func NewGenericCrudHandler[E Entity](decoder Decoder[E], store Store[E]) *GenericCrudHandler[E] {
	return &GenericCrudHandler[E]{
		decoder: decoder,
		store:   store,
	}
}

func (h *GenericCrudHandler[E]) Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got request: %v", r)

	body, _ := io.ReadAll(r.Body)

	log.Printf("Got body: %s", body)

	// Unmarshal

	entry, err := h.decoder.Decode(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	log.Printf("Got entry of type %s: %#v", reflect.TypeOf(entry), entry)

	// Add to store

	id, err := h.store.Add(entry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Created entry with id: %d", id)
}

func (h *GenericCrudHandler[E]) Read(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got request: %v", r)

	vars := mux.Vars(r)
	idString := vars["id"]

	id, _ := strconv.Atoi(idString)

	// Get from store

	entry, err := h.store.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entry)
}

func (h *GenericCrudHandler[E]) ReadAll(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got request: %v", r)

	// Get from store

	entries, err := h.store.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entries)
}

func (h *GenericCrudHandler[E]) Update(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got request: %v", r)

	vars := mux.Vars(r)
	idString := vars["id"]

	id, _ := strconv.Atoi(idString)

	body, _ := io.ReadAll(r.Body)

	log.Printf("Got body: %s", body)

	// Unmarshal

	entry, err := h.decoder.Decode(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	log.Printf("Got entry of type %s: %#v", reflect.TypeOf(entry), entry)

	// Update in store

	err = h.store.Update(id, entry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated entry with id: %d", id)
}
func (h *GenericCrudHandler[E]) Delete(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got request: %v", r)

	vars := mux.Vars(r)
	idString := vars["id"]

	id, _ := strconv.Atoi(idString)

	// Delete in store

	err := h.store.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted entry with id: %d", id)
}
