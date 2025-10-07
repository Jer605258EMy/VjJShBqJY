// 代码生成时间: 2025-10-07 18:29:40
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// TestData is a simple struct to hold test data.
type TestData struct {
    ID       int    "json:"id""
    Name     string "json:"name""
    Email    string "json:"email""
    Password string "json:"password""
}

// NewTestData generates a new TestData instance with random data.
func NewTestData() *TestData {
    rand.Seed(time.Now().UnixNano())
    data := &TestData{
        ID:       rand.Intn(1000),
        Name:     fmt.Sprintf("Name%d", rand.Intn(100)),
        Email:    fmt.Sprintf("%d@example.com", rand.Intn(100)),
        Password: fmt.Sprintf("password%d", rand.Intn(100)),
    }
    return data
}

// GenerateTestData creates a slice of TestData with the specified number of entries.
func GenerateTestData(count int) ([]*TestData, error) {
    if count <= 0 {
        return nil, fmt.Errorf("count must be greater than 0")
    }

    var testData []*TestData
    for i := 0; i < count; i++ {
        data := NewTestData()
        testData = append(testData, data)
    }
    return testData, nil
}

func main() {
    // Generate 10 test data entries.
    testData, err := GenerateTestData(10)
    if err != nil {
        fmt.Println("Error generating test data: ", err)
        return
    }

    // Print out the generated test data.
    for _, data := range testData {
        fmt.Printf("ID: %d, Name: %s, Email: %s, Password: %s
", data.ID, data.Name, data.Email, data.Password)
    }
}