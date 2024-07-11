package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"golang.org/x/exp/rand"
)

func main() {
	var choice int
	fmt.Println("Pilih jenis operasi:")
	fmt.Println("1. UpdateStock")
	fmt.Println("2. UpdateStock with WG")
	fmt.Println("3. UpdateStock With Mutex")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		RunUpdateStock()
	case 2:
		RunUpdateStockWithWG()
	case 3:
		RunUpdateStockWithMutex()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func RunUpdateStock() {
	url := "http://localhost:8080/update-stock"
	numRequests := 5 // Ubah sesuai dengan jumlah request yang ingin Anda kirimkan

	var wg sync.WaitGroup
	wg.Add(numRequests)

	for i := 0; i < numRequests; i++ {
		go func(requestNumber int) {
			defer wg.Done()

			data := rand.Intn(20)
			updateData := map[string]int{"amount": data} // Gunakan nilai random untuk amount
			jsonData, _ := json.Marshal(updateData)

			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("Error making request:", err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				return
			}

			fmt.Printf("Response from %s for request #%d:\n", url, requestNumber)
			fmt.Println(string(body))
		}(i)
	}

	wg.Wait()
	fmt.Println("All requests completed")
}

func RunUpdateStockWithWG() {
	url := "http://localhost:8080/update-stock-wg"
	numRequests := 5 // Ubah sesuai dengan jumlah request yang ingin Anda kirimkan

	var wg sync.WaitGroup
	wg.Add(numRequests)

	for i := 0; i < numRequests; i++ {
		go func(requestNumber int) {
			defer wg.Done()

			data := rand.Intn(20)
			updateData := map[string]int{"amount": data} // Gunakan nilai random untuk amount
			jsonData, _ := json.Marshal(updateData)

			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("Error making request:", err)
				return
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				return
			}

			fmt.Printf("Response from %s for request #%d:\n", url, requestNumber)
			fmt.Println(string(body))
		}(i)
	}

	wg.Wait()
	fmt.Println("All requests completed")
}

func RunUpdateStockWithMutex() {
	url := "http://localhost:8080/update-stock-mutex"
	numRequests := 5 // Change it according to the number of requests you want to send

	var wg sync.WaitGroup
	wg.Add(numRequests)

	for i := 0; i < numRequests; i++ {
		go func(requestNumber int) {
			defer wg.Done()
			data := rand.Intn(20)
			updateData := map[string]int{"amount": data} // Example of updated data, adjust to needs
			jsonData, _ := json.Marshal(updateData)

			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("Error making request:", err)
				return
			}

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				return
			}

			fmt.Printf("Response from %s for request #%d:\n", url, requestNumber)
			fmt.Println(string(body))
		}(i)
	}

	wg.Wait()
	fmt.Println("All requests completed")
}
