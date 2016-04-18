FROM scratch

ADD bin/aws-vpc-hosts-daemon /

CMD ["/aws-vpc-hosts-daemon"]
