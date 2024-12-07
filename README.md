# Advent of Code Solutions in Go

This repository contains solutions to the [Advent of Code](https://adventofcode.com) challenges implemented in Go. It is designed to streamline the process of running and generating solutions for specific days and years of the event.

## Features
- Solutions implemented in Go for various Advent of Code challenges.
- A Makefile for convenient build, run, and clean operations.
- Dynamic generation and execution of solution files.

---

## Getting Started

### Prerequisites
- [Go](https://golang.org/dl/) (version 1.22 or later is recommended)
- A Unix-like shell with `make` installed.

### Setting Up
1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/advent-of-code-go.git
   cd advent-of-code-go
   ```

2. Install required Go modules:
   ```bash
   go mod tidy
   ```

---

## Usage

### Build the Project
Before running or generating solutions, build the project using:
```bash
make build
```

This compiles the `aoc` binary into the root directory.

### Run a Solution
To run a specific solution, use the following format:
```bash
make run YEAR=<year> DAY=<day> [ARGS...]
```

Example:
```bash
make run 2024 1
```
This runs the solution for Day 1 of Advent of Code 2024.

You can optionally pass additional arguments required by your solution:
```bash
make run 2024 1 arg1 arg2
```

### Generate a New Solution File
To generate a new solution file, use:
```bash
make generate YEAR=<year> DAY=<day>
```

Example:
```bash
make generate 2024 1
```
This prepares a skeleton for the solution of Day 1 of 2024.

### Clean Up
To clean up the compiled binary:
```bash
make clean
```

---

## Project Structure
- `cmd/aoc`: Contains the main application logic for managing solutions.
- `inputs/`: Stores input files for specific days and years.
- `internal/`: Contains Go files implementing solutions for each day.

---

## Environment Configuration
Ensure you have the session cookie for Advent of Code stored in an `.env` file in the root directory:

```
SESSION_COOKIE=your-session-cookie-here
```

This allows the program to fetch puzzle inputs automatically.