NOTE: This package is forked from github.com/ains/go-test-html

# go-test-html
Converts `go test` output into a prettified HTML summary.

## Installation

Go version 1.1 or higher is required. Install or update using the `go get`
command:

	go get -u github.com/antonio-salieri/go-test-html

## Usage

The go-test-html command takes in files containing the stdout and stderr from the `go test` command and the location
of the HTML file to write.

    go-test-html [gotest_stdout_file] [gotest_stderr_file] [output_file]

To produce the `gotest_stdout_file` and `gotest_stderr_file` without changing the output of your `go test` runs, use
the following command.

    go test ./... -v 2> >(tee [gotest_stderr_file]) | tee [gotest_stdout_file]
