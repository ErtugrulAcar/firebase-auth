package handler

import (
	"FirebaseAuth/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	newUserUrl = "https://identitytoolkit.googleapis.com/v1/accounts:signUp?key=AIzaSyDJT_inXLYRMxjwBFtrWzC3ctMPFPnKOZk"
	signInUrl = "https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=AIzaSyDJT_inXLYRMxjwBFtrWzC3ctMPFPnKOZk"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var user entity.User

	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("email", user.Email)
	fmt.Println("password", user.Password)
	jsonizedUser, _ := json.Marshal(user)

	//properJson := []byte(`{"email":"denemsadex@denemex.com", "password" : "1as23d561ads1"}`)

	req, err := http.NewRequest("POST", newUserUrl, bytes.NewBuffer(jsonizedUser))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	if resp.Status == "200 OK" {
		w.WriteHeader(http.StatusOK)

		_, _ = fmt.Fprintf(w, "Success")
	}else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "Wrong")
	}


}

func SignIn(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var objmap map[string]json.RawMessage
	_ = decoder.Decode(&objmap)
	fmt.Println(string(objmap["email"]))
	fmt.Println(string(objmap["password"]))
	fmt.Println(string(objmap["returnSecureToken"]))

	jsonizedObj, _ := json.Marshal(objmap)

	req, err := http.NewRequest("POST", signInUrl, bytes.NewBuffer(jsonizedObj))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	respDecoder := json.NewDecoder(resp.Body)
	var respMap map[string]json.RawMessage
	_ = respDecoder.Decode(&respMap)
	fmt.Println(respMap["idToken"])


	if resp.Status == "200 OK" {
		w.WriteHeader(http.StatusOK)

		_, _ = fmt.Fprintf(w, "Success")
	}else {
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintf(w, "Wrong")
	}
}