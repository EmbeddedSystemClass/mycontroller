package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/zap"
)

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

	zap.L().Info("Http server is listening on the port 8080")
	return http.ListenAndServe(":8080", handler)
}
