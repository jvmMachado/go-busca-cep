package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Address struct {
	Cep string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf string `json:"uf"`
	Ibge string `json:"ibge"`
	Gia string `json:"gia"`
	Ddd string `json:"ddd"`
	Siafi string `json:"siafi"`
}

func main() {
		input := os.Args[1:]
		cep := input[0]
		endpoint := "http://viacep.com.br/ws/" + cep + "/json/"

		res, err := http.Get(endpoint)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer a requisição: %v\n", err)
		}
		defer res.Body.Close()

		data, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler o body da requisição: %v\n", err)
		}

		var address Address

		err = json.Unmarshal(data, &address)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer parse do json: %v\n", err)
		}

		fmt.Printf("CEP: %s\nLogradouro: %s\nBairro: %s\nLocalidade: %s\nUF: %s\n", address.Cep, address.Logradouro, address.Bairro, address.Localidade, address.Uf)
	}