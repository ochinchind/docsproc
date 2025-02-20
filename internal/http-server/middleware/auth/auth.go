package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/ochinchind/docsproc/internal/usecase"
	"log"
	"time"
)

func Auth(rd *redis.Client) gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "dto does not contain an access token"})
			context.Abort()
			return
		}

		// Check if it starts with "Bearer " and remove it
		const bearerPrefix = "Bearer "
		if len(tokenString) > len(bearerPrefix) && tokenString[:len(bearerPrefix)] == bearerPrefix {
			tokenString = tokenString[len(bearerPrefix):] // Extract the actual token
		}

		// Check if the token is blacklisted
		blacklistKey := fmt.Sprintf("blacklist_%s", tokenString)
		blacklisted, err := rd.Get(blacklistKey).Result()
		if blacklisted != "" {
			context.JSON(401, gin.H{"error": "Unauthenticated"})
			context.Abort()
			return
		}

		claims, err := usecase.ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		go func() {
			cacheKey := fmt.Sprintf("user_id_%d", claims.UserID)

			// Fetch existing tokens from Redis
			result, err := rd.Get(cacheKey).Result()
			var tokens []string

			if errors.Is(err, redis.Nil) {
				// Key does not exist, create a new slice with the token
				tokens = []string{tokenString}
			} else if err != nil {
				log.Println("Error fetching from Redis:", err)
				return
			} else {
				// Deserialize JSON array from Redis
				err = json.Unmarshal([]byte(result), &tokens)
				if err != nil {
					log.Println("Error unmarshalling JSON:", err)
					return
				}

				// Check if the token already exists
				for _, t := range tokens {
					if t == tokenString {
						return // Token already exists, no need to append
					}
				}

				// Append the new token
				tokens = append(tokens, tokenString)
			}

			// Serialize back to JSON and store in Redis with an expiry of 24 hours
			tokenBytes, _ := json.Marshal(tokens)
			rd.Set(cacheKey, tokenBytes, 24*time.Hour)
		}()

		context.Set("auth_user_id", claims.UserID)
		context.Set("auth_user_role", claims.Role)
		context.Set("auth_user_email", claims.Email)

		context.Next()
	}
}
