## AlfarPOS Backend Service

This is a backend service for Point of Sales system.

### How it works

TBD

### How to install

#### Prerequisite

- Golang >=1.15
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

### Dev Notes

#### Generate swagger docs

1. Install swag:
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

2. Generate docs:
```bash
swag init --parseDependency
```

3. Access the documentation:
Open [http://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html) in your browser.

#### Authentication
Most endpoints are protected by JWT. You need to login (`POST /api/users/login`) to get a token, then click the "Authorize" button in Swagger UI and enter `Bearer <your_token>`.
