Title: systemd service
Desc: Configuration systemd service for THUMBAI app
Keywords: systemd, systemctl, service, systemd service, systemctl service
---
Installing Service File|#installing-service-file
______________________________
Introduction|/docs/introduction
Configuration|/docs/configuration
Security|/docs/security
Upgrade|/docs/upgrade
---
# Configuring systemd service

If you would like to manage THUMBAI using `systemd`, follow the below documentation.


## Update Service File

THUMBAI service file is available in the download binary archive. Open the file `thumbai.service` in your favorite editor and update it per you environment

## Installing Service File

* Place file `thumbai.service` at `/etc/systemd/system/`
* Run `sudo systemctl daemon-reload` to reload the newly configured service file
* Run `sudo systemctl enable thumbai.service` to enable `thumbai.service` with systemd

Congrats, you have successfully configured thumbai service with systemd.

**Usage**

```bash
sudo systemctl start thumbai.service
# OR
sudo service thumbai start
```

```bash
sudo systemctl status thumbai.service
# OR
sudo service thumbai status
```