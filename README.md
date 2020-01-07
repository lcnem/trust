# Trust

## Docker

https://hub.docker.com/r/lcnem/trust/

## ノードの準備

testnetへ接続するための初期化を行います。

```Shell
$ docker run -it -p 26666:26656 -p 26667:26657 -v ~/.trustd:/root/.trustd -v ~/.trustcli:/root/.trustcli lcnem/trust sh init.sh
```

上記コマンドを実行すると、初期化に必要な情報の記入を求められるので、順に入力していきます。

1. `key-name`を要求されるので、アカウント名として、`alice` などと入力。
2. パスフレーズを要求されるので、適当に入力。
    アカウント復活のためのニーモニックフレーズが表示されるので、テキストファイルなどに記録してください。
3. `moniker` を要求されるので、 `zeus` などと入力。
4. `chain-id` を要求されるので、 `t` などと入力。
5. パスフレーズを要求されるので 2. と同じパスフレーズを入力。
    genesisを含めた必要な情報が、 `~/.trustd` `~/.trustcli` に保存されます。


### mainnet

mainnetに接続するための準備を行います。testnetへの接続が完了していることが前提です。開発中は実行しないでください。

```Shell
$ docker run -it -p 26656:26656 -p 26657:26657 -v ~/.trustd:/root/.trustd -v ~/.trustcli:/root/.trustcli lcnem/trust cp genesis.json ~/.trustd/config/genesis.json
```

## ノードの稼働開始

```Shell
$ docker run -it --name trust_node -p 26666:26656 -p 26667:26657 -p 1317:1317 -v ~/.trustd:/root/.trustd -v ~/.trustcli:/root/.trustcli lcnem/trust trustd start
```

## RESTサーバー起動

ノードが動作しているコンテナへ接続して実行します。

>sudo docker exec -it trust_node trustcli rest-server --chain-id t --trust-node=true --laddr tcp://0.0.0.0:1317


## CLIコマンド

ノードが動作しているコンテナへ接続して実行します。
>$ docker exec -it trust_node sh

シェルに接続しますので、CLIコマンドを直接操作できます。

CLIコマンドの詳細については `trustcli --help` をご覧ください。

## RESTサーバーへのアクセス

通常は [`trust-client-ts`](https://github.com/lcnem/trust-client-ts) クライアントを利用してRESTへアクセスしますが、ダイレクトにアクセスすることも可能です。

例）
>http://localhost:1317/node_info

## その他

trustd と trustcli 間の通信で使用されているポート番号（26656、26657）や外部に公開するRESTサーバのポート番号（1317）は適宜変更してください。

複数のノードやRESTサーバーを複数起動する場合などは、ポート番号の衝突が発生しないように調整してください。

