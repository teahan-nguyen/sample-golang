package utils

import (
	verifier "github.com/okta/okta-jwt-verifier-golang"
	"samples-golang/initializer"
)

func GetTokenPayload(tokenString string) (map[string]interface{}, error) {
	config, err := initializer.LoadConfig(".")
	if err != nil {
		return nil, err
	}

	tv := map[string]string{}
	tv["aud"] = "api://default"
	tv["cid"] = config.SpaClientId
	jv := verifier.JwtVerifier{
		Issuer:           config.Issuer,
		ClaimsToValidate: tv,
	}

	token, err := jv.New().VerifyAccessToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims := make(map[string]interface{})
	for key, value := range token.Claims {
		claims[key] = value
	}

	return claims, nil
}
