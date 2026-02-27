package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/leoscrowi/effective-mobile-test/internal/app/subscription/dto"
	"github.com/leoscrowi/effective-mobile-test/internal/utils"
	"log"
	"net/http"
)

// CreateSubscription godoc
// @Summary Создать новую подписку
// @Description Создает новую подписку с указанными параметрами
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param request body dto.CreateSubscriptionRequest true "Данные подписки"
// @Success 201 {object} dto.CreateSubscriptionResponse
// @Failure 400 {string} string "Ошибка валидации данных"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions [post]
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

// ReadSubscription godoc
// @Summary Получить подписку по ID
// @Description Получает информацию о подписке по её уникальному идентификатору
// @Tags subscriptions
// @Produce json
// @Param id path string true "ID подписки (UUID)"
// @Success 200 {object} dto.ReadSubscriptionResponse
// @Failure 400 {string} string "Отсутствует или неверный параметр id"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [get]
func (s *SubscriptionController) ReadSubscription(w http.ResponseWriter, r *http.Request) {
	op := `SubscriptionController.ReadSubscription`

	id := chi.URLParam(r, "id")
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

// EditSubscription godoc
// @Summary Обновить подписку
// @Description Обновляет информацию о подписке по её ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path string true "ID подписки (UUID)"
// @Param request body dto.EditSubscriptionRequset true "Данные для обновления подписки"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Ошибка валидации данных"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [patch]
func (s *SubscriptionController) EditSubscription(w http.ResponseWriter, r *http.Request) {
	op := `SubscriptionController.EditSubscription`
	var d dto.EditSubscriptionRequset

	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	urlID := chi.URLParam(r, "id")
	if urlID == "" {
		log.Printf("%s: %v;", op, errors.New("id parameter is missing"))
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	uuidID, err := uuid.Parse(urlID)
	if err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	d.ID = uuidID

	if err := dto.ValidateDto(d); err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, "Validation error: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = s.usecase.EditSubscription(r.Context(), d)
	if err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(w, "Error while editing subscription", http.StatusInternalServerError)
		return
	}

	utils.WriteHeader(w, http.StatusOK, map[string]string{"message": "Subscription edited successfully"})
}

// DeleteSubscription godoc
// @Summary Удалить подписку
// @Description Удаляет подписку по её ID
// @Tags subscriptions
// @Param id path string true "ID подписки (UUID)"
// @Success 204
// @Failure 400 {string} string "Отсутствует или неверный параметр id"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions/{id} [delete]
func (s *SubscriptionController) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	op := `SubscriptionController.DeleteSubscription`

	id := chi.URLParam(r, "id")
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

// ReadSubscriptionsList godoc
// @Summary Получить список всех подписок
// @Description Получает список всех подписок в системе
// @Tags subscriptions
// @Produce json
// @Success 200 {object} dto.ReadSubscriptionsListResponse
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions [get]
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

// GetSubscriptionsAmount godoc
// @Summary Получить суммарную стоимость подписок
// @Description Вычисляет сумму стоимости всех подписок с фильтрацией по user_id и service_name
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param request body dto.GetSubscriptionsAmountRequest true "Фильтры для подсчета"
// @Success 200 {object} dto.GetSubscriptionsAmountResponse
// @Failure 400 {string} string "Ошибка при чтении JSON"
// @Failure 500 {string} string "Внутренняя ошибка сервера"
// @Router /subscriptions/amount [get]
func (s *SubscriptionController) GetSubscriptionsAmount(writer http.ResponseWriter, request *http.Request) {
	op := `SubscriptionController.GetSubscriptionsAmount`
	var d dto.GetSubscriptionsAmountRequest

	if err := json.NewDecoder(request.Body).Decode(&d); err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(writer, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer request.Body.Close()

	resp, err := s.usecase.GetSubscriptionsAmount(request.Context(), d)
	if err != nil {
		log.Printf("%s: %v;", op, err)
		http.Error(writer, "Error while getting subscriptions amount", http.StatusInternalServerError)
		return
	}

	utils.WriteHeader(writer, http.StatusOK, resp)
}
