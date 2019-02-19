FROM library/centos:7.4.1708

RUN ["yum","install","-y","rsync","createrepo"]
COPY ./rpm-mirrors /usr/local/rpm-mirrors/
COPY ./rpm-mirrors.conf /usr/local/rpm-mirrors/conf/

WORKDIR /usr/local/rpm-mirrors
ENTRYPOINT ["./rpm-mirrors"]
