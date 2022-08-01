/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a simple gRPC client that demonstrates how to use gRPC-Go libraries
// to perform unary, client streaming, server streaming and full duplex RPCs.
//
// It interacts with the route guide service whose definition can be found in routeguide/route_guide.proto.
package main

import (
	"back-web/cache"
	"back-web/controlor"
	"back-web/middleware"

	"github.com/gin-gonic/gin"
)

// CollectRoute include all route that implemted, seprate three parts of Admin or User or Reviewer.
func CollectRoute(r *gin.Engine, foreIP string) *gin.Engine {
	admin := r.Group("/").Use(middleware.CORSMiddleware(foreIP))
	{
		admin.GET("/impact", controlor.ImpactEffect)
	}

	return r
}

func main() {
	redisClient := cache.NewRedisConnection()
	defer redisClient.Close()

	gin.SetMode(gin.DebugMode)
	r := CollectRoute(gin.New(), "http://127.0.0.1:9999")

	port := "50052"
	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run())
}
