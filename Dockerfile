# syntax=docker/dockerfile:1
ARG BUILD_COMMAND=build
ARG REDIRECT_URI=unknown
ARG CLIENT_ID=unknown

FROM node as node-build

ARG BUILD_COMMAND
ARG REDIRECT_URI
ARG CLIENT_ID

WORKDIR /src

ENV VITE_REDIRECT_URI=${REDIRECT_URI}
ENV VITE_CLIENT_ID=${CLIENT_ID}

COPY client/ .

RUN npm ci 
RUN npm run ${BUILD_COMMAND}

##############################################
FROM golang:alpine as go-build

WORKDIR /app

COPY server/ .
RUN go mod download && go mod verify
RUN go build -o ./server .


# ##############################################
ARG GIN_MODE=release
ARG PORT=unknown
FROM scratch
ARG PORT
ARG GIN_MODE

WORKDIR /splitify

# Copy the CA certificates from the previous build stage
COPY --from=go-build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=go-build app/server ./

COPY --from=node-build src/dist ./static

ENV PORT=${PORT}
ENV GIN_MODE=${GIN_MODE}

EXPOSE ${PORT}

ENTRYPOINT [ "./server" ]