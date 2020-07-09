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
CMD ["go", "run", "main.go"]