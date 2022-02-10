package compare

import (
	"image"

	"github.com/corona10/goimagehash"
)

var DefaultImageRange = 0.15

type Image interface {
	Height() int
	Width() int
	Hash() *goimagehash.ExtImageHash
}

func imageHash(img1, img2 Image, checkSize bool, limit float64) (int, bool) {
	if checkSize {
		if h1, h2 := float64(img1.Height()), float64(img2.Height()); h1 > h2*(1+limit) || h1 < h2*(1-limit) {
			return 0, false
		}
		if w1, w2 := float64(img1.Width()), float64(img2.Width()); w1 > w2*(1+limit) || w1 < w2*(1-limit) {
			return 0, false
		}
	}
	defer func() { recover() }()

	distance, err := img1.Hash().Distance(img2.Hash())
	if err != nil {
		return 0, false
	}
	return distance, true
}

func ImageHash(img1, img2 Image, checkSize bool) (int, bool) {
	return imageHash(img1, img2, checkSize, DefaultImageRange)
}

func ImageHashWithRange(img1, img2 Image, limit float64) (int, bool) {
	return imageHash(img1, img2, true, limit)
}

var _ Image = &img{}

type img struct {
	height, width int
	hash          *goimagehash.ExtImageHash
}

func (i *img) Height() int                     { return i.height }
func (i *img) Width() int                      { return i.width }
func (i *img) Hash() *goimagehash.ExtImageHash { return i.hash }

func NewImage(i image.Image, fn func(image.Image) image.Image) Image {
	if i == nil {
		panic("Image object can not be nil")
	}
	bounds := i.Bounds()
	if fn != nil {
		i = fn(i)
	}
	hash, _ := goimagehash.ExtPerceptionHash(i, 32, 32)
	return &img{bounds.Dx(), bounds.Dy(), hash}
}
