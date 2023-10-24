package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func main() {
	// Read the config file
	configFile, err := os.Open("config.txt")
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	// Extract the file paths from the config
	var todoFiles []string
	scanner := bufio.NewScanner(configFile)
	for scanner.Scan() {
		todoFiles = append(todoFiles, expandTilde(scanner.Text())) // Expand ~ to home directory
	}

	type item struct {
		line     string
		duration time.Duration
		file     string
	}

	// Iterate through the files
	var items []item
	for _, todoFile := range todoFiles {
		// Open the file
		file, err := os.Open(todoFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Read the file contents
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			// Check if the line has a priority of (W)
			if strings.HasPrefix(line, "(W)") {
				// Extract the creation date
				fields := strings.Fields(line)
				creationDate, err := time.Parse("2006-01-02", fields[1])
				if err != nil {
					panic(err)
				}

				// Calculate the duration since creation
				duration := time.Since(creationDate)

				// Store the waiting item and its duration
				items = append(items, item{line, duration, "+" + strings.TrimSuffix(filepath.Base(file.Name()), ".todo.txt")})
			}
		}
	}

	// Sort the waiting items by duration
	sort.Slice(items, func(i, j int) bool {
		return items[i].duration > items[j].duration
	})

	for _, item := range items {
		duration := timeToDays(item.duration)
		fmt.Printf("%s (%0.0f days, %s)\n", item.line, duration, item.file)
	}
}

func expandTilde(path string) string {
	if strings.HasPrefix(path, "~") {
		homeDir, _ := os.UserHomeDir()
		return strings.Replace(path, "~", homeDir, 1)
	}
	return path
}

func timeToDays(duration time.Duration) float64 {
	return float64(duration) / float64(time.Hour*24)
}
