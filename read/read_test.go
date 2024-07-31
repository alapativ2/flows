package main

import (
    "os"
    "path/filepath"
    "testing"

    "github.com/stretchr/testify/assert" // Ensure you have this package in your go.mod
)

func TestMain(t *testing.T) {
    assert := assert.New(t)

    // Get current working directory
    cwd, err := os.Getwd()
    assert.NoError(err, "failed to get current working directory")
    t.Logf("Current working directory: %s", cwd)

    // Define the file path
    filePath := filepath.Join(cwd, "..", "output1.txt")
    t.Logf("Checking file at path: %s", filePath)

    // Check if the file exists
    _, err = os.Stat(filePath)
    assert.NoError(err, "file path in  directory doest exist")
    if os.IsNotExist(err) {
        t.Fatalf("file does not exist: %s", filePath)
    }
    assert.NoError(err, "error checking file")

    // Read the file content
    data, err := os.ReadFile(filePath)
    assert.NoError(err, "failed reading file")

    t.Log("Contents of output.txt ", string(data))
}