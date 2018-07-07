HANDLER=save_bar read_bar
build: $(HANDLER)

$(HANDLER):
	dep ensure
	cd handler/$@; \
	env GOOS=linux go build -ldflags="-s -w" -o ../../bin/$@