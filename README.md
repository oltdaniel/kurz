# kurz :scissors:

another url shortener

## About

###### This project is mainly for research purposes. Feel free to use and change.

The idea behind this project is to create a fast reliable url shortener that
offers access to customizable urls, simple analytics and notifications.

This project has been implemented in Go, because it is amazing, and I choosed
Aerospike since it is an amazing Database with low latency.

## Installation

###### Long but nothing missed.

#### Aerospike

To install the aerospike server follow [these instructions](https://www.aerospike.com/docs/operations/install).
In order to have a good configuration, we need to change the instructions.

```shell
# Start aerospike if not done already
# - generate config files
$ sudo service aerospike start
# Open the configuration
$ vim /etc/aerospike/aerospike.conf
```

```
...
namespace kurz {
  memory-size 1G # In-memory size
  default-ttl 0  # Data will stay for ever
  storage-engine memory

  # Disk backup
  storage-device device {
    file /opt/aerospike/data/kurz.dat # Target file
    filesize 1G         # Backup file size
    data-in-memory true # In-memory and on disk
  }
}
```

```shell
# Close with :wq
# Restart aerospike
$ sudo service aerospike restart
# Create indexes for search
# Open aerospike client
$ aql
# Create email and id index
> CREATE INDEX email ON kurz.users (email) STRING
> CREATE INDEX id ON kurz.users (id) STRING
# Exit
> exit
$
```

#### kurz

```shell
# Clone the code
$ git clone https://oltdaniel/kurz
$ cd kurz
# Update makefile
$ vim Makefile
# Now change `HOME_DIR` to your home folder
# Then close with :wq
$ 
# Install dependencies
$ make util.install
# Generate API secret
$ make util.key
# Link kurz to go project folder
$ make link
# Add nginx configuration
$ make nginx
# Restart nginx
$ sudo service nginx restart
```

#### Run

```shell
# Start kurz server
$ make run
```

## Documentation

Documentation can be found [here in `/doc`](https://github.com/oltdaniel/kurz/blob/master/doc/README.md).

## Todo

Please view the todo list [here in `TODO.md`](https://github.com/oltdaniel/kurz/blob/master/TODO.md).

## License

_Just do what you'd like to_

[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/oltdaniel/kurz/blob/master/LICENSE)

#### Credit

[Daniel Oltmanns](https://github.com/oltdaniel) - creator
