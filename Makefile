HANDLER=save_bar read_bar
build: $(HANDLER)

all: init build deploy

init:
	dep ensure

deploy:
	serverless deploy -v

$(HANDLER):
	#dep ensure
	cd handler/$@; \
	env GOOS=linux go build -ldflags="-s -w" -o ../../bin/$@