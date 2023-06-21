package api

import (
	"net/http"
	"net/http/httptest"
	"test-go-server/pkg/resource"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

//go:generate mockgen -build_flags=--mod=mod -package api -destination ./mock_resources.go -source=../../pkg/resource/resource.go Resources
//go:generate mockgen -build_flags=--mod=mod -package api -destination ./mock_logger.go -source=../../logger/logger.go Logger

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	rs := NewMockResources(ctrl)
	l := NewMockLogger(ctrl)

	rs.EXPECT().List().Return([]*resource.Resource{{ID: "1234", Data: "data"}})

	a := NewAPI(rs, l)

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(a.List)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"ID":"1234","Data":"data"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
