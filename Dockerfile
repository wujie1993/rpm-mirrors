FROM library/centos:7.4.1708

RUN ["yum","install","-y","rsync","createrepo"]
COPY ./rpm-mirrors /usr/local/bin/
COPY ./rpm-mirrors.conf /etc/rpm-mirrors/

ENTRYPOINT ["/usr/local/bin/rpm-mirrors"]
