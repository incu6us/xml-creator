package api

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type ROOT struct {
	Shipment struct {
		Identifier int64 `xml:"IDENTIFIER", json:"identifier"`
		Booking    struct {
			Pax []struct {
				Name string `xml:"NAME", json:"name"`
			} `xml:"PAX", json:"pax"`
		} `xml:"BOOKING", json:"booking"`
	} `xml:"SHIPMENT", json:"shipment"`
}

type apiHandler struct {
	ContentType string
}

type response struct {
	Result interface{} `json:"Result"`
}

const (
	CONTENT_TYPE = "application/xml"
)

var handler *apiHandler

func (a *apiHandler) setJsonHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", a.ContentType)
	w.WriteHeader(http.StatusOK)
}

func (a apiHandler) print(w http.ResponseWriter, r *http.Request, message interface{}) {
	a.setJsonHeader(w)

	if encodeError := json.NewEncoder(w).Encode(response{message}); encodeError != nil {
		log.Println("Parse message error", encodeError)
	}
}

func (a apiHandler) CreateObject(w http.ResponseWriter, r *http.Request) {

	var err error
	var body, xmlOutput []byte

	if body, err = ioutil.ReadAll(r.Body); err != nil {
		a.print(w, r, err)
		return
	}

	var result ROOT
	if err = json.Unmarshal(body, &result); err != nil {
		log.Println(err)
	}

	log.Printf("JSON Request: %#v", result)

	if xmlOutput, err = xml.MarshalIndent(result, "", "  "); err != nil {
		log.Println(err)
	}

	w.Write(xmlOutput)
}

type ApiHandler interface {
	CreateObject(http.ResponseWriter, *http.Request)
}

func GetHandler() ApiHandler {

	if handler == nil {
		handler = &apiHandler{ContentType: CONTENT_TYPE}
	}

	return handler
}
