package photos

import (
	"github.com/disintegration/imaging"
	"image"
	"os"
	"path/filepath"
	"strings"
)

// Format is an image file format.
type PhotoType int

// Image file formats.
const (
	THUMB PhotoType = iota
	LANDSCAPE
	SQUARE
	PORTRAIT
)

var GenTypesNames = map[PhotoType]string{
	THUMB: "thumb",
	LANDSCAPE:  "landscape",
	SQUARE:  "square",
	PORTRAIT: "portrait",
}

func (gt PhotoType) String() string {
	return GenTypesNames[gt]
}

func CreateDirs(destDir string, subDir bool) error {
	var err error
	if err = os.MkdirAll(destDir, 0744); err != nil {
		return err
	}
	for _,name := range GenTypesNames {
		if err = os.MkdirAll(filepath.Join(destDir,name), 0744); err != nil {
			return err
		}
	}
	return nil
}

func (e Editor) GenerateAll(srcFile, destDir string, subDir bool) error {
	return e.Generate(srcFile, destDir, subDir, THUMB, LANDSCAPE, SQUARE, PORTRAIT)
}

func (e Editor) Generate(srcFile, destDir string, subDir bool, types ...PhotoType) error {
	img, err := imaging.Open(srcFile, imaging.AutoOrientation(true))
	if err != nil {
		return err
	}
	//First we should resize image to max size if needed:
	size := img.Bounds()
	if size.Dx() > e.maxWidth {
		img = imaging.Fit(img, e.maxWidth, size.Dy(), imaging.Lanczos)
	}


	//now generate the various images... (this should be parallelised)
	//generate thumb:
	for _, gt := range types {
		var rect image.Rectangle
		switch gt {
		case THUMB:
			rect = e.thumb
		case LANDSCAPE:
			rect = e.landscape
		case SQUARE:
			rect = e.square
		case PORTRAIT:
			rect = e.portrait
		}
		dstImg := resize(img, rect, imaging.Center)
		imaging.Save(dstImg, dstFile(srcFile, destDir, gt, subDir), imaging.JPEGQuality(e.quality))
	}
	return nil
}

func dstFile(srcFile, destDir string, gt PhotoType, subDir bool) string {
	base := filepath.Base(srcFile)
	if subDir {
		return filepath.Join(destDir, gt.String(), base)
	} else {
		suffix := filepath.Ext(srcFile)
		fname := strings.TrimSuffix(base, suffix)
		return filepath.Join(destDir, fname + "_" + gt.String() + suffix)
	}
}

func containsType(types []PhotoType, check PhotoType) bool {
	for _,v := range types {
		if v == check {
			return true
		}
	}
	return false
}

func resize(image image.Image, dst image.Rectangle, anchor imaging.Anchor) image.Image {
	return imaging.Fill(image, dst.Max.X, dst.Max.Y, anchor, imaging.Lanczos)
}


/*
func (e *Editor) createThumb(image image.Image, out io.WriteCloser, anchor imaging.Anchor) {
	defer out.Close()
	img := imaging.Fill(image,e.thumb.Max.X, e.thumb.Max.Y, anchor, imaging.Lanczos)
	imaging.Encode(out, img, imaging.JPEG, imaging.JPEGQuality(e.quality))
}

func (e *Editor) createLandscape(image image.Image, out io.WriteCloser) {
	img := imaging.Fill(image, e.landscape.Max.X, e.thumb.Max.Y, imaging.Center, imaging.Lanczos)
	defer out.Close()

}

func (e *Editor) CreateSquare(in io.ReadCloser, out io.WriteCloser, anchor CropAnchor) {

}

func (e *Editor) CreatePortrait(in io.ReadCloser, out io.WriteCloser, anchor CropAnchor) {

}
 */