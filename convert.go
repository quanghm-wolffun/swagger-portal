package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Get the directory of the current executable
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		os.Exit(1)
	}

	yamlFile := filepath.Join(dir, "docs", "swagger.yaml")
	jsonFile := filepath.Join(dir, "docs", "swagger.json")

	// Check if the YAML file exists
	if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
		fmt.Printf("YAML file not found: %s\n", yamlFile)
		os.Exit(1)
	}

	// Convert YAML to JSON using yq
	cmd := exec.Command("yq", "-o=json", ".", yamlFile)
	jsonOutput, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing yq command: %v\n", err)
		os.Exit(1)
	}

	// Validate the JSON by parsing it
	var jsonObj interface{}
	if err := json.Unmarshal(jsonOutput, &jsonObj); err != nil {
		fmt.Printf("Invalid JSON output from yq: %v\n", err)
		os.Exit(1)
	}

	// Pretty print the JSON for better readability
	prettyJSON, err := json.MarshalIndent(jsonObj, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting JSON: %v\n", err)
		os.Exit(1)
	}

	// Write the formatted JSON output to file
	err = os.WriteFile(jsonFile, prettyJSON, 0644)
	if err != nil {
		fmt.Printf("Error writing JSON file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully converted %s to %s\n", yamlFile, jsonFile)
}
