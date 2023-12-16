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

package rpc

import (
	"context"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
)

func TestHandler_WithValidAuthorization(t *testing.T) {
	secret := []byte("secret")
	hand := Handler(secret)

	ctx := context.Background()
	md := metadata.Pairs("Authorization", "Bearer "+createToken(secret))
	ctx = metadata.NewIncomingContext(ctx, md)

	_, err := hand(ctx, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	})
	assert.NoError(t, err)
}

func TestHandler_WithInvalidAuthorization(t *testing.T) {
	secret := []byte("secret")
	handler := Handler(secret)

	ctx := context.Background()
	md := metadata.Pairs("Authorization", "Bearer "+createToken([]byte("wrong secret")))
	ctx = metadata.NewIncomingContext(ctx, md)

	_, err := handler(ctx, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	})

	assert.Error(t, err)
}

func TestHandler_WithoutAuthorization(t *testing.T) {
	secret := []byte("secret")
	handler := Handler(secret)

	ctx := context.Background()

	_, err := handler(ctx, nil, nil, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	})

	assert.NoError(t, err)
}

func createToken(secret []byte) string {
	token := jwt.New(jwt.SigningMethodHS256)
	ss, _ := token.SignedString(secret)
	return ss
}
