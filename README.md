# :scroll: gh-stats
CLI GitHub user stats generator.

## Features
- CLI interface
- Authenticate to GitHub via OAuth2
- Use GitHub's GraphQL API to query and calculate stats
- Embed stats directly into files using Go templates

## Requirements
- go 1.20 to build from source
- a GitHub account and a [personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)
- (optional) an installed nerd font for fancy icons

## Installation
To install `gh-stats`, build it from source:
```
go install github.com/liouk/gh-stats/cmd/gh-stats
```

## Usage

### Authentication
In order to use `gh-stats` you need to authenticate to GitHub via OAuth2 using a [personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token). The tool reads the token from the environment variable `GITHUB_TOKEN`, so for example you can export it before using the tool:

```bash
$ export GITHUB_TOKEN="your token"
$ gh-stats all
```

### Basic usage
Once you've set your GitHub token to the env var, invoke the tool to obtain stats (omit `--no-icons` if you have a nerd font installed, for some fancy icons):
```
$ gh-stats all --no-icons
~~~~~~~~~~~~~~~~
  logged in as
  liouk
~~~~~~~~~~~~~~~~

Repos: 7
Forks: 13
Pulls: 16 (8 open; 0 closed; 8 merged)
Commits: 678
Reviews: 5
Language stats:
   Shell: 75.01%
   Lua: 13.93%
   Go: 3.75%
   Python: 3.39%
   Ruby: 2.59%
```
Have a look at `gh-stats help` for a list of available commands and options.

### Templates
This tool uses Go's [`text/template`](https://pkg.go.dev/text/template) package for templating. The tool exposes the following fields that can be used in a template:
```
{{ .User }}            // the logged in user's username
{{ .NumRepos }}        // the total number of public source (i.e. non-fork) repos of the user
{{ .NumForks }}        // the total number of public forks of the user
{{ .NumPulls }}        // the total number of public pull requests that the user has authored
{{ .NumOpenPulls }}    // the total number of open PRs
{{ .NumClosedPulls }}  // the total number of closed PRs
{{ .NumMergedPulls }}  // the total number of merged PRs
{{ .NumCommits }}      // the total number of total commits the user has authored in public repos and on their default branches
{{ .NumReviews }}      // the total number of reviews assigned to the user
{{ .Languages }}       // a slice containing language stats
{{ .Extras }}          // a special map for defining dynamic template fields from a JSON file
```

#### Languages
The `{{ .Languages }}` slice contains one element for each language in the statistics generated (the total number requested can be specified with the `--num` option; see `gh-stats help lang` or `gh-stats help all`).

Each element is an object with two fields:
- `{{ .Name }}`: the language name (e.g. "Go")
- `{{ .Perc }}`: the percentage of total bytes written, that are written in this language

#### Extras
:warning: TODO

#### Examples
The [examples/](https://github.com/liouk/gh-stats/tree/main/examples) dir of this repo contains an example template file, an example extras JSON file and the resulting markdown file, rendered with the following command:
```bash
$ gh-stats all --template examples/basic.tmpl --output examples/basic.md --template-extras examples/basic.json
```
