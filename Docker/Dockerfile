FROM ubuntu:18.04

RUN apt-get update && apt-get -y upgrade

RUN apt install -y curl python3.8 build-essential
RUN curl -fsSL https://deb.nodesource.com/setup_14.x | bash -
RUN apt install -y nodejs
RUN npm install --global yarn

RUN yarn global add wetty
RUN echo "root:root" | chpasswd
COPY --chown=root:root ./startup.sh /opt/
RUN chmod +x /opt/startup.sh
RUN yes | unminimize

ENTRYPOINT ["/opt/startup.sh"]
