all: prod

clean:
	rm -rf output

prod: clean output dep build
	mv ./ufs ./dev/bin

output:
	mkdir output
	mkdir -p output/bin
	mkdir -p output/log

build:
	go build -tags="Pprof" -gcflags="-N -l"

dep:
	export GO111MODULE=on
