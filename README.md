## File Downloader & Decompressor

A simple tool to download and extract files from a URL. E.g.,

```sh
fdd get https://some.url.with/a/data/file.ext
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

Notably, `lzw` is not currently supported.

## Installation

This tool can be installed by downloading one of the releases from [here](https://github.com/Chad-Glazier/fdd/releases/),
or by using `go install`:

```sh
go install github.com/Chad-Glazier/fdd@latest
```

Verify that the tool was installed with

```sh
fdd
```

which should describe the available commands.
