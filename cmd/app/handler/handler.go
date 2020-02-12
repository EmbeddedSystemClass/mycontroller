package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	st "github.com/mycontroller-org/mycontroller/pkg/storage"
	version "github.com/mycontroller-org/mycontroller/pkg/version"
	"github.com/rs/cors"
)

func storage() st.Client {
	return st.StorageClient
}

// StartHandler for http access
func StartHandler() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/version", versionData)
	mux.HandleFunc("/api/status", status)
	mux.HandleFunc("/api/gateways", listGateways)
	mux.HandleFunc("/api/gateways/update", updateGateway)

	// fs := http.FileServer(http.Dir("/app/web"))
	// mux.Handle("/", fs)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	// Insert the middleware
	handler := c.Handler(mux)
	//handler := cors.Default().Handler(mux)

	fmt.Println("Listening...")
	return http.ListenAndServe(":8080", handler)
}

func versionData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	v := version.Get()
	od, err := json.Marshal(&v)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(od)
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := map[string]interface{}{
		"time": time.Now(),
	}
	hn, err := os.Hostname()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	s["hostname"] = hn
	od, err := json.Marshal(&s)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Write(od)
}
