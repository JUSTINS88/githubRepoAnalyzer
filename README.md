# GitHub Repo Analyzer

A CLI tool written in Go that fetches public repository data for one or more
GitHub users concurrently, then displays aggregate statistics (total stars,
forks, and top languages used).

## Features

- Fetch repository data from the GitHub REST API
- Analyze multiple usernames **concurrently** using goroutines
- Aggregate statistics: total stars, total forks, most-used language
- Clean, idiomatic Go project structure (separated into packages)

## Why This Project?

This project was built to practice:

- Making HTTP requests and decoding JSON responses in Go
- Structuring a Go project across multiple packages
- Using goroutines and channels to fetch data concurrently
- Avoiding race conditions when multiple goroutines produce results

## Project Structure

\`\`\`
repo-analyzer/
├── go.mod
├── main.go
├── github/
│   └── client.go      # HTTP requests to the GitHub API, JSON decoding
├── stats/
│   └── analyzer.go    # Aggregation logic (totals, top language, etc.)
└── models/
    └── repo.go         # Repo struct definition
\`\`\`

## How It Works

1. The program accepts one or more GitHub usernames as input.
2. For each username, a goroutine is spawned to fetch that user's repos
   from the GitHub API.
3. Each goroutine sends its result (or error) back through a channel.
4. The main goroutine collects all results once every fetch completes.
5. Aggregate statistics are calculated and printed to the console.

## Concurrency Design

Fetching is done concurrently instead of sequentially to reduce total
wait time when analyzing multiple users.

\`\`\`
Sequential:  fetch(alice) -> fetch(bob) -> fetch(carol)   (slow, additive time)
Concurrent:  fetch(alice) ┐
             fetch(bob)   ├─ all run in parallel           (fast, ~max time)
             fetch(carol) ┘
\`\`\`

Each goroutine communicates its result back to the main goroutine via a
channel, rather than writing to shared memory directly. This avoids the
need for manual locking and keeps the data flow easy to reason about.

## Usage

\`\`\`bash
go run main.go alice bob carol
\`\`\`

### Example Output

\`\`\`
Fetching repositories for 3 users...

GitHub Stats Summary
=====================
Total Repos:  42
Total Stars:  1,204
Total Forks:  310
Top Language: Go

Per User:
  alice  - 15 repos, 500 stars
  bob    - 20 repos, 600 stars
  carol  - 7 repos,  104 stars
\`\`\`

## Installation

\`\`\`bash
git clone https://github.com/<your-username>/repo-analyzer.git
cd repo-analyzer
go mod tidy
\`\`\`

## Possible Future Improvements

- [ ] Cache API responses to avoid redundant requests
- [ ] Add filtering by language or last-updated date
- [ ] Export results to JSON or CSV
- [ ] Add rate-limit handling for the GitHub API
- [ ] Add unit tests for the `stats` package

## What I Learned

- How to design concurrent workflows in Go using goroutines and channels
- How to avoid goroutine leaks by matching the number of sends and receives
- How to structure a multi-package Go project
- How to decode JSON API responses into Go structs

## License

MIT
