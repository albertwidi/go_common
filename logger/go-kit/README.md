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
BenchmarkSimpleLogger-4                   200000              5457 ns/op            1192 B/op         28 allocs/op
BenchmarkLoggerWithFields-4               200000              7377 ns/op            1680 B/op         35 allocs/op
BenchmarkLoggerWithLongFields-4            50000             23500 ns/op            4776 B/op         90 allocs/op
BenchmarkErrors-4                         200000              7168 ns/op            1328 B/op         32 allocs/op
BenchmarkLongErrorsFields-4               100000             19297 ns/op            3627 B/op         68 allocs/op
BenchmarkErrorsWithFields-4               200000              8565 ns/op            1848 B/op         39 allocs/op
BenchmarkGokitLog-4                       500000              2725 ns/op             696 B/op         15 allocs/op
PASS
ok      github.com/albert-widi/go_common/logger/go-kit  11.032s
```

The benchmark is quite good, most of the time all type of log should be written below 0.01ms. `time.Now()` is actually killing most of the performance.

You can try to run the benchmark on your machine.

## Errors Type

Errors if faster than using `WithFields`, use `errors.Fields` in `errors` package to add more context to your `error`.

But you can also combone `WithFields` with `errors.Fields`. It is much slower, so it is recommended to just use `errors.Fields`

