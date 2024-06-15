# # Build stage
# FROM golang:1.21.4 as builder

# WORKDIR /app

# COPY . .

# # RUN go mod download
# RUN make install

# RUN ls -la .

# # Final stage
# FROM alpine:latest

# WORKDIR /root/

# # Copy the built binary from the builder stage
# COPY --from=builder /app/app .

# RUN ls -la /root

# EXPOSE 8345

# CMD ["./app"]

FROM golang:1.21.4

WORKDIR /app

COPY . .

RUN ls -la /app
COPY ./w.env .env

# Install make and other necessary packages
RUN apt-get update && apt-get install -y make

# Run make install to build the project
RUN make install

ENV ENV=docker

# List the contents of the directory to verify the build
RUN ls -la /app

# Create a script to source the .env file and run the application
RUN echo "#!/bin/sh\nexport \$(cat .env | xargs)\n./app" > start.sh
RUN chmod +x start.sh

EXPOSE 8345

# Run the application
CMD ["./start.sh"]