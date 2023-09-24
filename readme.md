# go-nibss-ssm: Interpolating Java through C in Golang

## Overview

go-nibss-ssm is a proof-of-concept project that demonstrates how to call Java code from a Go application by using GraalVM's `native-image` to generate a C shared library. The project utilizes a containerized environment defined by a Dockerfile, providing an abstraction over the native dependencies like GraalVM and Go.

## Project Structure

- `main.go`: The entry point to the Go application. Invokes functions exposed from `ssm/ssm.go`.
- `ssm/`: Contains the Go and Java source files related to the Java-C bridging.
  - `ssm.go`: Go file that uses CGO to call the generated C library functions.
  - `com/SSM.java`: Java source file that contains methods to be called from Go.

## Prerequisites

- Docker: Make sure Docker is installed on your machine.
- GraalVM 22.3.1: Tested but not restricted to this version.
- Go 1.20.6: Tested but not restricted to this version.

> Note: The `Dockerfile` will handle the installation of GraalVM and Go during the container build process.

## Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/the-wunmi/go-nibss-ssm.git
   ```

2. Change directory to the project root:
   ```bash
   cd go-nibss-ssm
   ```

3. Build the Docker image:
   ```bash
   docker build . -t=go-nibss-ssm
   ```

## Usage

To run the built Docker image:

```bash
docker run go-nibss-ssm
```

## API Overview

`ssm.go` exposes the following key methods which calls the corresponding Java method through the generated C shared library:

- `GenerateKeyPair(publicKeyStr string, privateKeyStr string, userIdStr string, passwordStr string)`

- `EncryptMessage(publicKeyStr string, messageStr string)`

- `GenerateKeyPair(privateKeyStr string, passwordStr string, encryptedStr string)`

## Known Limitations

None as of now. However, this is a proof of concept and is not intended for production usage.

## Contributing

If you find any issues or opportunities for improving this proof of concept, feel free to contribute by following the typical Git feature branch workflow.
