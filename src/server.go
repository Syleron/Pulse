/*
    PulseHA - HA Cluster Daemon
    Copyright (C) 2017  Andrew Zak <andrew@pulseha.com>

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Affero General Public License as published
    by the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU Affero General Public License for more details.

    You should have received a copy of the GNU Affero General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"context"
	"github.com/Syleron/PulseHA/proto"
	"github.com/coreos/go-log/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"os"
	"path/filepath"
	"sync"
	"github.com/Syleron/PulseHA/src/utils"
)

/**
 * Server struct type
 */
type Server struct {
	sync.Mutex
	Status        proto.HealthCheckResponse_ServingStatus
	//Last_response time.Time
	//Log log.Logger
	Server *grpc.Server
	Listener net.Listener
	Memberlist *Memberlist
}

/**
	Perform appr. health checks
 */
func (s *Server) Check(ctx context.Context, in *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	s.Lock()
	defer s.Unlock()
	switch in.Request {
	case proto.HealthCheckRequest_SETUP:
		log.Debug("Server:Check() - HealthCheckRequest Setup")
	case proto.HealthCheckRequest_STATUS:
		log.Debug("Server:Check() - HealthCheckRequest Status")
		return &proto.HealthCheckResponse{
			Status: proto.HealthCheckResponse_CONFIGURED,
		}, nil
	default:
	}
	return nil, nil
}

/**
	Join request for a configured cluster
 */
func (s *Server) Join(ctx context.Context, in *proto.PulseJoin) (*proto.PulseJoin, error) {
	s.Lock()
	defer s.Unlock()
	log.Debug("join test")
	return nil, nil
}

/**
 * Setup pulse server type
 */
func (s *Server) Setup() {
	configCopy := gconf.GetConfig()
	if !clusterCheck() {
		log.Info("PulseHA is currently un-configured.")
		return
	}
	var err error
	s.Listener, err = net.Listen("tcp", configCopy.LocalNode().IP+":"+configCopy.LocalNode().Port)
	if err != nil {
		log.Errorf("Failed to listen: %s", err)
		os.Exit(1)
	}
	if configCopy.Pulse.TLS {
		// Get project directory location
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Emergency(err)
		}
		if utils.CreateFolder(dir + "/certs") {
			log.Warning("TLS keys are missing! Generating..")
			GenOpenSSL()
		}
		creds, err := credentials.NewServerTLSFromFile(dir+"/certs/server.crt", dir+"/certs/server.key")
		if err != nil {
			log.Error("Could not load TLS keys.")
			os.Exit(1)
		}
		s.Server = grpc.NewServer(grpc.Creds(creds))
	} else {
		log.Warning("TLS Disabled! Pulse server connection unsecured.")
		s.Server = grpc.NewServer()
	}
	proto.RegisterServerServer(s.Server, s)
	s.Memberlist.Setup()
	log.Info("Pulse initialised on " + configCopy.LocalNode().IP + ":" + configCopy.LocalNode().Port)
	s.Server.Serve(s.Listener)
}

/**
 * Shutdown pulse server (not cli/cmd)
 */
func (s *Server) shutdown() {
	log.Debug("Shutting down server")
	s.Server.GracefulStop()
	s.Listener.Close()
}
