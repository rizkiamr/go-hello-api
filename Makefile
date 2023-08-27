docker-build:
	docker build -t go-hello-api:latest .

docker-run:
	docker run -d --name hello-api -p 8080:8080 go-hello-api:latest ./main

.PHONY:
	docker-run
