package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://api.github.com/users/okaykaylyn")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if resp.StatusCode == http.StatusOK {
		log.Printf("Response: %v", resp)
	} else {
		log.Fatalf("Error: %v", resp)
	}

	if _, err = io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatalf("Error: %v", err)
	}
	// fmt.Println(resp.Body)

	var reply Reply
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&reply); err != nil {
		log.Fatalf("Error: can't decode - %s", err)
	}

	// fmt.Printf("%#v", reply)

}

type Reply struct {
	Name         string
	Public_Repos int
}

/*
	JSON <-> GO
	T/F <-> T/F
	STRING <-> STRING
	NULL <-> nil
	NUMBER <-> FLOAT64, INT64, INT32, INT16, INT8...
	OBJECT <-> map[string]any (map[string]interface{})
	ARRAY <-> SLICE []any ([]interface{})
*/
