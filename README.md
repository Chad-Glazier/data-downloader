## Data Downloader

A simple tool to download and extract data (like compressed CSV files) from URLs. E.g.,

```shell
data-downloader https://some.url.with/a/data/file.ext
```

will download `file.ext`, uncompress it with an appropriate algorithm (if necessary).

Currently, the supported compression algorithms are:
- `tar`

## Installation

Currently, this tool can only be installed with [Go](https://go.dev). Verify that you have Go installed
with 

```shell
go version
```

If you get an error, then you need to install Go from their [website](https://go.dev/dl/). Once you have Go 
installed, you can install this tool with

```shell
go install github.com/Chad-Glazier/data-downloader@latest
```

in your command line. Verify that the tool was installed with

```shell
data-downloader version
```

which should print a number.
