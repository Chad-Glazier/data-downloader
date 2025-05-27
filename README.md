## File Downloader & Decompressor

A simple tool to download and extract files from a URL. E.g.,

```sh
fdd https://some.url.with/a/data/file.ext
```

will download `file.ext` and uncompress it with an appropriate algorithm (if
necessary).

Currently, the supported compression algorithms are:
- `tar`
- `zip`
- `bzip2`
- `flate`
- `gzip`
- `zlib`

## Installation

Currently, this tool can only be installed with [Go](https://go.dev). Verify
that you have Go installed with

```sh
go version
```

(Execute this in your CLI, which should be the command prompt or PowerShell if
you're using Windows, or the terminal on Mac or Linux.)

If you get an error, then you need to install Go from their
[website](https://go.dev/dl/). Once you have Go installed, you can install this
tool by running the following command.

```sh
go install github.com/Chad-Glazier/fdd
```

Verify that the tool was installed with

```sh
fdd version
```

which should display the current version.
