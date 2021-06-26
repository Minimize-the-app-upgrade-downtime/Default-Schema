package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

//User defines model for storing account details in database
type db_field struct {
	
	// office
	OfficeCode 		string `json:"officeCode,omitempty"`
	OfficeName 		string `json:"officeName,omitempty"`
	City 			string `json:"city,omitempty"`
	AddressLine1  	string `json:"addressLine1,omitempty"`
	Address 		string `json:"address,omitempty"`
	AddressLine2 	string `json:"addressLine2,omitempty"`
	PostalCode 		string `json:"postalCode,omitempty"`
	Street 			string `json:"street,omitempty"`
	Country 		string `json:"country,omitempty"`

	// employee
	EmployeeNumber  int    `json:"employeeNumber,omitempty",`
	FirstName       string `json:"firstName,omitempty",`
	LastName        string `json:"lastName,omitempty",`
	Email  			string `json:"email,omitempty"`
	JobTitle  		string `json:"job,omitempty",`
	PhoneNumber     string `json:"phonenumber,omitempty",`
}

func main(){
	fmt.Println("Default Schema Change UP");
	mux := http.NewServeMux()

	mux.HandleFunc("/", changeRequestForUpdatedVersion)
	
	http.ListenAndServe(":50002", mux)
}

func changeRequestForUpdatedVersion(w http.ResponseWriter, r *http.Request){
	log.Println("POST request received from localhost:50000")
	db := db_field{} //initialize 

	// log.Println(r)
	//Parse json request body and use it to set fields on db
	err := json.NewDecoder(r.Body).Decode(&db)
	
	if err != nil{
		panic(err)
	}
	log.Println("Apply Default Shcema for POST Request")
	
	//Marshal or convert user object back to json and write to response 
	dbJson, err := json.Marshal(db)
	if err != nil{
		panic(err)
	}
	
	responseBody := bytes.NewBuffer(dbJson)

	// create a post request
	url := "http://localhost:3000" + r.URL.String()
	//log.Println(url)
	resp, err := http.Post(url, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	
}
