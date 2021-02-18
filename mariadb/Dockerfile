FROM alpine
COPY target/linux/mariadb_test .
COPY passwd.minimal /etc/passwd
USER mariadb
CMD ["sh", "-c", "tail -f /dev/null"]