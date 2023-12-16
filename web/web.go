/*
 * Copyright (c) 2023 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package web

import (
	"net/http"
	"strings"

	"github.com/sixwaaaay/token"
)

// Middleware is a function that takes a secret key as a parameter and returns a new function.
// This new function takes an http.Handler and returns another http.Handler that wraps the provided http.Handler.
// The returned http.Handler checks for a JWT token in the Authorization header of the HTTP request.
// If a token is found, it is parsed and the claims are added to the request context.
// If no token is found, the request is passed through to the next handler without modification.
// If the token is invalid, a 401 Unauthorized response is returned.
//
// The secret parameter is used to validate the JWT token.
//
// This function typically is used in an HTTP server to protect routes that require authentication.
func Middleware(secret []byte) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// Get the JWT token from the request header
			tokenString := r.Header.Get("Authorization")
			// If the token is empty, let the application handle it
			if tokenString == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Remove the "Bearer " prefix from the token string
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
			// Parse the claims from the token
			c, err := token.ParseClaims(r.Context(), tokenString, secret)
			// If the token is invalid, return a 401 Unauthorized response
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r.WithContext(c))
		})
	}
}
