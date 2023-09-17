package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	familylistdto "test-gorilla-mux/dto/family-list"
	dto "test-gorilla-mux/dto/result"
	"test-gorilla-mux/models"
	"test-gorilla-mux/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerFamily struct {
	FamilyRepository repositories.FamilyRepository
}

func HandlerFamily(FamilyRepository repositories.FamilyRepository) *handlerFamily {
	return &handlerFamily{FamilyRepository}
}

func (h *handlerFamily) FindFamilies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	families, err := h.FamilyRepository.FindFamilies()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{
		Code: http.StatusOK,
		Data: families,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFamily) GetFamily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	family, err := h.FamilyRepository.GetFamily(id)
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
		Data: family,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFamily) CreateFamily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(familylistdto.CreateFamilyRequest)
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

	family := models.Family{
		FL_Relation: request.FL_Relation,
		FL_Name:     request.FL_Name,
		FL_Dob:      request.FL_Dob,
		CustomerID:  uint(request.CustomerID),
	}

	data, err := h.FamilyRepository.CreateFamily(family)
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
		Data: convertFamilyResponse(data),
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFamily) UpdateFamily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")

	request := new(familylistdto.UpdateFamilyRequest)
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
	family, err := h.FamilyRepository.GetFamily(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.FL_Relation != "" {
		family.FL_Relation = request.FL_Relation
	}

	if request.FL_Name != "" {
		family.FL_Name = request.FL_Name
	}

	if request.FL_Dob != "" {
		family.FL_Dob = request.FL_Dob
	}

	data, err := h.FamilyRepository.UpdateFamily(family)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertFamilyResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerFamily) DeleteFamily(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	family, err := h.FamilyRepository.GetFamily(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.FamilyRepository.DeleteFamily(family)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertFamilyResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertFamilyResponse(u models.Family) familylistdto.FamilyResponse {
	return familylistdto.FamilyResponse{
		FL_Relation: u.FL_Relation,
		FL_Name:     u.FL_Name,
		FL_Dob:      u.FL_Dob,
	}
}
