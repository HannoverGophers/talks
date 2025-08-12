# Build
FROM golang:1.23 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Statisch bauen + richtige Architektur!
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags lambda.norpc -o /out/bootstrap ./    # Name MUSS "bootstrap" sein

# Runtime
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /out/bootstrap /var/runtime/bootstrap
RUN chmod +x /var/runtime/bootstrap

# WICHTIG: Eigenes Entrypoint -> startet DEIN bootstrap direkt.
ENTRYPOINT ["/var/runtime/bootstrap"]
# Keine CMD nötig