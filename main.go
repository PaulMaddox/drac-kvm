package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/ogier/pflag"
)

var host *string = pflag.StringP("host", "h", "some.hostname.com", "The DRAC host (or IP)")
var username *string = pflag.StringP("username", "u", "root", "The DRAC username")
var password *string = pflag.StringP("password", "p", "calvin", "The DRAC password")
var javaws *string = pflag.StringP("javaws", "j", "/usr/bin/javaws", "The path to javaws binary")

func main() {

	// Parse the CLI flags
	pflag.Parse()

	// Check we have access to the javaws binary
	if _, err := os.Stat(*javaws); err != nil {
		log.Fatalf("No javaws binary found at %s", *javaws)
	}

	drac := &DRAC{
		Host:     *host,
		Username: *username,
		Password: *password,
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

	// Launch it!
	log.Printf("Launching DRAC KVM session to %s", drac.Host)
	output, err := exec.Command(*javaws, filename).Output()
	if err != nil {
		log.Fatalf("Unable to launch drac (%s)", err)
	}

	// Show any output from the javaws command
	fmt.Print(output)

	// Remove the generated viewer JNLP
	os.Remove(filename)

}
