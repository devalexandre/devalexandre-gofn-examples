package main

import (
	"github.com/devalexandre/gofn-examples/internal/http"
	"github.com/devalexandre/gofn/array"
	"strings"
)

func main() {

	api := http.NewApi()
	tacos, err := api.GetTacos()
	if err != nil {
		panic(err)
	}

	filteredTacos := array.Filter(tacos, func(taco http.Taco) bool {
		return strings.Contains(taco.Description, "doce")
	})

	for _, taco := range filteredTacos {
		println(taco.Description)
	}

	println(len(filteredTacos))

}
