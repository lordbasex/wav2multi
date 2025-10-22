# 📦 Installation Guide - wav2multi Transcoder

Complete installation guide for all supported platforms.

## 📋 Table of Contents

- [Docker Installation (Recommended)](#docker-installation-recommended)
- [Raspberry Pi 5 Installation](#raspberry-pi-5-installation)
- [Local Compilation](#local-compilation)
- [Using the Makefile](#using-the-makefile)
- [Verification](#verification)
- [Troubleshooting](#troubleshooting)

---

## 🐳 Docker Installation (Recommended)

### Prerequisites
- Docker installed on your system
- Internet connection

### Quick Start

```bash
# Pull the public image
docker pull cnsoluciones/wav2multi:latest

# Test the installation
docker run --rm cnsoluciones/wav2multi:latest --version

# Convert a file
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.g729 --format g729
```

### Build from Source

```bash
# Clone the repository
git clone https://github.com/lordbasex/wav2multi.git
cd wav2multi

# Build the Docker image
docker build -t wav2multi:latest .

# Test the build
docker run --rm wav2multi:latest --version
```

**Image Details:**
- **Size**: ~19MB (Alpine Linux)
- **Architecture**: Supports amd64 and arm64
- **Base**: golang:1.23-alpine (build) + alpine:latest (runtime)

---

## 🍓 Raspberry Pi 5 Installation

### Prerequisites
- Raspberry Pi 5 with Raspberry Pi OS 64-bit
- Docker installed on your PC (for compilation)
- Make installed
- SSH access to Raspberry Pi

### Option 1: Alpine/musl (Smaller)

```bash
# On your PC (not on Raspberry Pi):

# 1. Clone and navigate to project
git clone https://github.com/lordbasex/wav2multi.git
cd wav2multi

# 2. Build and extract binaries
make all

# 3. Copy files to Raspberry Pi
scp -r bin/ lib/ pi@raspberrypi.local:~/transcoder/

# On Raspberry Pi:

# 4. Connect to Raspberry Pi
ssh pi@raspberrypi.local

# 5. Install binaries and libraries
sudo mv ~/transcoder/bin/transcoding /usr/local/bin/
sudo chmod +x /usr/local/bin/transcoding
sudo mv ~/transcoder/lib/libbcg729.so* /usr/local/lib/
sudo ldconfig

# 6. Test installation
transcoding --version
transcoding --help
```

### Option 2: Debian/glibc (More Compatible)

```bash
# On your PC:

# 1. Build and extract binaries (Debian version)
make all-debian

# 2. Follow steps 3-6 from Option 1
```

**Which to Choose?**
- **Alpine**: Smaller binaries, good for custom systems
- **Debian**: Better compatibility with standard Raspberry Pi OS

### Makefile Commands

| Command | Description |
|---------|-------------|
| `make help` | Show all available commands |
| `make all` | Build + extract (Alpine) |
| `make all-debian` | Build + extract (Debian) |
| `make test` | Verify extracted binary |
| `make clean` | Clean all generated files |

---

## 🛠️ Local Compilation (Without Docker)

### Prerequisites

- Go 1.23 or higher
- CGO enabled
- C compiler (gcc)
- CMake
- Git
- Build tools

### Step 1: Install Dependencies

#### Ubuntu/Debian
```bash
sudo apt-get update
sudo apt-get install -y build-essential git cmake pkg-config
```

#### macOS
```bash
brew install cmake git go
```

#### Raspberry Pi OS
```bash
sudo apt-get update
sudo apt-get install -y build-essential git cmake golang
```

### Step 2: Install bcg729 Library

```bash
# Clone bcg729 repository
git clone https://github.com/BelledonneCommunications/bcg729
cd bcg729

# Build and install
cmake -S . -B build -DBUILD_SHARED_LIBS=ON
cmake --build build
sudo cmake --install build

# Update library cache (Linux)
sudo ldconfig

# Verify installation
ls -la /usr/local/lib/libbcg729*
ls -la /usr/local/include/bcg729/
```

### Step 3: Compile the Transcoder

```bash
# Clone the project
git clone https://github.com/lordbasex/wav2multi.git
cd wav2multi

# Set CGO environment
export CGO_ENABLED=1
export CGO_CFLAGS="-I/usr/local/include"
export CGO_LDFLAGS="-L/usr/local/lib -lbcg729"

# Build
go build -o transcoding transcoding.go

# Test
./transcoding --version
./transcoding --help
```

### Step 4: Install (Optional)

```bash
# Install binary
sudo cp transcoding /usr/local/bin/
sudo chmod +x /usr/local/bin/transcoding

# Verify
transcoding --version
```

---

## 🔧 Using the Makefile

The project includes a comprehensive Makefile for automated builds.

### Available Commands

```bash
# Show help
make help

# Build Docker image (Alpine/ARM64)
make build

# Build Docker image (Debian/ARM64)
make build-debian

# Extract binary and libraries (Alpine)
make extract

# Extract binary and libraries (Debian)
make extract-debian

# Build + Extract (Alpine) - Recommended
make all

# Build + Extract (Debian) - For Raspberry Pi OS
make all-debian

# Test extracted binary
make test

# Clean everything
make clean
```

### Makefile Variables

You can customize the build by modifying these variables in the Makefile:

```makefile
IMAGE_NAME := transcoding-arm64
IMAGE_NAME_DEBIAN := transcoding-arm64-debian
CONTAINER_NAME := transcoding-builder
BIN_DIR := bin
LIB_DIR := lib
```

### Example Workflow

```bash
# 1. View available commands
make help

# 2. Build everything for Alpine
make all

# 3. Verify the binary
make test

# 4. If you need to rebuild
make clean
make all-debian  # Build for Debian instead
```

---

## ✅ Verification

### Verify Installation

```bash
# Check if transcoding is in PATH
which transcoding

# Check version
transcoding --version

# Show help
transcoding --help

# Test with a sample file
transcoding input.wav output.g729 --format g729
```

### Verify Binary Dependencies (Linux)

```bash
# Check dynamic library dependencies
ldd $(which transcoding)

# Expected output should include:
# - libbcg729.so.0
# - libc.so.6
# - linux-vdso.so.1
```

### Verify bcg729 Library

```bash
# Check library installation
ls -la /usr/local/lib/libbcg729*

# Expected files:
# libbcg729.so
# libbcg729.so.0
# libbcg729.so.0.1.0 (or similar)
```

### Test Conversions

```bash
# Prepare a test WAV file (if you don't have one)
ffmpeg -f lavfi -i "sine=frequency=440:duration=5" -ar 8000 -ac 1 -sample_fmt s16 test.wav

# Test all formats
transcoding test.wav test.g729 --format g729
transcoding test.wav test.ulaw --format ulaw
transcoding test.wav test.alaw --format alaw
transcoding test.wav test.slin --format slin

# Verify conversions
ls -lh test.*

# Convert back to verify quality
ffmpeg -f g729 -i test.g729 -ar 8000 -ac 1 -c:a pcm_s16le test-verify.wav
```

---

## 🐛 Troubleshooting

### Docker Issues

#### "Cannot connect to Docker daemon"
```bash
# Start Docker service
sudo systemctl start docker

# Add user to docker group (Linux)
sudo usermod -aG docker $USER
# Log out and log back in
```

#### "Permission denied" when running Docker
```bash
# Use sudo
sudo docker run --rm cnsoluciones/wav2multi:latest --version

# Or add user to docker group (see above)
```

### Library Issues

#### "libbcg729.so: cannot open shared object file"
```bash
# Find the library
find /usr -name "libbcg729.so*" 2>/dev/null

# Update library cache
sudo ldconfig

# Add library path temporarily
export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

# Or add permanently to ~/.bashrc
echo 'export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH' >> ~/.bashrc
source ~/.bashrc
```

#### "bcg729/encoder.h: No such file or directory"
```bash
# Find the header
find /usr -name "encoder.h" 2>/dev/null

# If not found, reinstall bcg729
cd bcg729
sudo cmake --install build

# Set CGO flags
export CGO_CFLAGS="-I/usr/local/include"
```

### Build Issues

#### "CGO not enabled"
```bash
# Enable CGO
export CGO_ENABLED=1

# Verify
go env CGO_ENABLED
# Should output: 1
```

#### "Unsupported architecture"
```bash
# Check system architecture
uname -m

# For ARM64 (Raspberry Pi 5)
# Should show: aarch64

# For Intel/AMD
# Should show: x86_64
```

### WAV File Issues

#### "WAV is not PCM"
```bash
# Convert to PCM format
ffmpeg -i input.wav -acodec pcm_s16le output.wav
```

#### "mono required (1 channel)"
```bash
# Convert to mono
ffmpeg -i input.wav -ac 1 output.wav
```

#### "8000 Hz required"
```bash
# Change sample rate
ffmpeg -i input.wav -ar 8000 output.wav
```

#### "16-bit PCM required"
```bash
# Set bit depth
ffmpeg -i input.wav -sample_fmt s16 output.wav
```

#### All-in-one conversion
```bash
# Convert any audio to compatible format
ffmpeg -i input.mp3 -ar 8000 -ac 1 -sample_fmt s16 -acodec pcm_s16le output.wav
```

### Makefile Issues

#### "make: command not found"
```bash
# Install make

# Ubuntu/Debian
sudo apt-get install make

# macOS
xcode-select --install

# Raspberry Pi OS
sudo apt-get install make
```

#### "Docker not running"
```bash
# Check Docker status
docker info

# Start Docker
sudo systemctl start docker
```

---

## 📚 Additional Resources

- **Main Documentation**: [README.md](README.md)
- **Spanish Documentation**: [README.es.md](README.es.md)
- **Raspberry Pi Guide**: [README-RASPBERRYPI.md](README-RASPBERRYPI.md)
- **Library Documentation**: [wav2multi-lib/README.md](wav2multi-lib/README.md)

## 🔗 External References

- [bcg729 Library](https://github.com/BelledonneCommunications/bcg729)
- [Go Documentation](https://golang.org/doc/)
- [CGO Documentation](https://pkg.go.dev/cmd/cgo)
- [Docker Documentation](https://docs.docker.com/)

## 📞 Support

- **Author**: Federico Pereira
- **Email**: fpereira@cnsoluciones.com
- **Company**: CNSoluciones - Telecommunications and VoIP Solutions

---

**Need help?** Open an issue on GitHub or contact the author.

Made with ❤️ by CNSoluciones

