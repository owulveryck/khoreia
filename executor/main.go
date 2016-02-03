/*
Olivier Wulveryck - author of Gexecutor
Copyright (C) 2015 Olivier Wulveryck

This file is part of the Gexecutor project and
is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package executor

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	//"log"
	"net/http"
	//"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	//log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	//log.SetLevel(log.WarnLevel)

}

type Task struct {
	ID string `json:"id"`
}

func uuid() Task {
	var t Task
	u := make([]byte, 16)
	_, err := rand.Read(u)
	if err != nil {
		return t
	}
	u[8] = (u[8] | 0x80) & 0xBF // what does this do?
	u[6] = (u[6] | 0x40) & 0x4F // what does this do?
	t.ID = hex.EncodeToString(u)
	return t
}

// This will hold all the requested tasks
var tasks map[string](*node)

func Run() {

	tasks = make(map[string](*node), 0)
	router := NewRouter()

	caCert, err := ioutil.ReadFile("choreography.pem")
	if err != nil {
		log.Fatal(err)

	}
	caCertPool := x509.NewCertPool()
	ret := caCertPool.AppendCertsFromPEM(caCert)
	if ret == false {
		log.Fatal("No certificate found in the client file")
	}

	server := &http.Server{
		Addr:    ":8585",
		Handler: router,
		TLSConfig: &tls.Config{
			ClientCAs:  caCertPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	log.Println("Starting server on port 8585")
	log.Fatal(server.ListenAndServeTLS("executor.pem", "executor_key.pem"))
	//log.Fatal(http.ListenAndServe(":8585", router))
}
