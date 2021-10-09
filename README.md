# Error File API
This is a micro API written in Go (Golang/Gopher).

* Author: Jean Mattes
* Author-URI: https://risingcode.net/
* License: 2021 MIT (Code) and CC-BY-4.0 (Docs)

## Info
It responds to HTTP Requests on the Port specified by its commandline parameter 
(defaults to ``9876``). There are only two Response Codes used.
Either HTTP Status Code ``200 OK`` or ``418 I'm a teapot``.

This depends on a file named ``errorfile`` being present (418) in the 
same directory as the API executable or not (200).

The API can be used by any Program or Language supporting HTTP, e.g. ``curl``:
```bash
curl -w "Status %{http_code}\n" http://localhost:9876
```

*Note*: To simply display the Status Code with PowerShell curl.exe needs to be called instead;

## Building
For the common operating systems (OS) temporarily set the environment variable 
to the target OS and reset it afterwards. On the same OS simply use:
```bash
go build -o errorfile-api main.go
```

Example A: Target OS Linux; build from Windows
```
go env -w GOOS=linux; go build -o errorfile-api main.go; go env -w GOOS=windows
```

Example B: Target OS Windows (append `.exe`); build from Linux
```bash
go env -w GOOS=windows; go build -o errorfile-api.exe main.go; go env -w GOOS=linux
```

Example C: Target OS macOS (darwin); build from Linux
```bash
go env -w GOOS=darwin; go build -o errorfile-api main.go; go env -w GOOS=linux
```

## Usage
To check if it is working run the following lines (or something equivalent)

### Starting the API
Run ``errorfile-api`` or ``errorfile-api.exe`` on Windows

Run ``./errorfile-api`` on Linux/macOS/PowerShell

### Testing
Substitute your URI if running remotely.

Windows (for PowerShell replace `curl` with `curl.exe`)
```
curl -w "Status %{http_code}\n" http://localhost:9876
echo 2>errorfile
curl -w "Status %{http_code}\n" http://localhost:9876
del errorfile
curl -w "Status %{http_code}\n" http://localhost:9876
:: Example Output ("echo." instead of "echo" suppresses echo=on info):
:: Status 200
:: Status 418
:: Status 200
```

Linux
```bash
curl -w "Status %{http_code}\n" http://localhost:9876
touch errorfile
curl -w "Status %{http_code}\n" http://localhost:9876
rm errorfile
curl -w "Status %{http_code}\n" http://localhost:9876
# Example Output:
# Status 200
# Status 418
# Status 200
```
