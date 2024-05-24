package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleCafeListWhenDataCorrect(t *testing.T) {
	// Данные для запроса
	city := "moscow"
	count := 2
	reqURL := fmt.Sprintf("/cafe?count=%d&city=%s", count, city)

	req := httptest.NewRequest("GET", reqURL, nil)

	// Распаковка ответа
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handleCafeList)
	handler.ServeHTTP(responseRecorder, req)

	// Проверка наличия ошибки
	var err error
	require.NoError(t, err)
	// Проверка на статус 200
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	// Проверка на пусте тело ответа
	require.NotEmpty(t, responseRecorder.Body)
}

func TestHandleCafeListWhenCityIsWrong(t *testing.T) {
	// Данные для запроса
	city := "dubai"
	count := 1
	reqURL := fmt.Sprintf("/cafe?count=%d&city=%s", count, city)

	req := httptest.NewRequest("GET", reqURL, nil)

	// Распаковка ответа
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handleCafeList)
	handler.ServeHTTP(responseRecorder, req)

	// Проверка наличия ошибки
	var err error
	require.NoError(t, err)
	// Проверка наличия 400 ошибки
	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	// Проверка тела ответа
	actualBody := "wrong city value"
	assert.Equal(t, responseRecorder.Body.String(), actualBody)
}

func TestHandleCafeListWhenCountMoreThanTotal(t *testing.T) {
	// Данные для запроса
	city := "moscow"
	count := 100000
	reqURL := fmt.Sprintf("/cafe?count=%d&city=%s", count, city)

	req := httptest.NewRequest("GET", reqURL, nil)

	// Распаковка ответа
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handleCafeList)
	handler.ServeHTTP(responseRecorder, req)

	// Проверка наличия ошибки
	var err error
	require.NoError(t, err)
	// Проверка на пусте тело ответа
	require.NotEmpty(t, responseRecorder.Body)
	// Проверка на ожидаемую длину ответа
	actualAmountOfItems := strings.Split(responseRecorder.Body.String(), ",")
	assert.Len(t, actualAmountOfItems, len(cafeList[city]))
}
