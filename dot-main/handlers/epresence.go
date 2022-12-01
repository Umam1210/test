package handlers

import (
	"encoding/json"
	epresencesdto "journey/dto/epresence"
	dto "journey/dto/result"
	"journey/models"
	"journey/repositories"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/mux"
)

type handlerEpresence struct {
	EpresenceRepository repositories.EpresenceRepository
}

func HandlerEpresence(EpresenceRepository repositories.EpresenceRepository) *handlerEpresence {
	return &handlerEpresence{EpresenceRepository}
}

func (h *handlerEpresence) FindEpresences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	epresences, err := h.EpresenceRepository.FindEpresences()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: epresences}
	json.NewEncoder(w).Encode(response)
}
func (h *handlerEpresence) FindEpresencesbyUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user_id, _ := strconv.Atoi(mux.Vars(r)["user_id"])

	epresences, err := h.EpresenceRepository.FindEpresencesbyUserId(user_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: epresences}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEpresence) GetEpresence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var epresence models.Epresence

	epresence, err := h.EpresenceRepository.GetEpresence(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpresence(epresence)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEpresence) CreateEpresence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(epresencesdto.CreateEpresenceRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	date := time.Now()
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	userId := int(userInfo["id"].(float64))

	epresence := models.Epresence{
		Type:   request.Type,
		UserID: userId,
		Date:   date,
	}

	data, err := h.EpresenceRepository.CreateEpresence(epresence)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpresence(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEpresence) UpdateEpresence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(epresencesdto.UpdateEpresenceRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	epresence, err := h.EpresenceRepository.GetEpresence(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Type != "" {
		epresence.Type = request.Type
	}
	data, err := h.EpresenceRepository.UpdateEpresence(epresence)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpresence(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerEpresence) DeleteEpresence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	epresence, err := h.EpresenceRepository.GetEpresence(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.EpresenceRepository.DeleteEpresence(epresence)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseEpresence(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseEpresence(u models.Epresence) epresencesdto.EpresenceResponse {
	return epresencesdto.EpresenceResponse{
		ID:   u.ID,
		User: u.User,
		Date: u.Date,
		Type: u.Type,
	}
}
