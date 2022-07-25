package telbiz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func authToken() *telBizResOAuth {
	client := &http.Client{}
	data := telBizReqOAuth{
		ClientID:  os.Getenv("TELBIZ_CLIENT_ID"),
		Secret:    os.Getenv("TELBIZ_CLIENT_SECRET"),
		GrantType: "client_credentials",
		Scope:     "Telbiz_API_SCOPE profile openid",
	}
	jsonReq, err := json.Marshal(data)
	req, err := http.NewRequest(
		"POST",
		os.Getenv("TELBIZ_BASE_URI")+"connect/token",
		bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Print(err.Error())
		}
	}(resp.Body)
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject telBizResOAuth
	json.Unmarshal(bodyBytes, &responseObject)
	//fmt.Printf("API Response as struct %+v\n", responseObject)
	return &responseObject
}
