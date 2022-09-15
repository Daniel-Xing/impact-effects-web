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
		admin.POST("/impact", controlor.ImpactEffect)
		admin.POST("/simulator", controlor.SimulatorImpact)
		admin.POST("/simulatorWithRedis", controlor.SimulatorImpactWithRedis)
		admin.POST("/simulatorWithRedis2", controlor.SimulatorImpactWithRedis2)
	}

	return r
}

func main() {
	redisClient := cache.NewRedisConnection()
	defer redisClient.Close()

	gin.SetMode(gin.DebugMode)
	r := CollectRoute(gin.New(), "http://0.0.0.0:9999")

	port := "50052"
	if port != "" {
		panic(r.Run(":" + port))
	}

	panic(r.Run())
}
