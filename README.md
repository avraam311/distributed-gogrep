# Distributed utility grep on Golang

This is a distributed implementation of the `grep` utility in Go, designed to search for patterns across multiple servers for improved performance on large datasets. The system consists of a coordinator that parses input and distributes the search task to multiple worker servers, which process the data in parallel and return results for aggregation.

## Architecture

The project is structured as follows:

- **cmd/main.go**: Entry point that starts the worker servers and initiates the coordinator.
- **internal/coordinator/**: Handles parsing input, distributing tasks, collecting results, and aggregating output.
  - `parser.go`: Parses command-line flags and input data.
  - `separator.go`: Divides the input data into chunks for distribution.
  - `client.go`: Sends tasks to worker servers and receives results.
  - `agregator.go`: Aggregates and formats the final output.
- **pkg/grepper/**: Contains the worker server implementation.
  - `app.go`: Main application logic for the worker.
  - `server.go`: HTTP server handling search requests.
  - `grepper.go`: Core grep functionality.

## Installation

### Prerequisites
- Go 1.19 or later
- Make (for using the Makefile)

### Building
Clone the repository and build the binary:

```bash
git clone <repository-url>
cd distributed-gogrep
make build
```

This creates the `bin/gogrep` executable.

## Usage

The tool mimics standard `grep` behavior with support for common flags. It reads from a file or standard input and searches for the specified pattern.

### Command Syntax
```bash
./bin/gogrep [OPTIONS] PATTERN [FILE]
```

If no FILE is specified, reads from standard input.

### Supported Flags
- `-A NUM`: Print NUM lines of trailing context after matching lines.
- `-B NUM`: Print NUM lines of leading context before matching lines.
- `-C NUM`: Print NUM lines of context around matching lines.
- `-c`: Print only the count of matching lines.
- `-i`: Perform case-insensitive matching.
- `-v`: Invert the match (print non-matching lines).
- `-F`: Treat PATTERN as a fixed string (not regex).
- `-n`: Print line numbers with output.

### Running
For testing purposes, all 5 worker servers are started on localhost (ports 8080-8084). The coordinator distributes the search across these local servers.

To run:
```bash
make run
# or
./bin/gogrep PATTERN FILE
```

### Examples

1. **Basic search in a file:**
   ```bash
   ./bin/gogrep "error" log.txt
   ```

2. **Search from standard input:**
   ```bash
   cat file.txt | ./bin/gogrep "pattern"
   ```

3. **Case-insensitive search:**
   ```bash
   ./bin/gogrep -i "Error" log.txt
   ```

4. **Print line numbers:**
   ```bash
   ./bin/gogrep -n "pattern" file.txt
   ```

5. **Invert match:**
   ```bash
   ./bin/gogrep -v "pattern" file.txt
   ```

6. **Count matches:**
   ```bash
   ./bin/gogrep -c "pattern" file.txt
   ```

7. **Context lines:**
   ```bash
   ./bin/gogrep -C 2 "pattern" file.txt
   ```

## Deployment for Production

For real-world use with performance benefits, deploy the worker servers (`pkg/grepper`) on 5 separate physical or virtual machines. Update the coordinator's client configuration to point to the actual server addresses instead of localhost.

This distributed setup allows parallel processing across multiple nodes, potentially offering significant speedups for large files or high-volume searches compared to single-threaded grep.

## Performance Comparison

Comparative tests were conducted with the standard `grep` utility on a test file containing 100,000 lines (numbers 1 to 100,000).

### Test Case: Search for "50000"

- **Standard grep:**
  - Output: `50000`
  - Time: 0.003s

- **Distributed gogrep (localhost servers):**
  - Output: `50000`
  - Time: 2.035s

### Notes
- The distributed version has higher startup overhead due to launching multiple servers, making it slower for small files or local testing.
- For larger datasets or when deployed across multiple real servers, the distributed version should provide performance benefits by parallelizing the search across nodes.
- Network latency and data transfer overhead must be considered in distributed deployments.
