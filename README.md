# Distributed utility grep on Golang

This is a distributed implementation of the `grep` utility in Go, designed to search for patterns across multiple servers for improved performance.

## Usage

The tool supports standard `grep` flags and can search in files or from standard input.

### Building and Running

To build the project:
```bash
make build
```

To run directly:
```bash
make run
```

### Examples

1. **Search for a pattern in a file:**
   ```bash
   ./bin/gogrep "pattern" file.txt
   ```

2. **Search from standard input:**
   ```bash
   echo "some text" | ./bin/gogrep "text"
   ```

3. **Case-insensitive search:**
   ```bash
   ./bin/gogrep -i "Pattern" file.txt
   ```

4. **Print line numbers:**
   ```bash
   ./bin/gogrep -n "pattern" file.txt
   ```

5. **Invert match (show lines that do not match):**
   ```bash
   ./bin/gogrep -v "pattern" file.txt
   ```

6. **Count matches:**
   ```bash
   ./bin/gogrep -c "pattern" file.txt
   ```

7. **Print context lines (before and after):**
   ```bash
   ./bin/gogrep -C 2 "pattern" file.txt
   ```

The tool automatically starts distributed servers on ports 8080-8084 and coordinates the search across them.
