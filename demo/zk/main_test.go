package main

import (
	gocircom "github.com/iden3/go-circom-witnesscalc"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
)

func TestCompile(t *testing.T) {
	wasmBytes, err := ioutil.ReadFile("./circuit.wasm")
	require.NoError(t, err)

	inputBytes, err := ioutil.ReadFile("./input.json")
	require.NoError(t, err)

	calc, err := gocircom.NewCircom2WitnessCalculator(wasmBytes, true)
	require.NoError(t, err)
	require.NotEmpty(t, calc)

	inputs, err := gocircom.ParseInputs(inputBytes)
	require.NoError(t, err)

	witness, err := calc.CalculateWitness(inputs, true)
	require.NoError(t, err)
	require.NotEmpty(t, witness)
}
