package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	_ "net/http"
	"net/http/httptest"
	"testing"

	"github.com/getground/tech-tasks/backend/router"
	"github.com/stretchr/testify/assert"
)

func TestAddTable(t *testing.T) {
	router := router.SetupRouter()

    requestBody := map[string]interface{}{
        "capacity": 10,
        
    }

    // Convert the request body to JSON
    requestBodyBytes, _ := json.Marshal(requestBody)

    req := httptest.NewRequest("POST", "/tables", bytes.NewBuffer(requestBodyBytes))
    req.Header.Set("Content-Type", "application/json")


    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    // Assert the response status code
    if resp.Code != http.StatusOK {
        t.Errorf("Expected status %d but got %d", http.StatusOK, resp.Code)
    }

    var responseBody map[string]interface{}
    err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
    if err != nil {
        t.Fatalf("Failed to unmarshal response body: %v", err)
    }

     expectedResponseBody := map[string]interface{}{
        "id": float64(1),
        "capacity": float64(10),
    }

    assert.Equal(t, expectedResponseBody, responseBody, "Response body mismatch.\nExpected: %v\nActual: %v", expectedResponseBody, responseBody)

}


func TestAddGuestToList(t *testing.T) {
	router := router.SetupRouter()

    requestBody := map[string]interface{}{
        "table": 1,
        "accompanying_guests": 1,
        
    }

    // Convert the request body to JSON
    requestBodyBytes, _ := json.Marshal(requestBody)

    req := httptest.NewRequest("POST", "/guest_list/john", bytes.NewBuffer(requestBodyBytes))
    req.Header.Set("Content-Type", "application/json")


    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    // Assert the response status code
    if resp.Code != http.StatusOK {
        t.Errorf("Expected status %d but got %d", http.StatusOK, resp.Code)
    }

    var responseBody map[string]interface{}
    err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
    if err != nil {
        t.Fatalf("Failed to unmarshal response body: %v", err)
    }

     expectedResponseBody := map[string]interface{}{
        "name": "john",
    }

    assert.Equal(t, expectedResponseBody, responseBody, "Response body mismatch.\nExpected: %v\nActual: %v", expectedResponseBody, responseBody)

}

func TestGetGuestList(t *testing.T) {
	router := router.SetupRouter()

    req := httptest.NewRequest("GET", "/guest_list", nil)
    req.Header.Set("Content-Type", "application/json")


    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    // Assert the response status code
    if resp.Code != http.StatusOK {
        t.Errorf("Expected status %d but got %d", http.StatusOK, resp.Code)
    }

    var responseBody map[string]interface{}
    err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
    if err != nil {
        t.Fatalf("Failed to unmarshal response body: %v", err)
    }

     expectedResponseBody := map[string]interface{}{
          "guests": []interface{}{
              map[string]interface{}{
                   "name": "john",
                   "table": float64(1),
                   "accompanying_guests": float64(1),

              },
              
          },
       
    }

    assert.Equal(t, expectedResponseBody, responseBody, "Response body mismatch.\nExpected: %v\nActual: %v", expectedResponseBody, responseBody)

}


       
func TestCheckInGuest(t *testing.T) {
	router := router.SetupRouter()

    requestBody := map[string]interface{}{
        "accompanying_guests": 1,
        
    }

    // Convert the request body to JSON
    requestBodyBytes, _ := json.Marshal(requestBody)

    req := httptest.NewRequest("PUT", "/guests/john", bytes.NewBuffer(requestBodyBytes))
    req.Header.Set("Content-Type", "application/json")


    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    // Assert the response status code
    if resp.Code != http.StatusOK {
        t.Errorf("Expected status %d but got %d", http.StatusOK, resp.Code)
    }

    var responseBody map[string]interface{}
    err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
    if err != nil {
        t.Fatalf("Failed to unmarshal response body: %v", err)
    }

     expectedResponseBody := map[string]interface{}{
        "name": "john",
    }

    assert.Equal(t, expectedResponseBody, responseBody, "Response body mismatch.\nExpected: %v\nActual: %v", expectedResponseBody, responseBody)

}

func TestGetCheckedInGuest(t *testing.T) {
	router := router.SetupRouter()

    req := httptest.NewRequest("GET", "/guests", nil)
    req.Header.Set("Content-Type", "application/json")


    resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    // Assert the response status code
    if resp.Code != http.StatusOK {
        t.Errorf("Expected status %d but got %d", http.StatusOK, resp.Code)
    }

    var responseBody map[string]interface{}
    err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
    if err != nil {
        t.Fatalf("Failed to unmarshal response body: %v", err)
    }

    guests, ok := responseBody["guests"].([]interface{})
    assert.True(t, ok, "guests field is missing or not a slice")
    assert.NotEmpty(t, guests, "guests field is empty")

    guest, ok := guests[0].(map[string]interface{})
    assert.True(t, ok, "guests[0] is missing or not a map")


    timeArrived, ok := guest["time_arrived"]
    assert.True(t, ok, "time_arrived field is missing or not a string")
    assert.NotEmpty(t, timeArrived, "time_arrived field is empty")

    accompanyingGuests, ok := guest["accompanying_guests"].(float64)
    assert.True(t, ok, "accompanying_guests field is missing or not an int")
    assert.Equal(t, float64(1), accompanyingGuests, "accompanying_guests field does not have expected value")

    name, ok := guest["name"].(string)
    assert.True(t, ok, "name field is missing or not a string")
    assert.Equal(t, "john", name, "name field does not have expected value")

   


}


func TestCheckOutGuest(t *testing.T) {
    router := router.SetupRouter()

    req := httptest.NewRequest("DELETE", "/guests/john", nil)
    req.Header.Set("Content-Type", "application/json")

     resp := httptest.NewRecorder()
    router.ServeHTTP(resp, req)

    // Assert the response status code
    if resp.Code != http.StatusNoContent {
        t.Errorf("Expected status %d but got %d", http.StatusOK, resp.Code)
    }

}

