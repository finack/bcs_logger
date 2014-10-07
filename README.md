## BCS Logger

The Brewery Control System (BCS) by embeddedcc.com allows brewers to automate processes such as brewing and fermentation.

This application extends the BCS by recording state of the device at a given moment and recording it for longer term analysis.

Currently this logger only supports logging to Librato.com

### Application Overview

This is an application that is designed to run from a small embedded computer such as the raspberry pi. Typically this is run on the same network as the BCS and all of this is behind a firewall.

`bin/bcs_logger` starts up and repeatably polls a number of BCS device's statistics and publishes those statistics to Librato.

The application `bin/bcs_logger` takes a number of [environment variables](http://12factor.net/config) that control its behavior.

```sh
# How often we should poll the different BCS hosts in seconds
export BCS_POLL=60

# Comma separated list of hostnames
export BCS_HOSTS='hostname1,hostname2'

# Web basic auth username (if enabled)
export BCS_USER='admin'

# Web basic auth password (if enabled)
export BCS_PASS='password'

# API Key for your librato account
export LIBRATO_API_KEY='1234'
```

### Internals

http://wiki.embeddedcc.com/index.php/Ultemp.dat
