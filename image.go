package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strings"
)

type ImageResponse struct {
	Images []Image
}
type Image struct {
	URL string
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

// AdminUploadImage - Creates a new file upload http request
func (g *Ghost) AdminUploadImage(path string) (imageURL string, err error) {
	const paramName = "file"
	var uri = fmt.Sprintf("%s/ghost/api/v3/admin/images/upload/", g.url)

	params := map[string]string{
		"purpose": "image",
		"ref":     path,
	}
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	fileContents, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	fi, err := file.Stat()
	if err != nil {
		return "", err
	}
	if err := file.Close(); err != nil {
		return "", err
	}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Type", "image/jpeg")
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
		escapeQuotes(paramName), escapeQuotes(fi.Name())))
	part, err := writer.CreatePart(h)
	if err != nil {
		return "", err
	}
	_, err = part.Write(fileContents)
	if err != nil {
		return "", err
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}

	err = writer.Close()
	if err != nil {
		return "", err
	}

	if err := g.checkAndRenewJWT(); err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Ghost"+" "+g.jwtToken)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			fmt.Printf("cannot close body reader: %v\n", err)
		}
	}()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBody := string(content[:])
	if response.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("wrong http code: %d (%q)",
			response.StatusCode, responseBody)
	}

	var images ImageResponse
	if err := json.Unmarshal(content[:], &images); err != nil {
		return "", err
	}

	return images.Images[0].URL, nil
}
