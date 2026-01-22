package main

import (
	"context"
	"io"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/oh-tarnished/elevenlabs-go"
	"github.com/oh-tarnished/elevenlabs-go/option"
)

func main() {
	client := elevenlabs.NewClient(
		option.WithAPIKey("ELEVENLABS_API_KEY"),
	)
	resp, err := client.TextToSpeech.Convert(
		context.Background(),
		"JBFqnCBsd6RMkjVDRZzb", // voice ID
		elevenlabs.TextToSpeechConvertParams{
			Text:         "Hello Srikanth This is Ria from Machani Robotics",
			ModelID:      elevenlabs.String("eleven_multilingual_v2"),
			OutputFormat: elevenlabs.TextToSpeechConvertParamsOutputFormatWav16000,
		},
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	file, err := os.Create("output.wav")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
}
