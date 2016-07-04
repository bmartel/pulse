install:
	rm -rf .git
	git init
	cp example.env .env
	glide init
	glide get github.com/bmartel/joust
	glide get github.com/onsi/ginkgo/ginkgo
	glide get github.com/onsi/gomega
	glide up
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/pilu/fresh
	go get bitbucket.org/liamstask/goose/cmd/goose
	git add --all
	git commit -m"Initial Commit"

development:
	fresh

test:
	go test $(shell glide novendor)

production:
	go build -tags production
