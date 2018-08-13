FROM golang:latest

MAINTAINER cloudy

# config file path
ENV CONFIG /config
VOLUME /config
ENV workspace /workspace
WORKDIR ${workspace}
ADD gin-base-framework ${workspace}


EXPOSE 8000
ENTRYPOINT ["./gin-base-framework","-c","/config"]