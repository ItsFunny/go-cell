package main

import (
	"fmt"
	gocircom "github.com/iden3/go-circom-witnesscalc"
	"github.com/iden3/go-wasm3"
	"io/ioutil"
)

func main() {
	runtime := wasm3.NewRuntime(&wasm3.Config{
		Environment: wasm3.NewEnvironment(),
		StackSize:   64 * 1024,
	})
	defer runtime.Destroy()

	wasmBytes, err := ioutil.ReadFile("./circuit.wasm")
	noErr(err)

	module, err := runtime.ParseModule(wasmBytes)
	noErr(err)
	module, err = runtime.LoadModule(module)
	noErr(err)

	inputBytes, err := ioutil.ReadFile("./input.json")
	noErr(err)

	calc, err := gocircom.NewWitnessCalculator(runtime, module)
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
