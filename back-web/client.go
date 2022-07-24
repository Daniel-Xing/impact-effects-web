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
	"context"
	"flag"
	"log"
	"time"

	"back-web/google.golang.org/grpc/impactEffect/impactEffect"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/examples/data"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

//
func printEnergy(client impactEffect.ImpactEffectServiceClient, req *impactEffect.Cal_KineticEnergyRequest) {
	log.Printf("Getting feature for point (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feature, err := client.Cal_KineticEnergy(ctx, req)
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}
	log.Println(feature)
}

func printIFactor(client impactEffect.ImpactEffectServiceClient, req *impactEffect.CalIFactorRequest) {
	log.Printf("Getting feature for point (%f, %f)", req.Impactor.Diameter, req.Impactor.Density)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	feature, err := client.CalIFactor(ctx, req)
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}
	log.Println(feature)
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = data.Path("x509/ca_cert.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := impactEffect.NewImpactEffectServiceClient(conn)

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

	// printEnergy(client, req)
	printIFactor(client, req)
}
