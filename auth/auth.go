package auth


import (
  "strings"
	"time"
  "context"

  "github.com/gin-gonic/gin"
  "github.com/dgrijalva/jwt-go"
)



var (
  secretKey = "secretekey"
)


func GenerateToken(username string) (string, error) {
  token := jwt.New(jwt.SigningMethodHS256)
  	/* Create a map to store our claims */
  claims := token.Claims.(jwt.MapClaims)
  claims["username"] = username
  claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
  return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenStr string) (string, error) {
  token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    username := claims["username"].(string)
    return username,nil
  }
  return "", err
}


func Authenticate() gin.HandlerFunc {

  return func (c *gin.Context)  {
     header := c.Request.Header.Get("Authorization")
     if header == "" {
       c.AbortWithStatusJSON(401,gin.H{"error": "Login required"})
       return
     }
    headerArr := strings.Split(header," ")
    if len(headerArr) < 2 {
      c.AbortWithStatusJSON(400,gin.H{"error": "Bad request"})
      return
    }

    name, err := ParseToken(headerArr[1])

    if err != nil {
      c.AbortWithStatusJSON(400,gin.H{"error": "token not found"})
      return
    }
    c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(),"user",name))
    c.Next()
  }
}
