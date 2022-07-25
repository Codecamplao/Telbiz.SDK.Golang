package telbiz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("please consider environment variables: %s\n", err)
	}
}

func SMSService(title Title, phone string, message string) *TelBizResSMS {
	var responseObject TelBizResSMS
	responseObject.Response.Success = false
	responseObject.Response.Message = "SMS Failed"
	reqTitle := "Telbiz"
	switch title {
		case Default:
		case News:
		case Promotion:
		case OTP:
		case Info:
		case Unknown:
			reqTitle = string(title)
			break
		default:
			responseObject.Response.Message = "Invalid sms title"
			return &responseObject
	}
	// Call Token service
	token := authToken()
	// Call SMS service
	client := &http.Client{}
	data := TelBizReqSMS{
		Title:   reqTitle,
		Phone:   phone,
		Message: message,
	}
	jsonReq, err := json.Marshal(data)
	req, err := http.NewRequest(
		"POST",
		os.Getenv("TELBIZ_BASE_URI")+"smsservice/newtransaction",
		bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Print(err.Error())
		responseObject.Response.Message = "SMS Request error"
		responseObject.Response.Detail = err.Error()
		return &responseObject
	}
	req.Header.Add("Authorization", "Bearer "+token.AccessToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		responseObject.Response.Message = "Client Request error"
		responseObject.Response.Detail = err.Error()
		return &responseObject
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
		responseObject.Response.Message = "SMS Read Bytes error"
		responseObject.Response.Detail = err.Error()
		return &responseObject
	}

	json.Unmarshal(bodyBytes, &responseObject)
	return &responseObject
}

func TopUpService() {
	fmt.Printf("TopUp Service In Developing\n")
}
