package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/kevinburke/ssh_config"
)

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}
}

func dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}

func mountRepoUrl(remote string) (string, error) {
	repo := strings.Split(remote, ":")
	gitRemoteEndpoint := strings.Split(string(repo[0]), "@")
	gitRemoteUrl := strings.Split(gitRemoteEndpoint[1], " ")
	ssh_host := strings.Split(gitRemoteEndpoint[1], " ")

	f, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "config"))
	if err != nil {
		log.Fatal(err)
		return "", errors.New("Could not open config file")
	}

	cfg, err := ssh_config.Decode(f)
	if err != nil {
		log.Fatal(err)
		return "", errors.New("Could not decode the ssh_config file")
	}

	hostname, err := cfg.Get(ssh_host[0], "Hostname")
	if err != nil {
		log.Fatal(err)
		return "", errors.New("Could not get hostname")
	}

	if hostname == "" {
		hostname = gitRemoteUrl[0]
	}

	return "https://" + hostname + "/" + repo[1], nil
}

func main() {
	out, err := exec.Command("git", "remote", "get-url", "origin").CombinedOutput()

	if err != nil {
		fmt.Printf("%s\n", errors.New("Not a git repository"))
		os.Exit(1)
	}

	if strings.Contains(string(out), "https") {
		openbrowser(string(out))
		os.Exit(0)
	}

	repositoryUrl, err := mountRepoUrl(string(out))

	if err != nil {
		log.Fatalf("%s", "Could not retrieve the repositoryUrl correctly")
		os.Exit(1)
	}

	openbrowser(repositoryUrl)
}
