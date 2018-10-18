Title: Get Started
Desc: Getting started with THUMBAI application
Keywords: get started, getting started with thumbai
---
Download|#download
Update Configuration|#update-configuration
Run THUMBAI app|#run-thumbai-app
Using systemd|#using-systemd-to-manage-thumbai
---
## Download

Download appropriate platform targeted binaries from here https://thumbai.app/releases/latest.

Binary archive contains thumbai binary, License, `thumbai.service` and `thumbai.conf`.


## Update Configuration

Open file `thumbai.conf` in your favorite editor and update configuration as per your use case.

Refer to [Configuration](/docs/configuration) and [Security](/docs/security) docs as needed.

## Run THUMBAI app

```bash
# Let's say you have extract thumbai into /home/app
/home/app/thumbai -config thumbai.conf
```

## Using systemd to manage THUMBAI

If you would like to manage THUMBAI using `systemd` then update file `thumbai.service` and install it.

Installation Steps-

* Place your `thumbai.service` file at `/etc/systemd/system/`
* Execute `sudo systemctl daemon-reload` to reload the newly configured service file
* Execute `sudo systemctl enable thumbai.service` to enable `thumbai.service` with systemd

Congrats, you have successfully configured thumbai service with systemd.

Now you can do -

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