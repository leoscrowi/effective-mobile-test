package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/leoscrowi/effective-mobile-test/internal/app/subscription/dto"
	"github.com/leoscrowi/effective-mobile-test/internal/utils"
	"log"
	"net/http"
)

func (s *SubscriptionController) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	op := `SubscriptionController.CreateSubscription`
	var d dto.CreateSubscriptionRequest

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}
	defer r.Body.Close()

	if err := dto.ValidateDTO(d); err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := s.usecase.CreateSubscription(r.Context(), d)
	if err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, fmt.Sprint("Error while creating subscription"), http.StatusInternalServerError)
		return
	}

	utils.WriteHeader(w, http.StatusCreated, resp)
}

func (s *SubscriptionController) ReadSubscription(w http.ResponseWriter, r *http.Request) {
	op := `SubscriptionController.ReadSubscription`

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Printf("%s: %v;", op, errors.New("id parameter is missing"))
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	subscriptionID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	resp, err := s.usecase.ReadSubscription(r.Context(), dto.ReadSubscriptionRequest{ID: subscriptionID})
	if err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, "Error while reading subscription", http.StatusInternalServerError)
		return
	}

	utils.WriteHeader(w, http.StatusOK, resp)
}

func (s *SubscriptionController) EditSubscription(w http.ResponseWriter, r *http.Request) {
	op := `SubscriptionController.EditSubscription`
	var d dto.EditSubscriptionRequset

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := dto.ValidateDto(d); err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := s.usecase.EditSubscription(r.Context(), d)
	if err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, "Error while editing subscription", http.StatusInternalServerError)
		return
	}

	utils.WriteHeader(w, http.StatusOK, map[string]string{"message": "Subscription edited successfully"})
}

func (s *SubscriptionController) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	op := `SubscriptionController.DeleteSubscription`

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Printf("%s: %v;", op, errors.New("id parameter is missing"))
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	subscriptionID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	if err = s.usecase.DeleteSubscription(r.Context(), dto.DeleteSubscriptionRequest{ID: subscriptionID}); err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, "Error while deleting subscription", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *SubscriptionController) ReadSubscriptionsList(w http.ResponseWriter, r *http.Request) {
	op := `SubscriptionController.ReadSubscriptionsList`

	resp, err := s.usecase.ReadSubscriptionsList(r.Context())
	if err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, "Error while reading subscriptions list", http.StatusInternalServerError)
		return
	}

	utils.WriteHeader(w, http.StatusOK, resp)
}

func (s *SubscriptionController) GetSubscriptionsAmount(writer http.ResponseWriter, request *http.Request) {
	// TODO implement me
	panic("implement me")
}
