# Dump plugin

## Overview

This plguin just dumps current input to `file`.

## Example Configuration

```
{
        "cniVersion": "0.3.1",
        "name": "mynet",
        "plugins": [
                {
                        "type": "ptp",
                        "ipMasq": true,
                        "ipam": {
                                "type": "host-local",
                                "subnet": "172.16.30.0/24",
                                "routes": [
                                        {
                                                "dst": "0.0.0.0/0"
                                        }
                                ]
                        }
                },
                {
                        "type": "portmap",
                        "capabilities": {"portMappings": true},
                        "externalSetMarkChain": "KUBE-MARK-MASQ"
                },
                {
                        "type": "dump",
                        "file": "/tmp/cni_dump.log"
                }
        ]
}
```

## Config Reference

* `file` (string, optional): file path to output. default is `/tmp/cni_dump`
