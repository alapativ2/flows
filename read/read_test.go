package main

import (
    "testing"
    "io/ioutil"
    "log"
)

func TestMain(t *testing.T) {
    data, err := ioutil.ReadFile("../output1.txt")
    if err != nil {
        log.Fatalf("failed reading file: %s", err)
    }
    t.Log("Contents of output.txt ", string(data))
}