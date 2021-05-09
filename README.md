# Preparation

	go mod vendor

	echo 'This is not a PDF.' > zdogzipexample.pdf


# Run

	go run -mod=vendor ./zipTest.go

The expected output is:

	Successfully written:19	

The file compressed with zstandard should be found at `ziptest.zstd`.
