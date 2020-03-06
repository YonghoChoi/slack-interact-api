FROM debian:jessie

RUN apt-get update -y && apt-get install -y vim dos2unix

RUN mkdir /opt/slack-interact-api
COPY ./output/linux/slack-interact-api /opt/slack-interact-api/slack-interact-api
ADD ./cmd/api/conf /opt/slack-interact-api/
RUN chmod +x /opt/slack-interact-api/slack-interact-api

WORKDIR /opt/slack-interact-api
CMD /opt/slack-interact-api/slack-interact-api