Title: Security Configuration
Desc: THUMBAI security configurations
Keywords: security, security configurations
---
Overview|#overview
Sample|#sample
Permisions|#permissions
---
# Overview

Security configurations used to configure user datastore which contains user and user permissions for THUMBAI application.

Sample configuration in the distributed binary archive has two users (admin and readonly) defined in it for quick start.

# Sample configuration

User datastore congiuration goes into section `thumbai { ... }`.

```bash
# -----------------------------------------------------------------------------
# THUMBAI User Datastore
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
