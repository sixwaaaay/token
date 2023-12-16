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

package token

import (
	"context"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

// KeyFunc is a function that returns a jwt.Keyfunc, which is a function that takes a JWT token and returns the secret used to sign the token and an error.
// It takes a secret as a parameter, which is a byte slice.
// The returned jwt.Keyfunc is a function that takes a JWT token as a parameter and returns two values:
//   - The first return value is the secret used to sign the token. This is the same as the secret provided to KeyFunc.
//   - The second return value is an error. Since the function always returns nil for the error, this indicates that the function never fails.
func KeyFunc(secret []byte) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	}
}

// ExToken is a generic function that extracts a claim from a JWT token.
// It takes a JWT token and a key as parameters.
// The function returns two values:
//   - The first return value is the claim associated with the provided key.
//     The type of this value is determined by the type parameter T.
//   - The second return value is a boolean that indicates whether the operation was successful.
//     If the claim does not exist or cannot be cast to the type T, the function returns false.
func ExToken[T any](token *jwt.Token, key string) (T, bool) {
	// Declare a variable of type T
	var t T
	// Attempt to cast the token's claims to a MapClaims
	m, ok := token.Claims.(jwt.MapClaims)
	// If the cast was not successful, return the zero value of T and false
	if !ok {
		return t, false
	}
	return ExClaim[T](ok, m, key, t)
}

// ExClaim is a generic function that extracts a claim from a MapClaims.
func ExClaim[T any](ok bool, m jwt.MapClaims, key string, t T) (T, bool) {
	// Attempt to retrieve the claim associated with the provided key
	v, ok := m[key]
	// If the claim does not exist, return the zero value of T and false
	if !ok {
		return t, false
	}
	// Attempt to cast the claim to the type T
	t, ok = v.(T)
	// Return the claim and whether the operation was successful
	return t, ok
}

// Token is a generic function read a *jwt.Token from context.
func Token(ctx context.Context) (*jwt.Token, bool) {
	// Attempt to retrieve the token from context
	t, ok := ctx.Value(CtxKeyTokenV1).(*jwt.Token)
	// Return the token and whether the operation was successful
	return t, ok
}

// ClaimStrI64 is a function that extracts a claim from a JWT token and converts it to an int64.
// It takes a context and a key as parameters.
// The function returns two values:
//   - The first return value is the claim associated with the provided key, converted to an int64.
//     If the claim does not exist, cannot be cast to a string, or cannot be parsed as an int64, the function returns 0.
//   - The second return value is a boolean that indicates whether the operation was successful.
//     If the claim does not exist, cannot be cast to a string, or cannot be parsed as an int64, the function returns false.
func ClaimStrI64(ctx context.Context, key string) (int64, bool) {
	// Attempt to retrieve the token from context
	token, ok := Token(ctx)
	// If the token does not exist, return 0 and false
	if !ok {
		return 0, false
	}
	// Attempt to extract the claim associated with the provided key from the token
	exToken, b := ExToken[string](token, key)
	// If the claim does not exist or cannot be cast to a string, return 0 and false
	if !b {
		return 0, false
	}
	// Attempt to parse the claim as an int64
	t, err := strconv.ParseInt(exToken, 10, 64)
	// If the claim cannot be parsed as an int64, return 0 and false
	if err != nil {
		return 0, false
	}
	// Return the claim and true to indicate that the operation was successful
	return t, true
}

// ParseClaims is a function that parses a JWT token and adds it to a context.
// It takes three parameters:
//   - A context, which is the context to which the parsed token will be added.
//   - A token, which is a string representation of the JWT token to be parsed.
//   - A secret, which is a byte slice that represents the secret used to sign the token.
//
// The function returns two values:
//   - The first return value is a new context that is the same as the provided context, but with the parsed token added.
//     If the token cannot be parsed or is not valid, the function returns the original context.
//   - The second return value is an error that indicates whether the operation was successful.
//     If the token cannot be parsed or is not valid, the function returns an error.
func ParseClaims(ctx context.Context, token string, secret []byte) (context.Context, error) {
	t, err := jwt.Parse(token, KeyFunc(secret))
	if err != nil {
		return ctx, err
	}
	if !t.Valid {
		return ctx, jwt.ErrSignatureInvalid
	}
	return context.WithValue(ctx, CtxKeyTokenV1, t), nil
}
