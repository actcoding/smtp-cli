#!/bin/sh

if [ "$(id -u)" -ne 0 ]
  then echo "Must be run as root"
  exit 1
fi

set -e


install -D -m 0750 -o root -g root smtp-cli /usr/local/bin/smtp-cli
install -D -m 0600 -o root -g root smtp-cli.json /usr/local/etc/smtp-cli/config.json
install -D -m 0600 -o root -g root smtp-cli.json /usr/local/etc/smtp-cli/template.gotmpl


echo " >> Installtion succeeded."
echo " >> Run the CLI using the following command:"
echo
echo " $ /usr/local/bin/smtp-cli -config /usr/local/etc/smtp-cli/config.json -template /usr/local/etc/smtp-cli/template.gotmpl"
