package property

import (
	"net/http"
	services "rentroom/internal/services/property"
	"rentroom/internal/validators"
	"rentroom/utils"

	"github.com/gorilla/mux"
)

type publichHandler struct {
	propertyService *services.PropertyService
}

func NewPublicHandler(propertyService *services.PropertyService) *publichHandler {
	return &publichHandler{propertyService: propertyService}
}

func (h *publichHandler) PublicList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// PARSE
		countryStr := r.URL.Query().Get("country")
		countryID, err := validators.ParseCountryID(countryStr)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// QUERY
		properties, err := h.propertyService.ListPublicProperties(countryID)
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "properties returned",
			Data:    properties,
		}, http.StatusOK)
	}
}

func (h *publichHandler) PublicGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// PARSE
		propertyID, err := validators.ParsePropertyID(mux.Vars(r))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusBadRequest)
			return
		}

		// QUERY
		property, err := h.propertyService.GetPublishedByID(uint(propertyID))
		if err != nil {
			utils.JSONError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// RESPONSE
		utils.JSONResponse(w, utils.Response{
			Success: true,
			Message: "property returned",
			Data:    property,
		}, http.StatusOK)
	}
}
