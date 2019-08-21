package clientloop

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getNewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 5,
	}
}

func (l *clientLooper) postHTTP(path string, JSON interface{}) error {
	url := l.serverLocation.String() + "/api/" + path

	body, err := json.Marshal(JSON)
	if err != nil {
		return fmt.Errorf("Unable to marshal body -- %s", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Phoenix-Id", l.phoenixID)

	resp, err := l.http.Do(req)
	if err != nil {
		return fmt.Errorf("Unable to send body to server -- %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return nil
	}

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("Unable to read resp body -- %s", err)
	}

	return fmt.Errorf("Non-200 response -- (%d)%s", resp.StatusCode, respBody)

}
