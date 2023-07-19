##Guests List

This project implements the guest list service (the REST API) for the year-end party!

Before the party, this service allow the functionality to add and remove guests with their entourages from the guest list and generate an invitation for the invited guests.

When the party begins, guests will arrive with an entourage. This party may not be the size indicated on the guest list. If the guest's table can accommodate the extra people, then the whole party should be let in. Otherwise, they will be turned away. Guests will also leave throughout the course of the party. When a guest leaves, their accompanying guests will leave with them.

At any point in the party, we should be able to know:

- Our guests at the party
- How many empty seats there are
- Here, we assume that each guest (along with their entourage) will have a separate table in the party and the table will not be shared between two guests.


## Functionalities

The API provides the following key features:

### BEFORE PARTY
- Add table to list of tables
- Add a guest to the guest list
- Remove a guest from the guest list
- Get the list of guests in the guest list
### DURING PARTY
- Check in guest
- Get list of checked in guest
- Get empty seats after check in
- Check out guest
### Implementation Details

Programming Language: GoLang 1.14 (refer to go.mod file)

Database: MySQL (Version 5.7) - can be installed from here

###Requirements:
- github.com/go-sql-driver/mysql v1.5.0 (refer to go.mod file)
- github.com/stretchr/testify v1.6.1 (refer to go.mod file)
- github.com/gin-gonic/gin v1.7.4 (refer to go.mod file)
  
###Future Improvements

In the future, I would consider the following improvements in the system:

- Implement secure authentication and authorization mechanisms to protect sensitive operations and restrict access to certain endpoints or data.

- Make use of GORM(https://gorm.io/index.html) which is a fully-featured orm with Auto migrations, Associations, Logger etc. This will make the system much more modular and extensible

- Here, we are assuming that the tables are not shared between guests. We can modify this service to allow the table sharing between guests and their entourage.
-  Implement rate limiting to prevent abuse or excessive usage of the API. This can help maintain system performance, ensure fair resource allocation, and mitigate the risk of denial-of-service attacks.
-  Incorporate logging and monitoring mechanisms to track API usage, detect anomalies, and gather insights about performance, errors, and usage patterns. This information can help with troubleshooting, capacity planning, and identifying areas for improvement.

- Adding more unittests and end-to-end integration tests.

## How to use
- use make command to start the docker-compose
```bash
make docker-up
```


## API demo

### 1. Add a table to the table list
- `POST` /tables
```bash
curl -X POST -H "Content-Type: application/json" -d '{"capacity" : 10}' http://localhost:3000/tables
```
- Output
```bash
{
    "id": "1",
    "capacity": 10
}
```
### 2.  Add a guest to the guest list
- `POST` /guest_list/:name
```bash
curl -X POST -H "Content-Type: application/json" -d '{"table" : 1, "accompanying_guests":1}' http://localhost:3000/guest_list/john
```
- Output
```bash
{
    "name": "John"
}
```

### 3. Get the guest list
- `GET`  /guest_list
```bash
curl -X GET http://localhost:3000/guest_list
```

- Output
```bash
{
     "guests": [
        {
            "name": "john",
            "table": 1,
            "accompanying_guests": 1
        }
    ]
}
```

### 4. Check in guests
- `PUT` /guests/:name
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"accompanying_guests":2}' http://localhost:3000/guests/john
```
- Output
```bash
{
    "name": "john"
}
```

### 5. Get arrived guests
- `GET` /guests
```bash
curl -X GET http://localhost:3000/guests
```

- Output
```bash
{
    "guests": [
        {
            "name": "john",
            "accompanying_guests": 2,
            "time_arrived": "2023-07-18 17:04:01"
        }
    ]
}
```

### 6. Count number of empty seats
- `GET` /seats_empty
```bash
curl -X GET http://localhost:3000/seats_empty
```
- Output
```bash

{
    "seats_empty": 7
}

```


### 7. Guest Leaves
- `DELETE` /guests/:name
```bash
curl -X DELETE http://localhost:3000/guests/john
```
- Output
  
HTTP Response Status Code: 204 No Content


### API Tests

## TestHealth
This function tests the `/health` endpoint of the router by sending a GET request to the endpoint. It verifies that the response status code is `http.StatusOK` (200). This test ensures that the health check endpoint is working correctly.

## TestAddTable
This function tests the `/tables` endpoint of the router by sending a POST request to the endpoint with a JSON payload representing a table object. It verifies that the response status code is `http.StatusOK` (200) and the response body matches the expected JSON object. This test ensures that a table can be successfully added using the API.

## TestAddGuestToList
This function tests the `/guest_list/{name}` endpoint of the router by sending a POST request to the endpoint with a JSON payload representing a guest object. It verifies that the response status code is `http.StatusOK` (200) and the response body matches the expected JSON object. This test ensures that a guest can be added to the guest list using the API.

## TestGetGuestList
This function tests the `/guest_list` endpoint of the router by sending a GET request to the endpoint. It verifies that the response status code is `http.StatusOK` (200) and the response body matches the expected JSON object representing a list of guests. This test ensures that the guest list can be retrieved successfully from the API.

## TestCheckInGuest
This function tests the `/guests/{name}` endpoint of the router by sending a PUT request to the endpoint with a JSON payload representing the accompanying guests for a guest. It verifies that the response status code is `http.StatusOK` (200) and the response body matches the expected JSON object. This test ensures that a guest can be checked-in using the API.

## TestGetCheckedInGuest
This function tests the `/guests` endpoint of the router by sending a GET request to the endpoint. It verifies that the response status code is `http.StatusOK` (200) and the response body contains a list of checked-in guests. It further checks the fields of the first guest in the list to ensure they have the expected values. This test ensures that the checked-in guest list can be retrieved successfully from the API.

## TestCheckOutGuest
This function tests the `/guests/{name}` endpoint of the router by sending a DELETE request to the endpoint. It verifies that the response status code is `http.StatusNoContent` (204), indicating successful guest checkout. This test ensures that a guest can be checked out using the API.






