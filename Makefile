name=rpm-mirrors
version=0.1.0
release=1
docker_registry=192.168.1.100:5000

.PHONY: build_bin build_docker

all: build_bin

pre_build: 
	mkdir -p ./build

build_bin: pre_build
	go build -v -o ./build/$(name)

build_docker: build_bin
	cp ./Dockerfile ./build/
	docker build -t $(name):$(version)-$(release) ./build/

build_chart:
	helm package -d ./build/ ./chart/$(name)

push_docker: build_docker
	docker tag $(name):$(version)-$(release) $(docker_registry)/gzsunrun/$(name):$(version)-$(release)
	docker push $(docker_registry)/gzsunrun/$(name):$(version)-$(release)
