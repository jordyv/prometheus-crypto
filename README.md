# Prometheus exporter for cryptocurrency rates from the Litebit API #

## Installation ##

### Using Docker ###

```
$ docker run -p 8888:8888 jordyversmissen/prometheus-crypto -interval 10s lsk part btc
```

## Usage ##

```
./prometheus-crypto-sync -interval 10s lsk btc part
```

Go to 'http://localhost:8888/metrics' to see the metrics.

Example:

```
$ curl http://localhost:8888/metrics
# TYPE crypto_buy_rate gauge
crypto_buy_rate{coin="btc"} 5439.38
crypto_buy_rate{coin="lsk"} 2.59161
crypto_buy_rate{coin="part"} 2.23512
# HELP crypto_sell_rate Sell rate
# TYPE crypto_sell_rate gauge
crypto_sell_rate{coin="btc"} 5333.23
crypto_sell_rate{coin="lsk"} 2.38199
crypto_sell_rate{coin="part"} 2.02208
```
