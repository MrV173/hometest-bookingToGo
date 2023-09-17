package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	customerdto "test-gorilla-mux/dto/customer"
	dto "test-gorilla-mux/dto/result"
	"test-gorilla-mux/models"
	"test-gorilla-mux/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerCustomer struct {
	CustomerRepository repositories.CustomerRepository
}

func HandlerCustomer(CustomerRepository repositories.CustomerRepository) *handlerCustomer {
	return &handlerCustomer{CustomerRepository}
}

func (h *handlerCustomer) FindCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	customers, err := h.CustomerRepository.FindCustomers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{
		Code: http.StatusOK,
		Data: customers,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCustomer) GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	customer, err := h.CustomerRepository.GetCustomer(id)
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
		Data: customer,
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCustomer) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(customerdto.CreateCustomerRequest)
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

	customer := models.Customer{
		CST_Name:      request.CST_Name,
		CST_Dob:       request.CST_Dob,
		CST_PhoneNum:  request.CST_PhoneNum,
		NationalityID: request.NationalityID,
		CST_Email:     request.CST_Email,
	}

	data, err := h.CustomerRepository.CreateCustomer(customer)
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
		Data: convertResponse(data),
	}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCustomer) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")

	request := new(customerdto.UpdateCustomerRequest)
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
	customer, err := h.CustomerRepository.GetCustomer(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.CST_Name != "" {
		customer.CST_Name = request.CST_Name
	}

	if request.CST_Dob != "" {
		customer.CST_Dob = request.CST_Dob
	}

	if request.CST_PhoneNum != "" {
		customer.CST_PhoneNum = request.CST_PhoneNum
	}

	if request.CST_Email != "" {
		customer.CST_Email = request.CST_Email
	}

	data, err := h.CustomerRepository.UpdateCustomer(customer)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCustomer) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	customer, err := h.CustomerRepository.GetCustomer(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.CustomerRepository.DeleteCustomer(customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponse(u models.Customer) customerdto.CustomerResponse {
	return customerdto.CustomerResponse{
		ID:           u.ID,
		CST_Name:     u.CST_Name,
		CST_Dob:      u.CST_Dob,
		CST_PhoneNum: u.CST_PhoneNum,
		CST_Email:    u.CST_Email,
	}
}
