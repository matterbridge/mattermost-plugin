.PHONY: all build plugin
config=matterbridge.toml
all: build plugin

plugin:
ifeq ("$(wildcard $(config))","")
	$(error you need to create a matterbridge.toml file, look at matterbridge.toml.sample and README.md)
endif
	tar zcf plugin.tar.gz plugin.exe plugin.yaml
	@echo " "
	@echo "finished, upload plugin.tar.gz to mattermost"
	@ls -al plugin.tar.gz

build:
	@echo "building plugin.exe"
	@echo "-------------------"
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o plugin.exe

clean:
	rm plugin.tar.gz plugin.exe
