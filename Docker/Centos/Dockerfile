FROM centos:8

RUN yum -y update && yum -y upgrade

RUN yum install -y curl python3.8 gcc-c++ make
RUN curl -fsSL https://rpm.nodesource.com/setup_14.x | bash -
RUN yum install -y nodejs
RUN npm install --global yarn

RUN yarn global add wetty
RUN echo "root:root" | chpasswd
COPY --chown=root:root ./startup.sh /opt/
RUN chmod +x /opt/startup.sh

ENTRYPOINT ["/opt/startup.sh"]
