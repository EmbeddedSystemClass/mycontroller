package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	mc "github.com/mycontroller-org/mycontroller/pkg"
	"github.com/mycontroller-org/mycontroller/pkg/model"
)

func listGateways(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var results []model.Gateway
	err := storage().Find(mc.EntGateway, &results)
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

func updateGateway(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	d, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var g model.Gateway
	err = json.Unmarshal(d, g)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	storage().Save(g)
}
