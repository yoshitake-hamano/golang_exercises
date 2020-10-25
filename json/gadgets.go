package main

import (
	"encoding/json"
	"log"
	"reflect"
)

const gadgetsJson = `
[
  {"Type": "button",
    "Gadget": {
      "Port": "/dev/ttyACM0"}},
  {"Type": "hello",
    "Gadget": {
      "Message1": "str b 1",
      "Message2": "str b 2"}}
]
`

// Gadget
type Gadget interface {
	Execute([]byte) (bool, []byte)
}

// GadgetEnvelope
type GadgetEnvelope struct {
	Type   string
	Gadget Gadget
}

// ButtonGadget is a Gadget which accutuates a solenoid
type ButtonGadget struct {
	Port string
}

func (s *ButtonGadget) Execute(data []byte) (bool, []byte) {
	log.Printf("[%s] Button\r\n", reflect.TypeOf(*s))
	return true, data
}

// HelloGadget is a Gadget which says hello
type HelloGadget struct {
	Message1 string
	Message2 string
}

func (h *HelloGadget) Execute(data []byte) (bool, []byte) {
	log.Printf("[%s] Hello World\r\n", reflect.TypeOf(*h))
	return true, data
}

type GadgetDecorder struct {
	Type   string
	Gadget json.RawMessage
}

func buildGadgets() []GadgetEnvelope {
	var f []GadgetDecorder
	if err := json.Unmarshal([]byte(gadgetsJson), &f); err != nil {
		log.Fatalf("[gadget json]: %s", err)
	}

	gadgets := make([]GadgetEnvelope, 0)
	for _, e := range f {
		var g Gadget
		switch e.Type {
		case "hello":
			g = &HelloGadget{}
		case "button":
			g = &ButtonGadget{}
		}
		if err := json.Unmarshal(e.Gadget, g); err != nil {
			log.Fatal(err)
		}
		env := GadgetEnvelope{Type: e.Type, Gadget: g}
		gadgets = append(gadgets, env)
	}
	return gadgets
}

func main() {
	gadgets := buildGadgets()
	b, _ := json.Marshal(gadgets)
	log.Print(string(b))
}
