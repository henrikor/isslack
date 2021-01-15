package isslack

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// Isslack takes two arguments (service, and command)
// Example:
// Isslack("ismetering", "install")
// func postSlack(msg string, surl string) string {
func Isslack(msg string, surl string) string {
	message := map[string]interface{}{
		"text": msg,
	}
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post(surl, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return err.Error()
	} else {
		return resp.Status
	}
}
