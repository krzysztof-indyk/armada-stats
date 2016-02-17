# Armada-stats
It is a running daemon that collects, aggregates, processes, and exports to influxdb information about running armada containers.

### Collected inflormations

	ID           ID
	Name         name
	Address      listen address
	Status       container status [0 - passing, 1 - warning, 2 - critical]
	Tags         container tags
	Uptime       uptime in seconds
	Hostname

    CPUPercentage    average CPU per core usage
    Memory           memory usage, without cache, in bytes
    MemoryLimit      momory limit in bytes
    MemoryPercentage percent memory usage
    Swap             swap usage in bytes
    NetworkRx        total number of network Rx in bytes
    NetworkTx        total number of network Tx in bytes
    BlockRead        total number of BlockRead in bytes
    BlockWrite       total number of BlockWrite in bytes


### How to start developing
- Start vagrant `vagrant up`
- Log into it `vagrant ssh`
- Run influx `armada run influxdb -r armada-stats-influxdb -v '<local dir>:/var/influxdb'`
- Create a *custom.yml* in the conf directory to override default configuration options. Especially *armada_host*
- Build container `cd /projects/grafana && armada build`
- Run container `armada run --env dev -v /var/run/docker.sock:/var/run/docker.sock`
- Log into `armada ssh`
- Build armada-stats `go build .`
- Restart armada-stats `supervisorctl restart armada-stats`

To see changes after developing run `go build . && supervisorctl restart armada-stats`

### How to build deb package
`./packaging/build_package.sh``
