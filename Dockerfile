# Base container image
FROM golang:1.8-alpine AS builder

# Using Alpine's apk tool, install git which
# is used by Go to download packages
RUN apk --no-cache -U add git

# Install package manager
RUN go get -u github.com/kardianos/govendor

# Copy app files into container
WORKDIR /go/src/app
COPY . .

# Install dependencies
RUN govendor sync

# Build the app
RUN govendor build -o /go/src/app/GUC-SwitchHub

#------------------------------------------------#

# Smallest container image
FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

# Copy built executable from base image to here
COPY --from=builder /go/src/app/GUC-SwitchHub /
COPY templates/SignIn.html /go/src/app/templates/
COPY templates/SignUp.html /go/src/app/templates/
COPY templates/SignUp.css /go/src/app/templates/
COPY templates/SignUp.scss /go/src/app/templates/
COPY templates/SignUp.css.map /go/src/app/templates/
COPY templates/Home.html /go/src/app/templates/
COPY templates/Final.html /go/src/app/templates/
COPY templates/Warning.html /go/src/app/templates/
COPY config.json /go/src/app/

# Run the app
CMD ["/GUC-SwitchHub" ]