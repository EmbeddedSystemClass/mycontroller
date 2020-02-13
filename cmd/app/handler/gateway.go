package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	mc "github.com/mycontroller-org/mycontroller/pkg"
	"github.com/mycontroller-org/mycontroller/pkg/model"
)

func registerGatewayRoutes(router *mux.Router) {
	router.HandleFunc("/api/gateways", listGateways).Methods(http.MethodGet)
	router.HandleFunc("/api/gateways/{uuid}", getGateway).Methods(http.MethodGet)
	router.HandleFunc("/api/gateways", updateGateway).Methods(http.MethodPost)
}

func listGateways(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	results := make([]model.Gateway, 0)
	err := storage().Find(mc.EntGateway, nil, &results)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	od, err := json.Marshal(&results)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(od)
}

func getGateway(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := params(r)

	var result model.Gateway
	err := storage().FindOne(mc.EntGateway, p, &result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	od, err := json.Marshal(&result)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(od)
}

func updateGateway(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	d, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var g model.Gateway
	err = json.Unmarshal(d, &g)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = g.Save()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
