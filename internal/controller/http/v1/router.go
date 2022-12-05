package v1

import (
	"encoding/json"
	"net/http"

	"github.com/ilyasbulat/rest_api/internal/usecase"
	"github.com/julienschmidt/httprouter"

	_ "github.com/ilyasbulat/rest_api/docs"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(router *httprouter.Router, l *zap.Logger, t usecase.Translation) {

	// Swagger
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	// Routers

	newTranslationRoutes(router, t, l)
}

func respond(w http.ResponseWriter, code int, data interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
