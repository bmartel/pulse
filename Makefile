install:
	glide init
	glide get github.com/bmartel/pulse-core

development:
	go build

test:
	go test $(shell glide novendor)

production:
	go build -tags production
