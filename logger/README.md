# Logger

Another log library based on `go-kit/log`

## Why another log library?

Because I need a simpler and tinier log library

## But we usually use logrus, what is the benefit?

Well, you have most of the features of `Logrus`

- Independent logger object
- Level
- WithFields
- Write to another `*os.File`

## But this is only a wrapper, is it fast?

Yeah this is basically only a wrapper of `go-kit/log`.

You can see the benchmark by yourself:
```
BenchmarkSimpleLogger-4           300000              5184 ns/op
BenchmarkLoggerWithFields-4       200000              7095 ns/op
BenchmarkGokitLog-4               500000              2539 ns/op
PASS
ok      github.com/tokopedia/user/lib/log/logger        4.427s
```

The benchmark is quite good, most of the time all type of log should be written below 0.01ms. `time.Now()` is actually killing most of the performance.

You can try to run the benchmark on your machine.