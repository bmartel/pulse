install:
	rm -rf .git
	git init
	cp example.env .env
	glide init
	glide get github.com/eknkc/amber
	glide get github.com/joho/godotenv
	glide get github.com/jinzhu/gorm
	glide get github.com/gin-gonic/gin
	glide get github.com/gin-gonic/gin/contrib
	glide get github.com/facebookgo/inject
	glide get github.com/utrack/gin-csrf
	glide get github.com/Sirupsen/logrus
	glide get github.com/bmartel/zero
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
