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

type DRAC struct {
	Host     string
	Username string
	Password string
}

var templates map[int]string = map[int]string{
	6: viewer6,
	7: viewer7,
}

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

func (d *DRAC) Viewer() (string, error) {

	// Check we have a valid DRAC viewer template for this DRAC version
	version := d.GetVersion()
	if version < 0 {
		return "", errors.New("unable to detect DRAC version")
	}

	log.Printf("Found iDRAC version %d", version)

	if _, ok := templates[version]; !ok {
		msg := fmt.Sprintf("no support for DRAC v%d", version)
		return "", errors.New(msg)
	}

	// Generate a JNLP viewer from the template
	// Injecting the host/user/pass information
	buff := bytes.NewBufferString("")
	err := template.Must(template.New("viewer").Parse(templates[version])).Execute(buff, d)
	return buff.String(), err

}
