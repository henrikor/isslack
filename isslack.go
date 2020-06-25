package isslack

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// isslack takes two arguments (service, and command)
// Example:
// isslack("ismetering", "install")
// func postSlack(msg string, surl string) string {
func Isslack(msg string, surl string) string {
	// surl := "https://hooks.slack.com/services/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"
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
