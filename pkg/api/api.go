package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ea3hsp/book/pkg/models"
	"github.com/ea3hsp/book/pkg/render"
	"github.com/ea3hsp/book/pkg/store"

	"github.com/go-kit/kit/log"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// API struct definition
type API struct {
	logger  log.Logger
	address string
	passwd  string
	render  *render.Render
	store   store.IStore
}

func (a *API) returnError(err error, w http.ResponseWriter, r *http.Request) {
	// response holder
	res := new(models.APIResponse)
	a.logger.Log("[error]", fmt.Sprintf("method: %s", r.RequestURI), "msg", err.Error())
	res.Msg = err.Error()
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&res)
	return
}

func (a *API) returnOk(msg string, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := &models.APIResponse{
		Msg: msg,
	}
	// encode rec
	json.NewEncoder(w).Encode(&res)
}

// NewAPI creates new API
func NewAPI(logger log.Logger, addr, passwd string, render *render.Render, store store.IStore) *API {
	return &API{
		logger:  logger,
		address: addr,
		passwd:  passwd,
		render:  render,
		store:   store,
	}
}

// Init initializes API REST Service
func (a *API) Init() {
	// creates mux
	router := mux.NewRouter()
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	// Login handle func
	router.HandleFunc("/login", a.authentificate).Methods("POST")
	router.HandleFunc("/", a.homePage).Methods("GET")
	// library
	router.HandleFunc("/book/isbn/{isbn}", a.viewBook).Methods("GET")
	router.HandleFunc("/books", a.viewBooks).Methods("GET")
	router.HandleFunc("/book/create", a.viewCreateBook).Methods("GET")
	// Authors
	router.HandleFunc("/author/authid/{authid}", a.viewAuthor).Methods("GET")
	router.HandleFunc("/authors", a.viewAuthors).Methods("GET")
	//router.HandleFunc("/book/create", a.viewCreateAuthor).Methods("GET")
	// Publishers
	router.HandleFunc("/publisher/pubid/{pubid}", a.viewPublisher).Methods("GET")
	router.HandleFunc("/publishers", a.viewPublishers).Methods("GET")
	//router.HandleFunc("/book/create", a.viewCreatePublisher).Methods("GET")
	// Api handle funcs
	api := router.PathPrefix("/api").Subrouter()
	// IBA REST function handlers
	api.HandleFunc("/book", a.createBook).Methods("POST")
	api.HandleFunc("/author", a.createAuth).Methods("POST")
	api.HandleFunc("/publisher", a.createPublisher).Methods("POST")
	// Used middlewares
	// logging middleware
	router.Use(a.loggingMiddleware)
	// api auth middleware
	api.Use(a.authMiddleware)
	// set address
	address := fmt.Sprintf("0.0.0.0:%s", a.address)
	// init rest api server ...
	err := http.ListenAndServe(address, handlers.RecoveryHandler()(router))
	if err != nil {
		a.logger.Log("[error]", "Error Initializing API", "msg", err)
		os.Exit(0)
	}
}
