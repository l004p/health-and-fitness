package middleware

import (
	"fmt"
	//"log"
	"net/http"
	//"server/core/user"
	"server/services/authentication"
	"context"
	"server/services/user"
	//"context"
	"log"
	"github.com/99designs/gqlgen/graphql"
	"server/api/graph/model"
	//"server/api/rest"
	"github.com/vektah/gqlparser/v2/gqlerror"
	//"server/core/user"
)

type CtxKey string

func (mh *MiddlewareHandler) AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//verify token

			header := r.Header.Get("Authorization")
			if header == "" {
				http.Error(w, "no header", http.StatusForbidden)
				return
			}

			bearer := "Bearer "
			tokenString := header[len(bearer):]

			token, err := authentication.ValidateToken(tokenString)
			if err != nil || !token.Valid {
				http.Error(w, "error validating token", http.StatusForbidden)
				return
			}

			claim, _ := token.Claims.(*authentication.JwtCustomClaims)

			_, err = userservice.UserByID(mh.ur, r.Context(), claim.UserID)
			if err != nil {
				http.Error(w, "error with user", http.StatusForbidden)
				return
			}
			const userID CtxKey = "userID"
			var userIDValue string = claim.UserID
			ctx := context.WithValue(r.Context(), userID, userIDValue)

			r = r.WithContext(ctx)
			fmt.Println("Auth Middleware!")
			next.ServeHTTP(w, r)
		})
	}
}

func (mh *MiddlewareHandler) DirectiveMiddleware(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error){
	//fmt.Printf("hasRole Directive %s\n", role.String())
	//userservice.HasRole(repo, ctx, , role.String())
	log.Printf("hasRole Directive %s\n", role)
	getID := ctx.Value(CtxKey("userID")).(string)
	hasRole, err := mh.ur.HasRole(ctx, getID, role.String())
	if err != nil || !hasRole {
		return nil, &gqlerror.Error{
			Message: "Access Denied",
		}
	}
	return next(ctx)
}