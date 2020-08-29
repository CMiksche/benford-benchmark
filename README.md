# benford-benchmark

Testing the performance of checking some numbers in Benford's Law in different languages.

Blog Post: https://blog.m5e.de/post/benchmarking-go-crystal-python-and-js/

## Used libraries

| Language | GitHub Repo |
| --- | --- |
| Python | https://github.com/target/huntlib |
| Crystal | https://github.com/CMiksche/benfordslaw.cr |
| Go | https://github.com/CMiksche/benford |
| JavaScript (node) | https://github.com/CMiksche/benfordslaw |

More information about the test setup can be found in the corresponding language test directory.

## Tests

| Version | Command | Duration |
| --- | --- | --- |
| Python 3.8.5 | time python python/test.py | 1,86 secs |
| Crystal 0.35.1 | time crystal/test | 14,20 millis |
| go version go1.15 linux/amd64 | time go/test | 7,03 millis |
| v14.8.0 | time node javascript/test.js | 164,71 millis |

A note here: Crystal takes currently very very very long to compile. Especially with the `--release` flag. The build with `--release` took too long for me.

### Crystal and Go non compiled

You really should compile crystal and Go projects:

| Command | Duration |
| --- | --- |
| time crystal run test.cr | 4,89 secs |
| time go run test.go | 631,67 millis |