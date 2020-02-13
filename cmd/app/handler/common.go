package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mycontroller-org/mycontroller/pkg/interfaces"
	st "github.com/mycontroller-org/mycontroller/pkg/storage"
)

func storage() st.Client {
	return st.StorageClient
}

// WebConfig input
type WebConfig struct {
	BindAddress  string `yaml:"bindAddress"`
	Port         uint   `yaml:"port"`
	WebDirectory string `yaml:"webDirectory"`
}

func params(r *http.Request) map[string]string {
	return mux.Vars(r)
}

func findOne(w http.ResponseWriter, r *http.Request, entityName string, entity interface{}) {
	w.Header().Set("Content-Type", "application/json")

	p := params(r)

	err := storage().FindOne(entityName, p, entity)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	od, err := json.Marshal(entity)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(od)
}

func findMany(w http.ResponseWriter, r *http.Request, entityName string, entities interface{}) {
	w.Header().Set("Content-Type", "application/json")

	p := params(r)

	err := storage().Find(entityName, p, entities)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	od, err := json.Marshal(entities)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(od)
}

func saveEntity(w http.ResponseWriter, r *http.Request, entity interfaces.Entity) {
	w.Header().Set("Content-Type", "application/json")

	d, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = json.Unmarshal(d, &entity)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = entity.Save()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
