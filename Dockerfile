FROM golang:1.17
# Install Bazel
# Install Bazel
RUN apt-get update && \
    apt-get install -y curl gnupg && \
    curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor > bazel.gpg && \
    mv bazel.gpg /etc/apt/trusted.gpg.d/ && \
    echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list && \
    apt-get update && \
    apt-get install -y bazel

# Set the working directory
WORKDIR /app

# Copy the project files into the container
COPY . .

# Build the project using Bazel
RUN bazel build //cmd/server:server

# Expose the port that the server will listen on
EXPOSE 5000

# Set the entrypoint for the container to the binary
CMD ["./bazel-bin/cmd/server"]