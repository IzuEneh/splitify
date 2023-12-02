# syntax=docker/dockerfile:1
ARG BUILD_COMMAND=build
ARG REDIRECT_URI=unknown
ARG CLIENT_ID=unknown

FROM node as node-build

ARG BUILD_COMMAND
ARG REDIRECT_URI
ARG CLIENT_ID

WORKDIR /src

COPY client/ .
RUN npm ci 
RUN VITE_REDIRECT_URI=${REDIRECT_URI} VITE_CLIENT_ID=${CLIENT_ID} npm run ${BUILD_COMMAND}

##############################################
ARG GIN_MODE=release
FROM golang:alpine as go-build
ARG GIN_MODE

WORKDIR /app

COPY server/ .
RUN go mod download && go mod verify
RUN GIN_MODE=${GIN_MODE} go build -o ./server .


# ##############################################
FROM scratch

WORKDIR /splitify

# Copy the CA certificates from the previous build stage
COPY --from=go-build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=go-build app/server ./

COPY --from=node-build src/dist ./static

EXPOSE 8080

ENTRYPOINT [ "./server" ]