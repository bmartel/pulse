install:
	rm -rf .git
	git init
	cp example.env .env
	glide init
	glide get gopkg.in/mgo.v2
	glide get github.com/bmartel/joust
	glide get github.com/onsi/ginkgo/ginkgo
  glide get github.com/onsi/gomega
  glide up
  go get github.com/onsi/ginkgo/ginkgo
	go get github.com/pilu/fresh
	git add --all
	git commit -m"Initial Commit"

development:
	go build

test:
	go test $(shell glide novendor)

production:
	go build -tags production
