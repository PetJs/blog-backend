package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type elevenLabsRequest struct {
	Text          string              `json:"text"`
	ModelID       string              `json:"model_id"`
	VoiceSettings elevenLabsVoice     `json:"voice_settings"`
}

type elevenLabsVoice struct {
	Stability       float64 `json:"stability"`
	SimilarityBoost float64 `json:"similarity_boost"`
}

// GenerateElevenLabsAudio sends text to ElevenLabs, uploads the returned audio
// to Cloudinary, and returns the Cloudinary URL.
func GenerateElevenLabsAudio(text string) (string, error) {
	apiKey := os.Getenv("ELEVENLABS_API_KEY")
	voiceID := os.Getenv("ELEVENLABS_VOICE_ID")

	if apiKey == "" || voiceID == "" {
		return "", fmt.Errorf("ELEVENLABS_API_KEY and ELEVENLABS_VOICE_ID must be set")
	}

	payload := elevenLabsRequest{
		Text:    text,
		ModelID: "eleven_monolingual_v1",
		VoiceSettings: elevenLabsVoice{
			Stability:       0.5,
			SimilarityBoost: 0.75,
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://api.elevenlabs.io/v1/text-to-speech/%s", voiceID)
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("xi-api-key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "audio/mpeg")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		errBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ElevenLabs error %d: %s", resp.StatusCode, string(errBody))
	}

	audioBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Upload audio to Cloudinary (audio lives under "video" resource type in Cloudinary)
	return UploadFile(bytes.NewReader(audioBytes), "video")
}
