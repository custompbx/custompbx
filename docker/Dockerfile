#docker build . -t cpbx

### STAGE 1: Build Angular frontend ###
FROM node:22 AS frontend
WORKDIR /app

# Copy package.json and package-lock.json
COPY src/cweb-app/package*.json ./

# Install frontend dependencies
RUN npm ci

# Copy frontend files
COPY src/cweb-app/src ./src
COPY src/cweb-app/.browserslistrc ./
COPY src/cweb-app/.postcssrc.json ./
COPY src/cweb-app/angular.json ./
COPY src/cweb-app/tsconfig.json ./
COPY src/cweb-app/scripts ./scripts

RUN npm run test:i18n
RUN npm run build

FROM node:22 AS frontend-test
WORKDIR /app

ENV CHROME_BIN=/usr/bin/chromium

RUN apt-get update && \
    apt-get install -y --no-install-recommends chromium && \
    rm -rf /var/lib/apt/lists/*

COPY src/cweb-app/package*.json ./
RUN npm ci

COPY src/cweb-app/src ./src
COPY src/cweb-app/.browserslistrc ./
COPY src/cweb-app/.postcssrc.json ./
COPY src/cweb-app/angular.json ./
COPY src/cweb-app/tsconfig.json ./
COPY src/cweb-app/scripts ./scripts

RUN npm run test:i18n
RUN npm test -- --watch=false --browsers=ChromeHeadlessNoSandbox

### STAGE 2: Build Golang backend ###
FROM golang:1.25.11 AS backend-base
WORKDIR /app

# Copy go.mod and go.sum
COPY src/custompbx/go.mod src/custompbx/go.sum ./

# Copy backend files
COPY src/custompbx/ ./

# Download backend dependencies
#RUN go mod edit -replace github.com/custompbx/customorm=./customorm/
RUN go mod download

FROM backend-base AS backend-fmt
RUN files="$(find . -name '*.go' ! -path './cweb/cweb.go' -print0 | xargs -0 gofmt -l)" && \
    test -z "$files" || (echo "$files" && exit 1)

FROM backend-base AS backend-vet
COPY config.example.json /test/config.json
ENV CUSTOMPBX_CONFIG=/test/config.json
RUN go vet ./...

FROM backend-base AS backend-test
COPY config.example.json /test/config.json
ENV CUSTOMPBX_CONFIG=/test/config.json
RUN go test ./...

FROM backend-base AS backend-race
COPY config.example.json /test/config.json
ENV CUSTOMPBX_CONFIG=/test/config.json
RUN go test -race ./web ./webStruct ./cfg ./db

FROM backend-base AS integration-test
COPY docker/test-config.json /test/config.json
ENV CUSTOMPBX_CONFIG=/test/config.json
CMD ["go", "test", "-tags=integration", "./db"]

FROM backend-base AS backend
# Copy frontend assets and compile them into the binary.
COPY --from=frontend /app/dist/* ./dist/cweb-app
RUN go install github.com/go-bindata/go-bindata/go-bindata@v3.1.2+incompatible
RUN /go/bin/go-bindata -pkg cweb -prefix dist/cweb-app/browser -o ./cweb/cweb.go ./dist/cweb-app/browser/...
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o cpbx .

### STAGE 3: Final image ###
FROM alpine:latest AS release
WORKDIR /app

# Copy Golang backend from previous stage
COPY --from=backend /app/cpbx /app/cpbx

# Set environment variables if necessary
# ENV ...

VOLUME ["/app"]

# Expose the port that the Golang backend listens on
EXPOSE 8080

# Run the Golang backend
CMD ["/app/cpbx"]
