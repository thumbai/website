Title: Introduction
Desc: THUMBAI application introduction and overview.
Keywords: introduction, overview, intro, features
---
Overview|#overview
Features|#features
Get Started|/docs/get-started
Configuration|/docs/configuration
Security|/docs/security
Upgrade|/docs/upgrade
---
# Overview

Initially THUMBAI app features are developed for [aah framework](https://aahframework.org) then I have decedied to make it generalized form with simple web interface for Go community.

Description  | URL Format
----------- | -----------
Dashboard | `<host:port>/thumbai/dashboard`
Go Mod Proxy | `GOPROXY=<host:port>/repo`


# Features

THUMBAI provides an ability to have your own- 

* Go Modules Respository - for organization/individual, never loose dependency libraries.
* Go Vanity Server - Decouple SCM repository path from Go import path.
* Simple Proxy Server - Hosted applications on the same box then simply proxy it behind.

Every feature is on-demand and could be used individually. All THUMBAI features fully utilized by `aahframework.org`, `aahframe.work` and also THUMBAI proxy replaced my nginx proxy server.


