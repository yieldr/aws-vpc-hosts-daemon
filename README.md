# aws-vpc-hosts-daemon

Update `/etc/hosts` file with private DNS host names.

## Usage

	aws-vpc-hosts-daemon -hosts-file=/etc/hosts -original-hosts-file=/etc/hosts.orig

Using the docker image:

	docker run -v /usr/share/ca-certificates:/etc/ssl/certs:ro yieldr/aws-vpc-hosts-daemon

A common use case is to have the daemon run on a schedule such as cron or systemd timer. You can find an example in the `dist/systemd` folder.
