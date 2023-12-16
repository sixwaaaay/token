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
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestParseClaims_ValidToken(t *testing.T) {
	secret := []byte("secret")
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString(secret)

	ctx, err := ParseClaims(context.Background(), tokenString, secret)

	assert.NoError(t, err)
	assert.NotNil(t, ctx)
}

func TestParseClaims_InvalidToken(t *testing.T) {
	secret := []byte("secret")
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString(secret)

	ctx, err := ParseClaims(context.Background(), tokenString, []byte("wrong secret"))

	assert.Error(t, err)
	assert.Equal(t, context.Background(), ctx)
}

func TestParseClaims_EmptyToken(t *testing.T) {
	secret := []byte("secret")

	ctx, err := ParseClaims(context.Background(), "", secret)

	assert.Error(t, err)
	assert.Equal(t, context.Background(), ctx)
}

func TestClaimStrI64_ValidClaim(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["claim"] = "12345"
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte("secret"))

	ctx := context.Background()
	ctx, err := ParseClaims(ctx, tokenString, []byte("secret"))
	assert.NoError(t, err)

	claim, ok := ClaimStrI64(ctx, "claim")

	assert.True(t, ok)
	assert.Equal(t, int64(12345), claim)
}

func TestClaimStrI64_ClaimNotInt(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["claim"] = "not an int"
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte("secret"))

	ctx := context.Background()
	ctx, err := ParseClaims(ctx, tokenString, []byte("secret"))
	assert.NoError(t, err)

	claim, ok := ClaimStrI64(ctx, "claim")

	assert.False(t, ok)
	assert.Equal(t, int64(0), claim)
}

func TestClaimStrI64_ClaimNotExist(t *testing.T) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString([]byte("secret"))

	ctx := context.Background()
	//ctx = context.WithValue(ctx, CtxKeyTokenV1, token)
	ctx, err := ParseClaims(ctx, tokenString, []byte("secret"))
	assert.NoError(t, err)

	claim, ok := ClaimStrI64(ctx, "claim")

	assert.False(t, ok)
	assert.Equal(t, int64(0), claim)
}

func TestClaimStrI64_NoToken(t *testing.T) {
	ctx := context.Background()

	claim, ok := ClaimStrI64(ctx, "claim")

	assert.False(t, ok)
	assert.Equal(t, int64(0), claim)
}
