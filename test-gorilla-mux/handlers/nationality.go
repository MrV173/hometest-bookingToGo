package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	nationalitydto "test-gorilla-mux/dto/nationality"
	dto "test-gorilla-mux/dto/result"
	"test-gorilla-mux/models"
	"test-gorilla-mux/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerNationality struct {
	NationalityRepository repositories.NationalityRepository
}

func HandlerNationality(NationalityRepository repositories.NationalityRepository) *handlerNationality {
	return &handlerNationality{NationalityRepository}
}

func (h *handlerNationality) FindNationalities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nationalities, err := h.NationalityRepository.FindNationalities()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{
		Code: http.StatusOK,
		Data: nationalities,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerNationality) GetNationality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	nationality, err := h.NationalityRepository.GetNationality(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{
		Code: http.StatusOK,
		Data: nationality,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerNationality) CreateNationality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(nationalitydto.CreateNationalityRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	nationality := models.Nationality{

		Nationality_Name: request.Nationality_Name,
		Nationality_Code: request.Nationality_Code,
	}

	data, err := h.NationalityRepository.CreateNationality(nationality)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{
		Code: http.StatusOK,
		Data: convertNationalityResponse(data),
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerNationality) UpdateNationality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")

	request := new(nationalitydto.CreateNationalityRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	nationality, err := h.NationalityRepository.GetNationality(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Nationality_Name != "" {
		nationality.Nationality_Name = request.Nationality_Name
	}

	if request.Nationality_Code != "" {
		nationality.Nationality_Code = request.Nationality_Code
	}

	data, err := h.NationalityRepository.UpdateNationality(nationality)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertNationalityResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerNationality) DeleteNationality(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	nationality, err := h.NationalityRepository.GetNationality(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.NationalityRepository.DeleteNationality(nationality)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertNationalityResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertNationalityResponse(u models.Nationality) nationalitydto.NationalityResponse {
	return nationalitydto.NationalityResponse{
		Nationality_Name: u.Nationality_Name,
		Nationality_Code: u.Nationality_Code,
	}
}
