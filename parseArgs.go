package main

import (
	"errors"
	"flag"
	"net/url"
	"strings"
)

const (
	unknown uint8 = iota
	bzip2
	flate
	gzip
	lzw
	zlib
)

type Args struct {
	url                  string
	output               string
	compressionAlgorithm uint8
	decompress           bool
}

// Get the relevant CLI arguments and flags.
func parseArgs() (args *Args, err error) {
	var urlFlag string
	var outputFile string
	var compressionAlgorithm string
	var decompress bool

	flag.StringVar(&urlFlag, "url", "",
		"specifies the URL of the file to download",
	)
	flag.StringVar(&outputFile, "o", "",
		"specifies the location where the file will be downloaded to. If "+
			"not specified, the filename will be based on the URL provided.",
	)
	flag.StringVar(&compressionAlgorithm, "compression", "",
		"used to explicitly state the compression algorithm to use. If not "+
			"specified, the compression algorithm will be inferred. Accepted "+
			"values are:\n\t- bzip2,\n\t- flate,\n\t- gzip,\n\t- lzw, and"+
			"\n\t- zlib.",
	)
	flag.BoolVar(&decompress, "decompress", true,
		"set this to false if you don't want the file to be decompressed.",
	)

	flag.Parse()

	// handle unacceptable arguments

	if urlFlag == "" {
		urlFlag = flag.Arg(0)
		if (urlFlag == "") {
			return nil, errors.New(
				"a URL is required, try something like 'data-downloader " +
				"-url https://your.url/here.zip'",
			)
		}
	}
	
	url, err := url.Parse(urlFlag)
	if err != nil {
		return nil, errors.New("the provided URL is invalid")
	}

	// instantiate the output

	args = &Args{
		url:                  urlFlag,
		output:               outputFile,
		compressionAlgorithm: 0,
		decompress:           decompress,
	}

	// create defaults

	if outputFile == "" {
		pathComponents := strings.Split(url.Path, "/")
		args.output = pathComponents[len(pathComponents)-1]
	}

	switch compressionAlgorithm {
	case "bzip2":
		args.compressionAlgorithm = bzip2
	case "flate":
		args.compressionAlgorithm = flate
	case "gzip":
		args.compressionAlgorithm = gzip
	case "lzw":
		args.compressionAlgorithm = lzw
	case "zlib":
		args.compressionAlgorithm = zlib
	default:
		args.compressionAlgorithm = unknown
	}

	return args, nil
}
