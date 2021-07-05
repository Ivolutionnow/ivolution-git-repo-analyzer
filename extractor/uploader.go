package extractor

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// Upload result to Ivolution
func Upload(path, repoName string) (string, error) {

	url := "https://api.ivolution.ai/git-analyzer"

	// Read file
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Add file as multipart/form-data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return "", err
	}
	io.Copy(part, file)
	writer.Close()

	// Create and make the request
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	if response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("Server returned non 200 response. Error code: %d", response.StatusCode))
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// Get response and return resulting token
	var result uploadResponse
	err = json.Unmarshal(content, &result)
	if err != nil {
		return "", err
	}

	processURL := fmt.Sprintf("https://api.ivolution.ai/git-analyzer-verification?token=%s&git-repository-name=%s", result.Token, repoName)
	return processURL, nil
}

type uploadResponse struct {
	Token string `json:"token"`
}
