package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ea3hsp/book/pkg/models"
	"github.com/gorilla/mux"
)

func (a *API) viewAuthor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["authid"]
	// get book from data store
	auth, err := a.store.GetAuthor(id)
	if err != nil {
		a.returnError(err, w, r)
	}
	err = a.render.RenderTemplate(w, "vauthor.html", auth)
	if err != nil {
		a.returnError(err, w, r)
	}
}

func (a *API) viewAuthors(w http.ResponseWriter, r *http.Request) {
	// get book from data store
	auths, err := a.store.GetAuthors()
	if err != nil {
		a.returnError(err, w, r)
	}
	err = a.render.RenderTemplate(w, "vauthors.html", auths)
	if err != nil {
		a.returnError(err, w, r)
	}
}

func (a *API) createAuth(w http.ResponseWriter, r *http.Request) {
	// auth holder
	auth := &models.Author{}
	// Content type header
	w.Header().Set("Content-Type", "application/json")
	// Unmarshal body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.returnError(err, w, r)
	}
	// unmarshall metadata
	err = json.Unmarshal(bodyBytes, &auth)
	if err != nil {
		a.returnError(err, w, r)
	}
	// put book into data store
	res, err := a.store.CreateAuth(auth)
	if err != nil {
		a.returnError(err, w, r)
	}
	a.returnOk(res, w, r)
}
