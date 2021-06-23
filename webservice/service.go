package webservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"example.com/asmaraputra/learn_golang/palindrome"
)

/*
	make function to handle request with parameter w and r
	w = response that return to client
	r = response that client get after make request to server
*/

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// print request method
	log.Println(r.Method)

	// check incoming request using switch case
	switch r.Method {
	// method POST
	case http.MethodPost:
		// allow post method
		// parsing the form
		if err := r.ParseForm(); err != nil {
			// return error with bad status
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// get value from input
		inputNumber := r.FormValue("input-number")

		// print input
		fmt.Println(inputNumber)

		// call palindrom func from palindrome package
		palindrome, err := palindrome.Palindrome(string(inputNumber))

		// check if any error
		if err != nil {
			// return http error with message and status
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// create response body
		response := ResBody{
			Message: "Input Number : " + inputNumber + "Palindrome Found Is : " + strconv.FormatFloat(palindrome, 'f', 0, 64),
			Data:    palindrome,
			Status:  201,
			IsError: false,
		}

		// add Content-Type
		w.Header().Add("Content-Type", "application/json")
		// status code
		w.WriteHeader(http.StatusCreated)
		// encode response
		json.NewEncoder(w).Encode(response)

		// method GET
	case http.MethodGet:
		// return  message
		w.Write([]byte("Find total palindrome number that found between 2 numbers."))
		break
	default:
		// return error 403
		w.WriteHeader(403)
		w.Write([]byte("Method Not Allowed"))
		break
	}
}

// service that run request to find palindrome number
func Service() {

	// set route to /
	http.HandleFunc("/", handleRequest)

	// print notification service is running
	fmt.Println("Your App Is Running")

	// run service on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
