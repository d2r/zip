package main

import (
	"fmt"
	"io"
	"os"

	"github.com/klauspost/compress/zstd"
	//"github.com/DataDog/zstd"
)

// Compress input to output.
func Compress(in io.Reader, out io.Writer) error {
	enc, err := zstd.NewWriter(out)
	if err != nil {
		fmt.Printf("newWriter error:%v\n", err)
		return err
	}
	written, err := io.Copy(enc, in)
	if err != nil {
		fmt.Printf("io.Copy error:%v\n", err)
		enc.Close()
		return err
	}
	fmt.Printf("Successfully written:%v\n", written)
	return enc.Close()
}

func Decompress(in io.Reader, out io.Writer) error {
    d, err := zstd.NewReader(in)
    if err != nil {
        return err
    }
    defer d.Close()
    
    // Copy content...
    _, err = io.Copy(out, d)
    return err
}

func main() {
	file, err := os.Open("zdogzipexample.pdf")
	if err != nil {
		fmt.Println(err)
	}

	outfile, err := os.Create("ziptest.zstd")
	if err != nil {
		fmt.Println(err)
	}

	//	test, err := ioutil.ReadAll(file)

	Compress(file, outfile)

	file.Close()
	outfile.Close()

	outfile2, err := os.Open("ziptest.zip")
	if err != nil {
		fmt.Println(err)
	}

	outtest, err := os.Create("ziptestdecompress.pdf")
	if err != nil {
		fmt.Println(err)
	}


	Decompress(outfile2, outtest)
}
