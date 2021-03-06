/*
   PulseHA - HA Cluster Daemon
   Copyright (C) 2017-2020  Andrew Zak <andrew@linux.com>

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
package pulseha

import (
	"context"
	"crypto/rand"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/syleron/pulseha/rpc"
	"google.golang.org/grpc/peer"
	"net"
	"runtime"
	"time"
)

/**
Networking - Bring up the groups on the current node
*/
func MakeLocalActive() {
	log.Debug("Utils:MakeMemberActive() Local node now active")
	localNode := DB.Config.GetLocalNode()
	for _, node := range DB.Config.Nodes {
		if node.Hostname == localNode.Hostname {
			for iface, groups := range node.IPGroups {
				for _, groupName := range groups {
					makeGroupActive(iface, groupName)
				}
			}
		}
	}
}

/**
Networking - Bring down the ip groups on the current node
*/
func MakeLocalPassive() {
	DB.Logging.Debug("Utils:MakeMemberPassive() Making local node passive")
	localNode := DB.Config.GetLocalNode()
	for _, node := range DB.Config.Nodes {
		if node.Hostname == localNode.Hostname {
			for iface, groups := range node.IPGroups {
				for _, groupName := range groups {
					makeGroupPassive(iface, groupName)
				}
			}
		}
	}
}

/**
Bring up an []ips for a specific interface
*/
func BringUpIPs(iface string, ips []string) error {
	plugin := DB.Plugins.GetNetworkingPlugin()
	if plugin == nil {
		DB.Logging.Debug("Utils:BringUpIps() No networking plugin.. skipping network action")
		return nil
	}
	err := plugin.Plugin.(PluginNet).BringUpIPs(iface, ips)
	return err
}

/**
Bring down an []ips for a specific interface
*/
func BringDownIPs(iface string, ips []string) error {
	plugin := DB.Plugins.GetNetworkingPlugin()
	if plugin == nil {
		DB.Logging.Debug("Utils:BringDownIps() No networking plugin.. skipping network action")
		return nil
	}
	err := plugin.Plugin.(PluginNet).BringDownIPs(iface, ips)
	return err
}

/**
Inform our plugins that our member list state has changed
 */
func InformMLSChange() {
	plugins := DB.Plugins.GetGeneralPlugin()
	if plugins == nil {
		DB.Logging.Debug("Utils:InformMLSChange() No plugins found. Skipping member list state change inform.")
		return
	}

	var safeMemberList []Member

	for _, m := range DB.MemberList.Members {
		safeMemberList = append(safeMemberList, *m)
	}

	for _, p := range plugins {
		p.Plugin.(PluginGen).OnMemberListStatusChange(safeMemberList)
	}
}

/**

 */
func MyCaller() string {
	// we get the callers as uintptrs - but we just need 1
	fpcs := make([]uintptr, 1)
	// skip 3 levels to get to the caller of whoever called Caller()
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return "n/a" // proper error her would be better
	}
	// get the info of the actual function that's in the pointer
	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		return "n/a"
	}
	// return its name
	return fun.Name()
}

/**
Determine who is the correct active node if more than one active is brought online
TODO: Note: THIS ONLY WORKS WITH TWO NODES ATM
*/
func GetFailOverCountWinner(members []*rpc.MemberlistMember) string {
	// GO through our members, are we failing back or not?
	for i, member := range members {
		if member.Status != rpc.MemberStatus_UNAVAILABLE {

			tym, _ := time.Parse(time.RFC1123, member.LastReceived)
			if tym == (time.Time{}) && DB.Config.Pulse.AutoFailback {
				return member.Hostname
			} else {
				// If we are the second member, get the previous node
				if i > 1 {
					return members[i-1].Hostname
				}
				// Otherwise get the second node
				return members[i+1].Hostname
			}
		}
	}
	return ""
}

/**
Determine if a connection is coming in is a member of our config
*/
func CanCommunicate(ctx context.Context) bool {
	pr, ok := peer.FromContext(ctx)
	if !ok {
		DB.Logging.Warn("Unable to get address details for context")
		return false
	}
	// check to make sure the peer IP
	_, err := DB.Config.GetNodeHostnameByAddress(pr.Addr.(*net.TCPAddr).IP.String())
	if err != nil {
		DB.Logging.Warn(err.Error() + ". Communication received from another node not in cluster")
		return false
	}
	return true
}

// generateRandomString -  Generate a random string of length len
func generateRandomString(len int) string {
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}
