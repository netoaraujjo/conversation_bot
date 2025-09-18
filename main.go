package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"sync"
	"time"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	cpf := "07958758359"
	var wg sync.WaitGroup
	wg.Add(3)
	start := time.Now()
	go buscaProcedimentos(cpf, &wg)
	go buscaIBIOSEG(cpf, &wg)
	go buscaRENACH(cpf, &wg)
	wg.Wait()
	elapsed := time.Since(start)
	// fmt.Println(math.Ceil(elapsed.Seconds()))
	fmt.Printf("%v segundos\n", elapsed.Seconds())
}

func buscaProcedimentos(cpf string, wg *sync.WaitGroup) {
	fmt.Println("Iniciando busca de antecedentes")
	baseURL := "https://lupa.ssp.pi.gov.br/api/v1/"

	req, err := http.NewRequest("GET", baseURL+"procedimentos-policiais/suposto-autor-infrator/", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("cpf", cpf)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", "Api-Key " + os.Getenv("LUPA_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	defer wg.Done()

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(result))
	fmt.Println("Finalizada busca de antecedentes")
}

func buscaIBIOSEG(cpf string, wg *sync.WaitGroup) {
	fmt.Println("Iniciando busca no IBIOSEG")
	baseURL := "https://lupa.ssp.pi.gov.br/api/v1/"

	req, err := http.NewRequest("GET", baseURL+"ibioseg/pessoa/", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("cpf", cpf)
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", "Api-Key 4bcrvRhr.Gw6FKS6hZAhyuboVyK7pK6xFsbh5JUNb")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	defer wg.Done()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(result))
	fmt.Println("Finalizada busca no IBIOSEG")
}

func buscaRENACH(cpf string, wg *sync.WaitGroup) {
	fmt.Println("Iniciando busca no RENACH")
	baseURL := "https://lupa.ssp.pi.gov.br/api/v1/"

	req, err := http.NewRequest("GET", baseURL+"detran/veiculo/renach/", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("tipo", "CPF")
	q.Add("valor", cpf)

	req.URL.RawQuery = q.Encode()

	req.Header.Set("Authorization", "Api-Key 4bcrvRhr.Gw6FKS6hZAhyuboVyK7pK6xFsbh5JUNb")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	defer wg.Done()
	_, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(result))
	fmt.Println("Finalizada busca no RENACH")
}
