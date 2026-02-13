package utils

import (
	"encoding/json"
	"log"
)

func PreparePayloadToStruct[T any](payload interface{}, data *T) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshaling payload: %v", err)
		return err
	}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Printf("Error unmarshaling to AnswerPayload: %v", err)
		return err
	}
	return nil
}
