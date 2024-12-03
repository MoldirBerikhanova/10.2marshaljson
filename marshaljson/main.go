package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Course struct {
	Name           string `json:"name"`
	Studentsnumber int    `json:"studentsnumber"`
}

func main() {
	p := Course{Name: "GO: Путь Джуна", Studentsnumber: 10}

	p.Name = "GO: Путь Джуна Backend Developer"
	p.Studentsnumber = 15
	fmt.Println(p)

	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Couldnt open", err)
		return
	}

	filename := "example.json"
	err = os.WriteFile(filename, data, 0644) //////////чтение Json файла
	if err != nil {
		fmt.Println("error openin file", err)
	}
	fmt.Printf("Data written to %s successfully\n", filename)
}

// func TestAssertJSON(t *testing.T) {
// 	// Valid JSON example
// 	validJSON := `{"name": "John", "age": 30}`
// 	assert.True(t, IsValidJSON(validJSON), "The string should be valid JSON")

// 	// Invalid JSON example
// 	invalidJSON := `{"name": "John", "age": }`
// 	assert.False(t, IsValidJSON(invalidJSON), "The string should not be valid JSON")
// }
