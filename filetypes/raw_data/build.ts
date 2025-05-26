const NEW_FILE_HEAD = `package filetypes

`

async function main() {
	const text = await Deno.readTextFile("./raw_data/mimeData.json")
	const data: {
		name: string
		fileTypes: string[]
	}[] = JSON.parse(text)

	//
	// ExtToMime.go
	//

	const extToMimeFile = await Deno.open("ExtToMime.go", {
		create: true,
		write: true
	})
	const encoder = new TextEncoder()

	extToMimeFile.writeSync(encoder.encode(NEW_FILE_HEAD))
	extToMimeFile.writeSync(encoder.encode(
		"var ExtToMime = map[string]string{\n"
	))

	const extsIncluded = new Set<string>()

	for (const mime of data) {
		for (const extension of mime.fileTypes) {
			if (extsIncluded.has(extension)) {
				continue
			}
			extToMimeFile.writeSync(encoder.encode(
				`\t"${extension}": "${mime.name}",\n`
			))
			extsIncluded.add(extension)
		}
	}
	extToMimeFile.writeSync(encoder.encode(
		"}\n"
	))
	extToMimeFile.close()

	//
	// MimeToExt.go
	//

	const mimeToExtFile = await Deno.open("MimeToExt.go", {
		create: true,
		write: true
	})

	mimeToExtFile.writeSync(encoder.encode(NEW_FILE_HEAD))
	mimeToExtFile.writeSync(encoder.encode(
		"var MimeToExt = map[string]string{\n"
	))
	for (const mime of data) {
		const lastExt = mime.fileTypes.length == 0 ?
			""
			: mime.fileTypes[mime.fileTypes.length - 1]
		mimeToExtFile.writeSync(encoder.encode(
			`\t"${mime.name}": "${lastExt}",\n`
		))
	}
	mimeToExtFile.writeSync(encoder.encode(
		"}\n"
	))
	mimeToExtFile.close()

	//
	// MimeToCompression.go
	//

	const compressionCsv = await Deno.readTextFile("raw_data/mimeCompression.csv")
	let lines = compressionCsv.split("\r\n")
	if (lines.length == 1) {
		lines = compressionCsv.split("\n")
	}
	if (lines.length == 1) {
		lines = compressionCsv.split("\r")
	}
	lines.shift()
	lines.pop()

	const mimeToCompressionFile = await Deno.open("MimeToCompression.go", {
		create: true,
		write: true
	})

	mimeToCompressionFile.writeSync(encoder.encode(NEW_FILE_HEAD))
	mimeToCompressionFile.writeSync(encoder.encode(
		"var MimeToCompression = map[string]string{\n"
	))

	for (const line of lines) {
		const [ mime, compressed, goPackage ] = line.split(",")
		if (compressed === "FALSE") continue
		mimeToCompressionFile.writeSync(encoder.encode(
			`\t"${mime}": "${goPackage === "" ? "unknown" : goPackage}",\n`
		))
	}
	mimeToCompressionFile.writeSync(encoder.encode(
		"}\n"
	))
	mimeToCompressionFile.close()
}

main()
