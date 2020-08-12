# go-git-open

This is an extension for the `git` cli, that allows you to open any git repository in your browser.

Just type `git open` to open the repo website.

## Usage

```sh
# Open the page for this branch on the repo website
git open
```

## Installation

todo..

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
