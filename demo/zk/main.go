package main

import (
	"fmt"
	gocircom "github.com/iden3/go-circom-witnesscalc"
	"io/ioutil"
)

func main() {
	wasmBytes, err := ioutil.ReadFile("./circuit.wasm")
	noErr(err)

	inputBytes, err := ioutil.ReadFile("./input.json")
	noErr(err)

	calc, err := gocircom.NewCircom2WitnessCalculator(wasmBytes, true)
	noErr(err)

	inputs, err := gocircom.ParseInputs(inputBytes)
	noErr(err)

	witness, err := calc.CalculateWitness(inputs, true)
	noErr(err)
	fmt.Println(witness)
}

func noErr(e error) {
	if e != nil {
		panic(e)
	}
}
