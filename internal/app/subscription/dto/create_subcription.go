package dto

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strconv"
	"strings"
)

// CreateSubscriptionRequest описывает запрос на создание подписки
// @Description Структура для создания новой подписки
type CreateSubscriptionRequest struct {
	ServiceName string    `json:"service_name" example:"Yandex Music"`
	Price       int64     `json:"price" example:"400"`
	UserID      uuid.UUID `json:"user_id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string    `json:"start_date" example:"07-2026"`
	EndDate     *string   `json:"end_date" example:"08-2026"`
}

// CreateSubscriptionResponse описывает ответ при создании подписки
type CreateSubscriptionResponse struct {
	ID uuid.UUID `json:"id" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
}

func ValidateDTO(dto CreateSubscriptionRequest) error {
	if dto.ServiceName == "" {
		return fmt.Errorf("service_name: cannot be empty")
	}

	if dto.Price == 0 {
		return fmt.Errorf("price: cannot be zero")
	}

	if dto.Price < 0 {
		return fmt.Errorf("price: cannot be negative")
	}

	if dto.UserID == uuid.Nil {
		return fmt.Errorf("user_id: cannot be empty")
	}

	if dto.StartDate == "" {
		return fmt.Errorf("start_date: cannot be empty")
	}

	if !isValidDateFormat(dto.StartDate) {
		return fmt.Errorf("start_date: invalid format, expected MM-YYYY")
	}

	if dto.EndDate != nil && !isValidDateFormat(*dto.EndDate) {
		return fmt.Errorf("end_date: invalid format, expected MM-YYYY")
	}

	if dto.EndDate != nil && !isEndDateAfterStartDate(dto.StartDate, *dto.EndDate) {
		return fmt.Errorf("end_date: cannot be before or equal to start_date")
	}

	return nil
}

func isValidDateFormat(date string) bool {
	pattern := `^(0[1-9]|1[0-2])-\d{4}$`
	matched, _ := regexp.MatchString(pattern, date)
	if !matched {
		return false
	}

	parts := strings.Split(date, "-")
	if len(parts) != 2 {
		return false
	}

	month, err := strconv.Atoi(parts[0])
	year, errYear := strconv.Atoi(parts[1])

	if err != nil || errYear != nil {
		return false
	}

	if month < 1 || month > 12 {
		return false
	}

	if year < 1 {
		return false
	}

	return true
}

func isEndDateAfterStartDate(startDate, endDate string) bool {
	startParts := strings.Split(startDate, "-")
	endParts := strings.Split(endDate, "-")

	startYear, _ := strconv.Atoi(startParts[1])
	startMonth, _ := strconv.Atoi(startParts[0])

	endYear, _ := strconv.Atoi(endParts[1])
	endMonth, _ := strconv.Atoi(endParts[0])

	if endYear > startYear {
		return true
	}

	if endYear == startYear && endMonth > startMonth {
		return true
	}

	return false
}
