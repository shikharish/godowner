#!/bin/sh

cleanup() {
	echo "Container stopped."
}
trap 'cleanup' SIGQUIT SIGTERM SIGHUP
"${@}" &
wait $!
