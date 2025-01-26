package handler_test

import (
	"net/http"
	"net/http/httptest"
	"new-api/internal/handler"
	"testing"
)

func Test_PostNews(t *testing.T) {
	testCase := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)
			handler.PostNews()(w, r)

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected :%d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}
func Test_GetAllNews(t *testing.T) {
	testCase := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)
			handler.GetAllNews()(w, r)

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected :%d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}
func Test_GetNewsById(t *testing.T) {
	testCase := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)
			handler.GetNewsById()(w, r)

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected :%d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}
func Test_UpdateNewsById(t *testing.T) {
	testCase := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)
			handler.UpdateNewsById()(w, r)

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected :%d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}
func Test_DeleteNewsById(t *testing.T) {
	testCase := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)
			handler.DeleteNewsById()(w, r)

			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected :%d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}
