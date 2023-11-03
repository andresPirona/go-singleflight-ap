package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func setupTestServer(endpoint string, handler gin.HandlerFunc) *httptest.Server {
	r := gin.Default()
	r.GET(endpoint, handler)
	return httptest.NewServer(r)
}

func runConcurrentRequests(t *testing.T, endpoint string, handler gin.HandlerFunc) {
	ts := setupTestServer(endpoint, handler)

	var wg sync.WaitGroup

	numRequests := 10
	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			start := time.Now()
			res, err := http.Get(ts.URL + endpoint)
			if err != nil {
				t.Errorf("Error al realizar la solicitud: %v", err)
				return
			}
			defer res.Body.Close()

			if res.StatusCode != http.StatusOK {
				t.Errorf("Respuesta inesperada del servidor: %d", res.StatusCode)
				return
			}
			elapsed := time.Since(start)
			t.Logf("Solicitud completada en %s", elapsed)
		}()
	}

	wg.Wait()
	ts.Close()
}

func TestGetAll(t *testing.T) {
	runConcurrentRequests(t, "/countries", GetCountryList)
}

func TestSingleflightGetAll(t *testing.T) {
	runConcurrentRequests(t, "/countriesSingleflight", GetCountrySingleflightList)
}
