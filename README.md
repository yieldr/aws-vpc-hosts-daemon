# aws-vpc-hosts-daemon

[![wercker status](https://app.wercker.com/status/dc44203d288114092e6cc1fa57f3cf68/s/master "wercker status")](https://app.wercker.com/project/bykey/dc44203d288114092e6cc1fa57f3cf68)

Update `/etc/hosts` with private DNS host names.

## Usage

	docker run -v /usr/share/ca-certificates:/etc/ssl/certs:ro yieldr/aws-vpc-hosts-daemon

A common use case is to have the daemon run on a schedule such as cron or systemd timer. You can find an example in the `dist/systemd` folder.
