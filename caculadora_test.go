package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculadora(t *testing.T) {
	operacion := []struct {
		nombreOperacion string
		x               int
		y               int
		z               int
	}{
		{
			nombreOperacion: "suma",
			x:               10,
			y:               15,
			z:               25,
		},
		{
			nombreOperacion: "resta",
			x:               25,
			y:               15,
			z:               10,
		},
		{
			nombreOperacion: "multiplicacion",
			x:               8,
			y:               5,
			z:               40,
		},
		{
			nombreOperacion: "division",
			x:               32,
			y:               8,
			z:               4,
		},
	}

	for _, numero := range operacion {
		total := calculadora(numero.nombreOperacion, numero.x, numero.y)
		assert.Equal(t, numero.z, total)
	}
}
