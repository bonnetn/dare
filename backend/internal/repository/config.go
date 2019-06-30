package repository

import (
	"github.com/bonnetn/dare/backend/internal/entity"
	"os"
	"encoding/json"
	"fmt"
)

func LoadConfiguration(path string) (entity.Configuration, error) {
	file, err := os.Open(path)
	if err != nil {
		return entity.Configuration{}, fmt.Errorf("could not load configuration: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := entity.Configuration{}
	err = decoder.Decode(&configuration)
	return configuration, err
}
