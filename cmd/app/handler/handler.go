package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	st "github.com/mycontroller-org/mycontroller/pkg/storage"
	"github.com/rs/cors"
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

// StartHandler for http access
func StartHandler(config *WebConfig) error {
	router := mux.NewRouter()

	// register routes
	registerStatusRoutes(router)
	registerGatewayRoutes(router)

	// fs := http.FileServer(http.Dir("/app/web"))
	// router.Handle("/", fs)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	// Insert the middleware
	handler := c.Handler(router)
	//handler := cors.Default().Handler(router)

	fmt.Println("Listening...")
	return http.ListenAndServe(":8080", handler)
}

func params(r *http.Request) map[string]string {
	return mux.Vars(r)
}
