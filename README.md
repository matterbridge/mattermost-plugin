# mattermost-plugin
Matterbridge mattermost plugin
WIP

## Configuration
You have to create a `matterbridge.toml` configuration file before running `make`. Because the configuration file will be added to the plugin.

Look at the [wiki](https://github.com/42wim/matterbridge/wiki/How-to-create-your-config) on how to create your configuration.
You can also take a look at the `matterbridge.toml.sample` file.

:warning: _IMPORTANT_ :warning:
* the mattermost bridge to work with this plugin must be called `[mattermost.plugin]`
* the `server` directive must be `server="plugin"`
* the `password` directory must be `password="plugin"` 
* the `login` directive must be a user that exists on the channel you're bridging
* the `team` directive must be the team you want to bridge the channel for

```
[mattermost.plugin]
team="yourteam"
login="youruser"
server="plugin"
password="plugin"
```

## Binaries
You can find binaries on https://github.com/matterbridge/mattermost-plugin/releases/latest
Only tested on linux 64-bit!

### Notes :warning:
If you choose to install these preconfigured plugins you'll have to put manually a `matterbridge.toml` in the same directory as the plugin BEFORE enabling the plugin. See configurion above.

See https://docs.mattermost.com/administration/plugins.html#plugin-uploads for more information about the plugins structure.


## Build
Use go 1.11 if possible (only tested with this version).

You have to create a `matterbridge.toml` configuration file before running `make`. Because the configuration file will be added to the plugin.

Look at the [wiki](https://github.com/42wim/matterbridge/wiki/How-to-create-your-config) on how to create your configuration.  
You can also take a look at the `matterbridge.toml.sample` file. Keep in mind the IMPORTANT notice above.

run `make`


```
$ make
building plugin.exe
-------------------
CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' -o plugin.exe

creating plugin.tar.gz archive
------------------------------
tar zcf plugin.tar.gz plugin.exe plugin.yaml

finished, upload plugin.tar.gz to mattermost
-rw-r--r-- 1 wim wim 8183192 Nov 11 22:13 plugin.tar.gz
```

## Misc
~~I'm using a mattermost-server fork github.com/42wim/mattermost-server because the upstream one isn't working correctly with go 1.11 modules yet.~~
