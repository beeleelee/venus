# Filecoin (venus)

[![CircleCI](https://circleci.com/gh/filecoin-project/venus.svg?style=svg)](https://circleci.com/gh/filecoin-project/venus)
[![User Devnet Release](https://img.shields.io/endpoint.svg?color=brightgreen&style=flat&logo=GitHub&url=https://raw.githubusercontent.com/filecoin-project/go-filecoin-badges/master/user-devnet.json)](https://github.com/filecoin-project/venus/releases/latest)
[![Nightly Devnet Release](https://img.shields.io/endpoint.svg?color=blue&style=flat&logo=GitHub&url=https://raw.githubusercontent.com/filecoin-project/go-filecoin-badges/master/nightly-devnet.json)](https://github.com/filecoin-project/venus/releases)
[![Staging Devnet Release](https://img.shields.io/endpoint.svg?color=brightgreen&style=flat&logo=GitHub&url=https://raw.githubusercontent.com/filecoin-project/go-filecoin-badges/master/staging-devnet.json)](https://github.com/filecoin-project/venus/releases)

venus is an implementation of the Filecoin Distributed Storage Network. For more details  about Filecoin, checkout the [Filecoin Spec](https://spec.filecoin.io).

venus was the first Filecoin implementation originially initiated and developed by Protocol Labs, and now is maintained by the Filecoin community. See [maintenance](#maintenance) for more information.

__Questions or problems with venus? [Ask the community first](#community)__. Your problem may already be solved.

## Table of Contents
<!--
    TOC generated by https://github.com/thlorenz/doctoc
    Install with `npm install -g doctoc`.
    Regenerate with `doctoc README.md`.
    It's ok to edit manually if you don't have/want doctoc.
 -->
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [What is Filecoin?](#what-is-filecoin)
- [Install](#install)
  - [System Requirements](#system-requirements)
  - [Install from Source](#install-from-source)
    - [Install Go](#install-go)
    - [Install Dependencies](#install-dependencies)
    - [Build and run tests](#build-and-run-tests)
- [Usage](#usage)
  - [Advanced usage](#advanced-usage)
    - [Setting up a localnet](#setting-up-a-localnet)
- [Contributing](#contributing)
- [Community](#community)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## What is Filecoin?
Filecoin is a decentralized storage network that turns the world’s unused storage into an algorithmic market, creating a permanent, decentralized future for the web.
**Miners** earn the native protocol token (also called “Filecoin”) by providing data storage and/or retrieval. 
**Clients** pay miners to store or distribute data and to retrieve it.
Check out the [Filecoin website](https://filecoin.io/) and [Filecoin Documentation](https://docs.filecoin.io/) for more.

## Install

👋 Welcome to venus!

This README outlines the basics for building and running venus.
**For more background, configuration, and troubleshooting information check out the [venus Docs](https://go.filecoin.io/)**.

### System Requirements

venus can build and run on most Linux and MacOS systems. 
Windows is not yet supported.

A validating node can run on most systems with at least 8GB of RAM. 
A mining node requires significant RAM and GPU resources, depending on the sector configuration to be used.

### Install from Source

Clone this git repository to your machine:

```sh
mkdir -p /path/to/filecoin-project
git clone https://github.com/filecoin-project/venus.git /path/to/filecoin-project/venus
```

#### Install Go

The build process for venus requires [Go](https://golang.org/doc/install) >= v1.13

> Installing Go for the first time? We recommend [this tutorial](https://www.ardanlabs.com/blog/2016/05/installing-go-and-your-workspace.html) which includes environment setup.

Due to our use of `cgo`, you'll need a C compiler to build venus whether you're using a prebuilt library or building it yourself from source. 
If you want to use `gcc` (e.g. `export CC=gcc`) when building venus, you will need to use v7.4.0 or higher.

The build process will download a static library containing the [Filecoin proofs implementation](https://github.com/filecoin-project/rust-fil-proofs) (which is written in Rust).

> If instead you wish to build proofs from source, you'll need (1) Rust development environment and (2) to set the environment variable `FFI_BUILD_FROM_SOURCE=1`.
More info at [filecoin-ffi](https://github.com/filecoin-project/filecoin-ffi).

#### Install Dependencies

First, load all the Git submodules.

```sh
git submodule update --init --recursive
```

Initialize build dependencies.

```sh
make deps
```

Note: The first time you run `deps` can be **slow** as very large parameter files are either downloaded or generated locally in `/var/tmp/filecoin-proof-parameters`.
Have patience; future runs will be faster.

#### Build and run tests

```sh
# First, build the binary
make

# Then, run the unit tests.
go run ./build test

# Build and test can be combined!
go run ./build best
```

Other handy build commands include:

```sh
# Check the code for style and correctness issues
go run ./build lint

# Run different categories of tests by toggling their flags
go run ./build test -unit=false -integration=true -functional=true

# Test with a coverage report
go run ./build test -cover

# Test with Go's race-condition instrumentation and warnings (see https://blog.golang.org/race-detector)
go run ./build test -race

# Deps, Lint, Build, Test (any args will be passed to `test`)
go run ./build all
```

Note: Any flag passed to `go run ./build test` (e.g. `-cover`) will be passed on to `go test`.

If you have **problems with the build**, please consult the [Troubleshooting](https://go.filecoin.io/go-filecoin-tutorial/Troubleshooting-&-FAQ.html) section of the [venus Documentation](https://go.filecoin.io/).

## Usage

For a complete step-by-step tutorial, see [Getting Started](https://go.filecoin.io/venus-tutorial/Getting-Started.html).

#### Quick start:

```sh
# Remove any existing symlink to a repo directory
rm ~/.filecoin

# Initialize a new repository, downloading a genesis file and setting network parameters (in this case, for the Testnet network)
./venus init --genesisfile=https://ipfs.io/ipfs/QmXZQeezX1x8uRQX9EUaYxnyivUpTfJqQTvszk3c8SnFPN/testnet.car --network=testnet

# Start the daemon.  It will block until it connects to at least one bootstrap peer.
./venus daemon
```

Your node should now be connected to some peers, and begin downloading and validating the blockchain.

Open a new terminal to interact with your node:

```sh
# Print the node's connection information
./venus id

# Show chain sync status
./venus chain status
```

To see a full list of commands, run `./venus --help`.

### Advanced usage

#### Setting up a localnet

The localnet FAST binary tool allows users to quickly and easily setup a local network on the users computer. 
Please refer to the [localnet README](https://github.com/filecoin-project/venus/tree/master/tools/fast/bin/localnet#localnet) for more information. 
The localnet tool is only compatible when built from the same git ref as the targeted `venus` binary.

## Contributing

We ❤️ all our contributors; this project wouldn’t be what it is without you! If you want to help out, please see [CONTRIBUTING.md](CONTRIBUTING.md).

Check out the [venus code overview](CODEWALK.md) for a brief tour of the code.

## Community

Here are a few places to get help and connect with the Filecoin community:
- [venus Documentation](http://go.filecoin.io/) — for tutorials, troubleshooting, and FAQs
- The `#fil-venus` channel on [Filecoin Project Slack](https://filecoinproject.slack.com/messages/CEHHJNJS3/) or [Matrix/Riot](https://riot.im/app/#/room/#fil-dev:matrix.org) - for live help and some dev discussions
- [Filecoin Community Forum](https://discuss.filecoin.io) - for talking about design decisions, use cases, implementation advice, and longer-running conversations
- [GitHub issues](https://github.com/filecoin-project/venus/issues) - to report bugs, and view or contribute to ongoing development.
- [Filecoin Specification](https://github.com/filecoin-project/specs) - how Filecoin is supposed to work

Looking for even more? See the full rundown at [filecoin-project/community](https://github.com/filecoin-project/community).

## Maintenance

Venus (previously called `venus`) is now maintained by [IPFS-Force Community](https://github.com/ipfs-force-community)

Maintainers: @steven004, @diwufeiwen, @hunjixin, @felix00000

This repo is open for anyone to submit issues and PRs.

## License

The Filecoin Project is dual-licensed under Apache 2.0 and MIT terms:

- Apache License, Version 2.0, ([LICENSE-APACHE](https://github.com/filecoin-project/venus/blob/master/LICENSE-APACHE) or http://www.apache.org/licenses/LICENSE-2.0)
- MIT license ([LICENSE-MIT](https://github.com/filecoin-project/venus/blob/master/LICENSE-MIT) or http://opensource.org/licenses/MIT)
