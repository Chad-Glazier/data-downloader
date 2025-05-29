The `mimeData.json` file is from [here](https://github.com/patrickmccallum/mimetype-io/blob/master/src/mimeData.json). The file is not original, credit
should be given to Patrick McCallum and the other contributors to that repository for assembling the data.

The `mimeCompression.csv` file was put together by yours truly, so it might not be perfect and should not be treated as an authoritative source.

The `build.ts` file is written for the [Deno runtime](https://deno.com) and builds a group of `.go` files (under the package name `filetypes`) with some `map[string]string` objects that relate:
- MIME types to file extensions
- file extensions to MIME types
- MIME types to appropriate compression algorithms
