package facts

import (
	"encoding/json"
	"fmt"
	"time"

	// "fmt"
	"io"
	"os"
	// "time"
)

type Fact struct {
	ID   int    `json:"id"`
	Fact string `json:"fact"`
}

var facts []Fact

func LoadFacts() error {
	file, err := os.Open("./facts/facts.json")
	if err != nil {
		return nil
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &facts)
	if err != nil{
		return err
	}

	return nil
}

func GetFacts() ([]Fact, error) {
	startTime := time.Now()
	fmt.Println(time.Since(startTime))
	
	return facts, nil
}

func GetFactByID(id int) (*Fact, error) {
	for _, fact := range facts {
		if fact.ID == id {
			return &fact, nil
		}
	}

	return nil, fmt.Errorf("fact with ID %d not found", id)
}
