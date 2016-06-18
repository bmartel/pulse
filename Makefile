install:
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
	glide up
	go get bitbucket.org/liamstask/goose/cmd/goose

development:
	go build

test:
	go test $(shell glide novendor)

production:
	go build -tags production
