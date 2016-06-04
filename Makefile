install:
	glide init
	glide up

development:
	go build

test:
	go test $(shell glide novendor)

production:
	go build -tags production
