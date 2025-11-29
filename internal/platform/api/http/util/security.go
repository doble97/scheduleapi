package util

import (
	"fmt"
	"strconv"
	"time"

	"github.com/doble97/scheduleapi/config"
	"github.com/doble97/scheduleapi/internal/core/domain"
	"github.com/doble97/scheduleapi/internal/platform/api/http/dto"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(user domain.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := dto.CustomClaims{
		UserID:    user.Email,
		UserRole:  "user",
		UserEmail: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   strconv.FormatInt(int64(user.ID), 10),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.SecretKey))
	if err != nil {
		return "", fmt.Errorf("error al firmar el token %w", err)
	}
	return tokenString, nil
}

func VerifyJWT(tokenString string) (*dto.CustomClaims, error) {
	// El 'parser' toma el token en string y las claims en una nueva instancia de CustomClaims.
	jwtSecretKey := []byte(config.Config.SecretKey)
	token, err := jwt.ParseWithClaims(
		tokenString,
		&dto.CustomClaims{},
		// 1. Función para obtener la clave de verificación (keyFunc)
		// Esta función se ejecuta después del parsing inicial para obtener
		// la clave necesaria para validar la firma.
		func(token *jwt.Token) (interface{}, error) {
			// Confirmamos que el algoritmo usado en el token (del Header)
			// coincide con el esperado (por ejemplo, HS256).
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
			}
			// Retornamos la clave secreta (simétrica)
			return jwtSecretKey, nil
		},
	)
	if err != nil {
		// Manejo de errores comunes (expiración, firma inválida, etc.)
		return nil, fmt.Errorf("error al parsear o validar el token: %w", err)
	}

	// 2. Extracción y Verificación de Claims
	claims, ok := token.Claims.(*dto.CustomClaims)
	if !ok || !token.Valid {
		// El token no es válido o las claims no son del tipo esperado
		return nil, fmt.Errorf("token inválido o claims incorrectas")
	}

	// 3. Verificación de Claims Específicas (Opcional pero Recomendado)
	// Aquí puedes añadir verificaciones adicionales, como 'Issuer' o 'Audience'.
	// if claims.Issuer != "my-auth-service" { ... }

	return claims, nil
}
