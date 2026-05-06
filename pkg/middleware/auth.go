package middleware

import(
	"net/http"
	"strings"

	"go-chat-app/pkg/utils"
	"go-chat-app/pkg/models"
)

func ExtractTokenFromHeader(r *http.Request) (string, error){
	authHeader := r.Header.Get("Authorization")
	if authHeader == ""{
		http.Error(w, `{"error":"header not found"}`, http.StatusUnauthorized)
		return  
	}
}

func AuthMiddleWare(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
		if authHeader == ""{
			http.Error(w, `{"error":"header not found"}`, http.StatusUnauthorized)
			return 
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer"{
			http.Error(w, `{"error":"invalid format for authorization header"}`, http.StatusUnauthorized)
			return 
		}

		claims, err := utils.VerifyToken(parts[0])
		if err!=nil{
			http.Error(w, `{"error":"invalid token"}`, http.StatusUnauthorized)
			return 
		}

		r.Header.Set("X-User-Id", string(rune(claims.UserID)))
		r.Header.Set("X-Username", claims.Username)
		r.Header.Set("X-Email", claims.Email)

		next.ServeHTTP(w,r)
	})
}