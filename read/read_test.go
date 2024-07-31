package main

import (
    "os"
    "path/filepath"
    "testing"
)

func TestMain(t *testing.T) {
    // Get current working directory
    cwd, err := os.Getwd()
    if err != nil {
        t.Fatalf("failed to get current working directory: %s", err)
    }
    t.Logf("Current working directory: %s", cwd)

    // Define the file path
    filePath := filepath.Join(cwd, "..", "output1.txt")
    t.Logf("Checking file at path: %s", filePath)

    // Check if the file exists
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        t.Fatalf("file does not exist: %s", filePath)
    }

    // Read the file content
    data, err := os.ReadFile(filePath)
    if err != nil {
        t.Fatalf("failed reading file: %s", err)
    }

    t.Log("Contents of output.txt ", string(data))
}