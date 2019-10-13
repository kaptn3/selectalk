# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.13
ARG GO_VERSION=1.13

# First stage: build the executable
FROM golang:${GO_VERSION}-alpine as builder

# Path to a project directory relative $GOPATH/src
ARG PROJECT_PATH=github.com/rodsher/selectel

# Application name. Executable binary with have the same name.
ARG APP_NAME=hackinnovation

# Create the user and group files that will be used in the
# running container to run the process as an uprivileged user.
RUN mkdir /user && \
    echo 'app:x:65534:65534:app:/:' > /user/passwd && \
    echo 'app:x:65534:' > /user/group

# Install the Certificate-Authority certificates for the app to be able
# to make calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git

# Set the working directory. Next commands will execting in a directory that specified bellow.
WORKDIR ${GOPATH}/src/${PROJECT_PATH}

# Fetch dependencies first. They are less susceptible to change on every build
# and will therefore be cached for speeding up the next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import the code from the context
COPY . .

# Build the executable to `/app`.
# Mark the build as statically linked.
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app cmd/${APP_NAME}/main.go

# Final stage: the running container.
FROM scratch as final

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the first stage.
COPY --from=builder /app /app

# Import the migrations for deployment.
COPY  deployments/migrations/postgres /deployments/migrations/postgres

# Declare the port on which the webserver will be exposed.
# As we're going to run as an unprivileged user, we can't bind
# to ports below 1024.
EXPOSE 8080

# Perform any further action as an unprivileged user.
USER app:app

# Run the compiled binary.
ENTRYPOINT ["/app"]