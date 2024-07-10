package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	var choice int
	fmt.Println("Pilih jenis operasi:")
	fmt.Println("1. Without Mutex (Race Condition)")
	fmt.Println("2. With Mutex")
	fmt.Print("Pilihan: ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		RunWithoutMutex()
	case 2:
		RunWithMutex()
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func RunWithoutMutex() {
	url := "http://localhost:8080/update-stock-wg"
	numRequests := 5 // Ubah sesuai dengan jumlah request yang ingin Anda kirimkan

	var wg sync.WaitGroup
	wg.Add(numRequests)

	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()

			updateData := map[string]int{"amount": 10} // Contoh data update, sesuaikan dengan kebutuhan
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

			fmt.Printf("Response from %s:\n", url)
			fmt.Println(string(body))
		}()
	}

	wg.Wait()
	fmt.Println("All requests completed")
}

func RunWithMutex() {
	url := "http://localhost:8080/update-stock-mutex"
	numRequests := 5 // Ubah sesuai dengan jumlah request yang ingin Anda kirimkan

	var wg sync.WaitGroup
	wg.Add(numRequests)

	var mu sync.Mutex

	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()

			updateData := map[string]int{"amount": 10} // Contoh data update, sesuaikan dengan kebutuhan
			jsonData, _ := json.Marshal(updateData)

			mu.Lock()
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				fmt.Println("Error making request:", err)
				mu.Unlock()
				return
			}
			mu.Unlock()

			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response body:", err)
				return
			}

			fmt.Printf("Response from %s:\n", url)
			fmt.Println(string(body))
		}()
	}

	wg.Wait()
	fmt.Println("All requests completed")
}
