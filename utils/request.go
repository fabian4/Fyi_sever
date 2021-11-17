package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Post(url string, data string, token string) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, strings.NewReader(data))
	req.Header.Set("Authorization", "Bearer "+token)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
}
