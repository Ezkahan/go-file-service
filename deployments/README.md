# Meditation Service Setup

This guide explains how to set up the **Meditation system service** and
**PostgreSQL database** on Ubuntu 22.04.

------------------------------------------------------------------------

## 1. System Service

Create a custom systemd service for Meditation.

``` bash
# Create or edit service file
sudo vi /etc/systemd/system/meditation.service

# Reload systemd manager configuration
sudo systemctl daemon-reload

# Enable and start the service
sudo systemctl enable --now meditation.service

# Check service status
sudo systemctl status meditation.service
```

------------------------------------------------------------------------

## 2. PostgreSQL Setup

Install and configure PostgreSQL.

``` bash
# Update system packages
sudo apt update && sudo apt upgrade -y

# Install PostgreSQL and contrib package
sudo apt install postgresql postgresql-contrib -y

# Start and enable PostgreSQL service
sudo systemctl start postgresql
sudo systemctl enable postgresql

# Verify status
sudo systemctl status postgresql
```

### Database & User Setup

Switch to the default Postgres user and configure the database:

``` bash
# Switch to postgres user
sudo -i -u postgres

# Open PostgreSQL shell
psql
```

Inside the `psql` shell, run:

``` sql
CREATE USER dev WITH ENCRYPTED PASSWORD 'R49syTa' createdb;
SET ROLE dev;
CREATE DATABASE meditation;
GRANT ALL PRIVILEGES ON DATABASE meditation TO dev;
```

Exit `psql` with:

``` sql
\q
```

------------------------------------------------------------------------

✅ Your Meditation service and PostgreSQL database are now set up.

<!-- 
sudo vi /etc/systemd/system/meditation.service
sudo systemctl daemon-reload
sudo systemctl enable --now meditation.service
sudo systemctl status meditation.service
sudo apt update && sudo apt upgrade -y
sudo apt install postgresql postgresql-contrib -y
sudo systemctl start postgresql
sudo systemctl enable postgresql
sudo systemctl status postgresql
sudo -i -u postgres
psql
CREATE USER dev WITH ENCRYPTED PASSWORD 'R49syTa' createdb;
SET ROLE dev;
CREATE DATABASE meditation;
GRANT ALL PRIVILEGES ON DATABASE meditation TO dev;
-->