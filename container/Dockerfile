FROM golang:1.18-bullseye

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/spo-iitk/ras-backend

RUN apt-get update
RUN apt-get install -y vim nginx git
# Set timezone to Asia/Kolkata
ENV TZ=Asia/Kolkata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN git config --global user.name "SPO Web Team"
RUN git config --global user.email "pas@iitk.ac.in"

RUN git clone https://github.com/spo-iitk/ras-backend.git .

RUN cp $GOPATH/src/github.com/spo-iitk/ras-backend/secret.yml.template $GOPATH/src/github.com/spo-iitk/ras-backend/secret.yml

# Configure nginx
RUN rm /etc/nginx/sites-enabled/default
RUN ln -s  $GOPATH/src/github.com/spo-iitk/ras-backend/container/nginx.conf /etc/nginx/sites-enabled/default

# This container exposes port 80 to the outside world
EXPOSE 80

# Run the executable
CMD ["./scripts/production.sh"]
