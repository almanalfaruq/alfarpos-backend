#!/bin/bash

## Check root
if [[ $EUID -ne 0 ]]; then
  echo "Must be run as a root user"
  exit 1
fi

## Step 1 - Instal binary to $GOPATH folder
echo $GOPATH
BINARY="$GOPATH/bin/alfarpos-backend"
INSTALLMESSAGE="Installing alfarpos-backend"
UPDATE=0 ## Update if any
if [ -f $BINARY ]; then
  INSTALLMESSAGE="Updating alfarpos-backend"
  UPDATE=1
fi
echo $INSTALLMESSAGE
sudo -u $SUDO_USER /usr/local/go/bin/go install -v
if [ ! -f $BINARY ]; then
  echo "There's an error when installing the binary"
  exit 1
fi
if [ $UPDATE -eq 1 ]; then
  echo "Success updating alfarpos-backend binary"
  exit 0
fi

## Steps below for installing a new binary
## Step 2 - Setting log directory
echo "Setting up alfarpos log dir"
LOGDIR="/var/log/alfarpos/"
if [ ! -d "$LOGDIR" ]; then
  echo "Log directory didn't exist, creating the directory"
  mkdir -p $LOGDIR
  chown $SUDO_USER:$SUDO_USER $LOGDIR
fi

## Step 3 - Setting config file
DIRCONFIG="/etc/alfarpos/"
CONFIGFILE="config.yaml"
echo "Copying config file. Don't forget to change the config in $DIRCONFIG$CONFIGFILE"
if [ ! -d "$DIRCONFIG" ]; then
  echo "Config directory didn't exist, creating the directory"
  mkdir -p $DIRCONFIG
fi
if [ -f $DIRCONFIG$CONFIGFILE ]; then
  echo "Found existing config file, will creating a backup"
  mv "$DIRCONFIG$CONFIGFILE" "$DIRCONFIG$CONFIGFILE.bk"
  echo "Backup created at $DIRCONFIG$CONFIGFILE.bk"
fi
cp "./files/etc/alfarpos/config.yaml" $DIRCONFIG
echo "Please edit config file as needed in $DIRCONFIG$CONFIGFILE"

## Step 4 - Setting service file
echo "Setting up service file. You should start and enable it by yourself"
cat > /etc/systemd/system/alfarpos.service << EOL
[Unit]
Description=AlfarPOS Backend Service
After=postgresql.service
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=1
User=${SUDO_USER}
ExecStart=/home/${SUDO_USER}/go/bin/alfarpos-backend

[Install]
WantedBy=multi-user.target
EOL
