build: build.linux
build.linux:
	mkdir -p target/linux
	cargo build --target x86_64-linux-kernel --release --target-dir target/linux --bin mariadb_test it/src/main.rs
	chmod 755 target/darwin/mariadb_test
build.darwin:
	mkdir -p target/darwin
	cargo release --sign
	cargo release --upload-doc
	cargo build --target-dir target/darwin --bin mariadb_test it/src/main.rs
	chmod 755 target/darwin/mariabdb_test
