org=gzsunrun
name=rpm-mirrors
version=0.1.1
release=0

# harbor 
harbor_registry=192.168.0.100:30002
harbor_repo=helm-sunrun-charts
harbor_username=admin
harbor_password=Harbor12345

# docker
docker_registry=$(harbor_registry)
docker_registry_username=$(harbor_username)
docker_registry_password=$(harbor_password)

# helm
release_name=$(name)
namespace=bu

.PHONY: bin docker chart

all: bin

pre_build: 
	mkdir -p ./build

bin: pre_build
	go build -v -o ./build/$(name)

install: bin
	mkdir -p /etc/$(name)/
	cp ./conf/$(name).conf /etc/$(name)/$(name).conf.example
	cp ./build/$(name) /usr/local/bin/$(name) 

image: bin
	cp ./Dockerfile ./build/
	cp ./conf/$(name).conf ./build/
	docker build -t $(name):latest ./build/

push_image: image
	docker login -u $(docker_registry_username) -p $(docker_registry_password) $(docker_registry)
	docker tag $(name):latest $(docker_registry)/$(org)/$(name):$(version)-$(release)
	docker push $(docker_registry)/$(org)/$(name):$(version)-$(release)

chart:
	sed -i 's/^version: .*/version: $(version)/' ./chart/$(name)/Chart.yaml
	sed -i 's/^image: $(org)\/$(name).*/image: $(org)\/$(name):$(version)-$(release)/' ./chart/$(name)/values.yaml
	helm package -d ./build/ ./chart/$(name)

push_chart: chart
	helm push -u $(harbor_username) -p $(harbor_password)  ./build/$(name)-$(version).tgz $(harbor_repo)

helm_install: chart
	helm install --name $(release_name) --namespace $(namespace) ./build/$(name)-$(version).tgz

helm_upgrade: chart
	helm upgrade $(release_name) ./build/$(name)-$(version).tgz

