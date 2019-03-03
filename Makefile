name=rpm-mirrors
version=0.1.0
release=2

# harbor 
harbor_registry=192.168.1.100:5000
harbor_repo=helm-sunrun-charts-100
harbor_username=viva
harbor_password=Q1w2e3r4t5

# docker
docker_registry=$(harbor_registry)

# helm
release_name=$(name)
namespace=bu


.PHONY: build_bin build_docker

all: build_bin

pre_build: 
	mkdir -p ./build

build_bin: pre_build
	go build -v -o ./build/$(name)

install: build_bin
	cp -f ./build/$(name) /usr/local/bin/$(name) 

build_image: build_bin
	cp ./Dockerfile ./build/
	docker build -t $(name):$(version)-$(release) ./build/

push_image: build_image
	docker tag $(name):$(version)-$(release) $(docker_registry)/gzsunrun/$(name):$(version)-$(release)
	docker push $(docker_registry)/gzsunrun/$(name):$(version)-$(release)

build_chart:
	sed -i 's/^image: gzsunrun\/$(name).*/image: gzsunrun\/$(name):$(version)-$(release)/' ./chart/$(name)/values.yaml
	helm package -d ./build/ ./chart/$(name)

push_chart: build_chart
	helm push -u $(harbor_username) -p $(harbor_password)  ./build/$(name)-$(version).tgz $(harbor_repo)

helm_install: build_chart
	helm install --name $(release_name) --namespace $(namespace) ./build/$(name)-$(version).tgz

helm_upgrade: build_chart
	helm upgrade $(release_name) ./build/$(name)-$(version).tgz

