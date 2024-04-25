package example

import (
	"fmt"
	"io"
	"testing"
)

func Example_crop() {
	var r io.Reader
	img, err := Decode(r)
	if err != nil {
		panic(err)
	}
	err = Crop(img, 0, 0, 20, 20)
	if err != nil {
		panic(err)
	}
	var w io.Writer
	err = Encode(img, w)
	if err != nil {
		panic(err)
	}
	fmt.Println("Wrote out.png")
	//Output:
	// Wrote out.png
}

func TestDecode(t *testing.T) {

}

func TestCrop(t *testing.T) {

}

func TestEncode(t *testing.T) {

}
