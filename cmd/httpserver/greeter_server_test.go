package main_test

import (
	go_specs_greet "github.com/cossinner/go_specs_greet"
)

func TestGreeterServer(t *testing.T) {
	driver := go_specs_greet.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}