FROM golang:1.18-bullseye

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/go/src/github.com/spo-iitk/ras-backend

RUN apt-get update
RUN apt-get install -y sudo vim nginx git

RUN git clone https://github.com/spo-iitk/ras-backend.git .

# Configure nginx
RUN rm /etc/nginx/nginx.conf
RUN ln -s ./container/nginx.conf /etc/nginx/nginx.conf

# This container exposes port 80 to the outside world
EXPOSE 80

# Run the executable
CMD ["./scripts/production.sh"]
