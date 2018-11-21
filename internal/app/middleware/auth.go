// Copyright 2018 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/clivern/beaver/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

// Auth middleware
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method

		if strings.Contains(path, "/api/") {
			authToken := c.GetHeader("X-AUTH-TOKEN")
			if authToken != os.Getenv("APIToken") {
				logger.Infof(
					"Unauthorized access to %s:%s with token %s",
					method,
					path,
					authToken,
				)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		}
	}
}
