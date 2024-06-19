package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Start the server
	http.HandleFunc("/lldp", lldpHandler)
	fmt.Println("Server listinig at port 8080...")
	http.ListenAndServe(":8080", nil)
}

func sendResponse(w http.ResponseWriter, statusCode int, message string){
	response := map[string]string{"status_code": fmt.Sprintf("%d", statusCode), "message": message}
	jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintf(w, "Error marshalling response: %v", err)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}

func lldpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendResponse(w, http.StatusMethodNotAllowed, "Only POST method allowed for lldp endpoint")
        return
	}

	//Read the request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, fmt.Sprintf("Error reading request body: %v", err))
        return
	}

	// unmarshall the json data from the request body
	var requestbody lldp
	err = json.Unmarshal(body, &requestbody)
	if err != nil {
		sendResponse(w, http.StatusBadRequest, fmt.Sprintf("Error unmarshalling JSON data: %v", err))
        return
	}
	
	//store the data
	_, err = storeOrUpdateData(requestbody)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// fmt.Fprintf(w, "Error saving LLDP data: %v", err)
		// return
		sendResponse(w, http.StatusInternalServerError, fmt.Sprintf("Error saving LLDP data: %v", err))
        return
	}
	
	// send successful response
	sendResponse(w, http.StatusOK, "LLDP Data inserted successfully")
}
