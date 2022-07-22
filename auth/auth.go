package auth

import (
	"time"

	jwt "github.com/form3tech-oss/jwt-go"
)

func GenerateToken() string {
	//headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	//claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "123456789"
	claims["name"] = "yamato"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	//電子著名
	tokenString, _ := token.SignedString([]byte("sercret"))

	//JWTを返却
	return tokenString
}

//jwtmiddleware check token
// func checkToken() {
// 	var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
// 		ValidationKeyGetter: func(token *jwt.Token)(interface{}, error){
// 			return []byte("sercret"),nil
// 		},
// 		SigningMethod: jwt.SigningMethodHS256,
// 	})
// }
