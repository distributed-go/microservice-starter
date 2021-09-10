package authmodel

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/distributed-go/microservice-starter/models/recruitermodel"
)

// AppClaims represent the claims parsed from JWT access token.
type AppClaims struct {
	ID    string                `json:"ID,omitempty"`
	Sub   string                `json:"Sub,omitempty"`
	Roles []recruitermodel.Role `json:"Roles,omitempty"`
	jwt.StandardClaims
}

// RefreshClaims represents the claims parsed from JWT refresh token.
type RefreshClaims struct {
	ID        string `json:"ID,omitempty"`
	TokenUUID string `json:"TokenUUID,omitempty"`
	jwt.StandardClaims
}

// ParseClaims parses JWT claims into AppClaims.
func (c *AppClaims) ParseClaims(claims jwt.MapClaims) error {
	// parse ID
	id, ok := claims["ID"]
	if !ok {
		return errors.New("could not parse claim id")
	}
	c.ID = id.(string)

	// parse Sub
	sub, ok := claims["Sub"]
	if !ok {
		return errors.New("could not parse claim sub")
	}
	c.Sub = sub.(string)

	// parse Roles
	rl, ok := claims["Roles"]
	if !ok {
		return errors.New("could not parse claims roles")
	}
	var roles []recruitermodel.Role
	if rl != nil {
		for _, v := range rl.([]interface{}) {
			r := v.(string)
			roles = append(roles, recruitermodel.Role(r))
		}
	}
	c.Roles = roles

	return nil
}

// ParseClaims parses the JWT claims into RefreshClaims.
func (c *RefreshClaims) ParseClaims(claims jwt.MapClaims) error {
	// parse ID
	id, ok := claims["ID"]
	if !ok {
		return errors.New("could not parse claim id")
	}
	c.ID = id.(string)

	// parse Token
	token, ok := claims["TokenUUID"]
	if !ok {
		return errors.New("could not parse token uuid")
	}
	c.TokenUUID = token.(string)
	return nil
}
