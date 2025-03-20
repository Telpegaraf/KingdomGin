package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func CheckWebAppSignatureMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		xInitData := c.GetHeader("x-initData")
		if xInitData == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Missing init data"})
			return
		}

		parsedData, err := parseInitData(xInitData)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid init data format"})
			return
		}

		hash, ok := parsedData["hash"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Hash not found in init data"})
			return
		}

		delete(parsedData, "hash")
		dataCheckString := createDataCheckString(parsedData)
		secretKey := hmac.New(sha256.New, []byte("WebAppData"))
		secretKey.Write([]byte(os.Getenv("BOT_API_TOKEN")))
		calculatedHash := hmac.New(sha256.New, secretKey.Sum(nil))
		calculatedHash.Write([]byte(dataCheckString))
		expectedHash := hex.EncodeToString(calculatedHash.Sum(nil))

		if !hmac.Equal([]byte(expectedHash), []byte(hash)) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Invalid signature"})
			return
		}

		c.Next()
	}
}

func parseInitData(initData string) (map[string]string, error) {
	parsedData := make(map[string]string)
	pairs := strings.Split(initData, "&")
	for _, pair := range pairs {
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) != 2 {
			return nil, errors.New("invalid format")
		}
		parsedData[kv[0]] = kv[1]
	}
	return parsedData, nil
}

func createDataCheckString(data map[string]string) string {
	var dataCheckString strings.Builder
	for k, v := range data {
		dataCheckString.WriteString(k + "=" + v + "\n")
	}
	return strings.TrimSuffix(dataCheckString.String(), "\n")
}
