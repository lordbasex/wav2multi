# 🎵 wav2multi - Multi-Format Audio Transcoder v1.0.0

Federico Pereira <fpereira@cnsoluciones.com>

Conversor de audio multi-formato de WAV a G.729, μ-law, A-law y SLIN. Soporta codecs de telefonía VoIP con excelente compresión y compatibilidad.

## 📋 Descripción

Este proyecto proporciona una herramienta de línea de comandos que convierte archivos de audio en formato WAV (PCM) a múltiples codecs de telefonía: G.729, μ-law, A-law y SLIN. Perfecto para aplicaciones VoIP que requieren diferentes niveles de compresión y estándares de compatibilidad.

### Formatos Soportados

| Formato | Bitrate | Caso de Uso | Compresión |
|---------|---------|-------------|------------|
| **G.729** | 8 kbps | VoIP, máxima compresión | 94% |
| **μ-law** | 64 kbps | Telefonía estándar (US) | 50% |
| **A-law** | 64 kbps | Telefonía europea | 50% |
| **SLIN** | 128 kbps | PCM 16-bit crudo | 0% |

### Características

- ✅ **Soporte multi-formato**: G.729, μ-law, A-law, SLIN
- ✅ **Conversión WAV**: Entrada mono, 8kHz, 16-bit PCM
- ✅ **Codecs de alta calidad**: Usa `libbcg729` para G.729, Go nativo para otros
- ✅ **Compatibilidad hacia atrás**: Formato legacy `.g729` aún soportado
- ✅ **Dockerizado**: Despliegue fácil con imagen optimizada
- ✅ **Build multi-stage**: Imagen ultra-ligera ~19MB Alpine
- ✅ **Soporte ARM64**: Compatibilidad nativa con Raspberry Pi 5
- ✅ **Automatización Makefile**: Flujo de trabajo de compilación y extracción fácil

## 🔧 Requisitos

### Para usar con Docker (Recomendado)
- Docker instalado en tu sistema
- Archivos WAV con las siguientes especificaciones:
  - **Formato**: PCM (AudioFormat = 1)
  - **Canales**: Mono (1 canal)
  - **Sample Rate**: 8000 Hz
  - **Bits por muestra**: 16-bit

### Para compilación local
- Go 1.23 o superior
- CGO habilitado (`CGO_ENABLED=1`)
- `libbcg729` instalada en el sistema
- Herramientas de compilación (gcc, cmake, git)

### Para Raspberry Pi 5
- Raspberry Pi OS 64-bit
- Arquitectura ARM64
- Make instalado
- Docker (para compilación)

## 🚀 Uso rápido con Docker

### 📦 Imagen pública disponible

La imagen está disponible públicamente en Docker Hub como `cnsoluciones/wav2multi:latest`. No necesitas construir la imagen localmente.

### 1. Conversión multi-formato (Recomendado)

```bash
# Conversión G.729 (alta compresión)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.g729 --format g729

# Conversión μ-law (telefonía estándar)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.ulaw --format ulaw

# Conversión A-law (estándar europeo)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.alaw --format alaw

# Conversión SLIN (PCM crudo)
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.slin --format slin
```

### 2. Formato legacy (Compatibilidad hacia atrás)

```bash
# Aún funciona - detecta formato por extensión de archivo
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.g729
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.ulaw
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.alaw
docker run --rm -v $PWD:/work cnsoluciones/wav2multi:latest input.wav output.slin
```

**Explicación del comando:**
- `docker run`: Ejecuta un contenedor desde la imagen
- `--rm`: Elimina automáticamente el contenedor después de la ejecución
- `-v $PWD:/work`: Monta el directorio actual en `/work` dentro del contenedor
- `cnsoluciones/wav2multi:latest`: Imagen pública de Docker Hub
- `input.wav`: Archivo de entrada (WAV)
- `output.audio`: Archivo de salida con formato especificado por `--format` o extensión

### 3. Construir la imagen localmente (Opcional)

```bash
docker build -t wav2multi:latest .
```

Este comando:
- Descarga e instala todas las dependencias necesarias
- Compila la librería `bcg729` desde el código fuente
- Compila el programa Go con soporte CGO
- Crea una imagen optimizada de **~19MB** (Alpine Linux)

### 4. Obtener ayuda

```bash
# Mostrar ayuda completa
docker run --rm cnsoluciones/wav2multi:latest --help

# Mostrar versión
docker run --rm cnsoluciones/wav2multi:latest --version

# O simplemente ejecutar sin argumentos
docker run --rm cnsoluciones/wav2multi:latest
```

El helper incluye (en inglés):
- ✅ **Soporte multi-formato**: G.729, μ-law, A-law, SLIN
- ✅ **Ejemplos de uso**: Todos los formatos soportados
- ✅ **Comandos Docker**: Guía completa de uso con Docker
- ✅ **Comandos FFmpeg**: Para conversión de archivos incompatibles
- ✅ **Comandos de verificación**: Para validar todas las conversiones
- ✅ **Detalles técnicos**: Especificaciones y bitrates de codecs

### 5. Ejemplo con ruta completa

```bash
docker run --rm -v /ruta/a/tus/archivos:/work cnsoluciones/wav2multi:latest audio.wav audio.g729 --format g729
```

## 🍓 Compilar para Raspberry Pi 5 usando Makefile

El proyecto incluye un poderoso Makefile para facilitar la compilación y despliegue en Raspberry Pi 5.

### 📋 Comandos Disponibles del Makefile

| Comando | Descripción | Dependencias |
|---------|-------------|--------------|
| `make help` | Muestra todos los comandos disponibles | Ninguna |
| `make build` | Construye imagen Docker para ARM64 (Alpine/musl) | Docker |
| `make build-debian` | Construye imagen Docker para ARM64 (Debian/glibc) | Docker |
| `make extract` | Extrae binario y librerías (Alpine) | `build` |
| `make extract-debian` | Extrae binario y librerías (Debian) | `build-debian` |
| `make all` | Ejecuta `build` + `extract` (Alpine) | `extract` |
| `make all-debian` | Ejecuta `build` + `extract` (Debian) | `extract-debian` |
| `make clean` | Limpia archivos generados e imágenes Docker | Ninguna |
| `make test` | Prueba el binario extraído | `extract` |

### 🚀 Inicio Rápido para Raspberry Pi

#### Opción 1: Alpine Linux (musl libc)
```bash
# Compilar y extraer todo (Alpine)
make all

# Resultado: bin/transcoding y lib/libbcg729.so*
```

#### Opción 2: Debian (glibc)
```bash
# Compilar y extraer todo (Debian/glibc)
make all-debian

# Resultado: bin/transcoding y lib/libbcg729.so*
```

### 📦 Archivos Generados

Después de ejecutar `make all` o `make all-debian`, obtendrás:

```
.
├── bin/
│   └── transcoding          # Binario ARM64 (~2.4MB)
└── lib/
    ├── libbcg729.so         # Librería principal (~148KB)
    └── libbcg729.so.0       # Enlace de versión (~148KB)
```

### 📤 Desplegar en Raspberry Pi

```bash
# 1. Copiar archivos a Raspberry Pi
scp -r bin/ lib/ pi@raspberrypi.local:~/transcoder/

# 2. Conectar a Raspberry Pi
ssh pi@raspberrypi.local

# 3. Instalar binario y librerías
sudo mv ~/transcoder/bin/transcoding /usr/local/bin/
sudo chmod +x /usr/local/bin/transcoding
sudo mv ~/transcoder/lib/libbcg729.so* /usr/local/lib/
sudo ldconfig

# 4. Probar
transcoding input.wav output.g729 --format g729
```

### 🧹 Limpiar

```bash
# Eliminar todos los archivos generados e imágenes Docker
make clean
```

### 🧪 Probar Binario

```bash
# Verificar el binario extraído
make test
```

Este comando:
- ✅ Verifica si el binario existe
- ✅ Muestra información de arquitectura (comando `file`)
- ✅ Muestra dependencias de librerías (comando `ldd`)

## ✅ Verificar la conversión

Para validar que los archivos de salida se crearon correctamente, puedes convertirlos de vuelta a WAV con FFmpeg:

### Verificación G.729
```bash
ffmpeg -f g729 -i output.g729 -ar 8000 -ac 1 -c:a pcm_s16le output.wav
```

### Verificación μ-law/A-law
```bash
ffmpeg -f mulaw -i output.ulaw -ar 8000 -ac 1 -c:a pcm_s16le output.wav
ffmpeg -f alaw -i output.alaw -ar 8000 -ac 1 -c:a pcm_s16le output.wav
```

### Verificación SLIN
```bash
ffmpeg -f s16le -ar 8000 -ac 1 -i output.slin -c:a pcm_s16le output.wav
```

**Explicación del comando:**
- `-f g729/mulaw/alaw/s16le`: Especifica el formato de entrada
- `-i output.audio`: Archivo de entrada (el audio generado)
- `-ar 8000`: Frecuencia de muestreo de salida (8000 Hz)
- `-ac 1`: Número de canales de salida (1 = mono)
- `-c:a pcm_s16le`: Codec de audio de salida (PCM 16-bit little-endian)
- `output.wav`: Archivo WAV resultante

Ahora puedes reproducir `output.wav` con cualquier reproductor de audio para verificar la calidad de la conversión. Si escuchas el audio correctamente, ¡la conversión fue exitosa! 🎵

## 📁 Estructura del proyecto

```
.
├── Dockerfile                      # Imagen Docker multi-stage (Alpine)
├── Dockerfile-for-raspberrypi5    # Imagen ARM64 para Raspberry Pi (Alpine/musl)
├── Dockerfile-for-raspberrypi-debian # Imagen ARM64 para Raspberry Pi (Debian/glibc)
├── Makefile                        # Automatización de compilación para Raspberry Pi
├── go.mod                          # Dependencias del proyecto Go
├── go.sum                          # Checksums de las dependencias
├── transcoding.go                  # Código principal del conversor
├── bin/                            # Binarios generados (después de make)
│   └── transcoding                 # Binario ARM64
├── lib/                            # Librerías generadas (después de make)
│   ├── libbcg729.so
│   └── libbcg729.so.0
├── wav2multi-lib/                  # Librería Go para uso programático
│   ├── codecs.go                   # Implementaciones de codecs
│   ├── transcoder.go               # Lógica del transcoder
│   ├── types.go                    # Definiciones de tipos
│   ├── g729_codec.go               # G.729 con CGO
│   ├── g729_codec_nocgo.go         # G.729 stub sin CGO
│   └── example/                    # Ejemplos de uso de la librería
├── README.md                       # Documentación en inglés
├── README.es.md                    # Este archivo (Español)
└── README-RASPBERRYPI.md          # Guía específica para Raspberry Pi
```

## 🏗️ Arquitectura técnica

### Dockerfile Multi-stage optimizado

El proyecto utiliza un Dockerfile de dos etapas optimizado con **Alpine Linux**:

1. **Stage 1 (build)**: Imagen basada en `golang:1.23-alpine`
   - Instala herramientas de compilación (build-base, cmake, git)
   - Clona y compila `bcg729` como librería compartida (`libbcg729.so`)
   - Descarga dependencias Go
   - Compila el binario con CGO habilitado

2. **Stage 2 (runtime)**: Imagen basada en `alpine:latest`
   - Solo contiene el binario compilado y librerías necesarias
   - Copia la librería `libbcg729.so` y dependencias mínimas
   - **Resultado: imagen ultra-ligera de ~19MB** 🚀

### 🎯 Optimizaciones implementadas:

- ✅ **Alpine Linux**: Base mínima (~3MB) vs Debian (~80MB)
- ✅ **Librería compartida**: `libbcg729.so` en lugar de estática
- ✅ **Dependencias mínimas**: Solo `ca-certificates` y `libc6-compat`
- ✅ **Multi-stage build**: Compilación separada del runtime
- ✅ **Sin herramientas de desarrollo**: Solo lo necesario para ejecutar
- ✅ **ARM64 nativo**: Soporte multi-plataforma para Raspberry Pi

### Opciones de Compilación para Raspberry Pi

El proyecto ofrece **dos opciones de compilación** para Raspberry Pi:

#### 1. Alpine/musl (Dockerfile-for-raspberrypi5)
- **Ventajas**: Binario más pequeño, dependencias mínimas
- **Desventajas**: Requiere musl libc (viene con sistemas basados en Alpine)
- **Uso**: Cuando el tamaño es lo más importante

#### 2. Debian/glibc (Dockerfile-for-raspberrypi-debian)
- **Ventajas**: Mejor compatibilidad con Raspberry Pi OS
- **Desventajas**: Ligeramente más grande
- **Uso**: Para instalaciones estándar de Raspberry Pi OS

### Código Go con CGO

El programa utiliza CGO para llamar funciones C de `libbcg729`:

```go
/*
#cgo CFLAGS: -I/usr/local/include
#cgo LDFLAGS: -L/usr/local/lib -lbcg729 -Wl,-rpath,/usr/local/lib
#include <bcg729/encoder.h>
*/
import "C"
```

**Proceso de conversión:**
1. Lee el archivo WAV usando `github.com/youpy/go-wav`
2. Valida el formato (mono, 8kHz, 16-bit PCM)
3. Procesa el audio en frames de 80 muestras (10ms @ 8kHz)
4. Codifica cada frame con el codec seleccionado
5. Escribe el bitstream codificado al archivo de salida

### 🆘 Sistema de ayuda integrado

El programa incluye un helper completo que se activa cuando:
- Se ejecuta sin argumentos: `docker run --rm cnsoluciones/wav2multi:latest`
- Se solicita ayuda explícita: `docker run --rm cnsoluciones/wav2multi:latest --help`

**Características del helper (en inglés):**
- 📋 **Descripción completa** del programa y su propósito
- 📝 **Requisitos técnicos** del archivo WAV de entrada
- 💡 **Ejemplos prácticos** de uso con Docker
- 🔧 **Comandos FFmpeg** para convertir archivos incompatibles
- ✅ **Comandos de verificación** para validar la conversión
- 📊 **Información técnica** de todos los codecs soportados
- 🔗 **Enlaces a documentación** adicional

## 🔍 Validación del formato WAV

El programa valida automáticamente que el archivo WAV cumpla con los requisitos:

```
✅ AudioFormat = 1 (PCM)
✅ NumChannels = 1 (Mono)
✅ SampleRate = 8000 Hz
✅ BitsPerSample = 16
```

Si tu archivo no cumple estos requisitos, puedes convertirlo con FFmpeg:

```bash
# Convertir cualquier archivo de audio a formato compatible
ffmpeg -i entrada.mp3 -ar 8000 -ac 1 -sample_fmt s16 -acodec pcm_s16le salida.wav
```

## 🛠️ Compilación local (sin Docker)

Si prefieres compilar localmente sin Docker:

### 1. Instalar bcg729

```bash
git clone https://github.com/BelledonneCommunications/bcg729
cd bcg729
cmake -S . -B build
cmake --build build --target install
sudo ldconfig
```

### 2. Compilar el programa

```bash
export CGO_ENABLED=1
go build -o transcoding transcoding.go
```

### 3. Ejecutar

```bash
./transcoding input.wav output.g729 --format g729
```

## 📊 Detalles técnicos de los codecs

### G.729 (Alta Compresión)
- **Bitrate**: 8 kbps (muy eficiente)
- **Frame size**: 10ms (80 muestras @ 8kHz)
- **Frame encoding**: ~10 bytes por frame de voz
- **Uso**: VoIP, telefonía IP, videoconferencia
- **Ventaja**: Excelente calidad de voz con mínimo ancho de banda
- **Implementación**: Usa librería `libbcg729` vía CGO

### μ-law (Telefonía Estándar)
- **Bitrate**: 64 kbps (estándar)
- **Frame size**: 1 muestra (8kHz)
- **Frame encoding**: 1 byte por muestra
- **Uso**: Telefonía US, sistemas legacy
- **Ventaja**: Simple, ampliamente soportado
- **Implementación**: Algoritmo nativo en Go

### A-law (Estándar Europeo)
- **Bitrate**: 64 kbps (estándar)
- **Frame size**: 1 muestra (8kHz)
- **Frame encoding**: 1 byte por muestra
- **Uso**: Telefonía europea, internacional
- **Ventaja**: Mejor rango dinámico que μ-law
- **Implementación**: Algoritmo nativo en Go

### SLIN (PCM Crudo)
- **Bitrate**: 128 kbps (sin compresión)
- **Frame size**: 1 muestra (8kHz)
- **Frame encoding**: 2 bytes por muestra (little-endian)
- **Uso**: Audio crudo, debugging, alta calidad
- **Ventaja**: Sin artefactos de compresión
- **Implementación**: Extracción directa de PCM

### VAD (Voice Activity Detection) - Solo G.729

El encoder G.729 está configurado con VAD deshabilitado (`enableVAD = 0`):
- **VAD = 0**: Todos los frames se codifican como voz (más simple)
- **VAD = 1**: Detecta silencios y los codifica eficientemente (ahorra bandwidth)

Puedes modificar esta configuración en `transcoding.go` línea 19.

## 🐛 Solución de problemas

### Error: "WAV no es PCM"
Tu archivo está en formato comprimido. Conviértelo con FFmpeg:
```bash
ffmpeg -i archivo.wav -acodec pcm_s16le salida.wav
```

### Error: "se requiere mono (1 canal)"
Tu archivo es estéreo. Conviértelo a mono:
```bash
ffmpeg -i archivo.wav -ac 1 salida.wav
```

### Error: "se requiere 8000 Hz"
Cambia la frecuencia de muestreo:
```bash
ffmpeg -i archivo.wav -ar 8000 salida.wav
```

### Error: "se requiere 16-bit PCM"
Ajusta el formato de muestra:
```bash
ffmpeg -i archivo.wav -sample_fmt s16 salida.wav
```

### Conversión todo-en-uno con FFmpeg
```bash
ffmpeg -i entrada.mp3 -ar 8000 -ac 1 -sample_fmt s16 -acodec pcm_s16le salida.wav
```

### Raspberry Pi: "libbcg729.so no encontrada"
```bash
# Agregar ruta de librería
export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

# O actualizar caché de librerías
sudo ldconfig
```

### Makefile: Docker no está corriendo
```bash
# Iniciar servicio Docker
sudo systemctl start docker

# O verificar estado de Docker
docker info
```

## 📝 Notas importantes

- ⚠️ El archivo de salida `.g729` es un **raw bitstream** sin contenedor
- ⚠️ Para reproducir archivos G.729, necesitas un reproductor compatible o convertirlos de vuelta a WAV
- 💡 **Tip**: Usa `ffmpeg -f g729 -i output.g729 -ar 8000 -ac 1 -c:a pcm_s16le output.wav` para convertir G.729 a WAV
- ⚠️ Algunos codecs G.729 están sujetos a patentes (verifica en tu jurisdicción)
- ⚠️ `bcg729` es una implementación de código abierto y libre de regalías
- 🍓 Para Raspberry Pi, elige Alpine (más pequeño) o Debian (más compatible)
- 🔧 El Makefile automatiza todo el proceso de compilación y despliegue

## 📚 Referencias

- [bcg729 - Biblioteca codec G.729](https://github.com/BelledonneCommunications/bcg729)
- [go-wav - Parser WAV para Go](https://github.com/youpy/go-wav)
- [ITU-T G.729 Specification](https://www.itu.int/rec/T-REC-G.729)
- [CGO Documentation](https://pkg.go.dev/cmd/cgo)
- [GNU Make Manual](https://www.gnu.org/software/make/manual/)

## 📄 Licencia

**wav2multi** está licenciado bajo la **Apache License 2.0**.

Copyright © 2025 Federico Pereira <fpereira@cnsoluciones.com>

### Qué Significa Para Ti

✅ **PUEDES:**
- ✓ Usar comercialmente (gratis para código abierto)
- ✓ Modificar el código
- ✓ Distribuir el software
- ✓ Usar en proyectos privados
- ✓ Sublicenciar tus modificaciones

⚠️ **DEBES:**
- ✓ Incluir aviso de copyright
- ✓ Incluir archivo LICENSE en distribuciones
- ✓ Indicar cambios significativos realizados
- ✓ Incluir archivo NOTICE
- ✓ Proporcionar atribución

❌ **NO PUEDES:**
- ✗ Responsabilizar al autor
- ✗ Usar el nombre del autor para promoción sin permiso
- ✗ Eliminar avisos de copyright
- ✗ Cambiar la licencia del código original

### Componentes de Terceros

Este software utiliza los siguientes componentes:

- **bcg729** - GPLv3 - Librería codec G.729
  - Copyright © Belledonne Communications
  - https://github.com/BelledonneCommunications/bcg729
  - Las patentes de G.729 expiraron en 2017 - libre de usar mundialmente

- **go-wav** - Licencia MIT - Análisis de archivos WAV
  - Copyright © youpy
  - https://github.com/youpy/go-wav

- **go-audio** - Apache 2.0 - Procesamiento de audio
  - Copyright © Matt Aimonetti
  - https://github.com/go-audio

### Licencia Comercial

¿Necesitas términos de licencia diferentes? Ofrecemos licencias comerciales con:

- ✓ Sin requisitos de atribución en tu producto
- ✓ Soporte prioritario y consultoría
- ✓ Características personalizadas y desarrollo
- ✓ Acuerdos de Nivel de Servicio (SLA)
- ✓ Términos de licencia flexibles

**Contacto para Licencia Comercial:**
- Email: fpereira@cnsoluciones.com
- Empresa: CNSoluciones - Soluciones de Telecomunicaciones y VoIP
- Sitio web: https://cnsoluciones.com

### Información de Patentes

Las patentes del codec G.729 expiraron en 2017. Este software es libre de usar en todo el mundo sin pagos de regalías. La librería `bcg729` es una implementación de código abierto libre de regalías.

Para los términos completos de la licencia, consulta los archivos [LICENSE](LICENSE) y [NOTICE](NOTICE).

## 🤝 Contribuciones

Las contribuciones son bienvenidas. Por favor:
1. Haz fork del repositorio
2. Crea una rama para tu feature (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add some AmazingFeature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## 👨‍💻 Autor

**Federico Pereira** <fpereira@cnsoluciones.com>

Proyecto de conversión de audio WAV a múltiples formatos usando Go y CGO.

### 🏢 CNSoluciones

Este proyecto es parte de CNSoluciones, especializada en soluciones de telecomunicaciones y VoIP.

---

**¿Preguntas o problemas?** Abre un issue en el repositorio.

## 🌐 Versiones de idioma

- 🇺🇸 [English](README.md)
- 🇪🇸 [Español](README.es.md) (Actual)
- 🍓 [Guía Raspberry Pi](README-RASPBERRYPI.md)
