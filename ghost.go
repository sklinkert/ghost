package ghost

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gbrlsnchs/jwt/v3"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"strings"
	"time"
)

type Ghost struct {
	adminAPIToken      string
	contentAPIToken    string
	jwtToken           string
	jwtTokenExpiration time.Time
	url                string
}

// New creates new instance of ghost API client
func New(url, contentAPIToken, adminAPIToken string) (*Ghost, error) {
	return &Ghost{
		adminAPIToken:   adminAPIToken,
		contentAPIToken: contentAPIToken,
		url:             url,
	}, nil
}

func (g *Ghost) checkAndRenewJWT() error {
	if g.jwtToken == "" || g.jwtTokenExpiration.Before(time.Now()) {
		jwtToken, jwtTokenExpiration, err := generateJWT(g.adminAPIToken)
		if err != nil {
			return err
		}
		g.jwtToken = jwtToken
		g.jwtTokenExpiration = jwtTokenExpiration
	}
	return nil
}

func generateJWT(keyID string) (string, time.Time, error) {
	var now = time.Now()
	var expiration = now.Add(5 * time.Minute)

	keyParts := strings.Split(keyID, ":")
	if len(keyParts) != 2 {
		return "", expiration, fmt.Errorf("invalid Client.Key format")
	}
	id := keyParts[0]
	rawSecret := []byte(keyParts[1])
	secret := make([]byte, hex.DecodedLen(len(rawSecret)))
	_, err := hex.Decode(secret, rawSecret)
	if err != nil {
		return "", expiration, err
	}

	hs256 := jwt.NewHS256(secret)
	p := jwt.Payload{
		Audience:       jwt.Audience{"/v3/admin/"},
		ExpirationTime: jwt.NumericDate(expiration),
		IssuedAt:       jwt.NumericDate(now),
	}
	token, err := jwt.Sign(p, hs256, jwt.KeyID(id))
	if err != nil {
		return "", expiration, err
	}
	return string(token), expiration, nil
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func (g *Ghost) getJson(url string, target interface{}) error {
	if err := g.checkAndRenewJWT(); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Ghost"+" "+g.jwtToken)
	resp, err := myClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Cannot close body reader: %v\n", err)
		}
	}()

	content, _ := ioutil.ReadAll(resp.Body)
	responseBody := string(content[:])
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, responseBody)
	}

	err = json.Unmarshal(content, target)
	if err != nil {
		return err
	}
	return nil
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
	fileContents, err := ioutil.ReadAll(file)
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

	content, err := ioutil.ReadAll(response.Body)
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
