package openai

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type OpenAI struct {
	Organization string
}

func (openai *OpenAI) CallCompletions(completionsAPI string, requestBody RequestBody) string {
	data, err := json.Marshal(requestBody)
	if err != nil {
		return err.Error()
	}

	req, err := http.NewRequest("POST", completionsAPI, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	req.Header.Set("Content-Type", "application/data")
	req.Header.Set("Authorization", "Bearer "+openai.Organization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	var response ResponseBody
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	return response.Choices[0].Message.Content
}
