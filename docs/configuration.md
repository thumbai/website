Title: Configuration
Desc: THUMBAI application configurations
Keywords: thumbai configurations, configurations
---
Overview|#overview
Section: Admin|#section-admin
Section: Server|#section-server
Section: Logging|#section-log
---
# Overview

THUMBAI configuration file is same as flexible, friendly configuration format as aah framework.  Binary archive comes with sample configuration file and it explained detailed on this page.

All the THUMBAI defined inside -

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

Admin section used to configure contact email for your support purpose, host limit to access admin UI, data store and disabling Go Mod repository when needed.

Add admin config into section `thumbai { ... }`.

```bash
# Thumbai admin configurations
admin {
    # THUMBAI administrator or support team email address of your organization.
    contact_email = ""

    # Host value is used to enable access to thumbai admin interface.
    # For e.g.: https://example.com/thumbai
    host = "example.com"

    # Added IP's to limit thumbai admin access, by default thumbai admin could
    # be accessed from anywhere with vaild application credentials.
    #
    # Note: By default 127.0.0.1 and ::1 gets added to the allow list on-startup.
    #allow_only = ["192.168.1.1"]

    data_store {
        # Default value is <thumbai-base-directory/data>
        # On-startup thumbai creates the db file 'thumbai.db' if not eixsts.
        #location = "/path/to/datastore"
    }

    # Disable Go Mod repository on-demand.
    disable {
      # Default value is `false`.
      #gomod_repo = true
    }
}
```

# Section: server

Server section used to configure server listen port, timeouts, SSL certs, Let's Encrypt and access log.

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
        #
        # Doc: https://docs.aahframework.org/app-config.html#section-server-ssl-lets-encrypt
        #
        # NOTE: Let’s Encrypt does not provide certificates for `localhost`.
        # ----------------------------------------------------------------------------------
        lets_encrypt {
            # enable = true
            # host_policy = ["aahframe.work", "aahframework.org"]
            # cache_dir = "/home/thumbai/lets-encrypt-certs"
            # email = ""
        }
    }

    # Header value is written as HTTP header `Server: thumbai`.
    #
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
