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
FROM golang:alpine as go-build

WORKDIR /app

COPY server/ .
RUN go mod download && go mod verify
RUN go build -o ./server .


# ##############################################
FROM scratch

WORKDIR /splitify

COPY --from=go-build app/server ./

COPY --from=node-build src/dist ./static

EXPOSE 8080

ENTRYPOINT [ "./server" ]