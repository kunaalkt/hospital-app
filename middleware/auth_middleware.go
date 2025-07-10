package middleware

import (
    "net/http"
    "strings"
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        h := c.GetHeader("Authorization")
        if h == "" || !strings.HasPrefix(h, "Bearer ") {
            fmt.Println("Missing or malformed Authorization header")
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            return
        }
        tokenStr := strings.TrimPrefix(h, "Bearer ")
        token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
            return []byte("supersecretkey"), nil
        })
        if err != nil || !token.Valid {
            fmt.Println("Invalid token:", err)
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            fmt.Println("Error parsing token claims")
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
            return
        }

        c.Set("claims", claims)
        c.Next()
    }
}


func RoleMiddleware(requiredRole string) gin.HandlerFunc {
    return func(c *gin.Context) {
        claims, exists := c.Get("claims")
        if !exists {
            fmt.Println("Claims are missing")
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing claims"})
            return
        }

        role, ok := claims.(jwt.MapClaims)["role"].(string)
        if !ok {
            fmt.Println("Role is missing or invalid")
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
            return
        }

        if role != requiredRole {
            fmt.Println("Role mismatch:", role)
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
            return
        }

        c.Next()
    }
}