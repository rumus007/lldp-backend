package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/samber/lo"
)

type chassis struct {
	Nodeid string `json:"node_id"`
	Nodeidtype string `json:"node_id_type"`
	Sysname string `json:"sys_name"`
	Sysdescription string `json:"sys_description"`
	Mgmtip string `json:"mgmt_ip"`
	Capability struct {
		Bridge bool `json:"Bridge"`
		Router bool `json:"Router"`
		Wlan bool `json:"Wlan"`
		Station bool `json:"Station"`
	}
}

type neighbor struct {
	Neighborid string `json:"neighbor_id"`
	Neighboridtype string `json:"neighbor_id_type"`
	Name string `json:"name"`
	Mgmtip string `json:"mgmt_ip"`
	Portidtype string `json:"port_id_type"`
	Portid string `json:"port_id"`
	Portdescription string `json:"port_description"`
	Portttl string `json:"port_ttl"`
	Capability struct {
		Bridge bool `json:"Bridge"`
		Router bool `json:"Router"`
		Wlan bool `json:"Wlan"`
		Station bool `json:"Station"`
	}
}

type lldp struct {
	Chassis chassis
	Neighbor []neighbor 
}

func createIfDoesnotExist()(filename string){
	filename = "lldp.json"
	// Check if file exists
	_, err := os.Stat(filename)

	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist create a new json file
			file, err := os.Create(filename)

			if err != nil {
				fmt.Println("Error creating file: ", err)
				return ""
			}
			defer file.Close()
			
			// Create an empty slice to represent the initial array
  			var emptyArray []interface{} // Use interface{} for flexibility
			
			// Marshal the empty array to JSON
			jsonData, err := json.Marshal(emptyArray)
			if err != nil {
				fmt.Println("Error marshalling empty array:", err)
				return ""
			}

			// Write the JSON data to the file
			_, err = file.Write(jsonData)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return ""
			}

			fmt.Println("JSON file created successfully")

		}
	} else {
		fmt.Println("JSON File already exists")
	}

	return filename
}

func storeOrUpdateData(data lldp)(string, error){
	var lldpdata []lldp
	filename := createIfDoesnotExist()

	fileBytes, err := os.ReadFile(filename)

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(fileBytes, &lldpdata)

	if err != nil {
		return "", err
	}

	_, index, ok := lo.FindIndexOf(lldpdata, func (i lldp) bool {
		return i.Chassis.Nodeid == data.Chassis.Nodeid
	})


	if ok {
		// data already exists need to update the index
		lldpdata[index] = data
	} else {
		// data does not exist so append it to the json
		lldpdata = append(lldpdata, data)
	}

	updatedJson, err := json.Marshal(lldpdata)

	if err != nil {
		return "", err
	}

	err = os.WriteFile(filename, updatedJson, 0644)

	if err != nil {
		return "", err
	}

	return "success", nil
}