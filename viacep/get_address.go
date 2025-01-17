package viacep

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var baseUrl string = "https://viacep.com.br/ws/"

const ServiceName string = "Via CEP"

func GetAddress(ctx context.Context, cep string) (*GetAddressResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s/json", baseUrl, cep), nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var address GetAddressResponse
	if err := json.Unmarshal(resBody, &address); err != nil {
		return nil, err
	}

	if address.Cep == "" {
		return nil, nil
	}

	return &address, nil
}

type GetAddressResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}
