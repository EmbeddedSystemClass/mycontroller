package handler

import (
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
	findMany(w, r, mc.EntGateway, &[]model.Gateway{})
}

func getGateway(w http.ResponseWriter, r *http.Request) {
	findOne(w, r, mc.EntGateway, &model.Gateway{})
}

func updateGateway(w http.ResponseWriter, r *http.Request) {
	saveEntity(w, r, &model.Gateway{})
}
