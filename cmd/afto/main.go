package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/docopt/docopt-go"
	"github.com/fatih/color"
	"github.com/gorilla/handlers"
	"github.com/hako/afto/afutil"
	//"github.com/hako/afto/control"
)

var (
	defaultport = "2468"
	version     = "0.1"
	repoPath    = ""
)

var usage = `afto ` + version + `

Usage:
  afto new <name> 
  afto [-c <file> | --control <file>]
  afto [-d <dir> | --dir <dir>]
  afto [-p <port> | --port <port>]

options:
  -c, --control  Specify control file to use.
  -p, --port     Specify port number for afto.
  -h, --help     Show this screen.
  --version      Show version.

commands:
  new             Generate a new cydia repo.`

func main() {
	// Parse flags.
	if len(os.Args) == 1 {
		fmt.Println(usage)
		os.Exit(1)
	}

	// Parse options with docopt.
	opts, _ := docopt.Parse(usage, nil, true, "afto "+version, false)
	log.SetPrefix("afto: ")
	log.SetFlags(2)

	if (opts["-d"] != true && opts["--dir"] != true) && opts["new"] != true {
		fmt.Println("afto: -d or --dir is required")
		os.Exit(1)
	}

	// New command function [not implemented]
	if opts["new"] == true {
		name := opts["<name>"].(string)
		newRepo(name)
		os.Exit(0)
	}

	var dir = opts["<dir>"].(string)

	// Check if the input directory has the valid files for a cydia repo.
	_, direrr := afutil.ParseDir(dir)
	if direrr != nil {
		fmt.Println("afto: " + direrr.Error())
		os.Exit(1)
	}

	// Get the absolute path from dir.
	finalPath, abserr := filepath.Abs(dir)
	if abserr != nil {
		fmt.Println("afto: " + abserr.Error())
		os.Exit(1)
	}

	repoPath = finalPath

	// afto watches, listens and takes action. (afto listens on 0.0.0.0:[port])
	color.Cyan("afto (αυτο) v" + version + " - the cydia repo generator/manager.")
	color.Cyan("(c) 2016 Wesley Hill (@hako/@hakobyte)")
	fmt.Println("afto is watching & listening for connections on port " + defaultport)

	// Add middleware.
	mx := http.FileServer(http.Dir(repoPath))
	loggingHandler := handlers.LoggingHandler(os.Stdout, mx)

	// AFTODO: Put watcher command here.

	// Spin up a goroutine and serve the repo.
	go func() {
		err := http.ListenAndServe(":"+defaultport, loggingHandler)
		if err != nil {
			fmt.Println("afto: error " + err.Error())
			os.Exit(1)
		}
	}()

	select {}
}

// walkRepo checks multiple directories to see if they have the required files of
// a cydia repo. (running afto on its own triggers this.)
func walkRepos() {

}

// newRepo generates a new cydia compatible repo.
func newRepo(name string) {
	log.Println("generating repo: \"" + name + "\"")
	os.Mkdir(name, 0755)
	// Check for the dpkg command.
	err := afutil.CheckDPKG()
	if err != nil {
		log.Println(err)
		// Now check for compatible platform.
		message, err := afutil.DetectPlatform()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(message)
	}
	// Check for deb files.
	log.Println("checking for deb files...")
	debs, err := afutil.CheckDeb()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(strconv.Itoa(len(debs)) + " deb file(s) found.")
	}
	// Execute dpkg script.
	direrr := executeDpkgScript()
	if direrr != nil {
		log.Fatalln(direrr)
	}
	log.Println("generated Packages file.")
}

// executeDpkgScript executes a commandline script which makes
func executeDpkgScript() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	// Restore required assets to current directory.
	path, _ := filepath.Abs(cwd)
	dataerr := RestoreAsset(path, ".dpkg-scanpackages")
	if dataerr != nil {
		return dataerr
	}
	txterr := RestoreAsset(path, ".dpkg-gettext.pl")
	if txterr != nil {
		return txterr
	}
	// Run dpkg-scanpackages -m . /dev/null >Packages and save the output.
	packages, cmderr := exec.Command("dpkg-scanpackages", "-m", ".", "/dev/null").Output()
	if cmderr != nil {
		return cmderr
	}
	file, err := os.Create("Packages")
	if err != nil {
		return err
	}
	// Write the Packages file.
	_, werr := file.Write(packages)
	if werr != nil {
		return werr
	}
	// Remove assets.
	for _, asset := range AssetNames() {
		os.Remove(asset)
	}
	return nil
}
