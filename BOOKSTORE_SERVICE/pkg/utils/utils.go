package utils

import (
	"encoding/json"
	"io" //io/ioutil is deprecated
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) { //interface{} type allows you to pass any Go type (whether a struct, slice, map, etc.) into the function
	if body, err := io.ReadAll(r.Body); err == nil { //reads the entire body of the HTTP request and returns it as a byte slice (body)
		if err := json.Unmarshal([]byte(body), x); err != nil { // convert the byte slice body (which we assume contains JSON data) into a Go data structure
			return
		}
	}
}
