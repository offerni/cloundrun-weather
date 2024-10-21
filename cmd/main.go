package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/offerni/cloundrun-weather/viacep"
	"github.com/offerni/cloundrun-weather/weatherapi"
)

func main() {
	if os.Getenv("GO_ENV") != "production" {
		if err := godotenv.Load(); err != nil {
			log.Println("Error loading .env file")
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("PORT not set, defaulting to %s", port)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		params := r.URL.Query()
		cep := params.Get("cep")
		if len(cep) != 8 {
			http.Error(w, "Invalid CEP", http.StatusBadRequest)
			return
		}

		address, err := viacep.GetAddress(ctx, cep)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if address == nil {
			http.Error(w, "Address not found", http.StatusNotFound)
			return
		}

		apiKey := os.Getenv("WEATHER_API_API_KEY")
		if apiKey == "" {
			http.Error(w, "Weather API key not set", http.StatusInternalServerError)
			return
		}

		client, err := weatherapi.NewAPIClient(weatherapi.NewAPIClientOpts{
			APIKey: apiKey,
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

	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
