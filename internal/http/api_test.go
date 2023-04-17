package http_test

import (
	"github.com/devalexandre/gofn-examples/internal/http"
	"testing"
)

func TestGetTacos(t *testing.T) {
	api := http.NewApi()
	tacos, err := api.GetTacos()
	if err != nil {
		t.Error(err)
	}

	if len(tacos) == 0 {
		t.Error("no tacos found")
	}
}
