[简体中文](README_ZH.md)

<h1> <img src="https://git.pcl.ac.cn/yoyoyard/opendata/raw/branch/develop/public/img/favicon.png" alt="logo" width="30" height="30">OpenData - open data project management</h1>

[![Build Status](https://drone.gitea.io/api/badges/go-gitea/gitea/status.svg)](https://drone.gitea.io/go-gitea/gitea)
[![Join the Discord chat at https://discord.gg/Gitea](https://img.shields.io/discord/322538954119184384.svg)](https://discord.gg/Gitea)
[![](https://images.microbadger.com/badges/image/gitea/gitea.svg)](https://microbadger.com/images/gitea/gitea "Get your own image badge on microbadger.com")
[![codecov](https://codecov.io/gh/go-gitea/gitea/branch/master/graph/badge.svg)](https://codecov.io/gh/go-gitea/gitea)
[![GoDoc](https://godoc.org/code.gitea.io/gitea?status.svg)](https://godoc.org/code.gitea.io/gitea)
[![GitHub release](https://img.shields.io/github/release/go-gitea/gitea.svg)](https://github.com/go-gitea/gitea/releases/latest)
[![Help Contribute to Open Source](https://www.codetriage.com/go-gitea/gitea/badges/users.svg)](https://www.codetriage.com/go-gitea/gitea)
[![Become a backer/sponsor of gitea](https://opencollective.com/gitea/tiers/backers/badge.svg?label=backers&color=brightgreen)](https://opencollective.com/gitea)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![Crowdin](https://badges.crowdin.net/gitea/localized.svg)](https://crowdin.com/project/gitea)

## Purpose

The goal of this project is to make the easiest, fastest, and most
painless way of setting up a self-hosted Git service.
Using Go, this can be done with an independent binary distribution across
**all platforms** which Go supports, including Linux, macOS, and Windows
on x86, amd64, ARM and PowerPC architectures.
Want to try it before doing anything else?
Do it [with the online demo](https://try.gitea.io/)!
This project has been
[forked](https://blog.gitea.io/2016/12/welcome-to-gitea/) from
[Gogs](https://gogs.io) since 2016.11 but changed a lot.

## Building

From the root of the source tree, run:

    TAGS="bindata" make build

or if sqlite support is required:

    TAGS="bindata sqlite sqlite_unlock_notify" make build

The `build` target is split into two sub-targets:

- `make backend` which requires [Go 1.12](https://golang.org/dl/) or greater.
- `make frontend` which requires [Node.js 10.13](https://nodejs.org/en/download/) or greater.

If pre-built frontend files are present it is possible to only build the backend:

		TAGS="bindata" make backend

More info: https://docs.gitea.io/en-us/install-from-source/

## Using

    ./opendata web

NOTE: If you're interested in using our APIs, we have experimental
support with [documentation](https://try.gitea.io/api/swagger).
