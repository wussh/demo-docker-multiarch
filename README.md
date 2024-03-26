# Multi-Arch Sample

This repository contains a sample Docker project demonstrating building and running a multi-architecture Docker image using Docker Buildx.

## Overview

The purpose of this project is to showcase how to build and run Docker images that support multiple CPU architectures, specifically targeting both `linux/amd64` and `linux/arm64` platforms.

## Prerequisites

Before you begin, ensure you have the following installed:

- Docker Desktop or Docker Engine installed on your local machine.
- Docker Buildx plugin enabled. If not, enable it using the following command:
  ```
  docker buildx install
  ```

## Build and Run

### Build the Multi-Arch Docker Image

To build the multi-architecture Docker image, run the following command in the project directory:

```
docker buildx build --push --platform linux/amd64,linux/arm64 --tag wushie/multi_arch_sample:v2 .
```

This command will build the Docker image for both `linux/amd64` and `linux/arm64` platforms and push it to the Docker registry with the tag `wushie/multi_arch_sample:v2`.

### Run the Docker Container

To run the Docker container, use the following command:

```
docker run -p 8000:8000 wushie/multi_arch_sample:v2
```

This command will start the Docker container, exposing port 8000 on the host machine to port 8000 in the container.

## Testing

Once the Docker container is running, you can test the application by accessing `http://localhost:8000` in your web browser or using tools like cURL or Postman.