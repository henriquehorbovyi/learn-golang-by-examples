package os

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	webcam "github.com/blackjack/webcam"
)

// TakePicture starts the webcam and saves the bytes to a png file
func TakePicture() {
	cam, err := webcam.Open("/dev/video0")

	if err != nil {
		fmt.Println("Error while opening the webcam", err)
	}

	defer cam.Close()
	cam.StartStreaming()

	cam.WaitForFrame(1000)

	frame, _ := cam.ReadFrame()
	saveToPng("picture", frame)
}

func saveToPng(fileName string, pictureData []byte) {

	img, _, err := image.Decode(bytes.NewReader(pictureData))
	if err != nil {
		log.Fatalln(err)
	}

	out, _ := os.Create("./" + fileName + ".jpeg")
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 80

	err = jpeg.Encode(out, img, &opts)
	if err != nil {
		log.Println(err)
	}
}

func check(err error) bool {
	if err != nil {
		fmt.Println("Error >> ", err)
	}
	return err != nil
}
