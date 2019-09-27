/*
   PulseHA - HA Cluster Daemon
   Copyright (C) 2017-2019  Andrew Zak <andrew@linux.com>

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
package server

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/Sirupsen/logrus"
	"github.com/Syleron/PulseHA/proto"
	"github.com/Syleron/PulseHA/src/client"
	"github.com/Syleron/PulseHA/src/config"
	"github.com/Syleron/PulseHA/src/security"
	"github.com/Syleron/PulseHA/src/utils"
	"google.golang.org/grpc"
	"net"
	"os"
	"sync"
	"time"
)

var (
	CLUSTER_REQUIRED_MESSAGE = "You must be in a configured cluster to complete this action."
)

/**
Server struct type
*/
type CLIServer struct {
	sync.Mutex
	Server   *Server
	Listener net.Listener
}

/**
Setup pulse cli type
*/
func (s *CLIServer) Setup() {
	log.Info("CLI server initialised on 127.0.0.1:49152")
	lis, err := net.Listen("tcp", "127.0.0.1:49152")
	if err != nil {
		log.Errorf("Failed to listen: %s", err)
		// TODO: Note: We exit because the service is useless without the CLI server running
		os.Exit(0)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterCLIServer(grpcServer, s)
	grpcServer.Serve(lis)
}

/**
Attempt to join a configured cluster
Notes: We create a new client in attempt to communicate with our peer.
       If successful we acknowledge it and update our memberlist.
*/
func (s *CLIServer) Join(ctx context.Context, in *proto.PulseJoin) (*proto.PulseJoin, error) {
	//log.Debug("CLIServer:Join() Join Pulse cluster")
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		// Create a new client
		c := &client.Client{}
		// Attempt to connect
		err := c.Connect(in.Ip, in.Port, in.Hostname, false)
		// Handle a client connection error
		if err != nil {
			return &proto.PulseJoin{
				Success: false,
				Message: err.Error(),
			}, nil
		}
		// Create new local node config to send
		newNode := &config.Node{
			IP:       in.BindIp,
			Port:     in.BindPort,
			IPGroups: make(map[string][]string, 0),
		}
		// Convert struct into byte array
		buf, err := json.Marshal(newNode)
		// Handle failure to marshal config
		if err != nil {
			log.Errorf("Join() Unable to marshal config: %s", err)
			return &proto.PulseJoin{
				Success: false,
				Message: "Join failure. Please check the logs for more information",
			}, nil
		}
		// Send our join request
		hostname, err := utils.GetHostname()
		if err != nil {
			return nil, errors.New("cannot to join because unable to get hostname")
		}
		r, err := c.Send(client.SendJoin, &proto.PulseJoin{
			Config:   buf,
			Hostname: hostname,
			Token: in.Token,
		})
		// Handle a failed request
		if err != nil {
			log.Errorf("Join() Request error: %s", err)
			return &proto.PulseJoin{
				Success: false,
				Message: "Join failure. Unable to connect to host.",
			}, nil
		}
		// Handle an unsuccessful request
		if !r.(*proto.PulseJoin).Success {
			log.Errorf("Join() Peer error: %s", err)
			return &proto.PulseJoin{
				Success: false,
				Message: r.(*proto.PulseJoin).Message,
			}, nil
		}
		// write CA keys
		utils.CreateFolder(security.CertDir)
		security.WriteCertFile("ca", []byte(r.(*proto.PulseJoin).CaCrt))
		security.WriteKeyFile("ca", []byte(r.(*proto.PulseJoin).CaKey))
		// Generate our new keys
		if err := security.GenTLSKeys(in.BindIp); err != nil {
			log.Errorf("Join() Unable to generate TLS keys: %s", err)
			return &proto.PulseJoin{
				Success: false,
				Message: err.Error(),
			}, nil
		}
		// Update our local config
		peerConfig := &config.Config{}
		err = json.Unmarshal(r.(*proto.PulseJoin).Config, peerConfig)
		// handle errors
		if err != nil {
			log.Error("Unable to unmarshal config node.")
			return &proto.PulseJoin{
				Success: false,
				Message: "Unable to unmarshal config node.",
			}, nil
		}
		// !!!IMPORTANT!!!: Do not replace our local config
		peerConfig.Pulse = DB.Config.Pulse
		// Set the config
		DB.SetConfig(peerConfig)
		// Save the config
		if err := DB.Config.Save(); err != nil {
			return &proto.PulseJoin{
				Success: false,
				Message: "Failed to write config. Joined failed.",
			}, nil
		}
		// Reload config in memory
		DB.Config.Reload()
		// Setup our daemon server
		go s.Server.Setup()
		// reset our HC last received time
		localMember, _ := DB.MemberList.GetLocalMember()
		localMember.SetLastHCResponse(time.Now())
		// Close the connection
		c.Close()
		log.Info("Successfully joined cluster with " + in.Ip)
		return &proto.PulseJoin{
			Success: true,
			Message: "Successfully joined cluster",
		}, nil
	}
	return &proto.PulseJoin{
		Success: false,
		Message: "Unable to join as PulseHA is already in a cluster.",
	}, nil
}

/**
Break cluster / Leave from cluster
TODO: Remember to reassign active role on leave
*/
func (s *CLIServer) Leave(ctx context.Context, in *proto.PulseLeave) (*proto.PulseLeave, error) {
	log.Debug("CLIServer:Leave() - Leave Pulse cluster")
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		return &proto.PulseLeave{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	// Check to see if we are not the only one in the "cluster"
	// Let everyone else know that we are leaving the cluster
	hostname := DB.Config.GetLocalNode()
	if DB.Config.NodeCount() > 1 {
		DB.MemberList.Broadcast(
			client.SendLeave,
			&proto.PulseLeave{
				Replicated: true,
				Hostname:   hostname,
			},
		)
	}
	// make oureselves passive
	MakeLocalPassive()
	// reset our memberlist
	DB.MemberList.Reset()
	// Clear our config
	nodesClearLocal()
	groupClearLocal()
	// Shutdown daemon server
	s.Server.Shutdown()
	// save
	if err := DB.Config.Save(); err != nil {
		return &proto.PulseLeave{
			Success: false,
			Message: "PulseHA successfully removed from cluster but could not update local config",
		}, nil
	}
	// Remove our generated keys
	if !utils.DeleteFolder(security.CertDir) {
		log.Warn("Failed to remove certs directory. Please manually remove hanging certs.")
	}
	// yay?
	log.Info("Successfully left configured cluster. PulseHA no longer listening..")
	if DB.Config.NodeCount() == 1 {
		return &proto.PulseLeave{
			Success: true,
			Message: "Successfully dismantled cluster",
		}, nil
	}
	return &proto.PulseLeave{
		Success: true,
		Message: "Successfully left from cluster",
	}, nil
}

// Remove - Remove node from cluster by hostname
func (s *CLIServer) Remove(ctx context.Context, in *proto.PulseRemove) (*proto.PulseRemove, error) {
	log.Debug("CLIServer:Leave() - Remove node from Pulse cluster")
	s.Lock()
	defer s.Unlock()
	// make sure we are in a cluster
	if !DB.Config.ClusterCheck() {
		return &proto.PulseRemove{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	// make sure we are not removing our active node
	activeHostname, _ := DB.MemberList.GetActiveMember()
	if in.Hostname == activeHostname {
		return &proto.PulseRemove{
			Success: false,
			Message: "Unable to remove active node. Please promote another node and try again",
		}, nil
	}
	localHostname, err := utils.GetHostname()
	if err != nil {
		return &proto.PulseRemove{
			Success: false,
			Message: "Unable to perform remove as unable to get local hostname",
		}, nil
	}
	// Tell everyone else to do the same
	// TODO: This must come first otherwise the node we remove later wont be included in the request
	if DB.Config.NodeCount() > 1 {
		DB.MemberList.Broadcast(
			client.SendRemove,
			&proto.PulseRemove{
				Replicated: true,
				Hostname:   in.Hostname,
			},
		)
	}
	// Set our member status
	member := DB.MemberList.GetMemberByHostname(in.Hostname)
	member.SetStatus(proto.MemberStatus_LEAVING)
	// Check if I am the node being removed
	if in.Hostname == localHostname {
		nodesClearLocal()
		groupClearLocal()
		s.Server.Shutdown()
		log.Info("Successfully removed " + in.Hostname + " from cluster. PulseHA no longer listening..")
	} else {
		// Remove from our memberlist
		DB.MemberList.MemberRemoveByName(in.Hostname)
		// Remove from our config
		err := nodeDelete(in.Hostname)
		if err != nil {
			return &proto.PulseRemove{
				Success: false,
				Message: err.Error(),
			}, nil
		}
	}
	DB.Config.Save()
	return &proto.PulseRemove{
		Success: true,
		Message: "Successfully left from cluster",
	}, nil
}

/**
Create new PulseHA cluster
*/
func (s *CLIServer) Create(ctx context.Context, in *proto.PulseCreate) (*proto.PulseCreate, error) {
	var token string
	s.Lock()
	defer s.Unlock()
	// Make sure we are not in a cluster before creating one.
	if !DB.Config.ClusterCheck() {
		// Get our local hostname
		hostname := DB.Config.GetLocalNode()
		// Remove any hanging nodes
		nodesClearLocal()
		groupClearLocal()
		// Create a new local node config
		if err := nodeCreateLocal(in.BindIp, in.BindPort); err != nil {
			return &proto.PulseCreate{
				Success: false,
				Message: "Failed write local not to config",
				Token: token,
			}, nil
		}
		// Generate new token
		token = generateRandomString(20)
		// Create a new hasher for sha 256
		token_hash := security.GenerateSHA256Hash(token)
		// Set our token in our config
		DB.Config.Pulse.ClusterToken = token_hash
		// Save back to our config
		DB.Config.Save()
		// Cert stuff
		security.GenerateCACert(in.BindIp)
		// Generate client server keys if tls is enabled
		security.GenTLSKeys(in.BindIp)
		// Setup our pulse server
		go s.Server.Setup()
		// Save our newly generated token
		return &proto.PulseCreate{
			Success: true,
			Message: `Pulse cluster successfully created!
	
You can now join any number of machines by running the following on each node:

pulseha join -bind-ip=<IP_ADDRESS> -bind-port=<PORT> -token=` + token + ` ` + in.BindIp + ` ` + in.BindPort + ` ` + hostname + `
			`,
		}, nil
	} else {
		return &proto.PulseCreate{
			Success: false,
			Message: "Pulse daemon is already in a configured cluster",
			Token: token,
		}, nil
	}
}

/**
Add a new floating IP group
*/
func (s *CLIServer) NewGroup(ctx context.Context, in *proto.PulseGroupNew) (*proto.PulseGroupNew, error) {
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		return &proto.PulseGroupNew{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	groupName, err := groupNew(in.Name)
	if err != nil {
		return &proto.PulseGroupNew{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	DB.Config.Save()
	DB.MemberList.SyncConfig()
	return &proto.PulseGroupNew{
		Success: true,
		Message: groupName + " successfully added.",
	}, nil
}

/**
Delete floating IP group
*/
func (s *CLIServer) DeleteGroup(ctx context.Context, in *proto.PulseGroupDelete) (*proto.PulseGroupDelete, error) {
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		return &proto.PulseGroupDelete{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	err := groupDelete(in.Name)
	if err != nil {
		return &proto.PulseGroupDelete{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	DB.Config.Save()
	DB.MemberList.SyncConfig()
	return &proto.PulseGroupDelete{
		Success: true,
		Message: in.Name + " successfully deleted.",
	}, nil
}

/**
Add IP to group
*/
func (s *CLIServer) GroupIPAdd(ctx context.Context, in *proto.PulseGroupAdd) (*proto.PulseGroupAdd, error) {
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		return &proto.PulseGroupAdd{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	err := groupIpAdd(in.Name, in.Ips)
	if err != nil {
		return &proto.PulseGroupAdd{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	DB.Config.Save()
	DB.MemberList.SyncConfig()
	// bring up the ip on the active appliance
	activeHostname, activeMember := DB.MemberList.GetActiveMember()
	// Connect first just in case.. otherwise we could seg fault
	activeMember.Connect()
	iface := DB.Config.GetGroupIface(activeHostname, in.Name)
	activeMember.Send(client.SendBringUpIP, &proto.PulseBringIP{
		Iface: iface,
		Ips:   in.Ips,
	})
	// respond
	return &proto.PulseGroupAdd{
		Success: true,
		Message: "IP address(es) successfully added to " + in.Name,
	}, nil
}

/**
Remove IP from group
*/
func (s *CLIServer) GroupIPRemove(ctx context.Context, in *proto.PulseGroupRemove) (*proto.PulseGroupRemove, error) {
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		return &proto.PulseGroupRemove{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	// TODO: Note: Validation! IMPORTANT otherwise someone could DOS by seg faulting.
	if in.Ips == nil || in.Name == "" {
		return &proto.PulseGroupRemove{
			Success: false,
			Message: "Unable to process RPC call. Required parameters: Ips, Name",
		}, nil
	}
	_, activeMember := DB.MemberList.GetActiveMember()
	if activeMember == nil {
		return &proto.PulseGroupRemove{
			Success: false,
			Message: "Unable to remove IP(s) to group as there no active node in the cluster.",
		}, nil
	}
	err := groupIpRemove(in.Name, in.Ips)
	if err != nil {
		return &proto.PulseGroupRemove{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	DB.Config.Save()
	DB.MemberList.SyncConfig()
	// bring down the ip on the active appliance
	activeHostname, activeMember := DB.MemberList.GetActiveMember()
	// Connect first just in case.. otherwise we could seg fault
	activeMember.Connect()
	iface := DB.Config.GetGroupIface(activeHostname, in.Name)
	activeMember.Send(client.SendBringDownIP, &proto.PulseBringIP{
		Iface: iface,
		Ips:   in.Ips,
	})
	return &proto.PulseGroupRemove{
		Success: true,
		Message: "IP address(es) successfully removed from " + in.Name,
	}, nil
}

/**
Assign group to interface
*/
func (s *CLIServer) GroupAssign(ctx context.Context, in *proto.PulseGroupAssign) (*proto.PulseGroupAssign, error) {
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		return &proto.PulseGroupAssign{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	err := groupAssign(in.Group, in.Node, in.Interface)
	if err != nil {
		return &proto.PulseGroupAssign{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	DB.Config.Save()
	DB.MemberList.SyncConfig()
	return &proto.PulseGroupAssign{
		Success: true,
		Message: in.Group + " assigned to interface " + in.Interface + " on node " + in.Node,
	}, nil
}

/**
Unassign group from interface
*/
func (s *CLIServer) GroupUnassign(ctx context.Context, in *proto.PulseGroupUnassign) (*proto.PulseGroupUnassign, error) {
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		return &proto.PulseGroupUnassign{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	err := groupUnassign(in.Group, in.Node, in.Interface)
	if err != nil {
		return &proto.PulseGroupUnassign{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	DB.Config.Save()
	DB.MemberList.SyncConfig()
	return &proto.PulseGroupUnassign{
		Success: true,
		Message: in.Group + " unassigned from interface " + in.Interface + " on node " + in.Node,
	}, nil
}

/**
Show all groups
*/
func (s *CLIServer) GroupList(ctx context.Context, in *proto.GroupTable) (*proto.GroupTable, error) {
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		return &proto.GroupTable{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	table := new(proto.GroupTable)
	for name, ips := range DB.Config.Groups {
		nodes, interfaces := getGroupNodes(name)
		row := &proto.GroupRow{Name: name, Ip: ips, Nodes: nodes, Interfaces: interfaces}
		table.Row = append(table.Row, row)
	}
	return table, nil
}

/**
Return the status for each node within the cluster
*/
func (s *CLIServer) Status(ctx context.Context, in *proto.PulseStatus) (*proto.PulseStatus, error) {
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		return &proto.PulseStatus{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	table := new(proto.PulseStatus)
	for _, member := range DB.MemberList.Members {
		details, _ := nodeGetByName(member.Hostname)
		tym := member.GetLastHCResponse()
		var tymFormat string
		if tym == (time.Time{}) {
			tymFormat = ""
		} else {
			tymFormat = tym.Format(time.RFC1123)
		}
		row := &proto.StatusRow{
			Hostname:     member.GetHostname(),
			Ip:           details.IP,
			Latency:      member.GetLatency(),
			Status:       member.GetStatus(),
			LastReceived: tymFormat,
		}
		table.Row = append(table.Row, row)
	}
	table.Success = true
	return table, nil
}

/**
Handle CLI promote request
*/
func (s *CLIServer) Promote(ctx context.Context, in *proto.PulsePromote) (*proto.PulsePromote, error) {
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		return &proto.PulsePromote{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	err := DB.MemberList.PromoteMember(in.Member)
	if err != nil {
		return &proto.PulsePromote{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	return &proto.PulsePromote{
		Success: true,
		Message: "Successfully promoted member " + in.Member,
	}, nil
}

/**
Handle CLI promote request
*/
func (s *CLIServer) TLS(ctx context.Context, in *proto.PulseCert) (*proto.PulseCert, error) {
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		return &proto.PulseCert{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	err := security.GenTLSKeys(in.BindIp)
	if err != nil {
		return &proto.PulseCert{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	return &proto.PulseCert{
		Success: true,
		Message: "Successfully generated new TLS certificates",
	}, nil
}

// Config - Update any key's value in the pulseha section of the config
func (s *CLIServer) Config(ctx context.Context, in *proto.PulseConfig) (*proto.PulseConfig, error) {
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		return &proto.PulseConfig{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	// If the value is hostname, update our node in our nodes section as well
	if in.Key == "local_node" {
		return &proto.PulseConfig{
			Success: false,
			Message: "Please restart the PulseHA daemon to update the local node hostname.",
		}, nil
	}
	// Update our key value
	if err := DB.Config.UpdateValue(in.Key, in.Value); err != nil {
		return &proto.PulseConfig{
			Success: false,
			Message: err.Error(),
		}, nil
	}
	DB.Config.Save()
	DB.Config.Reload()
	return &proto.PulseConfig{
		Success: true,
		Message: "Successfully updated pulseha config",
	}, nil
}

// Token - Generate a new cluster token
func (s *CLIServer) Token(ctx context.Context, in *proto.PulseToken) (*proto.PulseToken, error) {
	s.Lock()
	defer s.Unlock()
	if !DB.Config.ClusterCheck() {
		return &proto.PulseToken{
			Success: false,
			Message: "You must be in a configured cluster before completing this action.",
		}, nil
	}
	// Generate new token
	token := generateRandomString(20)
	// Create a new hasher for sha 256
	token_hash := security.GenerateSHA256Hash(token)
	// Set our token in our config
	DB.Config.Pulse.ClusterToken = token_hash
	// Sync our config with the cluster
	if err := DB.MemberList.SyncConfig(); err != nil {
		return &proto.PulseToken{
			Success: false,
			Message: CLUSTER_REQUIRED_MESSAGE,
		}, nil
	}
	// Save back to our config
	DB.Config.Save()
	// report back with success
	return &proto.PulseToken{
		Success: true,
		Message: "Success! Your new cluster token is " + token,
		Token: token,
	}, nil
}
