# pool
Type-safe generic sync.Pool wrapper

```go
var readerPool = pool.New[bufio.Reader]()
...
bufioReader := readerPool.Get()
defer readerPool.Put(bufioReader)
bufioReader.Reset(reader)

```
