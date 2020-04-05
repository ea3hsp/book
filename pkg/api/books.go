package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ea3hsp/book/pkg/models"
	"github.com/gorilla/mux"
)

func (a *API) homePage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/books", http.StatusTemporaryRedirect)
	/* err := a.render.RenderTemplate(w, "index.tmpl", nil)
	if err != nil {
		a.returnError(err, w, r)
	} */
}

func (a *API) viewBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["isbn"]
	// get book from data store
	book, err := a.store.GetBook(id)
	if err != nil {
		a.returnError(err, w, r)
	}
	err = a.render.RenderTemplate(w, "vbook.html", book)
	if err != nil {
		a.returnError(err, w, r)
	}
}

func (a *API) viewBooks(w http.ResponseWriter, r *http.Request) {
	// get book from data store
	books, err := a.store.GetBooks()
	if err != nil {
		a.returnError(err, w, r)
	}
	err = a.render.RenderTemplate(w, "vbooks.html", books)
	if err != nil {
		a.returnError(err, w, r)
	}
}

func (a *API) viewCreateBook(w http.ResponseWriter, r *http.Request) {
	auths, err := a.store.GetAuthors()
	if err != nil {
		a.returnError(err, w, r)
	}
	err = a.render.RenderTemplate(w, "vcrbook.html", auths)
	if err != nil {
		a.returnError(err, w, r)
	}
}

func (a *API) createBook(w http.ResponseWriter, r *http.Request) {
	// book holder
	book := &models.Book{}
	// Content type header
	w.Header().Set("Content-Type", "application/json")
	// Unmarshal body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		a.returnError(err, w, r)
	}
	// unmarshall metadata
	err = json.Unmarshal(bodyBytes, &book)
	if err != nil {
		a.returnError(err, w, r)
	}
	// put book into data store
	res, err := a.store.CreateBook(book)
	if err != nil {
		a.returnError(err, w, r)
	}
	a.returnOk(res, w, r)
}
