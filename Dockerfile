# Start from golang base image
FROM golang:1.14.6-alpine3.12 as builder

# Add Maintainer info
LABEL maintainer="Daniel De La Pava Suarez <daniel·samkit@gmail.com>"

## Install git.
## Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Add the keys
ARG gitlab_user
ARG gitlab_password
ARG gitlab_token

# Set the current working directory inside the container
WORKDIR /app

# Configuration of gitlab to be able to access the private library
RUN git config --global url."https://${gitlab_user}:${gitlab_token}@gitlab.com".insteadOf "https://gitlab.com"
#RUN git config --global url."https://oauth2:${gitlab_token}@gitlab.com".insteadOf "https://gitlab.com"

# Create a netrc file using the credentials specified using --build-arg
#RUN printf "machine gitlab.com\n\
#    login ${gitlab_user}\n\
#    password ${gitlab_password}\n"\
#    >> /root/.netrc

#RUN chmod 600 /root/.netrc

# Required to access go mod from private middleware-securit repo
ENV GOPRIVATE gitlab.com/daniel.delapava/middleware-securit

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/main .
COPY --from=builder /app/database/migrations /migrations
#RUN ls -l

# Expose port 9000 to the outside world
EXPOSE 9000

#Command to run the executable
CMD ["./main"]