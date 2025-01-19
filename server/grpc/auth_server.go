// Copyright 2023 Woodpecker Authors
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

package grpc

import (
	"context"
	"errors"
	"fmt"

	"github.com/crowci/crow/v3/pipeline/rpc/proto"
	"github.com/crowci/crow/v3/server/model"
	"github.com/crowci/crow/v3/server/store"
	"github.com/crowci/crow/v3/server/store/types"
	"github.com/rs/zerolog/log"
)

type CrowAuthServer struct {
	proto.UnimplementedCrowAuthServer
	jwtManager       *JWTManager
	agentMasterToken string
	store            store.Store
}

func NewCrowAuthServer(jwtManager *JWTManager, agentMasterToken string, store store.Store) *CrowAuthServer {
	return &CrowAuthServer{jwtManager: jwtManager, agentMasterToken: agentMasterToken, store: store}
}

func (s *CrowAuthServer) Auth(_ context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {
	agent, err := s.getAgent(req.AgentId, req.AgentToken)
	if err != nil {
		return nil, fmt.Errorf("agent could not auth: %w", err)
	}

	accessToken, err := s.jwtManager.Generate(agent.ID)
	if err != nil {
		return nil, err
	}

	return &proto.AuthResponse{
		Status:      "ok",
		AgentId:     agent.ID,
		AccessToken: accessToken,
	}, nil
}

func (s *CrowAuthServer) getAgent(agentID int64, agentToken string) (*model.Agent, error) {
	// global agent secret auth
	if s.agentMasterToken != "" {
		if agentToken == s.agentMasterToken && agentID == -1 {
			agent := &model.Agent{
				OwnerID:  model.IDNotSet,
				OrgID:    model.IDNotSet,
				Token:    s.agentMasterToken,
				Capacity: -1,
			}
			err := s.store.AgentCreate(agent)
			if err != nil {
				log.Error().Err(err).Msg("error creating system agent")
				return nil, err
			}
			return agent, nil
		}

		if agentToken == s.agentMasterToken {
			agent, err := s.store.AgentFind(agentID)
			if err != nil && errors.Is(err, types.RecordNotExist) {
				return nil, fmt.Errorf("AgentID not found in database")
			}
			return agent, err
		}
	}

	// individual agent token auth
	agent, err := s.store.AgentFindByToken(agentToken)
	if err != nil && errors.Is(err, types.RecordNotExist) {
		return nil, fmt.Errorf("individual agent not found by token: %w", err)
	}
	return agent, err
}
