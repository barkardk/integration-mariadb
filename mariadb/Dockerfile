FROM alpine
COPY target/x86_64-unknown-linux-musl/release/mariadb_test .
COPY passwd.minimal /etc/passwd
USER mariadb
CMD ["sh", "-c", "tail -f /dev/null"]