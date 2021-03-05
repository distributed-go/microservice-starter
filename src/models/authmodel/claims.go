package authmodel

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/jobbox-tech/recruiter-api/models/recruitermodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AppClaims represent the claims parsed from JWT access token.
type AppClaims struct {
	ID    primitive.ObjectID    `json:"ID,omitempty"`
	Sub   string                `json:"Sub,omitempty"`
	Roles []recruitermodel.Role `json:"Roles,omitempty"`
	jwt.StandardClaims
}

// RefreshClaims represents the claims parsed from JWT refresh token.
type RefreshClaims struct {
	ID    primitive.ObjectID `json:"id,omitempty"`
	Token string             `json:"token,omitempty"`
	jwt.StandardClaims
}

// ParseClaims parses JWT claims into AppClaims.
func (c *AppClaims) ParseClaims(claims jwt.MapClaims) error {
	id, ok := claims["id"]
	if !ok {
		return errors.New("could not parse claim id")
	}
	c.ID = id.(primitive.ObjectID)

	sub, ok := claims["sub"]
	if !ok {
		return errors.New("could not parse claim sub")
	}
	c.Sub = sub.(string)

	rl, ok := claims["roles"]
	if !ok {
		return errors.New("could not parse claims roles")
	}

	var roles []recruitermodel.Role
	if rl != nil {
		for _, v := range rl.([]interface{}) {
			roles = append(roles, v.(recruitermodel.Role))
		}
	}
	c.Roles = roles

	return nil
}

// ParseClaims parses the JWT claims into RefreshClaims.
func (c *RefreshClaims) ParseClaims(claims jwt.MapClaims) error {
	token, ok := claims["token"]
	if !ok {
		return errors.New("could not parse claim token")
	}
	c.Token = token.(string)
	return nil
}
