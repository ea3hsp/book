package api

import "github.com/dgrijalva/jwt-go"

const secretPass = "Sup3rSecr3tP4ssWord"

func (a *API) getPassword() []byte {
	return []byte(secretPass)
}

func (a *API) getToken(name string) (string, error) {
	signingKey := a.getPassword()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"role": "admin",
	})
	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func (a *API) verifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := a.getPassword()
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
