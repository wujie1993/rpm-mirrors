FROM library/centos:7.4.1708

RUN ["yum","install","-y","rsync","createrepo"]
COPY ./rpm-mirrors /usr/local/rpm-mirrors/

WORKDIR /usr/local/rpm-mirrors
ENTRYPOINT ["./rpm-mirrors"]
