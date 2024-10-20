package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/offerni/cloundrun-weather/viacep"
	"github.com/offerni/cloundrun-weather/weatherapi"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		params := r.URL.Query()
		cep := params.Get("cep")

		address, err := viacep.GetAddress(ctx, cep)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		aPIKey := os.Getenv("WEATHER_API_API_KEY")

		client, err := weatherapi.NewAPIClient(weatherapi.NewAPIClientOpts{
			APIKey: aPIKey,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		temperature, err := client.GetCurrentInfo(ctx, address.Localidade)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(temperature)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
