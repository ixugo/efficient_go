package screenshot

import (
	"fmt"
	"image/png"
	"os"
	"testing"

	"github.com/kbinani/screenshot"
)

func TestShot(t *testing.T) {
	// n := screenshot.NumActiveDisplays()

	for i := 0; i < 1; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}

}
