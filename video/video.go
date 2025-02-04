package video

import (
	"fmt"
	"os"

	"github.com/kkdai/youtube/v2"
)

type Format string

const (
	MP3  Format = "mp3"
	MP4  Format = "mp4"
	AVI  Format = "avi"
	WEBM Format = "webm"
)

func HandlerDownloadUrl(url string, videoFormat Format) *error {
	outputFileName := "video.mp4"
	// convertedFileName := "video_convertido." + videoFormat

	client := youtube.Client{}
	video, err := client.GetVideo(url)
	if err != nil {
		fmt.Println("Erro ao obter o vídeo:", err)
		return nil
	}

	formats := video.Formats.WithAudioChannels()
	if len(formats) == 0 {
		fmt.Println("Nenhum formato com áudio disponível.")
		return nil
	}

	format := formats[0]

	stream, _, err := client.GetStream(video, &format)
	if err != nil {
		fmt.Println("Erro ao obter o stream do vídeo:", err)
		return nil
	}

	file, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return nil
	}
	defer file.Close()

	_, err = file.ReadFrom(stream)
	if err != nil {
		fmt.Println("Erro ao salvar o vídeo:", err)
		return err
	}

	// Converter o vídeo para outro formato
	// err = ffmpeg.Input(outputFileName).Output(convertedFileName).Run()
	// if err != nil {
	// 	fmt.Println("Erro ao converter o vídeo:", err)
	// 	return
	// }

	fmt.Println("Vídeo baixado e convertido com sucesso!")
}
