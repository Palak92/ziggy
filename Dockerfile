FROM golang:1.19-buster as builder

# Install Bazel
RUN apt-get update && apt-get install -y curl gnupg
RUN curl https://bazel.build/bazel-release.pub.gpg | apt-key add -
RUN echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list
RUN apt-get update && apt-get install -y bazel

# Install Gazelle
RUN go get github.com/bazelbuild/bazel-gazelle/cmd/gazelle

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./
RUN gazelle -exclude vendor

# Build the project with Bazel
RUN bazel build //...

# Run the grpc server.
CMD ["bazel", "run", "//app/cmd/server:server"]