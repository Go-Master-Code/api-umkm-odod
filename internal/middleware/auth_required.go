package middleware

import (
	"context"
	"net/http"
	"strings"
	"umkm-odod/auth"
	"umkm-odod/internal/constants"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ambil token dari header authorization
		tokenString := c.GetHeader("Authorization")
		// validasi header bearer token
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not found / invalid token"})
			c.Abort()
			return
		}

		// menghilangkan suatu prefix / suatu string dari suatu string utuh
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// validasi jwt token
		token, err := auth.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort() // abort berguna agar handler selanjutnya tidak dieksekusi jika token tidak valid
			// contoh: handler GetItemById() tidak boleh dijalankan jika token tidak valid
			return
		}

		// tambahan untuk RBAC (Role Based Access Control)
		/*
			Masalah utama:
			Saat token sudah valid, kita harus menyimpan informasi claims ke dalam context.
			RBAC (dan handler) perlu akses ke role user di dalam JWT.
		*/

		/*
			Jika token valid:
			- ambil claims JWT
			- simpan ke gin context (opsional)
			- inject ke request context (BEST PRACTICE)
		*/
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// ===== OPTIONAL =====
			// tetap simpan ke gin context
			// agar middleware RBAC masih bisa pakai c.Get()
			c.Set("claims", claims) // bisa akses "role", "username", dsb dari middleware lain
			// simpan username ke context untuk simpan log ke database
			if username, exists := claims["username"].(string); exists {
				c.Set("username", username)
			}

			// ===== BEST PRACTICE =====
			// inject claims ke request context
			// supaya bisa diakses di service layer
			ctx := c.Request.Context()

			// inject tenant_id
			if tenantID, exists := claims["tenant_id"].(string); exists {
				ctx = context.WithValue(
					/*
						simpan tenantID ke context
						dengan key ContextTenantID
					*/
					ctx,
					constants.ContextTenantID,
					tenantID,
				)
			}

			// inject user_id
			if userID, exists := claims["user_id"].(string); exists {
				ctx = context.WithValue(
					ctx,
					constants.ContextUserID,
					userID,
				)
			}

			// inject role
			if role, exists := claims["role"].(string); exists {
				ctx = context.WithValue(
					ctx,
					constants.ContextRole,
					role,
				)
			}

			// inject username
			if username, exists := claims["username"].(string); exists {
				ctx = context.WithValue(
					ctx,
					constants.ContextUsername,
					username,
				)
			}

			// replace request context lama dengan context baru
			c.Request = c.Request.WithContext(ctx)

			// lanjut ke handler berikutnya
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// jika token valid
		c.Next()
	}
}
