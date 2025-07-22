package utils

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// 	"os"
// 	"sonit_server/constant/env"
// 	"sonit_server/constant/noti"
// 	"strings"
// 	"time"
// 	"unicode"

// 	"github.com/golang-jwt/jwt/v5"
// )

// // Generate tokens for login action
// func GenerateTokens(email, userId, role string, logger *log.Logger) (string, string, error) {
// 	var bytes = []byte(os.Getenv(env.SECRET_KEY))
// 	var errMsg string = "Error while generating tokens - "

// 	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id": userId,
// 		"email":   email,
// 		"role":    role,
// 		"expire":  time.Now().Add(AccessDuration).Unix(),
// 	}).SignedString(bytes)
// 	if err != nil {
// 		logger.Print(errMsg + fmt.Sprint(err))
// 		return "", "", errors.New(noti.INTERNALL_ERR_MSG)
// 	}

// 	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id": userId,
// 		"email":   email,
// 		"role":    role,
// 		"expire":  time.Now().Add(RefreshDuration).Unix(),
// 	}).SignedString(bytes)
// 	if err != nil {
// 		logger.Print(errMsg + fmt.Sprint(err))
// 		return "", "", errors.New(noti.INTERNALL_ERR_MSG)
// 	}

// 	return accessToken, refreshToken, nil
// }

// // Generate token for actions such as: register, verify, activate, ...
// func GenerateActionToken(email, userId, role string, logger *log.Logger) (string, error) {
// 	var bytes = []byte(os.Getenv(env.SECRET_KEY))
// 	var errMsg string = "Error while generating action token - "

// 	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id": userId,
// 		"email":   email,
// 		"role":    role,
// 		"expire":  time.Now().Add(NormalActionDuration).Unix(),
// 	}).SignedString(bytes)
// 	if err != nil {
// 		logger.Print(errMsg + fmt.Sprint(err))
// 		return "", errors.New(noti.INTERNALL_ERR_MSG)
// 	}

// 	return token, nil
// }

// func ExtractDataFromToken(tokenString string, logger *log.Logger) (string, string, time.Time, error) {
// 	var errRes error = errors.New(noti.GENERIC_ERROR_WARN_MSG)
// 	var errLogMsg string = "Error at ExtractDataFromToken - "

// 	// Check for empty token
// 	if tokenString == "" {
// 		logger.Println(errLogMsg + "empty token")
// 		return "", "", time.Time{}, errRes
// 	}

// 	// Remove "Bearer " prefix if present
// 	if len(tokenString) > 7 && strings.HasPrefix(tokenString, "Bearer ") {
// 		tokenString = strings.TrimSpace(tokenString[7:])
// 	}

// 	// Trim any whitespace and non-printable characters
// 	tokenString = strings.TrimSpace(tokenString)

// 	// Check if token contains invalid characters
// 	for i, c := range tokenString {
// 		if !unicode.IsPrint(c) || c == ' ' {
// 			logger.Printf(errLogMsg+"invalid character at position %d: %q\n", i, c)
// 			return "", "", time.Time{}, errRes
// 		}
// 	}

// 	// Parse the token
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Validate signing method
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		secretKey := os.Getenv(env.SECRET_KEY)
// 		if secretKey == "" {
// 			logger.Println(errLogMsg + "secret key not found in environment")
// 			return nil, errors.New("missing secret key")
// 		}

// 		return []byte(secretKey), nil
// 	})

// 	if err != nil {
// 		logger.Println(errLogMsg + err.Error())
// 		return "", "", time.Time{}, errRes
// 	}

// 	// Extract claims
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		logger.Println(errLogMsg + "invalid claims or token")
// 		return "", "", time.Time{}, errRes
// 	}

// 	// Extract user_id
// 	userId, ok := claims["user_id"].(string)
// 	if !ok || userId == "" {
// 		logger.Println(errLogMsg + "missing or invalid user_id claim")
// 		return "", "", time.Time{}, errRes
// 	}

// 	// Extract role
// 	role, ok := claims["role"].(string)
// 	if !ok || role == "" {
// 		logger.Println(errLogMsg + "missing or invalid role claim")
// 		return "", "", time.Time{}, errRes
// 	}

// 	// Extract expiration
// 	expFloat, ok := claims["expire"].(float64)
// 	if !ok {
// 		logger.Println(errLogMsg + "missing or invalid expiration claim")
// 		return "", "", time.Time{}, errRes
// 	}

// 	// Convert Unix timestamp to time.Time
// 	exp := time.Unix(int64(expFloat), 0)

// 	return userId, role, exp, nil
// }
