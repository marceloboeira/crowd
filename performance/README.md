# Performance
> Metrics and insights

## Running

* `make` - Runs a performance test with the profiles described below agains a running instance of `crowd`.


## SQS
> Pushing a message to AWS SQS wtih crowd

### 500 req/s
> Running crowd at 500 req/s

```
Latencies     [mean, 50, 95, 99, max]  5.425418ms, 7.350279ms, 20.691278ms, 86.623242ms, 250.673111ms
Success       [ratio]                  100.00%
```

(Running on an EC2 machine, for SQS latency puporses).
