#!/bin/bash
echo "Installing alfarpos-backend"
go install -v
echo "Setting up alfarpos log dir"
LOGDIR="/var/log/alfarpos/"
if [ -d "$LOGDIR" ]; then
  mkdir -p $LOGDIR
  chown root:$USER $LOGDIR
fi
echo "Copying config file"
mkdir -p "/etc/alfarpos/"
cp "./files/etc/alfarpos/config.yaml" "/etc/alfarpos/"
echo "Setting up service file"
cp "./files/etc/systemd/system/alfarpos.service" "/etc/systemd/system/"