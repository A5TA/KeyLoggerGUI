package main

import (
	"errors"
	"fmt"
	"log"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	charmaps "github.com/A5TA/keyBindViewer/char_maps"
	imagegen "github.com/A5TA/keyBindViewer/image_gen"
	"github.com/MarinX/keylogger"
)

func main() {
	//Create The images to be shown on the App
	err := imagegen.Generate_key_images()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("KeyLogger Images generated successfully!")

	// Create a new app and window using Fyne
	a := app.New()
	w := a.NewWindow("Keylogger App")

	// Create a label to display the key information
	keyInput := widget.NewLabel("Press keys and see them here!")

	// Set the window content to a VBox container with the label
	w.SetContent(container.NewVBox(
		keyInput,
		widget.NewButton("Start Keylogger", func() {
			go startKeylogger(keyInput) // Start keylogger in a goroutine
		}),
	))

	// Show the window and run the Fyne app
	w.ShowAndRun()
}

func startKeylogger(keyInput *widget.Label) {

	//find the keyboard device
	keyboard := keylogger.FindKeyboardDevice()

	//Check if a keyboard is located
	if len(keyboard) <= 0 {
		fmt.Println("Ther is no keyboard found")
		return
	}
	fmt.Println("A keyboard has been loacted at ", keyboard)

	//Time to init a keylogger for the keyboard
	kLogger, err := keylogger.New(keyboard)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer kLogger.Close() // close the logger instance

	events := kLogger.Read()
	var shiftPressed = false //boolean for knowing if the shift key is being pressed to include special characters
	var keyPressed = ""
	for e := range events {
		switch e.Type {
		case keylogger.EvKey:
			if e.KeyPress() {
				//This condition is only needed to enable the shift pressed condition
				if e.KeyString() == "L_SHIFT" {
					//lets enable the boolean
					shiftPressed = true
				}
			}

			if e.KeyRelease() {
				if e.KeyString() == "L_SHIFT" {
					//lets disable the boolean
					shiftPressed = false
				} else {
					keyPressed, err = displayImage(e.KeyString(), shiftPressed)

					if err == nil {
						// Display the pressed key (or special character) in the Fyne window
						updateLabel(keyInput, keyPressed)
					} else {
						fmt.Println("Error Occured, " + err.Error())
					}
				}
			}
			break
		}
	}
}

func displayImage(key string, shiftPressed bool) (string, error) {
	if shiftPressed {
		//This key needs to be converted to its special character
		convertedKey, exists := charmaps.CharConvertMap[key]

		if exists {
			fmt.Printf("the key %s is found, converting to %s\n", key, convertedKey)
			key = convertedKey //continue with the rest of the image path lookup
		}
	}
	imagePath, exists := charmaps.ImageMap[key]

	if !exists {
		fmt.Printf("the key %s cannot be found\n", key)
		return "", errors.New("key is not found")
	}

	return imagePath, nil
	// fmt.Println("[Output] Image for ", imagePath)
}

// Update the label with the pressed key or special character
func updateLabel(label *widget.Label, key string) {
	// Update the label's text to display the key
	label.SetText(fmt.Sprintf("Key Pressed: %s", key))
}
