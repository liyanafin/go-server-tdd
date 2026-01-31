# Go Server with TDD

This project is a simple Go web server that serves a static HTML form and handles form submissions.

## Features
- Serves a static form at `/hello` (renders `public/form.html`)
- Handles form submissions at `/form` (POST)
- Returns a thank you message on valid submission
- Returns appropriate HTTP status codes for invalid requests

## Project Structure
```
main.go           # Main server code
public/form.html  # HTML form served to users
public/index.html # (Optional) Home page
readme.md         # Project documentation
tests/            # (Optional) Test files
```

## Running the Server
1. Make sure you have Go installed (https://golang.org/dl/)
2. In the project directory, run:
   ```sh
   go run main.go
   ```
3. Open your browser and go to [http://localhost:8080/hello](http://localhost:8080/hello) to see the form.

## Testing
### Test Organization
- By default, Go expects test files (ending with `_test.go`) to be in the same directory as the code they test.
- This project keeps its main tests in the project root for compatibility with Go's test runner.

### Running Tests

#### Basic usage
To run all tests with standard output:
```sh
go test -v -count=1
```
This will run all tests in the current directory and print detailed output for each test run.

#### Advanced usage
- Run tests with code coverage:
   ```sh
   go test -cover

   PASS
coverage: 50.0% of statements
ok      github.com/liyanafin/go-server-tdd      0.373s
   ```
- Generate an HTML coverage report:
   ```sh
   go test -coverprofile=coverage.out
   go tool cover -html=coverage.out
   ```
- Run only tests matching a pattern (e.g., only tests with "Valid" in the name):
   ```sh
   go test -run Valid
   ```
- Run tests with the race detector (to find data races):
   ```sh
   go test -race
   ```
- Run tests in short mode (skip long-running tests):
   ```sh
   go test -short
   ```
=== RUN   TestFormHandler_MissingFields
--- PASS: TestFormHandler_MissingFields (0.00s)
=== RUN   TestFormHandler_GetMethod
--- PASS: TestFormHandler_GetMethod (0.00s)
PASS
ok      github.com/liyanafin/go-server-tdd      0.238s
## Endpoints
- `GET /` - Serves Home page
- `GET /contact` - Serves the HTML form
- `POST /form` - Handles form submission

## License
MIT
