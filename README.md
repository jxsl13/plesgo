# plesgo


Plesgo is a generated restful client for the netcup.de Plesk api.

It is a package to play around with the Plsk rest api.


# Updating

Example swagger file location: `https://a2e73.webhosting.systems/api/v2/swagger.yml`

Fetch a new swagger file and update the local one.

Execute to cleanup the old generated client and generate a new one:

```shell
make
```


# Requirements
- `curl`to download *go-swagger*
- `jq` to install *go-swagger*

Installing *go-swagger* as `swagger`
```shell
make install
```


