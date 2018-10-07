.PHONY: all be fe webpack clean

GOPATH :=
ifeq ($(OS),Windows_NT)
	GOPATH := $(CURDIR)/_vender;$(CURDIR)
else
	GOPATH := $(CURDIR)/_vender:$(CURDIR)
endif

export GOPATH

all: be fe

be:
	go install be/cmd/xksjht
	go install be/cmd/parser

webpack:
	cd src/fe && npm run build

fe: webpack
	echo '{{define "index"}}' > src/fe/dist/_index.html
	cat src/fe/dist/index.html >> src/fe/dist/_index.html
	echo "{{end}}" >> src/fe/dist/_index.html
	rm -rf src/fe/dist/index.html
	rm -rf src/fe/dist/www_base.html
	rm -rf src/fe/dist/www_articles.html
	rm -rf src/fe/dist/www_article.html
	mv src/fe/dist/_index.html src/fe/dist/index.html
	cp src/fe/src/assets/login.html src/fe/dist/login.html
	cp src/fe/src/assets/www_base.html src/fe/dist/www_base.html
	cp src/fe/src/assets/www_articles.html src/fe/dist/www_articles.html
	cp src/fe/src/assets/www_article.html src/fe/dist/www_article.html
	cp src/fe/src/assets/www.css src/fe/dist/static/www.css
	cp src/fe/src/assets/weibo.jpg src/fe/dist/static/weibo.jpg
	cp src/fe/src/assets/logo.png src/fe/dist/static/logo.png

clean:
	rm -rf bin
	rm -rf pkg
