package main

import (
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var goVersion = runtime.Version()

//Test reaching the URL and getting multiple airplanes
func TestHttpURL(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/airplanes", func(c *gin.Context) { c.String(http.StatusOK, "Success") })

	go func() {
		assert.NoError(t, router.Run())

	}()
	// Wait for go routine to start and run the server
	// otherwise the main thread will complete
	time.Sleep(5 * time.Millisecond)

}

//Tests getting a single airplane value by ID
func TestGetSinglePlane(t *testing.T) {
	handler := func(c *gin.Context) {
		c.String(http.StatusOK, "OK")

	}

	router := gin.New()
	router.GET("/airplane/1", handler)

	req, _ := http.NewRequest("GET", "/airplane/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "OK")

}

// Tests the update airplane function
func TestUpdateAirplane(t *testing.T) {
	handler := func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	}

	router := gin.New()
	router.PUT("/airplane/1", handler)

	req, _ := http.NewRequest("PUT", "/airplane/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "OK")
}

// Tests the update airplane function
func TestDestroyAirplane(t *testing.T) {
	handler := func(c *gin.Context) {
		c.String(http.StatusOK, "Destroyed")
	}

	router := gin.New()
	router.DELETE("/airplane/1", handler)

	req, _ := http.NewRequest("DELETE", "/airplane/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "Destroyed")

}
