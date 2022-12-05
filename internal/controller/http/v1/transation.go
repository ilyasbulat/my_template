package v1

import (
	"encoding/json"
	"net/http"

	"github.com/ilyasbulat/rest_api/internal/apperror"
	"github.com/ilyasbulat/rest_api/internal/entity"
	"github.com/ilyasbulat/rest_api/internal/usecase"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type translationRoutes struct {
	t usecase.Translation
	l *zap.Logger
}

func newTranslationRoutes(router *httprouter.Router, t usecase.Translation, l *zap.Logger) {
	r := &translationRoutes{t, l}

	{
		router.GET("/history", middleware(r.history))
		router.POST("/do-translate", middleware(r.doTranslate))
	}
}

type historyResponse struct {
	History []entity.Translation `json:"history"`
}

// @Summary     Show history
// @Description Show all translation history
// @ID          history
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} apperror.AppError
// @Router      /translation/history [get]
func (tr *translationRoutes) history(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	translations, err := tr.t.History(r.Context())
	if err != nil {
		tr.l.Error(err.Error())
		return err
	}

	respond(w, http.StatusOK, historyResponse{translations})
	return nil
}

type doTranslateRequest struct {
	Source      string `json:"source"       binding:"required"  example:"auto"`
	Destination string `json:"destination"  binding:"required"  example:"en"`
	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
}

// @Summary     Translate
// @Description Translate a text
// @ID          do-translate
// @Tags  	    translation
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} apperror.AppError
// @Failure     500 {object} apperror.AppError
// @Router      /translation/do-translate [post]
func (tr *translationRoutes) doTranslate(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	var request doTranslateRequest
	json.NewDecoder(r.Body).Decode(&request)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		tr.l.Error(err.Error(), zap.String("service:", "http - v1 - doTranslate"))
		return apperror.ErrBadRequest
	}

	translation, err := tr.t.Translate(
		r.Context(),
		entity.Translation{
			Source:      request.Source,
			Destination: request.Destination,
			Original:    request.Original,
		},
	)
	if err != nil {
		tr.l.Error(err.Error(), zap.String("service:", "http - v1 - doTranslate"))
		return err
	}
	respond(w, http.StatusOK, translation)
	return nil
}
