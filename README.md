# Trust

## Docker

https://hub.docker.com/r/lcnem/trust/

## Testnet

```Shell
$ docker run -it -p 26656:26656 -p 26657:26657 -v ~/.trustd:/root/.trustd -v ~/.trustcli:/root/.trustcli lcnem/trust bash init.sh
```

## Mainnet

```Shell
$ docker run -it -p 26656:26656 -p 26657:26657 -v ~/.trustd:/root/.trustd -v ~/.trustcli:/root/.trustcli lcnem/trust cp genesis.json ~/.trustd/config/genesis.json
```

## Start

```Shell
$ docker run -it -p 26656:26656 -p 26657:26657 -v ~/.trustd:/root/.trustd -v ~/.trustcli:/root/.trustcli lcnem/trust trustd start
```