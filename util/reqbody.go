package util 

import (
     "net/http"
	 "encoding/json"
)

func ReadJson(w http.ResponseWriter, r *http.Request, dst interface{}) error {
     
	//Initialize the decoder
	dec := json.NewDecoder(r.Body)
    // Configure the decoder to return an error if the JSON contains unknown fields.
    dec.DisallowUnknownFields()
	err := dec.Decode(dst)

	if err != nil {
		return err
	}

  return nil 
}