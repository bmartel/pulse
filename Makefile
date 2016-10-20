install:
	rm -rf .git
	chmod +x ./install.sh
	./install.sh
	git init
	cp example.env .env
	glide install
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/codegangsta/gin
	go get bitbucket.org/liamstask/goose/cmd/goose
	git add --all
	git commit -m"Initial Commit"

development:
	gin -a 8080 run

test:
	go test $(shell glide novendor)

production:
	docker build .
