package example

import "io"

// Image is a fake image type
type Image struct{}

func Decode(r io.Reader) (*Image, error) {
	return &Image{}, nil
}

func Crop(img *Image, x1, y1, x2, y2 int) error {
	return nil
}

func Encode(img *Image, w io.Writer) error {
	return nil
}
