FROM golang:1.18-bullseye

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/spo-iitk/ras-backend

RUN apt-get install -y sudo vim nginx git

RUN git config --global user.name "SPO Web Team"
RUN git config --global user.email "pas@iitk.ac.in"

RUN git clone https://github.com/spo-iitk/ras-backend.git .

# Configure nginx
RUN rm /etc/nginx/sites-enabled/default
RUN ln -s  $GOPATH/src/github.com/spo-iitk/ras-backend/container/nginx.conf /etc/nginx/sites-enabled/

# This container exposes port 80 to the outside world
EXPOSE 80

# Run the executable
CMD ["./scripts/production.sh"]
