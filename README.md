# go-git-open

[![Build Status](https://img.shields.io/endpoint.svg?url=https://actions-badge.atrox.dev/mogensen/go-git-open/badge)](https://actions-badge.atrox.dev/mogensen/go-git-open/goto)
[![Go Report Card](https://goreportcard.com/badge/github.com/mogensen/go-git-open)](https://goreportcard.com/report/github.com/mogensen/go-git-open)
[![codecov](https://codecov.io/gh/mogensen/go-git-open/branch/master/graph/badge.svg)](https://codecov.io/gh/mogensen/go-git-open)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmogensen%2Fgo-git-open.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmogensen%2Fgo-git-open?ref=badge_shield)

This is an extension for the `git` cli, that allows you to open any git repository in your browser.

Just type `git open` to open the repo website.

## Usage

```sh
# Open the page for this branch on the repo website
git open
```

## Installation

### Linux

```bash
curl -Lo /tmp/git-open.tar.gz https://github.com/mogensen/go-git-open/releases/download/v0.0.1/git-open_Linux_x86_64.tar.gz
tar xzvf /tmp/git-open.tar.gz -C /tmp
chmod +x /tmp/git-open 
sudo mv /tmp/git-open /usr/local/bin
```

### macOS
```bash
curl -Lo /tmp/git-open.tar.gz https://github.com/mogensen/go-git-open/releases/download/v0.0.1/git-open_Darwin_x86_64.tar.gz
tar xzvf /tmp/git-open.tar.gz -C /tmp
chmod +x /tmp/git-open 
sudo mv /tmp/git-open /usr/local/bin
```

### Windows

Download the binary from the [Latest Release](https://github.com/mogensen/go-git-open/releases/latest/) and add it to your path.


## Supported remote repositories

`go-git-open` can create correct browser urls for a range of different git repository providers.

The currently tested ones are:

- github.com
- gist.github.com
- bitbucket.org
- Azure DevOps

## Configuration 

```bash
# Overwrite the domain from the git remote url with a specific domain
git config open.domain dev.azure.co
```

## Contributing & Development

### Testing

```sh
go test ./...
```

## Related projects

This project is based on the idea from [paulirish/git-open](https://github.com/paulirish/git-open)

- [`git open`](https://github.com/paulirish/git-open) - Bash base version if this repo
- [`git recent`](https://github.com/paulirish/git-recent) - View your most recent git branches
- [`diff-so-fancy`](https://github.com/so-fancy/diff-so-fancy/) - Making the output of `git diff` so fancy

## License

Licensed under MIT. http://opensource.org/licenses/MIT


[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmogensen%2Fgo-git-open.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmogensen%2Fgo-git-open?ref=badge_large)