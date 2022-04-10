install:
	./bin/install.sh

build:
	./bin/build.sh

gui-nonix: build
	./target/hashmeme

# nixGL is required if you are running on nix
gui: build
	cd ./target && nixGL ./hashmeme

tests:
	./bin/tests.sh
