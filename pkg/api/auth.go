package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *API) authentificate(w http.ResponseWriter, r *http.Request) {
	// credentials holder
	var login credentials
	// body request decoder
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &login)
	// check error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	// check name and password length
	if len(login.Username) == 0 || len(login.Password) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please provide name and password to obtain the token"))
		return
	}
	// check username and password
	if login.Username == "admin" && login.Password == a.passwd {
		token, err := a.getToken(login.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error generating JWT token: " + err.Error()))
		} else {
			w.Header().Set("Authorization", "Bearer "+token)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Token: " + token))
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Name and password do not match"))
		return
	}

}

// Rest API auth middleware
func (a *API) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get token string
		tokenString := r.Header.Get("Authorization")
		// check token string length
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		// Remove Bearer part from token string
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := a.verifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		name := claims.(jwt.MapClaims)["name"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)

		r.Header.Set("name", name)
		r.Header.Set("role", role)
		next.ServeHTTP(w, r)
	})
}
