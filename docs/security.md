Title: Security Configuration
Desc: THUMBAI security configurations
Keywords: security, security configurations, thumbai security
---
Overview|#overview
Session & Anti-CSRF Keys|#session-anti-csrf-keys
User Datastore|#user-datastore
Sample|#sample
Permisions|#permissions
______________________________
Introduction|/docs/introduction
Get Started|/docs/get-started
Configuring systemd service|/docs/systemd
---
# Security Configuration

Security configurations used to configure Session, Anti-CSRF and User Datastore.

## Session & Anti-CSRF Keys

Use handy command `thumbai generate securekeys` to create secure keys for session and anti-csrf section.

```bash
# -----------------------------------------------------------------------------
# Security Configuration
#
# Refer documentation to explore and customize the configurations.
# Doc: https://docs.aahframework.org/security-config.html
# -----------------------------------------------------------------------------
security {
    session {
        sign_key = ""
        enc_key = ""
    } 

    anti_csrf {
        sign_key = ""
        enc_key = ""
    }
}
```

## User Datastore

User datastore is section which contains user and user permissions for THUMBAI application.

Sample configuration in the distributed binary archive has two users defined (`admin` and `readonly`) in it for a quick start.

**THUMBAI permission is structured as - Basically you could create a Permission Matrix easily.**

```bash
# Permission created with the power of https://docs.aahframework.org/security-permissions.html
app-name : module-name : access-level

# Full Access to all modules (so called admin)
thumbai:gomod,vanity,proxy,tools:read,write

# Full Access to only vanity service
thumbai:vanity:read,write
```

## Sample User Datastore Config

User datastore congiuration goes into section `thumbai { ... }`.

```bash
# -----------------------------------------------------------------------------
# User Datastore
# 
# Doc: https://docs.aahframework.org/security-permissions.html
# Doc: https://docs.aahframework.org/security-config.html
# -----------------------------------------------------------------------------
user_datastore {
    admin {
        # password value is 'welcome'
        password = "$2y$12$hqoJEVD9YnEUOPSrbQkjh.avcAQaj1hs6XHRjJuGpg/jZwFEQZA.i"
        permissions = [
            "thumbai:gomod,vanity,proxy,tools:read,write"
        ]
        # locked = true
        # expired = true
    }
    readonly {
        # password value is 'readonly'
        password = "$2a$12$pNCCu4OOh1Xj.fwaF8YeFORpWD/MlbHxteznC5RciRPoM9489aq/y"
        permissions = [
            "thumbai:gomod,vanity,proxy,tools:read"
        ]
        # locked = true
        # expired = true
    }
}
```

# Permissions

THUMBAI uses aah framework security feature. aah permission model provides full capability to define customized permissions for the application use case. 

THUMBAI Permission has three parts-

```cfg
app-name:feature-names:access-level
```
<br><br>
**For example:**

You would like to create an user called `user1` with access to only `vanity` feature. So configuration would look like -

Define inside the section `user_datastore { ... }`.

```bash
user1 {
    # password value is 'user1'
    password = "$2y$12$btFqdhUZcmidGL/nD3S/S.nrXqAk0mz5dH0Z2179a6eCZI/doobrC"
    permissions = [
        "thumbai:vanity:read,write"
    ]
    # locked = true
    # expired = true
}
```
