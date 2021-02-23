FROM ubuntu:18.04

RUN apt-get update && apt-get -y upgrade

RUN apt install -y curl python3.8 build-essential
RUN curl -fsSL https://deb.nodesource.com/setup_14.x | bash -
RUN apt install -y nodejs
RUN npm install --global yarn

RUN yarn global add wetty
RUN echo "root:root" | chpasswd
RUN useradd -r wettyhost
RUN echo "wettyhost:pass123" | chpasswd
RUN su wettyhost

CMD wetty --base=/abc123
