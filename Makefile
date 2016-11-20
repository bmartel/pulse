install:
	go get github.com/Masterminds/glide
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/sclevine/agouti
	go get github.com/codegangsta/gin
	rm -rf .git
	chmod +x ./install.sh
	./install.sh
	git init
	mv ./config/local.yml ./config/config.yml
	glide install
	go install ./vendor/github.com/mattn/go-sqlite3
	git add --all
	git commit -m"Initial Commit"

development:
	gin -a 8080 run

test:
	go test $(shell glide novendor)

production:
	chmod +x ./build.sh
	./build.sh
