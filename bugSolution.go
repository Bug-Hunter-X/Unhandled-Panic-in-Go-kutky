package main

import (

        "errors"
        "fmt"
)

func calculate(a, b int) (int, error) {
        if b == 0 {
                return 0, errors.New("division by zero")
        }
        return a / b, nil
}

func main() {
        defer func() {
                if r := recover(); r != nil {
                        fmt.Println("Recovered from panic:", r)
                }
        }()
        result, err := calculate(10, 0)
        if err != nil {
                fmt.Println("Error:", err)
        } else {
                fmt.Println("Result:", result)
        }
} 