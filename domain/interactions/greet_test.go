package go_specs_greet_test

import (
	"testing"
	go_specs_greet "github.com/cossinner/go_specs_greet"
	"github.com/cossinner/go_specs_greet/specifications"
)
func TestGreet(t *testing.T) {
	// This is a test function that runs the GreetSpecification test
	// with the go_specs_greet.Greet function as the implementation.
	specifications.GreetSpecification(t, specifications.GreetAdapter(go_specs_greet.Greet))
}