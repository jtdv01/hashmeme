install:
	./bin/install.sh

build: install
	./bin/build.sh

gui: build
	nixGL ./target/hashmeme

test:
	./bin/test.sh