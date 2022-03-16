# GOPEN
Go lang TOOL to open the GIT origin remote of a repo in the Browser

## Installing

You can install Gopen:
- by downloading the [latest release](https://github.com/rogerioefonseca/gopen)

## How it works
Sometimes you need to define different keys for different GIT Versioning systems and then you decide to use `ssh_config` to organize yours IdentifyFiles.
`gopen` checks if the actual git repo remote origin is using a `SSH_HOST` and then if yes it gets the URL of the host configured in the .ssh/config file.
If it does not exists there it will open the remote url configured in the git origin/remote.

```bash
Host github-rogerioefonseca
 HostName github.com
 IdentityFile ~/.ssh/github-rogerioefonseca

Host gitlab-company
 HostName gitlab.company.net
 IdentityFile ~/.ssh/id_rsa

Host github-other-private-project
 HostName gitlab-private.com
 IdentityFile ~/.ssh/id_rsa_other_project

Host gitlab-XYZ
 HostName gitlab.internal.xyz.com
 IdentityFile ~/.ssh/id_rsa_gitlab_internal
```
