# vcd-go (Still in progress)

Simple go cli tool to communicate with VCD (VMWare Cloud Director).

## Configuration

Create a config file with the following contents in $HOME/.config/vcd/config.yaml or specify a config file with the `--config-file` flag.

```yaml
user: "SOME_USER"
token: "SOME_ACCESS_TOKEN"
url: "SOME_API_URL"
organisation: "SOME_ORG"
vdc: "SOME_VDC"
```

## Features

- Authentication with access_token in configfile
- List VM's
- List VDC's
