package send

import (
	"bytes"
	"github.com/clong1995/go-config"
	"io"
	"log"
	"net/http"
)

var client *http.Client

func init() {
	client = &http.Client{}
}

func Send(url string, data []byte, headers map[string]string) (res []byte, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+config.Value("QWEN_KEY"))
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if res, err = io.ReadAll(resp.Body); err != nil {
		log.Println(err)
		return
	}
	return
}
