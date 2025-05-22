## Data Downloader

A simple tool to download and extract data (like compressed CSV files) from URLs. E.g.,

```sh
data-downloader https://some.url.with/a/data/file.ext
```

will download `file.ext`, uncompress it with an appropriate algorithm (if necessary) and put the
results into `./file.ext` on your machine. Alternatively, a custom filename can be specified with

```sh
data-downloader https://some.url.with/a/data/file.ext -o data.ext
```

## Installation

Currently, this tool can only be installed with [Go](https://go.dev). If you have Go installed, then
you can run

```sh
go install github.com/Chad-Glazier/data-downloader@latest
```

in any command line to install the tool.
