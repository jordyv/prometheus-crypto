# Prometheus exporter for cryptocurrency rates from the Litebit API #

## Installation ##

### Using Docker ###

```
$ docker run -p 8888:8888 jordyversmissen/prometheus-crypto
```

## Usage ##

```
./prometheus-crypto-sync -interval 10s lsk btc part
```

Go to 'http://localhost:8888/metrics' to see the metrics.

Example:

```
$ curl http://localhost:8888/metrics

# HELP crypto_buy_rate Buy rate
# TYPE crypto_buy_rate histogram
crypto_buy_rate_bucket{coin="btc",le="0.005"} 0
crypto_buy_rate_bucket{coin="btc",le="0.01"} 0
crypto_buy_rate_bucket{coin="btc",le="0.025"} 0
crypto_buy_rate_bucket{coin="btc",le="0.05"} 0
crypto_buy_rate_bucket{coin="btc",le="0.1"} 0
crypto_buy_rate_bucket{coin="btc",le="0.25"} 0
crypto_buy_rate_bucket{coin="btc",le="0.5"} 0
crypto_buy_rate_bucket{coin="btc",le="1"} 0
crypto_buy_rate_bucket{coin="btc",le="2.5"} 0
crypto_buy_rate_bucket{coin="btc",le="5"} 0
crypto_buy_rate_bucket{coin="btc",le="10"} 0
crypto_buy_rate_bucket{coin="btc",le="+Inf"} 30
crypto_buy_rate_sum{coin="btc"} 163073.2999999999
crypto_buy_rate_count{coin="btc"} 30
crypto_buy_rate_bucket{coin="lsk",le="0.005"} 0
crypto_buy_rate_bucket{coin="lsk",le="0.01"} 0
crypto_buy_rate_bucket{coin="lsk",le="0.025"} 0
crypto_buy_rate_bucket{coin="lsk",le="0.05"} 0
crypto_buy_rate_bucket{coin="lsk",le="0.1"} 0
crypto_buy_rate_bucket{coin="lsk",le="0.25"} 0
crypto_buy_rate_bucket{coin="lsk",le="0.5"} 0
crypto_buy_rate_bucket{coin="lsk",le="1"} 0
crypto_buy_rate_bucket{coin="lsk",le="2.5"} 0
crypto_buy_rate_bucket{coin="lsk",le="5"} 30
crypto_buy_rate_bucket{coin="lsk",le="10"} 30
crypto_buy_rate_bucket{coin="lsk",le="+Inf"} 30
crypto_buy_rate_sum{coin="lsk"} 78.06445999999998
crypto_buy_rate_count{coin="lsk"} 30
crypto_buy_rate_bucket{coin="part",le="0.005"} 0
crypto_buy_rate_bucket{coin="part",le="0.01"} 0
crypto_buy_rate_bucket{coin="part",le="0.025"} 0
crypto_buy_rate_bucket{coin="part",le="0.05"} 0
crypto_buy_rate_bucket{coin="part",le="0.1"} 0
crypto_buy_rate_bucket{coin="part",le="0.25"} 0
crypto_buy_rate_bucket{coin="part",le="0.5"} 0
crypto_buy_rate_bucket{coin="part",le="1"} 0
crypto_buy_rate_bucket{coin="part",le="2.5"} 30
crypto_buy_rate_bucket{coin="part",le="5"} 30
crypto_buy_rate_bucket{coin="part",le="10"} 30
crypto_buy_rate_bucket{coin="part",le="+Inf"} 30
crypto_buy_rate_sum{coin="part"} 66.98975999999998
crypto_buy_rate_count{coin="part"} 30
# HELP crypto_sell_rate Sell rate
# TYPE crypto_sell_rate histogram
crypto_sell_rate_bucket{coin="btc",le="0.005"} 0
crypto_sell_rate_bucket{coin="btc",le="0.01"} 0
crypto_sell_rate_bucket{coin="btc",le="0.025"} 0
crypto_sell_rate_bucket{coin="btc",le="0.05"} 0
crypto_sell_rate_bucket{coin="btc",le="0.1"} 0
crypto_sell_rate_bucket{coin="btc",le="0.25"} 0
crypto_sell_rate_bucket{coin="btc",le="0.5"} 0
crypto_sell_rate_bucket{coin="btc",le="1"} 0
crypto_sell_rate_bucket{coin="btc",le="2.5"} 0
crypto_sell_rate_bucket{coin="btc",le="5"} 0
crypto_sell_rate_bucket{coin="btc",le="10"} 0
crypto_sell_rate_bucket{coin="btc",le="+Inf"} 30
crypto_sell_rate_sum{coin="btc"} 158111.68999999997
crypto_sell_rate_count{coin="btc"} 30
crypto_sell_rate_bucket{coin="lsk",le="0.005"} 0
crypto_sell_rate_bucket{coin="lsk",le="0.01"} 0
crypto_sell_rate_bucket{coin="lsk",le="0.025"} 0
crypto_sell_rate_bucket{coin="lsk",le="0.05"} 0
crypto_sell_rate_bucket{coin="lsk",le="0.1"} 0
crypto_sell_rate_bucket{coin="lsk",le="0.25"} 0
crypto_sell_rate_bucket{coin="lsk",le="0.5"} 0
crypto_sell_rate_bucket{coin="lsk",le="1"} 0
crypto_sell_rate_bucket{coin="lsk",le="2.5"} 30
crypto_sell_rate_bucket{coin="lsk",le="5"} 30
crypto_sell_rate_bucket{coin="lsk",le="10"} 30
crypto_sell_rate_bucket{coin="lsk",le="+Inf"} 30
crypto_sell_rate_sum{coin="lsk"} 71.69573000000003
crypto_sell_rate_count{coin="lsk"} 30
crypto_sell_rate_bucket{coin="part",le="0.005"} 0
crypto_sell_rate_bucket{coin="part",le="0.01"} 0
crypto_sell_rate_bucket{coin="part",le="0.025"} 0
crypto_sell_rate_bucket{coin="part",le="0.05"} 0
crypto_sell_rate_bucket{coin="part",le="0.1"} 0
crypto_sell_rate_bucket{coin="part",le="0.25"} 0
crypto_sell_rate_bucket{coin="part",le="0.5"} 0
crypto_sell_rate_bucket{coin="part",le="1"} 0
crypto_sell_rate_bucket{coin="part",le="2.5"} 30
crypto_sell_rate_bucket{coin="part",le="5"} 30
crypto_sell_rate_bucket{coin="part",le="10"} 30
crypto_sell_rate_bucket{coin="part",le="+Inf"} 30
crypto_sell_rate_sum{coin="part"} 61.10154000000004
crypto_sell_rate_count{coin="part"} 30
```