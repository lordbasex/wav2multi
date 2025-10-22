// Copyright 2025 Federico Pereira <fpereira@cnsoluciones.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package main implements a multi-format audio transcoder.
//
// transcoding.go converts WAV audio files to multiple telephony codecs:
// G.729, μ-law, A-law, and SLIN (raw PCM).
//
// Usage:
//   ./transcoding input.wav output.audio --format g729|ulaw|alaw|slin
//
// The input WAV file must be:
//   - Mono (1 channel)
//   - 8000 Hz sample rate
//   - 16-bit PCM format
//
// Supported output formats:
//   - G.729: High compression (8 kbps) using libbcg729
//   - μ-law: Standard telephony (64 kbps)
//   - A-law: European telephony (64 kbps)
//   - SLIN: Raw 16-bit PCM (128 kbps)
//
// Author: Federico Pereira <fpereira@cnsoluciones.com>
// Company: CNSoluciones - Telecommunications & VoIP Solutions

package main

/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -lbcg729 -Wl,-rpath,/usr/local/lib

#include <stdlib.h>
#include <stdint.h>
#include <bcg729/encoder.h>

// Envoltorios mínimos para cgo
// enableVAD: 0 = deshabilitado (siempre voz), 1 = habilitado (detecta silencios)
static bcg729EncoderChannelContextStruct* enc_new() {
    return initBcg729EncoderChannel(0);  // VAD deshabilitado para simplicidad
}

static void enc_close(bcg729EncoderChannelContextStruct* ctx) {
    closeBcg729EncoderChannel(ctx);
}

// Codifica un frame de 80 muestras (10 ms @ 8kHz).
// Devuelve en outLen la cantidad efectiva de bytes (típicamente 10 para voz o 2 para SID).
static void enc_frame(bcg729EncoderChannelContextStruct* ctx,
                      const int16_t *in80,
                      uint8_t *outBytes,
                      uint8_t *outLen) {
    bcg729Encoder(ctx, (int16_t*)in80, outBytes, outLen);
}
*/
import "C"

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/go-audio/wav"
	youpywav "github.com/youpy/go-wav"
)

// AudioFormat represents supported output formats
type AudioFormat string

const (
	FormatG729 AudioFormat = "g729"
	FormatULaw AudioFormat = "ulaw"
	FormatALaw AudioFormat = "alaw"
	FormatSLIN AudioFormat = "slin"
)

func main() {
	// Verificar si se solicita ayuda o versión (solo si hay al menos 2 argumentos)
	if len(os.Args) >= 2 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" || os.Args[1] == "help" {
			showHelp()
			os.Exit(0)
		}
		if os.Args[1] == "--version" || os.Args[1] == "-v" || os.Args[1] == "version" {
			fmt.Println("WAV to Multi-Format Transcoder v1.0.0")
			os.Exit(0)
		}
	}

	// Parsear argumentos: input.wav output.audio --format g729|ulaw|alaw|slin
	var inPath, outPath string
	var format AudioFormat = FormatG729 // Default to G.729 for backward compatibility

	if len(os.Args) == 3 {
		// Backward compatibility: input.wav output.g729
		inPath = os.Args[1]
		outPath = os.Args[2]
		// Detect format from file extension
		if strings.HasSuffix(outPath, ".g729") {
			format = FormatG729
		} else if strings.HasSuffix(outPath, ".ulaw") {
			format = FormatULaw
		} else if strings.HasSuffix(outPath, ".alaw") {
			format = FormatALaw
		} else if strings.HasSuffix(outPath, ".slin") {
			format = FormatSLIN
		} else {
			showHelp()
			os.Exit(2)
		}
	} else if len(os.Args) == 5 && os.Args[3] == "--format" {
		// New format: input.wav output.audio --format g729
		inPath = os.Args[1]
		outPath = os.Args[2]
		format = AudioFormat(os.Args[4])

		// Validate format
		if format != FormatG729 && format != FormatULaw && format != FormatALaw && format != FormatSLIN {
			fmt.Printf("❌ Invalid format: %s\n", format)
			fmt.Printf("Supported formats: g729, ulaw, alaw, slin\n")
			os.Exit(2)
		}
	} else {
		showHelp()
		os.Exit(2)
	}

	if err := run(inPath, outPath, format); err != nil {
		log.Fatal(err)
	}
}

func showHelp() {
	fmt.Printf(`
🎵 wav2multi - Multi-Format Audio Transcoder v1.0.0 - Federico Pereira <fpereira@cnsoluciones.com>

DESCRIPTION:
  Converts WAV audio files to multiple formats: G.729, μ-law, A-law, and SLIN.
  Optimized for VoIP telephony with excellent compression and compatibility.

SUPPORTED FORMATS:
  • G.729  - High compression (8 kbps) for VoIP
  • μ-law  - Standard telephony format (64 kbps)
  • A-law  - European telephony format (64 kbps)
  • SLIN   - Raw 16-bit PCM (128 kbps)

USAGE:
  # New format (recommended)
  %s input.wav output.audio --format g729|ulaw|alaw|slin
  
  # Legacy format (backward compatibility)
  %s input.wav output.g729
  %s input.wav output.ulaw
  %s input.wav output.alaw
  %s input.wav output.slin

WAV FILE REQUIREMENTS:
  ✅ Format: PCM (AudioFormat = 1)
  ✅ Channels: Mono (1 channel)
  ✅ Sample Rate: 8000 Hz
  ✅ Bits per Sample: 16-bit

EXAMPLES:
  # G.729 conversion (high compression)
  %s audio.wav audio.g729 --format g729
  
  # μ-law conversion (standard telephony)
  %s audio.wav audio.ulaw --format ulaw
  
  # A-law conversion (European standard)
  %s audio.wav audio.alaw --format alaw
  
  # SLIN conversion (raw PCM)
  %s audio.wav audio.slin --format slin
  
  # With Docker
  docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.g729 --format g729

VERIFICATION:
  # Verify G.729 conversion
  ffmpeg -f g729 -i output.g729 -ar 8000 -ac 1 -c:a pcm_s16le output.wav
  
  # Verify μ-law/A-law conversion
  ffmpeg -f mulaw -i output.ulaw -ar 8000 -ac 1 -c:a pcm_s16le output.wav
  ffmpeg -f alaw -i output.alaw -ar 8000 -ac 1 -c:a pcm_s16le output.wav

CONVERTING INCOMPATIBLE FILES:
  # Convert any audio to compatible format
  ffmpeg -i input.mp3 -ar 8000 -ac 1 -sample_fmt s16 -acodec pcm_s16le output.wav

FEATURES:
  • Multiple codec support (G.729, μ-law, A-law, SLIN)
  • Frame size: 10ms (80 samples @ 8kHz) for G.729
  • VAD disabled for G.729 (all frames as voice)
  • Optimized Docker image (~19MB with Alpine)
  • Backward compatibility with existing scripts

MORE INFORMATION:
  GitHub: https://github.com/lordbasex/wav2multi
  Documentation: See README.md

`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}

func run(inPath, outPath string, format AudioFormat) error {
	// Validar y mostrar información del archivo WAV
	if err := validateAndShowWavInfo(inPath); err != nil {
		return err
	}

	// Abrir WAV para conversión
	in, err := os.Open(inPath)
	if err != nil {
		return fmt.Errorf("opening WAV: %w", err)
	}
	defer in.Close()

	r := youpywav.NewReader(in)

	// Validar formato (validación adicional con youpy/go-wav)
	f, err := r.Format()
	if err != nil {
		return fmt.Errorf("reading WAV format: %w", err)
	}
	if f.AudioFormat != 1 {
		return fmt.Errorf("WAV is not PCM (AudioFormat=%d). Convert to PCM s16le first", f.AudioFormat)
	}
	if f.NumChannels != 1 {
		return fmt.Errorf("mono required (1 channel), channels=%d", f.NumChannels)
	}
	if f.SampleRate != 8000 {
		return fmt.Errorf("8000 Hz required, sampleRate=%d", f.SampleRate)
	}
	if f.BitsPerSample != 16 {
		return fmt.Errorf("16-bit PCM required, bits=%d", f.BitsPerSample)
	}

	// Crear salida
	out, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("creating output: %w", err)
	}
	defer func() {
		_ = out.Sync()
		_ = out.Close()
	}()

	// Route to appropriate encoder based on format
	switch format {
	case FormatG729:
		return encodeG729(r, out)
	case FormatULaw:
		return encodeULaw(r, out)
	case FormatALaw:
		return encodeALaw(r, out)
	case FormatSLIN:
		return encodeSLIN(r, out)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}

// encodeG729 encodes audio using G.729 codec
func encodeG729(r *youpywav.Reader, w io.Writer) error {
	// Inicializar encoder
	ctx := C.enc_new()
	if ctx == nil {
		return errors.New("initBcg729EncoderChannel returned NULL")
	}
	defer C.enc_close(ctx)

	const samplesPerFrame = 80 // 10 ms @ 8kHz
	sampleBuf := make([]int16, 0, samplesPerFrame)

	// Leer y codificar en streaming
	for {
		samples, err := r.ReadSamples(1024) // ~buffer de lectura
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("reading WAV samples: %w", err)
		}

		for _, s := range samples {
			// Canal 0 (es mono, por eso usamos Values[0])
			// Values[0] contiene el valor de la muestra para el primer canal
			v := int16(s.Values[0])
			sampleBuf = append(sampleBuf, v)

			if len(sampleBuf) == samplesPerFrame {
				if err := encodeG729Frame(ctx, sampleBuf, w); err != nil {
					return err
				}
				sampleBuf = sampleBuf[:0]
			}
		}
	}

	// Si queda cola (<80), la completamos con silencio y codificamos un último frame
	if len(sampleBuf) > 0 {
		padded := make([]int16, samplesPerFrame)
		copy(padded, sampleBuf)
		if err := encodeG729Frame(ctx, padded, w); err != nil {
			return err
		}
	}

	return nil
}

func encodeG729Frame(ctx *C.bcg729EncoderChannelContextStruct, pcm80 []int16, w io.Writer) error {
	if len(pcm80) != 80 {
		return fmt.Errorf("encodeG729Frame: expected 80 samples, got %d", len(pcm80))
	}

	var out [10]byte     // tamaño máximo típico de un frame de voz G.729
	var outLen C.uint8_t // longitud efectiva devuelta por la lib

	// Pasamos punteros a los buffers C
	C.enc_frame(
		ctx,
		(*C.int16_t)(&pcm80[0]),
		(*C.uint8_t)(&out[0]),
		&outLen,
	)

	n := int(outLen)
	if n <= 0 || n > len(out) {
		return fmt.Errorf("encoder returned invalid length: %d", n)
	}

	_, err := w.Write(out[:n])
	return err
}

// encodeULaw encodes audio using μ-law (ulaw) codec
func encodeULaw(r *youpywav.Reader, w io.Writer) error {
	for {
		samples, err := r.ReadSamples(1024)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("reading WAV samples: %w", err)
		}

		for _, s := range samples {
			// Convert 16-bit PCM to μ-law
			pcmSample := int16(s.Values[0])
			ulawByte := pcmToULaw(pcmSample)

			if _, err := w.Write([]byte{ulawByte}); err != nil {
				return fmt.Errorf("writing μ-law sample: %w", err)
			}
		}
	}
	return nil
}

// encodeALaw encodes audio using A-law (alaw) codec
func encodeALaw(r *youpywav.Reader, w io.Writer) error {
	for {
		samples, err := r.ReadSamples(1024)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("reading WAV samples: %w", err)
		}

		for _, s := range samples {
			// Convert 16-bit PCM to A-law
			pcmSample := int16(s.Values[0])
			alawByte := pcmToALaw(pcmSample)

			if _, err := w.Write([]byte{alawByte}); err != nil {
				return fmt.Errorf("writing A-law sample: %w", err)
			}
		}
	}
	return nil
}

// encodeSLIN encodes audio as raw 16-bit PCM (SLIN)
func encodeSLIN(r *youpywav.Reader, w io.Writer) error {
	for {
		samples, err := r.ReadSamples(1024)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("reading WAV samples: %w", err)
		}

		for _, s := range samples {
			// Write 16-bit PCM directly (little-endian)
			pcmSample := int16(s.Values[0])
			// Convert to little-endian bytes
			bytes := []byte{
				byte(pcmSample & 0xFF),        // Low byte
				byte((pcmSample >> 8) & 0xFF), // High byte
			}

			if _, err := w.Write(bytes); err != nil {
				return fmt.Errorf("writing SLIN sample: %w", err)
			}
		}
	}
	return nil
}

// pcmToULaw converts 16-bit PCM to μ-law
func pcmToULaw(pcm int16) byte {
	// Get sign and magnitude
	sign := pcm < 0
	if sign {
		pcm = -pcm
	}

	// Clamp to 14-bit range
	if pcm > 32767 {
		pcm = 32767
	}

	// Add bias
	pcm += 33

	// Find segment
	segment := 0
	temp := pcm >> 7
	for temp > 0 {
		segment++
		temp >>= 1
	}
	if segment > 7 {
		segment = 7
	}

	// Quantize
	quantization := (pcm >> (segment + 3)) & 0x0F

	// Build μ-law byte
	ulaw := byte(segment<<4) | byte(quantization)
	if sign {
		ulaw |= 0x80
	}

	return ^ulaw // Invert all bits
}

// pcmToALaw converts 16-bit PCM to A-law
func pcmToALaw(pcm int16) byte {
	// Get sign and magnitude
	sign := pcm < 0
	if sign {
		pcm = -pcm
	}

	// Clamp to 13-bit range
	if pcm > 32767 {
		pcm = 32767
	}

	// Add bias
	pcm += 33

	// Find segment
	segment := 0
	temp := pcm >> 7
	for temp > 0 {
		segment++
		temp >>= 1
	}
	if segment > 7 {
		segment = 7
	}

	// Quantize
	quantization := (pcm >> (segment + 3)) & 0x0F

	// Build A-law byte
	alaw := byte(segment<<4) | byte(quantization)
	if sign {
		alaw |= 0x80
	}

	// A-law uses even bits for segment
	alaw ^= 0x55 // XOR with 0x55 to get even bits

	return alaw
}

func validateAndShowWavInfo(filePath string) error {
	// Abrir archivo para análisis
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("couldn't open audio file - %v", err)
	}
	defer f.Close()

	// Crear decoder
	d := wav.NewDecoder(f)

	// Obtener el buffer completo para calcular la duración
	buffer, err := d.FullPCMBuffer()
	if err != nil {
		return fmt.Errorf("couldn't decode audio file - %v", err)
	}

	// Obtener información del archivo
	typeFileTmp := fmt.Sprintf("%v", d)
	substrStart := strings.Index(typeFileTmp, "WAVE")
	var typeFile string
	if substrStart >= 0 {
		typeFile = string([]rune(typeFileTmp)[substrStart : substrStart+4])
	} else {
		typeFile = "UNKNOWN"
	}

	// Calcular la duración en segundos
	totalSamples := len(buffer.Data)
	sampleRate := int(d.SampleRate)
	channels := int(d.NumChans)
	bitDepth := int(d.BitDepth)

	// Calcular duración en segundos
	durationSeconds := float64(totalSamples) / float64(sampleRate*channels)

	// Calcular duración en minutos y segundos para mejor legibilidad
	minutes := int(durationSeconds) / 60
	seconds := int(durationSeconds) % 60
	milliseconds := int((durationSeconds - float64(int(durationSeconds))) * 1000)

	// Mostrar toda la información del archivo
	log.Printf("=== WAV FILE INFORMATION ===")
	log.Printf("File type: %v", typeFile)
	log.Printf("Bit depth: %v bits", bitDepth)
	log.Printf("Sample rate: %v Hz", sampleRate)
	log.Printf("Number of channels: %v", channels)
	log.Printf("Total samples: %v", totalSamples)
	log.Printf("Total duration: %.3f seconds", durationSeconds)
	log.Printf("Formatted duration: %d minutes, %d seconds and %d milliseconds", minutes, seconds, milliseconds)

	// Validar requisitos para G.729
	log.Printf("=== VALIDATION ===")

	// Validar formato PCM (go-audio/wav no expone AudioFormat directamente)
	// Verificamos que sea un archivo WAV válido
	if typeFile != "WAVE" {
		return fmt.Errorf("❌ Invalid file format: %s (required: WAVE). Convert to WAV first", typeFile)
	}
	log.Printf("✅ Audio format: WAVE")

	// Validar canales mono
	if channels != 1 {
		return fmt.Errorf("❌ Invalid channels: %d (required: mono = 1). Convert to mono first", channels)
	}
	log.Printf("✅ Channels: Mono")

	// Validar frecuencia de muestreo
	if sampleRate != 8000 {
		return fmt.Errorf("❌ Invalid sample rate: %d Hz (required: 8000 Hz). Convert to 8000 Hz first", sampleRate)
	}
	log.Printf("✅ Sample rate: 8000 Hz")

	// Validar profundidad de bits
	if bitDepth != 16 {
		return fmt.Errorf("❌ Invalid bit depth: %d bits (required: 16 bits). Convert to 16-bit first", bitDepth)
	}
	log.Printf("✅ Bit depth: 16 bits")

	log.Printf("✅ File is compatible with all supported formats (G.729, μ-law, A-law, SLIN)")
	log.Printf("📝 Note: Players may show %d seconds (rounded)", int(durationSeconds))

	return nil
}
