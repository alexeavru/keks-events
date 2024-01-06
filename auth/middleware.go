package auth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/alexeavru/keks-events/users"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type User struct {
	UserID   string `json:"userid"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Skip check bearer token for /login page
			if r.RequestURI == "/login" {
				next.ServeHTTP(w, r)
				return
			}

			header := r.Header.Get("Authorization")

			// Allow unauthenticated users in
			if header == "" {
				http.Error(w, "Authorization token must be present", http.StatusForbidden)
				return
			}

			//validate jwt token
			tokenStr := header
			username, err := ValidateJWT(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// Put user in context from token
			user := User{}
			// Тут нужно добавить выборку ID юзера из базы по username из токена
			user.UserID = "1"
			user.Username = username.(User).Username

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, &user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func getTokenSecretKey() []byte {
	// Get the TOKEN_SECRET environment variable
	tokenSecret, exists := os.LookupEnv("TOKEN_SECRET")

	if exists {
		return []byte(tokenSecret)
	} else {
		log.Fatal("Check .env file. Env TOKEN_SECRET not found!")
	}
	return nil
}

func CreateTokenEndpoint(response http.ResponseWriter, request *http.Request) {

	header := request.Header.Get("Authorization")

	// Allow unauthenticated users in
	if header == "" {
		http.Error(response, "Authorization token must be present", http.StatusForbidden)
		return
	}

	splitToken := strings.Split(header, "Basic ")
	if len(splitToken) != 2 {
		http.Error(response, "Authorization token not in proper format", http.StatusForbidden)
		return
	}

	// Decode Basic Auth
	data, err := base64.StdEncoding.DecodeString(splitToken[1])
	if err != nil {
		log.Fatal("error:", err)
	}

	splitBasicAuth := strings.Split(string(data), ":")

	// Check user password
	result, _ := users.CheckPassword(splitBasicAuth[0], splitBasicAuth[1])

	if !result {
		http.Error(response, "Error login or password", http.StatusForbidden)
		return
	}

	var user User
	user.Username = splitBasicAuth[0]

	_ = json.NewDecoder(request.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
	})
	tokenString, error := token.SignedString(getTokenSecretKey())
	if error != nil {
		fmt.Println(error)
	}
	response.Header().Set("content-type", "application/json")
	response.Write([]byte(`{ "token": "` + tokenString + `" }`))
}

func ValidateJWT(t string) (interface{}, error) {
	if t == "" {
		return nil, errors.New("authorization token must be present")
	}
	splitToken := strings.Split(t, "Bearer ")
	if len(splitToken) != 2 {
		return nil, errors.New("bearer token not in proper format")
	}
	token, _ := jwt.Parse(splitToken[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return getTokenSecretKey(), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var decodedToken User
		mapstructure.Decode(claims, &decodedToken)
		return decodedToken, nil
	} else {
		return nil, errors.New("invalid authorization token")
	}
}

func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}
