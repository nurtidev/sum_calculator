package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Object struct {
	A int `json:"a"`
	B int `json:"b"`
}

const MaxGoroutines = 1000

func generateJSONFile(filePath string, numObjects int) error {
	rand.Seed(time.Now().UnixNano())

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString("[\n")
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(writer)

	for i := 0; i < numObjects; i++ {
		obj := Object{
			A: rand.Intn(21) - 10,
			B: rand.Intn(21) - 10,
		}

		if i > 0 {
			_, err = writer.WriteString(",\n")
			if err != nil {
				return err
			}
		}

		if err := encoder.Encode(&obj); err != nil {
			return fmt.Errorf("failed to encode object: %v", err)
		}
	}

	_, err = writer.WriteString("]\n")
	if err != nil {
		return err
	}

	return writer.Flush()
}

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s <file-path> <num-workers> <generate>", os.Args[0])
	}

	filePath := os.Args[1]
	numWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid number of workers: %v", err)
	}
	generate := os.Args[3] == "true"

	if generate {
		fmt.Println("Generating JSON file...")
		if err := generateJSONFile(filePath, 1000000); err != nil {
			log.Fatalf("Error generating file: %v", err)
		}
		fmt.Println("JSON file generated successfully.")
	}

	var memStatsBefore runtime.MemStats
	runtime.ReadMemStats(&memStatsBefore)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	objectChan := make(chan Object, 1000)
	var totalSum int64
	var mutex sync.Mutex

	numWorkers = min(numWorkers, MaxGoroutines)
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			localSum := int64(0)
			for obj := range objectChan {
				localSum += int64(obj.A + obj.B)
			}
			mutex.Lock()
			totalSum += localSum
			mutex.Unlock()
		}()
	}

	fmt.Println("Processing JSON file...")
	decoder := json.NewDecoder(bufio.NewReader(file))

	token, err := decoder.Token()
	if err != nil || token != json.Delim('[') {
		log.Fatalf("Expected start of array '['")
	}

	for decoder.More() {
		var obj Object
		if err := decoder.Decode(&obj); err != nil {
			log.Fatalf("Failed to decode object: %v", err)
		}
		objectChan <- obj
	}

	close(objectChan)

	wg.Wait()

	// Collect memory statistics
	var memStatsAfter runtime.MemStats
	runtime.ReadMemStats(&memStatsAfter)

	fmt.Printf("Total sum: %d\n", totalSum)
	fmt.Printf("Memory used: %.2f MB\n", float64(memStatsAfter.Alloc-memStatsBefore.Alloc)/1024/1024)
	fmt.Printf("Number of allocations: %d\n", memStatsAfter.Mallocs-memStatsBefore.Mallocs)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
