package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

type Request struct {
	Url     *string
	Method  string
	Data    *map[string]string
	Headers *map[string]string
	RepInfo map[string]interface{}
}

func (r *Request) Body() {
	data, err := json.Marshal(r.Data)
	checkError(err)
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, *r.Url, bytes.NewBuffer(data))
	checkError(err)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	for k, v := range *r.Headers {
		req.Header.Add(k, v)
	}
	rep, err := client.Do(req)
	checkError(err)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(rep.Body)
	body, err := ioutil.ReadAll(rep.Body)
	checkError(err)
	if err = json.Unmarshal(body, &r.RepInfo); err != nil {
		fmt.Println(err)
	}
}

func (r *Request) Query() {
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, *r.Url, nil)
	checkError(err)
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	for k, v := range *r.Headers {
		req.Header.Add(k, v)
	}
	query := req.URL.Query()
	for k, v := range *r.Data {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()
	rep, err := client.Do(req)
	checkError(err)
	_ = rep.Body.Close()
	body, err := ioutil.ReadAll(rep.Body)
	checkError(err)
	err = json.Unmarshal(body, &r.RepInfo)
	checkError(err)
}
