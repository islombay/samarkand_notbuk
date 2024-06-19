FROM golang:1.21.4

WORKDIR /app

COPY . .

# Install make and other necessary packages
RUN apt-get update && apt-get install -y make

# Run make install to build the project
RUN make install

ENV ENV=docker

EXPOSE 8345

# Run the application
CMD ["make","start"]