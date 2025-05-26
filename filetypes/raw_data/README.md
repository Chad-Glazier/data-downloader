The `mimeData.json` file is from [here](https://github.com/patrickmccallum/mimetype-io/blob/master/src/mimeData.json). The file is not original.

The `mimeCompression.csv` file is original.

The `build.ts` file is written for the [Deno runtime](https://deno.com) and builds a group of `.go` files (under the package name `filetypes`) with some `map[string]string` objects that relate:
- MIME types to file extensions
- file extensions to MIME types
- MIME types to appropriate compression algorithms
