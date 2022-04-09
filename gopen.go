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
	"github.com/manifoldco/promptui"
)

var i int

type (
	openBrowser func(url string)
)

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

	url := strings.TrimSuffix("https://"+hostname+"/"+repo[1], "\n")

	return url, nil
}

func listRemotes() []string {
	out, err := exec.Command("git", "remote").CombinedOutput()

	if err != nil {
		fmt.Printf("%s\n", errors.New("Not a git repository"))
		os.Exit(1)
	}

	return strings.Split(string(out), "\n")
}

func getGitRemoteOrigin() []byte {
	prompt := promptui.Select{
		Label: "Select git remote",
		Items: listRemotes(),
	}

	_, result, _ := prompt.Run()

	out, err := exec.Command("git", "remote", "get-url", result).CombinedOutput()

	if err != nil {
		fmt.Printf("%s\n", errors.New("Not a git repository"))
		os.Exit(1)
	}

	return out
}

func execOpenCommand(open openBrowser, cmdOutput string) {
	open(cmdOutput)
}

func main() {
	cmdOutput := getGitRemoteOrigin()
	openFunc := func(url string) {
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

	if strings.Contains(string(cmdOutput), "https") {
		execOpenCommand(openFunc, string(cmdOutput))
		os.Exit(0)
	}

	repositoryUrl, err := mountRepoUrl(string(cmdOutput))

	if err != nil {
		log.Fatalf("%s", "Could not retrieve the repositoryUrl correctly")
		os.Exit(1)
	}

	execOpenCommand(openFunc, repositoryUrl)
}
