CONTAINER_TAG?=us.icr.io/attila-fabian/proxy-protocol-debugger

build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w' -o proxy-protocol-debugger -a .

container: build
	docker build -t ${CONTAINER_TAG} --no-cache .

push: container
	docker push ${CONTAINER_TAG}

run:
	go run main.go

deploy:
	kubectl apply -f deployment/proxy-protocol-debugger.yaml