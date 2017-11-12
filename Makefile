all: build

clean:
	rm -rf pkg bin

deploy: dockerbuild down
	docker run -d --name database.instance -p 8080:80 database.container

up: build
	./Database

dockerbuild:
	./dockerbuild.sh

build:
	go build -o ./Database

down:
	(docker stop database.instance || true) && (docker rm database.instance || true)
