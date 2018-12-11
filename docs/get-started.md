Title: Get Started
Desc: Getting started with THUMBAI application
Keywords: get started, getting started with thumbai
---
Download|#download
Update thumbai.conf|#update-configuration
Run THUMBAI app|#run-thumbai-app
Command Help|#command-help-to-know-more
______________________________
Introduction|/docs/introduction
Configuration|/docs/configuration
Security|/docs/security
Configuring systemd service|/docs/systemd
---
## Download

Download THUMBAI binary from <a href="https://thumbai.app/releases/latest" target="_blank">https://thumbai.app/releases/latest</a>.

Binary archive contains-

* THUMBAI binary
* thumbai.service
* thumbai.conf
* LICENSE.txt
* CREDITS.txt


## Update Configuration

Open the file `thumbai.conf` in your favorite editor and update required configuration section and keys.

Refer to [Configuration](/docs/configuration) and [Security](/docs/security) documentation as needed.

Section / Key | Description
------------- | -----------
`thumbai.admin.host` | Host address to be used to access to thumbai admin interface and Go Mod Repository. <br>For e.g.: `https://example.com/thumbai` and `GOPROXY=https://example.com/repo`
`thumbai.server { ... }` | THUMBAI server port no., SSL cert OR Let's Encrypt cert
`thumbai.security.session { ... }`<br>`thumbai.security.anti_csrf { ... }`| Use handy command `thumbai generate securekeys` to get secure keys for Sessiona & Anti-CSRF. 


## Run THUMBAI app

```bash
# Let's say, THUMBAI archive extracted into `/home/app`.
# You could run -
/home/app/thumbai run --config /home/app/thumbai.conf
```

## Command `help` to know more

```bash
/home/app/thumbai help
```
