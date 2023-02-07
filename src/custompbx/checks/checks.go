package checks

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type jsonRequest struct {
	Token string `json:"token"`
	Nonce string `json:"nonce"`
}

type ApiData struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (a *ApiData) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ApiChecks(w http.ResponseWriter, r *http.Request) ([]byte, bool) {
	var jsonRequest jsonRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		SendResponce(w, ApiData{3, "Cant read data", ""})
		return nil, false
	}
	err = json.Unmarshal([]byte(body), &jsonRequest)
	if err != nil {
		SendResponce(w, ApiData{3, "Please use JSON data", ""})
		return nil, false
	}

	return body, true
}

func CheckDomainName(name string) bool {
	match, err := regexp.MatchString(`^[a-z0-9]{3,26}\.[a-z]{2,4}$`, name)
	if err != nil {
		log.Print("debug 7")
		return false
	}
	return match
}

func SendResponce(w http.ResponseWriter, data ApiData) {
	w.Header().Set("Content-Type", "application/json")
	jresponse, err := json.Marshal(data)
	log.Print(string(jresponse))
	if err != nil {
		log.Print("cant marshal to send")
		w.Write([]byte(`{Status: 3, Msg: "Cant create responce", Data: err}`))
		return
	}
	w.Write(jresponse)
}
