package http

import (
	"encoding/json"
	"net/http"
)

type Api struct {
	httpClient *http.Client
}

func NewApi() *Api {
	return &Api{
		httpClient: &http.Client{},
	}
}

//get all tacos
func (a *Api) GetTacos() ([]Taco, error) {

	var tacos []Taco

	resp, err := a.httpClient.Get("http://localhost:4000/api/v1/food")
	if err != nil {
		return tacos, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&tacos); err != nil {
		return tacos, err
	}

	return tacos, nil
}
