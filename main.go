package main

import (
	"flag"
	"net/http"
	"os"
	"strconv"
	"time"

	api "github.com/jordyv/crypto-sync/litebit-api"

	log "github.com/Sirupsen/logrus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	coins           []string
	collectors      []CoinCollector
	refreshDuration time.Duration
	addr            string
)

type CoinCollector struct {
	CoinCode      string
	BuyHistogram  prometheus.Histogram
	SellHistogram prometheus.Histogram
}

func init() {
	flag.DurationVar(&refreshDuration, "interval", 5*time.Second, "Refresh interval for rates")
	flag.StringVar(&addr, "listen", "127.0.0.1:8888", "Listen address for HTTP endpoint")
	verbose := flag.Bool("verbose", false, "Verbose output")
	flag.Parse()
	coins := flag.Args()
	if len(coins) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	for _, coin := range coins {
		sell := prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:        "crypto_sell_rate",
			Help:        "Sell rate",
			ConstLabels: prometheus.Labels{"coin": coin},
		})
		buy := prometheus.NewHistogram(prometheus.HistogramOpts{
			Name:        "crypto_buy_rate",
			Help:        "Buy rate",
			ConstLabels: prometheus.Labels{"coin": coin},
		})
		collectors = append(collectors, CoinCollector{SellHistogram: sell, BuyHistogram: buy, CoinCode: coin})
	}

	if *verbose {
		log.SetLevel(log.DebugLevel)
	}

	log.Debugln("Interval", refreshDuration)
}

func main() {
	registry := prometheus.NewRegistry()
	for _, collector := range collectors {
		registry.MustRegister(collector.BuyHistogram)
		registry.MustRegister(collector.SellHistogram)
	}

	go func() {
		for {
			for _, collector := range collectors {
				result, err := api.GetMarket(collector.CoinCode)
				if err != nil {
					log.Error("Error: " + err.Error())
					continue
				}

				if buy, err := strconv.ParseFloat(result.Result.Buy, 64); err == nil {
					log.Debugf("Set buy metric to %f for coin '%s'", buy, collector.CoinCode)
					collector.BuyHistogram.Observe(buy)
				}
				if sell, err := strconv.ParseFloat(result.Result.Sell, 64); err == nil {
					log.Debugf("Set sell metric to %f for coin '%s'", sell, collector.CoinCode)
					collector.SellHistogram.Observe(sell)
				}
			}
			time.Sleep(refreshDuration)
		}
	}()

	log.Info("Listen at ", addr)
	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(addr, nil))
}
