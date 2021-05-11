#!/bin/bash
if [[ $EUID -ne 0 ]]; then
  echo "Must be run as a root user"
  exit 1
fi
BINARY="/home/$SUDO_USER/go/bin/alfarpos-backend"
if [ -f $BINARY ]; then
  echo "Removing existing binary and updating with the new one"
  rm $BINARY
fi
echo "Installing alfarpos-backend"
sudo -u $SUDO_USER /usr/local/go/bin/go install -v
if [ ! -f $BINARY ]; then
  echo "There's an error when installing the binary"
  exit 1
fi
echo "Setting up alfarpos log dir"
LOGDIR="/var/log/alfarpos/"
if [ -d "$LOGDIR" ]; then
  echo "Log directory didn't exist, creating the directory"
  mkdir -p $LOGDIR
  chown root:$SUDO_USER $LOGDIR
fi
echo "Copying config file"
mkdir -p "/etc/alfarpos/"
cp "./files/etc/alfarpos/config.yaml" "/etc/alfarpos/"
echo "Setting up service file"
cp "./files/etc/systemd/system/alfarpos.service" "/etc/systemd/system/"