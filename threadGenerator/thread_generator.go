package threadGenerator

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"math"
	"os"
	"time"

	"github.com/disintegration/imaging"
)

type (
	Nail = image.Point

	ThreadGenerator struct {
		nailsQuantity     int
		imgSize           int
		maxPaths          int
		startingNail      int
		minimumDifference int
		brightnessFactor  int
		imageName         string
		imageContrast     float64
		pathsDictionary   map[string][]Nail
		pathsList         []Path
		nailsList         []Nail
	}

	Path struct {
		StartingNail int
		EndingNail   int
	}

	Args struct {
		NailsQuantity     int
		ImgSize           int
		MaxPaths          int
		StartingNail      int
		MinimumDifference int
		BrightnessFactor  int
		ImageName         string
	}

	OutputStats struct {
		TotalLines   int
		ThreadLength int
		TotalTime    time.Duration
	}
)

func (tg *ThreadGenerator) getDefaults() {
	tg.nailsQuantity = 300
	tg.imgSize = 800
	tg.maxPaths = 10000
	tg.startingNail = 0
	tg.minimumDifference = 10
	tg.brightnessFactor = 50
	tg.imageContrast = 40
}

func (tg *ThreadGenerator) mergeArgs(args Args) error {
	tg.getDefaults()

	if args.NailsQuantity > 0 {
		tg.nailsQuantity = args.NailsQuantity
	}
	if args.ImgSize > 0 {
		tg.imgSize = args.ImgSize
	}
	if args.MaxPaths > 0 {
		tg.maxPaths = args.MaxPaths
	}
	if args.StartingNail > 0 {
		tg.startingNail = args.StartingNail
	}
	if args.MinimumDifference > 0 {
		tg.minimumDifference = args.MinimumDifference
	}
	if args.BrightnessFactor > 0 {
		tg.brightnessFactor = args.BrightnessFactor
	}

	if args.ImageName != "" {
		tg.imageName = args.ImageName
	} else {
		return errors.New("Image name is required")
	}

	return nil
}

func (tg *ThreadGenerator) Generate(args Args) (*OutputStats, error) {
	start := time.Now()
	err := tg.mergeArgs(args)
	if err != nil {
		return nil, err
	}

	sourceImage, err := tg.getSourceImage()
	if err != nil {
		return nil, err
	}

	nailsList := tg.getNailsListFromImage(sourceImage)

	tg.computePathsListFromImage(sourceImage, nailsList)

	return &OutputStats{
		TotalLines:   len(tg.pathsList),
		ThreadLength: 0,
		TotalTime:    time.Since(start),
	}, nil
}

func (tg *ThreadGenerator) getSourceImage() (*image.NRGBA, error) {
	file, err := os.Open(tg.imageName)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	imgGray := imaging.Grayscale(img)

	imgGray = imaging.AdjustContrast(imgGray, float64(tg.imageContrast))

	imgSquare := imgGray
	bounds := imgSquare.Bounds()
	if bounds.Dx() != bounds.Dy() {
		imgSquare = imaging.CropAnchor(imgSquare, bounds.Dx(), bounds.Dx(), imaging.Center)
	}

	// Crop it into a circle
	circleImg := image.NewRGBA(bounds)
	midPoint := bounds.Dx() / 2
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			xx, yy := float64(x-midPoint), float64(y-midPoint)
			if xx*xx+yy*yy <= float64(midPoint*midPoint) {
				circleImg.Set(x, y, imgSquare.At(x, y))
			} else {
				circleImg.Set(x, y, color.RGBA{255, 255, 255, 255})
			}
		}
	}

	circleImgMin := imaging.Resize(circleImg, tg.imgSize, tg.imgSize, imaging.Lanczos)

	return circleImgMin, nil
}

// getNailsListFromImage generates a list of nails from the source image in a circle
func (tg *ThreadGenerator) getNailsListFromImage(sourceImage image.Image) []Nail {
	centerX := sourceImage.Bounds().Dx() / 2
	centerY := sourceImage.Bounds().Dy() / 2
	radius := math.Min(float64(centerX), float64(centerY))
	tg.nailsList = make([]image.Point, tg.nailsQuantity)
	for i := 0; i < tg.nailsQuantity; i++ {
		alpha := float64(i) * 2 * math.Pi / float64(tg.nailsQuantity)
		x := centerX + int(radius*math.Cos(alpha))
		y := centerY + int(radius*math.Sin(alpha))
		tg.nailsList[i] = Nail{X: x, Y: y}
	}
	return tg.nailsList
}

// computePathsListFromImage generates a list of paths from the source image
func (tg *ThreadGenerator) computePathsListFromImage(sourceImage image.Image, nailsList []Nail) []Path {
	sourceImageBounds := sourceImage.Bounds()
	canvas := image.NewGray(sourceImageBounds)
	for y := sourceImageBounds.Min.Y; y < sourceImageBounds.Max.Y; y++ {
		for x := sourceImageBounds.Min.X; x < sourceImageBounds.Max.X; x++ {
			canvas.Set(x, y, sourceImage.At(x, y))
		}
	}

	tg.generateDictionary(nailsList)

	var nailIndex = tg.startingNail
	var pathsList = []Path{}
	usedPaths := make(map[string]bool)

	for i := 0; i < tg.maxPaths; i++ {
		var maxWeight = 0
		var maxLine = []image.Point{}
		var maxnailIndex = 0

		for nextnailIndex := 0; nextnailIndex < len(nailsList); nextnailIndex++ {
			if nailIndex == nextnailIndex {
				continue
			}

			difference := int(math.Abs(float64(nextnailIndex) - float64(nailIndex)))

			if difference < tg.minimumDifference || difference > (len(nailsList)-tg.minimumDifference) {
				continue
			}

			if _, exists := usedPaths[tg.getPairKey(nextnailIndex, nailIndex)]; exists {
				continue
			}

			line := tg.pathsDictionary[tg.getPairKey(nailIndex, nextnailIndex)]
			weight := len(line) * 255

			for _, pixelPosition := range line {
				pixelColor := canvas.GrayAt(pixelPosition.X, pixelPosition.Y).Y
				weight -= int(pixelColor)
			}

			weight = weight / len(line)

			if weight > maxWeight {
				maxWeight = weight
				maxLine = line
				maxnailIndex = nextnailIndex
			}

		}

		if nailIndex == maxnailIndex {
			break
		}

		usedPaths[tg.getPairKey(nailIndex, maxnailIndex)] = true
		pathsList = append(pathsList, Path{nailIndex, maxnailIndex})
		nailIndex = maxnailIndex

		// Brighthen brightness of chosen line
		for _, pixelPosition := range maxLine {
			var pixel = int(canvas.GrayAt(pixelPosition.X, pixelPosition.Y).Y)
			value := uint8(min(255, pixel+tg.brightnessFactor))
			canvas.SetGray(pixelPosition.X, pixelPosition.Y, color.Gray{value})
		}

	}
	tg.pathsList = pathsList
	return pathsList
}

// GenerateDictionary generates a dictionary of all possible lines between nails
// It's way faster to generate all possible lines at the beginning than to calculate them on the fly
func (tg *ThreadGenerator) generateDictionary(nailsList []image.Point) map[string][]Nail {
	nailsQuantity := len(nailsList)
	tg.pathsDictionary = make(map[string][]Nail, nailsQuantity*(nailsQuantity-1)/2)

	for i := 0; i < nailsQuantity; i++ {
		for j := i + 1; j < nailsQuantity; j++ {
			tg.pathsDictionary[tg.getPairKey(i, j)] = tg.bresenham(nailsList[i], nailsList[j])
		}
	}
	return tg.pathsDictionary
}

// Bresenham's line algorithm - https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
// Returns a list of points between two points
func (threadGen *ThreadGenerator) bresenham(startPoint, endPoint image.Point) []image.Point {
	xDifference := threadGen.abs(endPoint.X - startPoint.X)
	yDifference := -threadGen.abs(endPoint.Y - startPoint.Y)

	signX, signY := -1, -1

	// Determine direction for X
	if startPoint.X < endPoint.X {
		signX = 1
	}

	// Determine direction for Y
	if startPoint.Y < endPoint.Y {
		signY = 1
	}

	error := xDifference + yDifference

	var linePoints []image.Point
	// Continue until end point is reached
	for {
		linePoints = append(linePoints, startPoint)
		if startPoint == endPoint {
			break
		}
		errorDouble := 2 * error

		// Handle X direction
		if errorDouble >= yDifference {
			error += yDifference
			startPoint.X += signX
		}

		// Handle Y direction
		if errorDouble <= xDifference {
			error += xDifference
			startPoint.Y += signY
		}
	}
	return linePoints
}

func (tg *ThreadGenerator) abs(x int) int {
	return int(math.Abs(float64(x)))
}

// getPairKey returns a key for a map of lines between two points
func (tg *ThreadGenerator) getPairKey(a, b int) string {
	switch {
	case a < b:
		return fmt.Sprintf("%d:%d", a, b)
	case a > b:
		return fmt.Sprintf("%d:%d", b, a)
	default:
		return fmt.Sprintf("%d:%d", b, a)
	}
}

func (tg *ThreadGenerator) GeneratePathsImage() (image.Image, error) {
	if len(tg.pathsDictionary) == 0 {
		return nil, errors.New("Dictionary is empty")
	}

	pathsImage := image.NewGray(image.Rect(0, 0, tg.imgSize, tg.imgSize))

	for x := 0; x < tg.imgSize; x++ {
		for y := 0; y < tg.imgSize; y++ {
			pathsImage.SetGray(x, y, color.Gray{255})
		}
	}

	for i := 0; i < len(tg.pathsList); i++ {
		line := tg.pathsDictionary[tg.getPairKey(tg.pathsList[i].StartingNail, tg.pathsList[i].EndingNail)]
		for _, point := range line {
			currentValue := pathsImage.GrayAt(point.X, point.Y).Y
			newValue := max(int(currentValue)-20, 0)
			pathsImage.SetGray(point.X, point.Y, color.Gray{uint8(newValue)})
		}
	}

	return pathsImage, nil
}

func (tg *ThreadGenerator) GetPathsList() []Path {
	return tg.pathsList
}