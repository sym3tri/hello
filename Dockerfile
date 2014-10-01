FROM scratch

ADD bin/hello /opt/hello
EXPOSE 8080
CMD ["/opt/hello"]
