# mogan-mini

[![release](https://github.com/anondigriz/mogan-mini/actions/workflows/release.yml/badge.svg)](https://github.com/anondigriz/mogan-mini/actions/workflows/release.yml)

mogan-mini implements a local Editor of the Multidimensional Open Gnoseological Active Network (MOGAN) on macOS, Linux, and Windows.

!!! THIS PROJECT IS NOT READY FOR MASS USE YET. See [roadmap](#roadmap) about our plans.

## About

This is an open source implementation of a local editor of the MOGAN knowledge bases (aka the [Mivar](https://mivar.org/en) knowledge bases) editor.

The project is inspired by [Wi!Mi (KESMI)](https://mivar.org/en/services/projects-showcase/wimi-kesmi). You can get a demo of this product from [here](https://github.com/iu5git/MIVAR).

**Wi!Mi** (KESMI – Mivar expert system designer) is a tool for designing knowledge models with unlimited number of connections, parameters and relations, which has logical inference.

## Features

!!! IT IS NOT READY YET. These features are expected from the project as a result of creating an MVP.

mogan-mini provides a CLI for:

1. Work with **the local version of the knowledge base** (store and edit).
2. **Securely synchronize** the knowledge base with a cloud-based implementation of the knowledge base editor (this is a commercial project that is expected to be launched in the future).
3. Prepare of a **logical inference** from the knowledge base (this is a task after passing the MVP stage). 

mogan-mini also provides a user interface based on [Vue.js](https://vuejs.org/) for knowledge base management.


## Installation

Download the latest [here](https://github.com/anondigriz/mogan-mini/releases/latest).

Available platforms:

- macOS amd64 (Intel) and arm64 (M1)
- Linux amd64 and arm64
- Windows amd64


## Documentation

See tutorial [here](https://mini.mivar.org/docs/intro). 

## Demo

[![asciicast](https://asciinema.org/a/576444.svg)](https://asciinema.org/a/576444)

## Roadmap

The project is at the MVP preparation stage and is being developed by an in-house team of developers from the [Research Institute Mivar](https://mivar.org/en) and [Balabza lab](https://balabza.com/). [Here](https://github.com/anondigriz/mogan-mini/milestone/1) is the roadmap for preparing the MVP.

## Development

For CLI development:

- [Install Go](https://go.dev/).
- Run `go run cmd/mogan/main.go`.

For UI development see [here](ui/README.md).
