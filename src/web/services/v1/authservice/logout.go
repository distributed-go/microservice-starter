package authservice

import (
	"net/http"

	"github.com/go-chi/render"
)

func (as *authservice) Logout(w http.ResponseWriter, r *http.Request) {
	// rt := jwt.RefreshTokenFromCtx(r.Context())
	// token, err := as.Store.GetToken(rt)
	// if err != nil {
	// 	render.Render(w, r, ErrUnauthorized(jwt.ErrTokenExpired))
	// 	return
	// }
	// as.Store.DeleteToken(token)

	render.Respond(w, r, http.NoBody)
}
