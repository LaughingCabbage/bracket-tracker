package pga

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetCurrentTournament returns the json of the current tournament
func GetCurrentTournament() map[string]interface{} {
	resp, err := http.Get(CurrentEventEndpoint)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var f interface{}
	if err := json.Unmarshal(body, &f); err != nil {
		panic(err)
	}

	m := f.(map[string]interface{})

	return m
}

func printMap(m map[string]interface{}) {
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
