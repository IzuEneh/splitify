package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func Map[T interface{}, V interface{}](data []V, f func(int, V) T) []T {
	mapped := make([]T, len(data))

	for i, e := range data {
		mapped[i] = f(i, e)
	}

	return mapped
}

func LogReq(req *http.Request) {
	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("REQUEST:\n%s", string(reqDump))
}

func LogResp(resp *http.Response) {
	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("RESPONSE:\n%s", string(respDump))
}

func ReadFromFile[T interface{}](fileName string) *T {
	// Let's first read the `config.json` file
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload T
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return &payload
}
