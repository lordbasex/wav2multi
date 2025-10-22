# 🎵 wav2multi - Multi-Format Audio Transcoder v1.0.0

Federico Pereira <fpereira@cnsoluciones.com>

Multi-format audio transcoder from WAV to G.729, μ-law, A-law, and SLIN formats. Supports VoIP telephony codecs with excellent compression and compatibility.

## 📋 Description

This project provides a command-line tool that converts audio files from WAV (PCM) format to multiple telephony codecs: G.729, μ-law, A-law, and SLIN. Perfect for VoIP applications requiring different compression levels and compatibility standards.

### Supported Formats

| Format | Bitrate | Use Case | Compression |
|--------|---------|----------|-------------|
| **G.729** | 8 kbps | VoIP, maximum compression | 94% |
| **μ-law** | 64 kbps | Standard telephony (US) | 50% |
| **A-law** | 64 kbps | European telephony | 50% |
| **SLIN** | 128 kbps | Raw 16-bit PCM | 0% |

### Features

- ✅ **Multi-format support**: G.729, μ-law, A-law, SLIN
- ✅ **WAV conversion**: Mono, 8kHz, 16-bit PCM input
- ✅ **High-quality codecs**: Uses `libbcg729` for G.729, native Go for others
- ✅ **Backward compatibility**: Legacy `.g729` format still supported
- ✅ **Dockerized**: Easy deployment with optimized image
- ✅ **Multi-stage build**: Ultra-light ~19MB Alpine image
- ✅ **ARM64 support**: Native Raspberry Pi 5 compatibility
- ✅ **Makefile automation**: Easy build and extraction workflow

## 🔧 Requirements

### For Docker usage (Recommended)
- Docker installed on your system
- WAV files with the following specifications:
  - **Format**: PCM (AudioFormat = 1)
  - **Channels**: Mono (1 channel)
  - **Sample Rate**: 8000 Hz
  - **Bits per sample**: 16-bit

### For local compilation
- Go 1.23 or higher
- CGO enabled (`CGO_ENABLED=1`)
- `libbcg729` installed on the system
- Build tools (gcc, cmake, git)

### For Raspberry Pi 5
- Raspberry Pi OS 64-bit
- ARM64 architecture
- Make installed
- Docker (for compilation)

## 🚀 Quick usage with Docker

### 📦 Public image available

The image is publicly available on Docker Hub as `cnsoluciones/wav2multi:latest`. You don't need to build the image locally.

### 1. Multi-format conversion (Recommended)

```bash
# G.729 conversion (high compression)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.g729 --format g729

# μ-law conversion (standard telephony)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.ulaw --format ulaw

# A-law conversion (European standard)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.alaw --format alaw

# SLIN conversion (raw PCM)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.slin --format slin
```

### 2. Legacy format (Backward compatibility)

```bash
# Still works - detects format from file extension
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.g729
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.ulaw
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.alaw
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.slin
```

**Command explanation:**
- `docker run`: Runs a container from the image
- `--rm`: Automatically removes the container after execution
- `-v $PWD:/work`: Mounts the current directory to `/work` inside the container
- `cnsoluciones/wav2multi:latest`: Public Docker Hub image
- `input.wav`: Input file (WAV)
- `output.audio`: Output file with format specified by `--format` or file extension

### 3. Build the image locally (Optional)

```bash
docker build -t cnsoluciones/wav2multi:latest .
```

This command:
- Downloads and installs all necessary dependencies
- Compiles the `bcg729` library from source code
- Compiles the Go program with CGO support
- Creates an optimized image of **~19MB** (Alpine Linux)

### 4. Get help

```bash
# Show complete help
docker run --rm cnsoluciones/wav2multi:latest --help

# Show version
docker run --rm cnsoluciones/wav2multi:latest --version

# Or simply run without arguments
docker run --rm cnsoluciones/wav2multi:latest
```

The helper includes (in English):
- ✅ **Multi-format support**: G.729, μ-law, A-law, SLIN
- ✅ **Usage examples**: All supported formats
- ✅ **Docker commands**: Complete Docker usage guide
- ✅ **FFmpeg commands**: For incompatible file conversion
- ✅ **Verification commands**: To validate all conversions
- ✅ **Technical details**: Codec specifications and bitrates

### 5. Example with full path

```bash
docker run --rm -v /path/to/your/files:/work cnsoluciones/wav2multi:latest audio.wav audio.g729 --format g729
```

## 🍓 Build for Raspberry Pi 5 using Makefile

The project includes a powerful Makefile for easy compilation and deployment to Raspberry Pi 5.

### 📋 Available Makefile Commands

| Command | Description | Dependencies |
|---------|-------------|--------------|
| `make help` | Shows all available commands | None |
| `make build` | Builds Docker image for ARM64 (Alpine/musl) | Docker |
| `make build-debian` | Builds Docker image for ARM64 (Debian/glibc) | Docker |
| `make extract` | Extracts binary and libraries (Alpine) | `build` |
| `make extract-debian` | Extracts binary and libraries (Debian) | `build-debian` |
| `make all` | Runs `build` + `extract` (Alpine) | `extract` |
| `make all-debian` | Runs `build` + `extract` (Debian) | `extract-debian` |
| `make clean` | Cleans generated files and Docker images | None |
| `make test` | Tests the extracted binary | `extract` |

### 🚀 Quick Start for Raspberry Pi

#### Option 1: Alpine Linux (musl libc)
```bash
# Build and extract everything (Alpine)
make all

# Result: bin/transcoding and lib/libbcg729.so*
```

#### Option 2: Debian (glibc)
```bash
# Build and extract everything (Debian/glibc)
make all-debian

# Result: bin/transcoding and lib/libbcg729.so*
```

### 📦 Generated Files

After running `make all` or `make all-debian`, you'll get:

```
.
├── bin/
│   └── transcoding          # ARM64 binary (~2.4MB)
└── lib/
    ├── libbcg729.so         # Main library (~148KB)
    └── libbcg729.so.0       # Version symlink (~148KB)
```

### 📤 Deploy to Raspberry Pi

```bash
# 1. Copy files to Raspberry Pi
scp -r bin/ lib/ pi@raspberrypi.local:~/transcoder/

# 2. Connect to Raspberry Pi
ssh pi@raspberrypi.local

# 3. Install binary and libraries
sudo mv ~/transcoder/bin/transcoding /usr/local/bin/
sudo chmod +x /usr/local/bin/transcoding
sudo mv ~/transcoder/lib/libbcg729.so* /usr/local/lib/
sudo ldconfig

# 4. Test
transcoding input.wav output.g729 --format g729
```

### 🧹 Clean Up

```bash
# Remove all generated files and Docker images
make clean
```

### 🧪 Test Binary

```bash
# Verify the extracted binary
make test
```

This command:
- ✅ Checks if binary exists
- ✅ Displays architecture information (`file` command)
- ✅ Shows library dependencies (`ldd` command)

## ✅ Verify conversion

To validate that the output files were created correctly, you can convert them back to WAV with FFmpeg:

### G.729 verification
```bash
ffmpeg -f g729 -i output.g729 -ar 8000 -ac 1 -c:a pcm_s16le output.wav
```

### μ-law/A-law verification
```bash
ffmpeg -f mulaw -i output.ulaw -ar 8000 -ac 1 -c:a pcm_s16le output.wav
ffmpeg -f alaw -i output.alaw -ar 8000 -ac 1 -c:a pcm_s16le output.wav
```

### SLIN verification
```bash
ffmpeg -f s16le -ar 8000 -ac 1 -i output.slin -c:a pcm_s16le output.wav
```

**Command explanation:**
- `-f g729/mulaw/alaw/s16le`: Specifies input format
- `-i output.audio`: Input file (the generated audio)
- `-ar 8000`: Output sample rate (8000 Hz)
- `-ac 1`: Number of output channels (1 = mono)
- `-c:a pcm_s16le`: Output audio codec (PCM 16-bit little-endian)
- `output.wav`: Resulting WAV file

Now you can play `output.wav` with any audio player to verify the conversion quality. If you hear the audio correctly, the conversion was successful! 🎵

## 📁 Project structure

```
.
├── Dockerfile                      # Multi-stage Docker image (Alpine)
├── Dockerfile-for-raspberrypi5    # ARM64 image for Raspberry Pi (Alpine/musl)
├── Dockerfile-for-raspberrypi-debian # ARM64 image for Raspberry Pi (Debian/glibc)
├── Makefile                        # Build automation for Raspberry Pi
├── go.mod                          # Go project dependencies
├── go.sum                          # Dependency checksums
├── transcoding.go                  # Main transcoder code
├── bin/                            # Generated binaries (after make)
│   └── transcoding                 # ARM64 binary
├── lib/                            # Generated libraries (after make)
│   ├── libbcg729.so
│   └── libbcg729.so.0
├── wav2multi-lib/                  # Go library for programmatic use
│   ├── codecs.go                   # Codec implementations
│   ├── transcoder.go               # Transcoder logic
│   ├── types.go                    # Type definitions
│   ├── g729_codec.go               # G.729 with CGO
│   ├── g729_codec_nocgo.go         # G.729 stub without CGO
│   └── example/                    # Library usage examples
├── README.md                       # This file (English)
├── README.es.md                    # Spanish documentation
└── README-RASPBERRYPI.md          # Raspberry Pi specific guide
```

## 🏗️ Technical architecture

### Optimized Multi-stage Dockerfile

The project uses a two-stage Dockerfile optimized with **Alpine Linux**:

1. **Stage 1 (build)**: Image based on `golang:1.23-alpine`
   - Installs build tools (build-base, cmake, git)
   - Clones and compiles `bcg729` as shared library (`libbcg729.so`)
   - Downloads Go dependencies
   - Compiles binary with CGO enabled

2. **Stage 2 (runtime)**: Image based on `alpine:latest`
   - Contains only the compiled binary and necessary libraries
   - Copies `libbcg729.so` and minimal dependencies
   - **Result: ultra-light image of ~19MB** 🚀

### 🎯 Implemented optimizations:

- ✅ **Alpine Linux**: Minimal base (~3MB) vs Debian (~80MB)
- ✅ **Shared library**: `libbcg729.so` instead of static
- ✅ **Minimal dependencies**: Only `ca-certificates` and `libc6-compat`
- ✅ **Multi-stage build**: Separate compilation from runtime
- ✅ **No development tools**: Only what's necessary to run
- ✅ **ARM64 native**: Cross-platform support for Raspberry Pi

### Raspberry Pi Build Options

The project offers **two build options** for Raspberry Pi:

#### 1. Alpine/musl (Dockerfile-for-raspberrypi5)
- **Pros**: Smaller binary, minimal dependencies
- **Cons**: Requires musl libc (comes with Alpine-based systems)
- **Use**: When size matters most

#### 2. Debian/glibc (Dockerfile-for-raspberrypi-debian)
- **Pros**: Better compatibility with Raspberry Pi OS
- **Cons**: Slightly larger
- **Use**: For standard Raspberry Pi OS installations

### Go code with CGO

The program uses CGO to call C functions from `libbcg729`:

```go
/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -lbcg729 -Wl,-rpath,/usr/local/lib
#include <bcg729/encoder.h>
*/
import "C"
```

**Conversion process:**
1. Reads WAV file using `github.com/youpy/go-wav`
2. Validates format (mono, 8kHz, 16-bit PCM)
3. Processes audio in frames of 80 samples (10ms @ 8kHz)
4. Encodes each frame with the selected codec
5. Writes encoded bitstream to output file

### 🆘 Integrated help system

The program includes a complete helper that activates when:
- Run without arguments: `docker run --rm cnsoluciones/wav2multi:latest`
- Explicit help request: `docker run --rm cnsoluciones/wav2multi:latest --help`

**Helper features (in English):**
- 📋 **Complete description** of the program and its purpose
- 📝 **Technical requirements** for input WAV file
- 💡 **Practical examples** of Docker usage
- 🔧 **FFmpeg commands** to convert incompatible files
- ✅ **Verification commands** to validate conversion
- 📊 **Technical information** about all supported codecs
- 🔗 **Additional documentation** links

## 🔍 WAV format validation

The program automatically validates that the WAV file meets the requirements:

```
✅ AudioFormat = 1 (PCM)
✅ NumChannels = 1 (Mono)
✅ SampleRate = 8000 Hz
✅ BitsPerSample = 16
```

If your file doesn't meet these requirements, you can convert it with FFmpeg:

```bash
# Convert any audio file to compatible format
ffmpeg -i input.mp3 -ar 8000 -ac 1 -sample_fmt s16 -acodec pcm_s16le output.wav
```

## 🛠️ Local compilation (without Docker)

If you prefer to compile locally without Docker:

### 1. Install bcg729

```bash
git clone https://github.com/BelledonneCommunications/bcg729
cd bcg729
cmake -S . -B build
cmake --build build --target install
sudo ldconfig
```

### 2. Compile the program

```bash
export CGO_ENABLED=1
go build -o transcoding transcoding.go
```

### 3. Run

```bash
./transcoding input.wav output.g729 --format g729
```

## 📊 Codec technical details

### G.729 (High Compression)
- **Bitrate**: 8 kbps (very efficient)
- **Frame size**: 10ms (80 samples @ 8kHz)
- **Frame encoding**: ~10 bytes per voice frame
- **Usage**: VoIP, IP telephony, videoconferencing
- **Advantage**: Excellent voice quality with minimal bandwidth
- **Implementation**: Uses `libbcg729` library via CGO

### μ-law (Standard Telephony)
- **Bitrate**: 64 kbps (standard)
- **Frame size**: 1 sample (8kHz)
- **Frame encoding**: 1 byte per sample
- **Usage**: US telephony, legacy systems
- **Advantage**: Simple, widely supported
- **Implementation**: Native Go algorithm

### A-law (European Standard)
- **Bitrate**: 64 kbps (standard)
- **Frame size**: 1 sample (8kHz)
- **Frame encoding**: 1 byte per sample
- **Usage**: European telephony, international
- **Advantage**: Better dynamic range than μ-law
- **Implementation**: Native Go algorithm

### SLIN (Raw PCM)
- **Bitrate**: 128 kbps (uncompressed)
- **Frame size**: 1 sample (8kHz)
- **Frame encoding**: 2 bytes per sample (little-endian)
- **Usage**: Raw audio, debugging, high quality
- **Advantage**: No compression artifacts
- **Implementation**: Direct PCM extraction

### VAD (Voice Activity Detection) - G.729 only

The G.729 encoder is configured with VAD disabled (`enableVAD = 0`):
- **VAD = 0**: All frames are encoded as voice (simpler)
- **VAD = 1**: Detects silence and encodes it efficiently (saves bandwidth)

You can modify this configuration in `transcoding.go` line 19.

## 🐛 Troubleshooting

### Error: "WAV is not PCM"
Your file is in compressed format. Convert it with FFmpeg:
```bash
ffmpeg -i file.wav -acodec pcm_s16le output.wav
```

### Error: "mono required (1 channel)"
Your file is stereo. Convert to mono:
```bash
ffmpeg -i file.wav -ac 1 output.wav
```

### Error: "8000 Hz required"
Change the sample rate:
```bash
ffmpeg -i file.wav -ar 8000 output.wav
```

### Error: "16-bit PCM required"
Adjust the sample format:
```bash
ffmpeg -i file.wav -sample_fmt s16 output.wav
```

### All-in-one conversion with FFmpeg
```bash
ffmpeg -i input.mp3 -ar 8000 -ac 1 -sample_fmt s16 -acodec pcm_s16le output.wav
```

### Raspberry Pi: "libbcg729.so not found"
```bash
# Add library path
export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

# Or update library cache
sudo ldconfig
```

### Makefile: Docker not running
```bash
# Start Docker service
sudo systemctl start docker

# Or check Docker status
docker info
```

## 📝 Important notes

- ⚠️ The output `.g729` file is a **raw bitstream** without container
- ⚠️ To play G.729 files, you need a compatible player or convert them back to WAV
- 💡 **Tip**: Use `ffmpeg -f g729 -i output.g729 -ar 8000 -ac 1 -c:a pcm_s16le output.wav` to convert G.729 to WAV
- ⚠️ Some G.729 codecs are subject to patents (check in your jurisdiction)
- ⚠️ `bcg729` is an open-source and royalty-free implementation
- 🍓 For Raspberry Pi, choose Alpine (smaller) or Debian (more compatible) build
- 🔧 The Makefile automates the entire build and deployment process

## 📚 References

- [bcg729 - G.729 codec library](https://github.com/BelledonneCommunications/bcg729)
- [go-wav - WAV parser for Go](https://github.com/youpy/go-wav)
- [ITU-T G.729 Specification](https://www.itu.int/rec/T-REC-G.729)
- [CGO Documentation](https://pkg.go.dev/cmd/cgo)
- [GNU Make Manual](https://www.gnu.org/software/make/manual/)

## 📄 License

**wav2multi** is licensed under the **Apache License 2.0**.

Copyright © 2025 Federico Pereira <fpereira@cnsoluciones.com>

### What This Means For You

✅ **You CAN:**
- ✓ Use commercially (free for open source)
- ✓ Modify the code
- ✓ Distribute the software
- ✓ Use in private projects
- ✓ Sublicense your modifications

⚠️ **You MUST:**
- ✓ Include copyright notice
- ✓ Include LICENSE file in distributions
- ✓ State significant changes made
- ✓ Include NOTICE file
- ✓ Provide attribution

❌ **You CANNOT:**
- ✗ Hold the author liable
- ✗ Use author's name for endorsement without permission
- ✗ Remove copyright notices
- ✗ Change the license of the original code

### Third-Party Components

This software uses the following components:

- **bcg729** - GPLv3 - G.729 codec library
  - Copyright © Belledonne Communications
  - https://github.com/BelledonneCommunications/bcg729
  - G.729 patents expired in 2017 - free to use worldwide

- **go-wav** - MIT License - WAV file parsing
  - Copyright © youpy
  - https://github.com/youpy/go-wav

- **go-audio** - Apache 2.0 - Audio processing
  - Copyright © Matt Aimonetti
  - https://github.com/go-audio

### Commercial Licensing

Need different licensing terms? We offer commercial licenses with:

- ✓ No attribution requirements in your product
- ✓ Priority support and consulting
- ✓ Custom features and development
- ✓ Service Level Agreements (SLA)
- ✓ Flexible licensing terms

**Contact for Commercial Licensing:**
- Email: fpereira@cnsoluciones.com
- Company: CNSoluciones - Telecommunications & VoIP Solutions
- Website: https://cnsoluciones.com

### Patent Information

The G.729 codec patents expired in 2017. This software is free to use worldwide without royalty payments. The `bcg729` library is an open-source, royalty-free implementation.

For complete license terms, see the [LICENSE](LICENSE) and [NOTICE](NOTICE) files.

## 🤝 Contributing

Contributions are welcome. Please:
1. Fork the repository
2. Create a branch for your feature (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 👨‍💻 Author

**Federico Pereira** <fpereira@cnsoluciones.com>

WAV to Multi-Format audio conversion project using Go and CGO.

### 🏢 CNSoluciones

This project is part of CNSoluciones, specialized in telecommunications and VoIP solutions.

---

**Questions or issues?** Open an issue in the repository.

## 🌐 Language versions

- 🇺🇸 [English](README.md) (Current)
- 🇪🇸 [Español](README.es.md)
- 🍓 [Raspberry Pi Guide](README-RASPBERRYPI.md)
