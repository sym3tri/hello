FROM scratch

ADD bin/hello /opt/hello
CMD ["/opt/hello"]
