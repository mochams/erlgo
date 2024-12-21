MODULES := hello_world

build: $(MODULES)

$(MODULES):
	go build -o priv/go/$@ ./go_src/$@

clean:
	rm -rf priv/go/*
