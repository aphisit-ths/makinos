# Start from the latest golang base image
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go.mod and go.sum files first; they are less frequently changed than source code, so Docker can cache this layer
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

ARG APP_ENV
ENV GOLANG_ENVIRONMENT $APP_ENV

ARG API_VERSION="latest"
ENV API_VERSION ${API_VERSION}
ARG PROJECT_NAME="makino"
ENV PROJECT_NAME ${PROJECT_NAME}

# Build the Go app
RUN go build -o main .

RUN sed -i 's/\r$//' start-dev.sh
RUN chmod +x 'start-dev.sh'
