Title: Configuration
Desc: THUMBAI app configuration
Keywords: thumbai configuration, configuration
---
Overview|#overview
Section: Admin|#section-admin
Section: Server|#section-server
Section: Logging|#section-log
______________________________
Introduction|/docs/introduction
Get Started|/docs/get-started
Security|/docs/security
Configuring systemd service|/docs/systemd
---
# Overview

THUMBAI configuration file is flexible, friendly configuration format like aah framework configs.  Binary archive comes with sample configuration file and it explained in detail on this page.

THUMBAI configurations defined inside section `thumbai { ... }` -

```bash
# -----------------------------------------------------------------------------
# THUMBAI Application configuration (https://thumbai.app/docs/configuration)
#
# NOTE: THUMBAI uses many aah framework OOTB features and its configurations.
#       So you could refer to aah configuration as needed.
# -----------------------------------------------------------------------------

thumbai {
    # ...
}
```

# Section: admin

Admin section used to configure 

* Contact email for your support purpose
* Host name config and Limit to access admin UI
* Data store path
* Disabling Go Mod repository as needed

Add `admin` config into section `thumbai { ... }`.

```bash
# Thumbai admin configurations
admin {
    # THUMBAI administrator or support team email address of your organization.
    contact_email = "admin@example.com"

    # Host value is used to enable access to thumbai admin interface.
    # For e.g.: https://example.com/thumbai
    host = "example.com"

    # Added IP's to limit thumbai admin access, by default thumbai admin could
    # be accessed from anywhere with vaild application credentials.
    #
    # Note: By default 127.0.0.1 and ::1 gets added to the allow list on-startup.
    allow_only = ["192.168.1.1"]

    data_store {
        # Default value is <thumbai-base-directory/data/>
        # On-startup thumbai creates the db file 'thumbai.db' if not eixsts.
        #directory = "/path/to/datastore"
    }

    # Disable Go Mod repository on-demand.
    disable {
      # Default value is `false`.
      gomod_repo = false
    }

    # GoDoc server host address for Go Vanity service.
    # Default value is `https://godoc.org`.
    godoc_host = "https://godoc.org"
}
```

# Section: server

Server section used to configure server listen port, timeouts and SSL cert OR Let's Encrypt cert.

Add server config into section `thumbai { ... }`.

```bash
# -----------------------------------------------------------------------------
# Server configuration
# Doc: https://docs.aahframework.org/app-config.html#section-server
# -----------------------------------------------------------------------------
server {
    # Port is used to bind server listener on particular port.
    #
    # For standard port `80` and `443`, put empty string or a value
    # Default value is 8080.
    port = "443"

    timeout {
        write = "3m"
    }

    # -----------------------------------------------------------------------------
    # SSL configuration
    # Doc: https://docs.aahframework.org/app-config.html#section-server-ssl
    # -----------------------------------------------------------------------------
    ssl {
        enable = true
        #cert = "/path/to/server.cert"
        #key = "/path/to/server.key"

        # ----------------------------------------------------------------------------------
        # Let’s Encrypt.
        # Doc: https://docs.aahframework.org/app-config.html#section-server-ssl-lets-encrypt
        #
        # NOTE: Let’s Encrypt does not provide certificates for `localhost`.
        # ----------------------------------------------------------------------------------
        lets_encrypt {
            enable = false
            # host_policy = ["aahframe.work", "aahframework.org"]
            # cache_dir = "/home/app/lets-encrypt-certs"
            # email = "admin@example.com"
        }

        # Uncomment `redirect_http` section for Let's Encrypt cert usage
        #redirect_http {
        #  enable = true
        #  port = "80"
        #}
    }

    # Header value is written as HTTP header `Server: thumbai`.
    # To exclude header `Server` from writing, simply comment it out.
    header = "thumbai"
}
```

# Section: log 

Log section used to configure server log configuration.

Add log config into section `thumbai { ... }`.

```bash
# -----------------------------------------------------------------------------
# Log Configuration
#
# Doc: https://docs.aahframework.org/logging.html
# Doc: https://docs.aahframework.org/log-config.html
# -----------------------------------------------------------------------------
log {
    receiver = "file"
    level = "info"
}
```
