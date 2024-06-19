# This is a mock server using Golang

## Requirements
Golang >= v1.22.4

## Installation and usage
Clone the repo

Run command `go build` to create the binary file named `lldp-backend`

Run command `./lldp-backend` to run the binary file with creates a server running on port 8080 on the localhost

## Endpoints

Currently only has one endpoint `localhost:8080/lldp` with is a **POST** request that takes the lldp data and stores as a JSON file. 

## Example of Body for API request

The `chassis` key contains the data of the machine and the `neighbor` key contains an array of the devices the machine is connected with.

You can test the API after running the server with POSTMAN or Thunderclient or any API testing software of your choice. Add the follwing endpoint

```http://localhost:8080/lldp``` as a **POST** request and copy the following JSON object into the BODY of the request.

```
{
   "chassis":{
      "node_id":"0c:42:a1:36:b3:e8",
      "node_id_type":"mac",
      "sys_name":"ixp",
      "sys_description":"Cumulus Linux version 5.2.1 running on Mellanox Technologies Ltd. MSN2100",
      "mgmt_ip":"10.99.0.2",
      "capability":{
         "Bridge":true,
         "Router":true,
         "Wlan":false,
         "Station":false
      }
   },
   "neighbor":[
      {
         "neighbor_id":"b4:a9:fc:a7:07:02",
         "neighbor_id_type":"mac",
         "name":"mg-bridge1",
         "mgmt_ip":"10.0.0.254",
         "port_id_type":"ifname",
         "port_id":"swp48",
         "port_description":"swp48",
         "port_ttl":"120",
         "capability":{
            "Bridge":true,
            "Router":true,
            "Wlan":false,
            "Station":false
         }
      },
      {
         "neighbor_id":"3c:fd:fe:9e:f0:2b",
         "neighbor_id_type":"mac",
         "name":"None",
         "mgmt_ip":"None",
         "port_id_type":"mac",
         "port_id":"3c:fd:fe:9e:f0:2b",
         "port_description":"None",
         "port_ttl":"121",
         "capability":{
            "Bridge":false,
            "Router":false,
            "Wlan":false,
            "Station":false
         }
      },
      {
         "neighbor_id":"3c:fd:fe:9e:f0:2a",
         "neighbor_id_type":"mac",
         "name":"None",
         "mgmt_ip":"None",
         "port_id_type":"mac",
         "port_id":"3c:fd:fe:9e:f0:2a",
         "port_description":"None",
         "port_ttl":"121",
         "capability":{
            "Bridge":false,
            "Router":false,
            "Wlan":false,
            "Station":false
         }
      }
   ]
}
```