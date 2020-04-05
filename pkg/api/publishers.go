package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ea3hsp/book/pkg/models"
	"github.com/gorilla/mux"
)

func (a *API) viewPublisher(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["pubid"]
	// get publisher from data store
	publisher, err := a.store.GetPublisher(id)
	if err != nil {
		a.returnError(err, w, r)
	}
	err = a.render.RenderTemplate(w, "vpub.html", publisher)
	if err != nil {
		a.returnError(err, w, r)
	}
}

func (a *API) viewPublishers(w http.ResponseWriter, r *http.Request) {
	// get book from data store
	publishers, err := a.store.GetPublishers()
	if err != nil {
		a.returnError(err, w, r)
	}
	err = a.render.RenderTemplate(w, "vpubs.html", publishers)
	if err != nil {
		a.returnError(err, w, r)
	}
}

func (a *API) createPublisher(w http.ResponseWriter, r *http.Request) {
	// auth holder
	publisher := &models.Publisher{}
	// Content type header
	w.Header().Set("Content-Type", "application/json")
	// Unmarshal body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.returnError(err, w, r)
	}
	// unmarshall metadata
	err = json.Unmarshal(bodyBytes, &publisher)
	if err != nil {
		a.returnError(err, w, r)
	}
	// put book into data store
	res, err := a.store.CreatePublisher(publisher)
	if err != nil {
		a.returnError(err, w, r)
	}
	a.returnOk(res, w, r)
}
