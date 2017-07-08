GO=$(shell which go)

dev-build:
	CGO_ENABLED=0 GOOS=linux $(GO) build -a -installsuffix cgo -o flipped
	
dev-run:
	./flipped

