echo "Starting Database"
/docker-entrypoint.sh postgres &
echo "Updating CA certificates"
update-ca-certificates --fresh >/dev/null

echo "Starting mattermost"
cd /opt/mattermost
ln -s /lib/ld-musl-x86_64.so.1 /lib64/libc.musl-x86_64.so.1
exec /opt/mattermost/bin/mattermost