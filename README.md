# Mariadb

Client used to run integration and e2e tests in kubernetes

## Makefile
To see what options there are run
```bash
#> make 

cross                          Add dependencies needed for cross compilation
darwin                         Build a webserver debug binary on darwin
debug                          Create a docker container with a debug mode compiled binary and source code. Expose the binary via gdbserver on port 1234. Tag and push docker to registry
dev                            Cross compile a linux binary from darwin in debug mode
release                        Cross compile a linux binary in release mode (Takes longer)
run                            Build and run a docker container locally with debugger on port 1234
test                           Run cargo test on all integration tests

```
