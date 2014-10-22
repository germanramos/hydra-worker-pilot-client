hydra-worker-pilot-client
=========================

Piloting client request through specific instances.<br />
No new maps will be generated.
Worker for Hydra v3.2.x<br />

# Installation

## Ubuntu/Debian

Add PPAs for:  
https://launchpad.net/~chris-lea/+archive/libpgm  
https://launchpad.net/~chris-lea/+archive/zeromq  
  
and run:  
```
sudo dpkg -i hydra-worker-pilot-client-1-0.x86_64.deb
sudo apt-get install -f
```
## CentOS/RedHat/Fedora
```
sudo yum install libzmq3-3.2.2-13.1.x86_64.rpm hydra-worker-pilot-client-1-0.x86_64.rpm
```

# Configuration

In apps.json:

- Name: "PilotClient"
- Arguments:
  - clientFilterField: the query param sent by the client to drive its request
  - instanceFilterField: the instance attribute for filtering the compatible instances with client request
  - matchers: array of matchers
    - matcher: establishes an association between client and intances
      - instanceFilterPattern: regular expression for matching instances
      - clientFilterPatterns: regular expression for matching client requests

## Configuration example
```
{
  "worker": "PilotClient",
  "clientFilterField: "client_id",
  "instanceFilterField": "version",
  "matchers": [
  {
    "instanceFilterPattern": "1.*",
    "clientFilterPatterns": ["xe4([0-9]+)", "xe5([0-9]+)"]
  },
  {
    "instanceFilterPattern": "1\.0\.0",
    "clientFilterPatterns": [".*"]
  }
  ]
}
```
In the example, the client requests with, for example, client_id=xe47030 or client_id=xe59613 would be driven to instances with, for example, version=1.1.1 or version=1.9.5 and any other client_id would be driven to instances with version=1.0.0. If something is wrong or no instances have the associated version all initial instances will be returned.

## Service configuration

No additional configuration is needed if running in the same machine that Hydra.  
Tune start file at /etc/init.d/hydra-worker-pilot-client if you run it in a separate machine.

# Run
```
sudo /etc/init.d/hydra-worker-pilot-client start
```

# License

(The MIT License)
