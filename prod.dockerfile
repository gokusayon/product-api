# ------------- Builder Stage ---------------
FROM golang:1.14.3 as builder

# Setting up the env varaibles
ENV GO111MODULE=on
ARG SOURCE_LOCATION=/app

# Set up directories
RUN mkdir -p ${SOURCE_LOCATION}/products-api
RUN mkdir -p ${SOURCE_LOCATION}/currency

# Copy dependencies
COPY products-api/go.mod ${SOURCE_LOCATION}/products-api
COPY currency/ ${SOURCE_LOCATION}/currency

# Download modules
WORKDIR ${SOURCE_LOCATION}/products-api
RUN go mod download

# Build app
COPY products-api/ ${SOURCE_LOCATION}/products-api/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# ----------- Final Build ----------------
FROM alpine:latest  
ARG SOURCE_LOCATION=/app

# Install curl and set working dir
RUN apk --no-cache add curl
WORKDIR /root/

# Copy binaries from previous stage
COPY --from=builder ${SOURCE_LOCATION}/products-api .
# EXPOSE 8080

# Docker container entrypoint
ENTRYPOINT ["./app"]