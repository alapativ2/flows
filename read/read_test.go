package main

import (
    "os"
    "testing"
)

func TestMain(t *testing.T) {
    filePath := "../output1.txt"

    // Check if the file exists
    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        t.Fatalf("file does not exist: %s", filePath)
    }

    data, err := os.ReadFile(filePath)
    if err != nil {
        t.Fatalf("failed reading file: %s", err)
    }

    t.Log("Contents of output.txt ", string(data))
}