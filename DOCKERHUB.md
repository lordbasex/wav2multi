# 🎵 wav2multi - Multi-Format Audio Transcoder

[![Docker Image](https://img.shields.io/docker/v/cnsoluciones/wav2multi?label=version)](https://hub.docker.com/r/cnsoluciones/wav2multi)
[![Docker Pulls](https://img.shields.io/docker/pulls/cnsoluciones/wav2multi)](https://hub.docker.com/r/cnsoluciones/wav2multi)
[![Docker Image Size](https://img.shields.io/docker/image-size/cnsoluciones/wav2multi/latest)](https://hub.docker.com/r/cnsoluciones/wav2multi)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/lordbasex/wav2multi/blob/main/LICENSE)

**Professional multi-format audio transcoder for VoIP telephony codecs**

Convert WAV audio files to G.729, μ-law, A-law, and SLIN formats with excellent compression and compatibility. Optimized for VoIP applications, IP telephony, and telecommunications systems.

---

## 🚀 Quick Start

```bash
# G.729 conversion (8 kbps - high compression)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.g729 --format g729

# μ-law conversion (64 kbps - US standard)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.ulaw --format ulaw

# A-law conversion (64 kbps - European standard)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.alaw --format alaw

# SLIN conversion (128 kbps - raw PCM)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.slin --format slin
```

## 📋 Supported Formats

| Format | Bitrate | Use Case | Compression |
|--------|---------|----------|-------------|
| **G.729** | 8 kbps | VoIP, maximum compression | 94% |
| **μ-law** | 64 kbps | Standard telephony (US) | 50% |
| **A-law** | 64 kbps | European telephony | 50% |
| **SLIN** | 128 kbps | Raw 16-bit PCM | 0% |

## ✨ Key Features

- ✅ **Multi-format support**: G.729, μ-law, A-law, SLIN
- ✅ **High-quality codecs**: Uses `libbcg729` for G.729, native Go for others
- ✅ **Ultra-light image**: Only ~19MB (Alpine Linux)
- ✅ **Multi-architecture**: Supports amd64 and arm64
- ✅ **Easy to use**: Simple command-line interface
- ✅ **Automatic validation**: Checks WAV format requirements
- ✅ **Backward compatible**: Legacy format still supported

## 📦 Image Information

- **Base Image**: Alpine Linux (multi-stage build)
- **Size**: ~19MB compressed
- **Architectures**: linux/amd64, linux/arm64
- **Go Version**: 1.23
- **CGO**: Enabled (for G.729 support)

## 🔧 Requirements

Your WAV files must meet these specifications:
- **Format**: PCM (AudioFormat = 1)
- **Channels**: Mono (1 channel)
- **Sample Rate**: 8000 Hz
- **Bits per Sample**: 16-bit

## 💡 Usage Examples

### Basic Conversion

```bash
# Convert to G.729 (highest compression)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest audio.wav audio.g729 --format g729

# Legacy format (auto-detect from extension)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest audio.wav audio.g729
```

### Batch Processing

```bash
# Convert all WAV files in current directory
for file in *.wav; do
    docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest "$file" "${file%.wav}.g729" --format g729
done
```

### With Custom Path

```bash
# Mount specific directory
docker run --rm -v /path/to/audio:/work cnsoluciones/wav2multi:latest input.wav output.g729 --format g729
```

### Get Help

```bash
# Show complete help
docker run --rm cnsoluciones/wav2multi:latest --help

# Show version
docker run --rm cnsoluciones/wav2multi:latest --version
```

## ✅ Verify Conversions

Convert back to WAV using FFmpeg to verify quality:

```bash
# Verify G.729
ffmpeg -f g729 -i output.g729 -ar 8000 -ac 1 -c:a pcm_s16le verify.wav

# Verify μ-law
ffmpeg -f mulaw -i output.ulaw -ar 8000 -ac 1 -c:a pcm_s16le verify.wav

# Verify A-law
ffmpeg -f alaw -i output.alaw -ar 8000 -ac 1 -c:a pcm_s16le verify.wav

# Verify SLIN
ffmpeg -f s16le -ar 8000 -ac 1 -i output.slin -c:a pcm_s16le verify.wav
```

## 🔄 Converting Incompatible Files

If your WAV file doesn't meet requirements, convert it with FFmpeg:

```bash
# All-in-one conversion to compatible format
ffmpeg -i input.mp3 -ar 8000 -ac 1 -sample_fmt s16 -acodec pcm_s16le output.wav

# Then transcode
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest output.wav output.g729 --format g729
```

## 📊 Codec Technical Details

### G.729 (High Compression)
- **Bitrate**: 8 kbps
- **Frame size**: 10ms (80 samples @ 8kHz)
- **Usage**: VoIP, IP telephony, videoconferencing
- **Implementation**: libbcg729 library (CGO)

### μ-law (Standard Telephony)
- **Bitrate**: 64 kbps
- **Usage**: US telephony, legacy systems
- **Implementation**: Native Go algorithm

### A-law (European Standard)
- **Bitrate**: 64 kbps
- **Usage**: European telephony, international
- **Implementation**: Native Go algorithm

### SLIN (Raw PCM)
- **Bitrate**: 128 kbps
- **Usage**: Raw audio, debugging, high quality
- **Implementation**: Direct PCM extraction

## 🐛 Troubleshooting

### Error: "WAV is not PCM"
```bash
ffmpeg -i file.wav -acodec pcm_s16le output.wav
```

### Error: "mono required (1 channel)"
```bash
ffmpeg -i file.wav -ac 1 output.wav
```

### Error: "8000 Hz required"
```bash
ffmpeg -i file.wav -ar 8000 output.wav
```

### All-in-one fix
```bash
ffmpeg -i input.mp3 -ar 8000 -ac 1 -sample_fmt s16 -acodec pcm_s16le output.wav
```

## 🎯 Use Cases

- **VoIP Systems**: Prepare audio files for Asterisk, FreeSWITCH
- **IP Telephony**: Convert hold music, announcements, IVR prompts
- **Telecommunications**: Standard format conversion for telecom systems
- **Audio Optimization**: Reduce file sizes for bandwidth-constrained applications
- **Testing**: Generate test files in different codec formats

## 🏗️ Technical Architecture

This Docker image uses a **multi-stage build** for optimal size:

1. **Build Stage**: Compiles bcg729 library and Go binary
2. **Runtime Stage**: Only includes binary and essential libraries
3. **Result**: Ultra-light ~19MB image

### What's Inside
- Alpine Linux (minimal base)
- Go 1.23 compiled binary
- libbcg729.so (G.729 codec library)
- Minimal dependencies (ca-certificates, libc6-compat)

## 📚 Documentation

- **GitHub Repository**: [github.com/lordbasex/wav2multi](https://github.com/lordbasex/wav2multi)
- **Complete Documentation**: [README.md](https://github.com/lordbasex/wav2multi/blob/main/README.md)
- **Spanish Documentation**: [README.es.md](https://github.com/lordbasex/wav2multi/blob/main/README.es.md)
- **Raspberry Pi Guide**: [README-RASPBERRYPI.md](https://github.com/lordbasex/wav2multi/blob/main/README-RASPBERRYPI.md)
- **Installation Guide**: [INSTALL.md](https://github.com/lordbasex/wav2multi/blob/main/INSTALL.md)

## 🔗 Links

- **Docker Hub**: [hub.docker.com/r/cnsoluciones/wav2multi](https://hub.docker.com/r/cnsoluciones/wav2multi)
- **GitHub**: [github.com/lordbasex/wav2multi](https://github.com/lordbasex/wav2multi)
- **bcg729 Library**: [github.com/BelledonneCommunications/bcg729](https://github.com/BelledonneCommunications/bcg729)
- **G.729 Specification**: [ITU-T G.729](https://www.itu.int/rec/T-REC-G.729)

## 👨‍💻 Author

**Federico Pereira** - [fpereira@cnsoluciones.com](mailto:fpereira@cnsoluciones.com)

### 🏢 CNSoluciones

This project is part of **CNSoluciones**, specialized in telecommunications and VoIP solutions.

## 📄 License

**wav2multi** is licensed under the **Apache License 2.0**.

Copyright © 2025 Federico Pereira <fpereira@cnsoluciones.com>

### Quick Summary

✅ **You CAN**: Use commercially, modify, distribute, sublicense  
⚠️ **You MUST**: Include copyright, provide attribution, state changes  
❌ **You CANNOT**: Hold author liable, remove copyright notices

### Third-Party Components

- **bcg729** (GPLv3) - G.729 codec library
- **go-wav** (MIT) - WAV file parsing  
- **go-audio** (Apache 2.0) - Audio processing

### Patent Information

G.729 codec patents **expired in 2017**. Free to use worldwide without royalties.

### Commercial Licensing

Need different terms? Commercial licenses available:
- No attribution requirements
- Priority support
- Custom development
- Contact: fpereira@cnsoluciones.com

For complete terms, see [LICENSE](https://github.com/lordbasex/wav2multi/blob/main/LICENSE)

## 🤝 Contributing

Contributions are welcome! Please visit our [GitHub repository](https://github.com/lordbasex/wav2multi) to:
- Report issues
- Submit pull requests
- Request features
- Improve documentation

## 🌟 Show Your Support

If you find this project useful, please:
- ⭐ Star the [GitHub repository](https://github.com/lordbasex/wav2multi)
- 🐳 Pull and use the Docker image
- 📢 Share with your team
- 🤝 Contribute improvements

---

**Questions or issues?** Open an issue on [GitHub](https://github.com/lordbasex/wav2multi/issues)

Made with ❤️ by CNSoluciones | Telecommunications & VoIP Solutions

