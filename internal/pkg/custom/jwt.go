package custom

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	formsEntity "github.com/fiber-go-sis-app/internal/app/constant"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var privateKey *rsa.PrivateKey

func GenerateJWT(embedPrivatePEMFile []byte) error {
	var err error

	privateKeyDecode, _ := pem.Decode(embedPrivatePEMFile)
	privateKey, _ = x509.ParsePKCS1PrivateKey(privateKeyDecode.Bytes)
	return err
}

func GetPrivateKey() *rsa.PrivateKey {
	return privateKey
}

func CreateJWTToken(req formsEntity.JWTRequest) (formsEntity.JWTTokenKey, error) {
	var (
		err         error
		jwtTokenKey formsEntity.JWTTokenKey
	)

	// Create access token
	accessTokenClaims := jwt.MapClaims{
		"user_id":  req.UserID,
		"name":     req.Name,
		"is_admin": req.IsAdmin,
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodRS256, accessTokenClaims)
	jwtTokenKey.AccessToken, err = accessToken.SignedString(privateKey)
	if err != nil {
		return formsEntity.JWTTokenKey{}, err
	}

	// Create refresh token
	refreshTokenClaims := jwt.MapClaims{
		"user_id": req.UserID,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodRS256, refreshTokenClaims)
	jwtTokenKey.RefreshToken, err = refreshToken.SignedString(privateKey)
	if err != nil {
		return formsEntity.JWTTokenKey{}, err
	}

	return jwtTokenKey, nil
}
