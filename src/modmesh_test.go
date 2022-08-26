package main

import "fmt"
import "testing"

func TestMain(m *testing.M) {
    testVar := 1
    if testVar != 1 {
        fmt.Errorf("Problem.")
    }
}
