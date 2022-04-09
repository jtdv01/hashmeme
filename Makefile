install:
	./bin/install.sh

build: install
	./bin/build.sh

gui: build
	./target/hashmeme

# nixGL is required if you are running on nix
gui-nix: build
	cd ./target && nixGL ./hashmeme

tests:
	./bin/tests.sh