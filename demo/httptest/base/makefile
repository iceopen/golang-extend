build:
	@echo "--> Building..."
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
	docker build -t go-base-http .

runtest:
	docker run -it --rm  -p 8080:8080 go-base-http:latest