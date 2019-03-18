all: prod

clean:
	rm -rf output

prod: clean output dep build
	cp ./config/* ./output/config
	cp ./controll.sh ./output/
	mv ./app ./output/bin

output:
	mkdir output
	mkdir -p output/bin
	mkdir -p output/config
	mkdir -p output/log

build:
	go build  -o app -tags="Pprof" -gcflags="-N -l"

dep:
	export GO111MODULE=on
