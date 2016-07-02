install:
	rm -rf .git
	git init
	cp example.env .env
	glide init
	glide get github.com/eknkc/amber
	glide get github.com/joho/godotenv
	glide get gopkg.in/mgo.v2
	glide get github.com/gin-gonic/gin
	glide get github.com/gin-gonic/gin/contrib
	glide get github.com/facebookgo/inject
	glide get github.com/utrack/gin-csrf
	glide get github.com/Sirupsen/logrus
	glide get github.com/bmartel/joust
	glide up
	go get github.com/pilu/fresh
	git add --all
	git commit -m"Initial Commit"

development:
	go build

test:
	go test $(shell glide novendor)

production:
	go build -tags production
