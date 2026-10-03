# GitHub Repo Analyzer

A CLI tool written in Go that fetches public repository data for one or more GitHub users concurrently.

## Status

Work in progress. The program currently fetches the repository list for each username and prints it to the terminal. Aggregate statistics (total stars, total forks, top language) are still under development.

## Features

- Fetch repository data from the GitHub REST API (`/users/{username}/repos`)
- Process multiple usernames concurrently using goroutines and channels
- Display stars and language for each repository

### Planned

- Aggregate statistics: total stars, total forks, top language
- HTTP status checking and timeouts on the HTTP client
- Pagination for users with more than 100 repositories
- Context for limiting the duration of each request
- Worker pool for large numbers of usernames
- Unit tests with `httptest`
- Export results to JSON or CSV

## Why This Project?

This project was built to practice:

- Making HTTP requests and decoding JSON responses in Go
- Structuring a Go project across multiple packages
- Using goroutines and channels to fetch data concurrently
- Avoiding race conditions when multiple goroutines send results

## Project Structure

```
githubRepoAnalyzer/
├── go.mod
├── main.go
├── github/
│   └── client.go      # HTTP requests to the GitHub API and JSON decoding
├── stats/
│   └── analyzer.go    # Aggregate statistics (planned)
└── models/
    └── repo.go        # Repo struct definition
```

Note: `client.go` currently uses `package ghclient`. It will be renamed to `package github` to match the structure above.

## How It Works

1. The program accepts one or more GitHub usernames as command-line arguments.
2. Each username is processed in its own goroutine, which calls `FetchRepos`.
3. Each goroutine sends its result (repositories or an error) to a channel.
4. The main goroutine receives exactly one result per username from the channel.
5. The results are printed to the console.

## Concurrency Design

Data is fetched concurrently so that total wait time does not grow linearly with the number of usernames.

```
Sequential: fetch(alice) -> fetch(bob) -> fetch(carol)   (time adds up)
Concurrent: fetch(alice) ┐
            fetch(bob)   ├─ run in parallel               (time ≈ slowest fetch)
            fetch(carol) ┘
```

Goroutines do not write to shared memory. Each result is sent to the main goroutine through a channel, so no mutex is needed for this design.

## Installation

```bash
git clone https://github.com/JUSTINS88/githubRepoAnalyzer.git
cd githubRepoAnalyzer
go mod tidy
```

## Usage

```bash
go run . alice bob carol
```

## Example Output

Current output (per repository):

```
This is repos for alice:
1 (stars: 12) hello-go Language: Go
2 (stars: 0) dotfiles Language: Unknown
```

Target output once aggregation is complete:

```
GitHub Stats Summary

Total Repos: 42
Total Stars: 1,204
Total Forks: 310
Top Language: Go
```

## What I Learned

- Designing concurrent workflows in Go with goroutines and channels
- Matching the number of sends and receives to avoid deadlocks
- Structuring a Go project across multiple packages
- Decoding JSON API responses into Go structs

## License

MIT
