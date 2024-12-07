package input

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const baseURL = "https://adventofcode.com"

// Read reads the content of the input file for the given year and day.
// If the file does not exist, it fetches the input from the Advent of Code website and saves it.
func Read(year, day int) string {
	filename := fmt.Sprintf("inputs/%d/day_%02d.txt", year, day)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		input := fetchAndSaveInput(year, day, filename)
		return input
	}
	return strings.TrimSpace(string(data))
}

// ReadLines returns a channel that streams lines from the input file for the given year and day.
func ReadLines(year, day int) <-chan string {
	filename := fmt.Sprintf("inputs/%d/day_%02d.txt", year, day)
	lines := make(chan string, 100) // Buffered channel for better performance

	go func() {
		defer close(lines)
		file, err := os.Open(filename)
		if err != nil {
			log.Fatalf("Failed to open file: %v", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
	}()

	return lines
}

// fetchAndSaveInput fetches the input from the Advent of Code website and saves it locally.
func fetchAndSaveInput(year, day int, filename string) string {
	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the session cookie from the environment variables
	sessionCookie := os.Getenv("SESSION_COOKIE")
	if sessionCookie == "" {
		log.Fatalf("SESSION_COOKIE is not set in the environment variables")
	}

	// Construct the URL for the input
	url := fmt.Sprintf("%s/%d/day/%d/input", baseURL, year, day)

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Add the session cookie for authentication
	req.AddCookie(&http.Cookie{Name: "session", Value: sessionCookie})

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to fetch input: %v", err)
	}
	defer resp.Body.Close()

	// Check for errors in the response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch input: HTTP %d", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	input := string(body)

	// Save the input to the local file
	saveInputToFile(input, filename)

	return strings.TrimSpace(input)
}

// saveInputToFile saves the input to the specified file.
func saveInputToFile(input, filename string) {
	// Ensure the directory exists
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	// Write the input to the file
	if err := ioutil.WriteFile(filename, []byte(input), 0644); err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}
