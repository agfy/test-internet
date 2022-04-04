# test-internet
**Command Line Interface to Test Internet Speed using [speedtest.net](http://www.speedtest.net/) and [fast.com](http://www.fast.com/)**.

### Installation

```bash
$ go install github.com/agfy/test-internet@latest
```

### Usage

```bash
$ test-internet --help
Usage of test-internet:
  -fast
    	use -fast for www.fast.com speed evaluation
  -speedtest
    	use -speedtest for www.speedtest.net speed evaluation
```

### Test Internet Speed

```bash
$ test-internet -speedtest
Started evaluation on www.speedtest.net
Download: 53.43 Mbit/s
Upload: 85.48 Mbit/s
```