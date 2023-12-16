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
	"net/http"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/sixwaaaay/token"
)

// Handler returns a UnaryServerInterceptor that checks for a JWT token in the
// incoming request's metadata. If a token is found, it is parsed and the claims
// are added to the context. If no token is found or the token is invalid, the
// request is passed through without modification.
//
// The secret parameter is used to validate the JWT token.
//
// This function is typically used as a gRPC server interceptor.
func Handler(secret []byte) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		incomingContext, b := metadata.FromIncomingContext(ctx)
		if !b || len(incomingContext["authorization"]) == 0 {
			return handler(ctx, req) // no jwt, let app handle it
		}
		tokenString := strings.TrimPrefix(incomingContext["authorization"][0], "Bearer ")
		c, err := token.ParseClaims(ctx, tokenString, secret)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		return handler(c, req)
	}
}

// Ctx4H extracts the Authorization header from the HTTP request and appends it
// to the outgoing context. If no Authorization header is found, the request's
// context is returned unmodified.
//
// This function is typically used to propagate HTTP headers to gRPC metadata.
func Ctx4H(r *http.Request) context.Context {
	header := r.Header.Get("Authorization")
	if header == "" {
		return r.Context()
	}
	return metadata.AppendToOutgoingContext(r.Context(), "Authorization", header)
}
