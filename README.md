## AlfarPOS Backend Service

This is a backend service for Point of Sales system.

### How it works

### How to install

#### Prerequisite

- Golang >=1.14
- PostgreSQL >=9.6
- Linux system

#### Steps to install

1. Make sure the all of the prerequisite has been installed
2. Run

```
$ sudo -E install.sh
```

3. Change the config on `/etc/alfarpos/config.yaml` according to your setting
4. Start the service by running

```
$ sudo systemctl start alfarpos.service
```

5. Monitor the log in `/var/log/alfarpos/info.log`
