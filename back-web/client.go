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
	"back-web/controlor"
	"back-web/google.golang.org/grpc/impactEffect/impactEffect"
	"back-web/rpc"
)

func main() {

	impactor := &impactEffect.Impactor{}
	impactor.Density = 111
	impactor.Diameter = 111
	impactor.Velocity = 111
	impactor.Theta = 45

	target := &impactEffect.Targets{}
	target.Density = 111
	target.Depth = 0
	target.Distance = 111

	req := &impactEffect.CalIFactorRequest{
		Impactor: impactor,
		Targets:  target,
		Choice:   1,
	}

	service := rpc.NewImpactEffectRPCService()
	service.CalIFactor(req)

	controlor.ImpactEffect()
}
