#!/bin/bash
if [[ $EUID -ne 0 ]]; then
  echo "Must be run as a root user"
  exit 1
fi
BINARY="/home/$SUDO_USER/go/bin/alfarpos-backend"
INSTALLMESSAGE="Installing alfarpos-backend"
UPDATE=0
if [ -f $BINARY ]; then
  INSTALLMESSAGE="Updating alfarpos-backend"
  UPDATE=1
fi
echo "Installing alfarpos-backend"
sudo -u $SUDO_USER /usr/local/go/bin/go install -v
if [ ! -f $BINARY ]; then
  echo "There's an error when installing the binary"
  exit 1
fi
if [ $UPDATE -e 1]; then
  echo "Success updating alfarpos-backend binary"
  exit 0
fi
echo "Setting up alfarpos log dir"
LOGDIR="/var/log/alfarpos/"
if [ ! -d "$LOGDIR" ]; then
  echo "Log directory didn't exist, creating the directory"
  mkdir -p $LOGDIR
  chown $SUDO_USER:$SUDO_USER $LOGDIR
fi
DIRCONFIG="/etc/alfarpos/"
CONFIGFILE="config.yaml"
echo "Copying config file. Don't forget to change the config in $DIRCONFIG$CONFIGFILE"
if [ ! -d "$LOGDIR" ]; then
  echo "Log config didn't exist, creating the directory"
  mkdir -p $DIRCONFIG
fi
if [ -f $DIRCONFIG$CONFIGFILE ]; then
  echo "Found existing config file, will creating a backup"
  mv "$DIRCONFIG$CONFIGFILE" "$DIRCONFIG$CONFIGFILE.bk"
  echo "Backup created at $DIRCONFIG$CONFIGFILE.bk"
fi
cp "./files/etc/alfarpos/config.yaml" $DIRCONFIG
echo "Setting up service file. You should start and enable it by yourself"
cp "./files/etc/systemd/system/alfarpos.service" "/etc/systemd/system/"
