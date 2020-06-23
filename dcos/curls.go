package dcos

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"

	"github.com/savaki/jq"
	// you can remove but be sure to remove where it is used as well below.
)

// Payload used to login to the Injector service
type Payload struct {
	UID      string `json:"uid"`
	Password string `json:"password"`
}

// InjectorEntry is used to add/update items within Injector
type InjectorEntry struct {
	keyName  string
	value    string
	secret   bool
	path     string
	function string
}

// CurlItOne - test method is not longer needed.
func CurlItOne(masterURL string) {

	reader := strings.NewReader(`{"uid":"gsil-readonly","password":"Gmf5ZF7k9LYJTmpwKMeG"}`)
	request, err := http.NewRequest("POST", "http://gdmst001v.gsil.rri-usa.org/injector/acs/api/v1/auth/login", reader)

	if err != nil {
		fmt.Println("Could not create request - ", err)
	}

	// TODO: check err
	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		fmt.Println("Could not POST request", err)
	}

	responseBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("\n%v\n\n", string(responseBody))
}

// Login - is ued to, surprise! Login to the injector service and obtain a token
func Login(masterURL string, username string, password string) []byte {

	var payload Payload
	payload.UID = username
	payload.Password = password

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Could not marshall json - ", err)
	}

	fmt.Println("payload : ")
	fmt.Println(payloadBytes)

	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", masterURL+"acs/api/v1/auth/login", body)

	if err != nil {
		fmt.Println("Could not POST req", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Could not handle response - ", err)
	}

	defer resp.Body.Close()

	body2, _ := ioutil.ReadAll(resp.Body)
	token := extractByName(body2, ".token")

	return token
}

// IntentManagerURL - used to grab the URL of the intent manager from the Injector service.
func IntentManagerURL(masterURL string, token string) []byte {

	req, err := http.NewRequest("GET", masterURL+"injector/v1/config/intent_manager_url", nil)
	if err != nil {
		fmt.Println("Could not crate request : ", err)
	}

	token = removeQuotes(token)
	req.Header.Set("Authorization", "Bearer "+token) //os.ExpandEnv("Bearer $JWT"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Could not get proper response : ", err)
		//return ""
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

// RegisterKey with injector
func RegisterKey(masterURL string, token string, keyName string) []byte {
	body := strings.NewReader(`{"name": "` + keyName + `", "schema":{}}`)
	req, err := http.NewRequest("PUT", masterURL+"injector/v1/keys/"+keyName+"}", body)
	if err != nil {
		// handle err
	}

	token = removeQuotes(token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	body2, _ := ioutil.ReadAll(resp.Body)
	return body2
}

// SetInjectorEntry with injector
func SetInjectorEntry(masterURL string, token string, entry InjectorEntry) []byte {

	body := strings.NewReader(`{"value": "` + entry.value + `", "secret": ` + strconv.FormatBool(entry.secret) + `}`)
	//fmt.Println("body : " + body.)

	endpointURL := masterURL + "injector/v1/config" + entry.path + "/" + entry.keyName
	fmt.Println("url : " + endpointURL)

	req, err := http.NewRequest(entry.function, endpointURL, body)
	if err != nil {
		fmt.Println("Could not create new http request : ", err)
	}

	token = removeQuotes(token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	log.Println("[DEBUG] - setInjectEntry : " + string(requestDump))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	body2, _ := ioutil.ReadAll(resp.Body)
	return body2
}

func extractByName(input []byte, name string) []byte {
	op, err := jq.Parse(name) //.foo | ..")

	if err != nil {
		fmt.Println("Could not create query", err)
	}

	data, _ := op.Apply(input)

	return data

}

func removeQuotes(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}
