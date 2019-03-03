# rpm-mirrors

## Requirements

Golang: 1.10

## Run by binary

```
# build and install binary
make install

# copy and edit config file
cp /etc/rpm-mirrors/rpm-mirrors.conf.example /etc/rpm-mirrors/rpm-mirrors.conf
vi /etc/rpm-mirrors/rpm-mirrors.conf

# run by binary
rpm-mirrors
```

## Run by docker

```
# build docker image
make build_image

# edit config file
vi conf/rpm-mirrors.conf

# run by docker
docker run -d --rm -v /data/rpm-mirrors:/data/rpm-mirrors -v conf:/etc/rpm-mirrors rpm-mirrors:latest
```
