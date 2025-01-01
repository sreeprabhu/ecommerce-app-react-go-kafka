package helper

import (
	"errors"
	"fmt"
	"go-react-ecommerce-app/internal/domain"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Secret string
}

/*
*
Instantiates the Auth struct so that we can use the receiver functions below
*/
func SetupAuth(s string) Auth {
	return Auth{
		Secret: s,
	}
}

func (a Auth) CreateHashedPassword(password string) (string, error) {

	if len(password) < 6 {
		return "", errors.New("password length should be at least 6 characters long")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		// log actual error and report to a logging tool later
		return "", errors.New("password hash failed")
	}

	return string(hashPassword), nil
}

func (a Auth) GenerateToken(id uint, email string, role string) (string, error) {

	if id == 0 || email == "" || role == "" {
		return "", errors.New("required inputs are missing to generare token")
	}

	// generate token - go get github.com/golang-jwt/jwt/v4

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(a.Secret))

	if err != nil {
		// log actual error and report to a logging tool later
		return "", errors.New("unable to sign the token")
	}

	return tokenStr, nil
}

func (a Auth) VerifyPassword(plainPassword string, hashPassword string) error {

	if len(plainPassword) < 6 {
		return errors.New("password length should be at least 6 characters long")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(plainPassword))

	if err != nil {
		return errors.New("password does not match")
	}

	return nil
}

func (a Auth) VerifyToken(token string) (domain.User, error) {
	// Bearer df12dfsd3dsf124dsfds3dfs25345345
	tokenArr := strings.Split(token, " ")
	if len(tokenArr) != 2 {
		return domain.User{}, nil
	}

	tokenStr := tokenArr[1]

	if tokenArr[0] != "Bearer" {
		return domain.User{}, errors.New("invalid token")
	}

	// Parse the token
	parsedToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header)
		}

		// Return the secret key for verification
		return []byte(a.Secret), nil
	})

	if err != nil {
		return domain.User{}, errors.New("invalid signing method")
	}

	// Verify token validity
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {

		// Safely access and validate the "exp" claim
		exp, ok := claims["exp"].(float64)
		if !ok {
			return domain.User{}, errors.New("invalid or missing 'exp' claim")
		}
		if float64(time.Now().Unix()) > exp {
			return domain.User{}, errors.New("token is expired")
		}

		fmt.Printf("===== %v", exp)

		// Initialize the user object
		user := domain.User{}

		// Safely access and convert "user_id"
		userId, ok := claims["user_id"].(float64)
		if !ok {
			return domain.User{}, errors.New("invalid or missing 'user_id' claim")
		}
		user.ID = uint(userId)

		fmt.Printf("===== %v", user.ID)

		// Safely access and convert "email"
		email, ok := claims["email"].(string)
		if !ok {
			return domain.User{}, errors.New("invalid or missing 'email' claim")
		}
		user.Email = email

		fmt.Printf("===== Email %v ======", user.Email)

		// Safely access and convert "user_type"
		userType, ok := claims["role"].(string)
		if !ok {
			return domain.User{}, errors.New("invalid or missing 'user_type' claim")
		}
		user.UserType = userType

		fmt.Printf("===== UserType %v =====", user.UserType)

		// user.ID = uint(claims["user_id"].(float64))
		// user.Email = claims["email"].(string)
		// user.UserType = claims["user_type"].(string)

		return user, nil
	}

	return domain.User{}, errors.New("token verification failed")
}

func (a Auth) Authorize(ctx *fiber.Ctx) error {

	authHeaders := ctx.GetReqHeaders()["Authorization"]

	if len(authHeaders) == 0 {
		// Handle missing Authorization header
		return errors.New("missing Bearer token")
	}
	// Extract the first header value
	authHeader := authHeaders[0]
	user, err := a.VerifyToken(authHeader)

	if err == nil && user.ID > 0 {
		// setting the user to the context object, so that we can grab this user details in future
		ctx.Locals("user", user)

		return ctx.Next()
	} else {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": "authorization failed",
			"reason":  err,
		})
	}
}

func (a Auth) CurrentUser(ctx *fiber.Ctx) domain.User {

	user := ctx.Locals("user")

	// cast the object to domain.User{}
	return user.(domain.User)
}
