# etcd terminal user interface

What is [etcd](https://etcd.io/)?
> A distributed, reliable key-value store for the most critical data of a distributed system

## Usage

```bash
# version
$ etcd-tui --version
# ask for help
$ etcd-tui --help
```

### Terminal user interface

```bash
$ etcd-tui localhost:2379 --user <user> --password <password>
```

## Demo

```bash
$ make demo-etcd-docker-up
$ make demo-connect
```
