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
	"os"
	"path/filepath"
	"github.com/coreos/go-log/log"
)

/**
 * Generate new Cert/Key pair. No Cert Authority.
 */
func GenOpenSSL() {
	// Get project directory location
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Emergency(err)
	}
	_, err = Execute("openssl", "req", "-x509", "-newkey", "rsa:2048", "-keyout", dir + "/certs/server.key", "-out", dir + "/certs/server.crt", "-days", "365", "-subj", "/CN="+GetHostname(), "-sha256", "-nodes")

	if err != nil {
		log.Emergency("Failed to create private server key.")
		os.Exit(1)
	}
}