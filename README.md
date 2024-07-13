# grpc-web-template
This repository serves as a template for setting up and running a gRPC-Web service with an Envoy proxy.

Prerequisites
Before you begin, ensure you have the following installed:

```
Node.js
Docker
Envoy
grpc-web
```

## Installation
To install gRPC-Web, refer to the official documentation on the [grpc-web](https://github.com/grpc/grpc-web/tree/master?tab=readme-ov-file) GitHub page.

## Setup
1. Build Envoy
To build Envoy, execute the script buildEnvoy.sh. This script sets up and builds the Envoy proxy required for the gRPC-Web service.

`./script/buildEnvoy.sh`

2.  Beautify JS Code and Convert Proto Files to Go Code
To beautify your JavaScript code and convert your .proto files into Go code, run the run.sh script. This script will format your JavaScript code according to best practices, create a release version, and use protoc along with protoc-gen-go and protoc-gen-go-grpc plugins to generate Go code from your protocol buffer definitions.

`./run.sh`

3. Build Binary Executable File
To build the binary executable file for your gRPC-Web service, run the script/build.sh script. This script compiles the necessary code and creates the executable file.

`./script/build.sh`

## Admin Interface
This template also incorporates a Bootstrap Admin template for the frontend interface. We reference and use the [RoyalUI Free Bootstrap Admin Template](https://github.com/BootstrapDash/RoyalUI-Free-Bootstrap-Admin-Template.git). The necessary code has already been added to this project.

