// Copyright 2024 Woodpecker Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/crowci/crow/v3/pipeline/rpc/proto"
	"github.com/crowci/crow/v3/server"
	crowGrpcServer "github.com/crowci/crow/v3/server/grpc"
	"github.com/crowci/crow/v3/server/store"
)

func runGrpcServer(ctx context.Context, c *cli.Command, _store store.Store) error {
	lis, err := net.Listen("tcp", c.String("grpc-addr"))
	if err != nil {
		return fmt.Errorf("failed to listen on grpc-addr: %w", err)
	}

	jwtSecret := c.String("grpc-secret")
	jwtManager := crowGrpcServer.NewJWTManager(jwtSecret)

	authorizer := crowGrpcServer.NewAuthorizer(jwtManager)
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(authorizer.StreamInterceptor),
		grpc.UnaryInterceptor(authorizer.UnaryInterceptor),
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime: c.Duration("keepalive-min-time"),
		}),
	)

	crowServer := crowGrpcServer.NewCrowServer(
		server.Config.Services.Queue,
		server.Config.Services.Logs,
		server.Config.Services.Pubsub,
		_store,
	)
	proto.RegisterCrowServer(grpcServer, crowServer)

	crowAuthServer := crowGrpcServer.NewCrowAuthServer(
		jwtManager,
		server.Config.Server.AgentToken,
		_store,
	)
	proto.RegisterCrowAuthServer(grpcServer, crowAuthServer)

	grpcCtx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	go func() {
		<-grpcCtx.Done()
		if grpcServer == nil {
			return
		}
		log.Info().Msg("terminating grpc service gracefully")
		grpcServer.GracefulStop()
		log.Info().Msg("grpc service stopped")
	}()

	if err := grpcServer.Serve(lis); err != nil {
		// signal that we don't have to stop the server gracefully anymore
		grpcServer = nil

		// wrap the error so we know where it did come from
		return fmt.Errorf("grpc server failed: %w", err)
	}

	return nil
}
