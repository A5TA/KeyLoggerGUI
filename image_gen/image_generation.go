package imagegen

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	charmaps "github.com/A5TA/keyBindViewer/char_maps"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func Generate_key_images() error {
	// Define the folder where the images will be saved
	folder := "key_images"
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return fmt.Errorf("failed to create folder: %v", err)
	}

	// Load a font
	fontBytes, err := os.ReadFile("image_gen/fonts/Roboto-Regular.ttf") // Path to your .ttf font file
	if err != nil {
		fmt.Println("Error loading font:", err)
		return fmt.Errorf("failed to load font: %v", err)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("Error parsing font:", err)
		return fmt.Errorf("failed to parse font: %v", err)
	}

	// Generate images for each character
	for key := range charmaps.ImageMap {
		err := generateImageForCharacter(key, font, folder)
		if err != nil {
			fmt.Println("Error generating image for character ", key, " :", err)
		}
	}

	return nil
}

func generateImageForCharacter(char string, font *truetype.Font, folder string) error {
	// Create a blank image
	imgWidth, imgHeight := 200, 100
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	// Set background color to white
	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)

	// Prepare the font context
	ctx := freetype.NewContext()
	ctx.SetFont(font)
	ctx.SetFontSize(48)
	ctx.SetDst(img)
	ctx.SetClip(img.Bounds())
	ctx.SetSrc(image.Black) // Set text color to black

	// Measure text width by drawing it into a temporary context
	textWidth, _ := measureTextSize(char) // Measure width of the string

	textHeight := 48 // Approximate text height is the font size

	// Calculate the x and y positions to center the text
	x := (imgWidth - int(textWidth)) / 2
	y := (imgHeight + textHeight) / 2

	// Draw the character onto the image
	_, err := ctx.DrawString(char, freetype.Pt(x, y))
	if err != nil {
		return fmt.Errorf("failed to draw string: %v", err)
	}

	// Create the file name for the image
	filename := fmt.Sprintf("%s/%s.png", folder, charmaps.ImageMap[char])

	// Create the output PNG file
	outFile, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer outFile.Close()

	// Encode the image as PNG
	err = png.Encode(outFile, img)
	if err != nil {
		return fmt.Errorf("failed to encode PNG: %v", err)
	}

	fmt.Printf("Generated image for character '%s' at %s\n", char, filename)
	return nil
}

func measureTextSize(char string) (int, error) {
	textWidth := len(char) * 30 // Each character gets 20px width - estimated
	return textWidth, nil
}
