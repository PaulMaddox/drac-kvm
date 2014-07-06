package main

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"text/template"
	"time"
)

// DRAC contains all of the information required
// to connect to a Dell DRAC KVM
type DRAC struct {
	Host     string
	Username string
	Password string
	Version  int
}

// Templates is a map of each viewer.jnlp template for
// the various Dell iDRAC versions, keyed by version number
var Templates = map[int]string{
	6: viewer6,
	7: viewer7,
}

// GetVersion attempts to detect the iDRAC version by checking
// if various known libraries are available via HTTP GET requests.
// Retursn the version if found, or -1 if unknown
func (d *DRAC) GetVersion() int {

	log.Print("Detecting iDRAC version...")

	version := -1

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		Dial: func(netw, addr string) (net.Conn, error) {
			deadline := time.Now().Add(5 * time.Second)
			c, err := net.DialTimeout(netw, addr, time.Second*5)
			if err != nil {
				return nil, err
			}
			c.SetDeadline(deadline)
			return c, nil
		},
	}

	client := &http.Client{
		Transport: transport,
	}

	// Check for iDRAC7 specific libs
	if response, err := client.Get("https://" + d.Host + "/software/avctKVMIOMac64.jar"); err == nil {
		response.Body.Close()
		if response.StatusCode == 200 {
			return 7
		}
	}

	// Check for iDRAC6 specific libs
	if response, err := client.Get("https://" + d.Host + "/software/jpcsc.jar"); err == nil {
		response.Body.Close()
		if response.StatusCode == 200 {
			return 6
		}
	}

	return version

}

// Viewer returns a viewer.jnlp template filled out with the
// necessary details to connect to a particular DRAC host
func (d *DRAC) Viewer() (string, error) {

	var version int

	// Check we have a valid DRAC viewer template for this DRAC version
	if d.Version < 0 {
		version = d.GetVersion()
	} else {
		version = d.Version
	}
	if version < 0 {
		return "", errors.New("unable to detect DRAC version")
	}

	log.Printf("Found iDRAC version %d", version)

	if _, ok := Templates[version]; !ok {
		msg := fmt.Sprintf("no support for DRAC v%d", version)
		return "", errors.New(msg)
	}

	// Generate a JNLP viewer from the template
	// Injecting the host/user/pass information
	buff := bytes.NewBufferString("")
	err := template.Must(template.New("viewer").Parse(Templates[version])).Execute(buff, d)
	return buff.String(), err

}
