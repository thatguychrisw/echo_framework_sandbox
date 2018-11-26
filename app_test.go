package main

import (
	"os"
	"testing"

	"gopkg.in/h2non/baloo.v3"
)

func TestMain(m *testing.M) {
	e := setupRouter()
	e.HideBanner = true
	e.HidePort = true

	go startApp(e)

	os.Exit(m.Run())
}

func TestItHasAHealthCheckEndpoint(t *testing.T) {
	server := baloo.New("http://localhost:8000")

	_ = server.Get("/lb-status").
		Expect(t).
		Status(200).
		Done()
}
