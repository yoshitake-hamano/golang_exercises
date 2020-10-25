package main

import (
	"encoding/json"
	"log"
)

const input = `
[
  {"ID": "a", "Config": {
    "Message": "str a 1"}},
  {"ID": "b", "Config": {
    "Message1": "str b 1",
    "Message2": "str b 2"}}
]
`

type Envelope struct {
	ID     string
	Config json.RawMessage
	config IConfig
}

type IConfig interface {
}

type ConfigA struct {
	Message string
}

type ConfigB struct {
	Message1 string
	Message2 string
}

func main() {
	var envelopes []Envelope
	if err := json.Unmarshal([]byte(input), &envelopes); err != nil {
		log.Fatal(err)
	}
	for _, env := range envelopes {
		log.Print(env.ID)
		log.Print(string(env.Config))
		switch env.ID {
		case "a":
			var config ConfigA
			if err := json.Unmarshal([]byte(env.Config), &config); err != nil {
				log.Fatal(err)
			}
			log.Print("ConfigA : ", config.Message)
			env.config = config
		case "b":
			var config ConfigB
			if err := json.Unmarshal([]byte(env.Config), &config); err != nil {
				log.Fatal(err)
			}
			log.Print("ConfigB : ", config.Message1)
			log.Print("ConfigB : ", config.Message2)
			config.Message1 = "foo"
			env.Config, _ = json.Marshal(&config)
			env.config = config
		}
	}

	jsonData, err := json.Marshal(envelopes)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(input)
	log.Print(string(jsonData))
}
