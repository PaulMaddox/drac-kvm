package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"time"

	"github.com/Unknwon/goconfig"
	"github.com/ogier/pflag"
)

// CLI flags
var _host = pflag.StringP("host", "h", "some.hostname.com", "The DRAC host (or IP)")
var _username = pflag.StringP("username", "u", "", "The DRAC username")
var _password = pflag.BoolP("password", "p", false, "Prompt for password (optional, will use 'calvin' if not present)")
var javaws = pflag.StringP("javaws", "j", "/usr/bin/javaws", "The path to javaws binary")
var _delay = pflag.IntP("delay", "d", 10, "Number of seconds to delay for javaws to start up & read jnlp before deleting it")

func promptPassword() string {
	var pass string
	fmt.Print("Password: ")
	_, err := fmt.Scan(&pass)
	if err != nil {
		log.Fatalf("Unable to read password from CLI")
	}
	return pass
}

func main() {
	var host string
	var username = "root"
	var password = "calvin"

	// Parse the CLI flags
	pflag.Parse()

	// Check we have access to the javaws binary
	if _, err := os.Stat(*javaws); err != nil {
		log.Fatalf("No javaws binary found at %s", *javaws)
	}

	// Search for existing config file
	usr, _ := user.Current()
	cfg, _ := goconfig.LoadConfigFile(usr.HomeDir + "/.drackvmrc")

	// Finding host in config file or using the one passed in param
	host = *_host
	hostFound := false
	if cfg != nil {
		_, err := cfg.GetSection(*_host)
		if err == nil {
			value, err := cfg.GetValue(*_host, "host")
			if err == nil {
				hostFound = true
				host = value
			} else {
				hostFound = true
				host = *_host
			}
		}
	}

	if *_username != "" {
		username = *_username
	} else {
		if cfg != nil && hostFound {
			value, err := cfg.GetValue(*_host, "username")
			if err == nil {
				username = value
			}
		}
	}

	// If password not set, prompt
	if *_password {
		password = promptPassword()
	} else {
		if cfg != nil && hostFound {
			value, err := cfg.GetValue(*_host, "password")
			if err == nil {
				password = value
			}
		}
	}

	drac := &DRAC{
		Host:     host,
		Username: username,
		Password: password,
	}

	// Generate a DRAC viewer JNLP
	viewer, err := drac.Viewer()
	if err != nil {
		log.Fatalf("Unable to generate DRAC viewer for %s@%s (%s)", drac.Username, drac.Host, err)
	}

	// Write out the DRAC viewer to a temporary file so that
	// we can launch it with the javaws program
	filename := os.TempDir() + string(os.PathSeparator) + "drac_" + drac.Host + ".jnlp"
	ioutil.WriteFile(filename, []byte(viewer), 0600)
	defer os.Remove(filename)

	// Launch it!
	log.Printf("Launching DRAC KVM session to %s", drac.Host)
	if err := exec.Command(*javaws, filename).Start(); err != nil {
		os.Remove(filename)
		log.Fatalf("Unable to launch DRAC (%s)", err)
	}
	// Give javaws a few seconds to start & read the jnlp
	time.Sleep(time.Duration(*_delay) * time.Second)
}
