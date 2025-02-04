package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Definição dos formatos suportados como um tipo personalizado
type Format string

const (
	MP3  Format = "mp3"
	MP4  Format = "mp4"
	AVI  Format = "avi"
	WEBM Format = "webm"
)

// Estrutura para receber os dados da requisição
type RequestData struct {
	Format Format `json:"format"`
	Url    string `json:"url"`
}

// Função para validar se o formato recebido é válido
func isValidFormat(format Format) bool {
	switch format {
	case MP3, MP4, AVI, WEBM:
		return true
	default:
		return false
	}
}

func POST(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost { // Alterei para POST por ser mais adequado
		http.Error(w, "Método não permitido. Utilize POST.", http.StatusMethodNotAllowed)
		return
	}

	// Lê o body da requisição
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o body", http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	// Decodifica o JSON enviado no body
	var data RequestData
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Valida se o formato é um dos permitidos
	if !isValidFormat(data.Format) {
		http.Error(w, "Formato inválido. Use um dos seguintes: mp3, mp4, avi, webm.", http.StatusBadRequest)
		return
	}

	// Verifica se a URL foi informada
	if data.Url == "" {
		http.Error(w, "Parâmetro 'url' é obrigatório.", http.StatusBadRequest)
		return
	}

	// Exemplo de resposta
	fmt.Fprintf(w, "Formato válido: %s\nURL do YouTube: %s\n", data.Format, data.Url)
}

func HandleHealthCheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Servidor está funcionando!"))
}
