# 🍓 Transcoder para Raspberry Pi 5 - Guía Completa

**Instalación rápida y fácil del transcoder de audio multi-formato en Raspberry Pi 5**

Federico Pereira <fpereira@cnsoluciones.com>

## 📋 Descripción

Esta guía te ayudará a compilar e instalar el transcoder de audio multi-formato (G.729, μ-law, A-law, SLIN) en tu Raspberry Pi 5 usando el Makefile incluido.

## 🚀 Instalación Rápida (Desde tu PC)

### Prerrequisitos

- Docker instalado en tu PC
- Make instalado
- Acceso SSH a tu Raspberry Pi 5
- Raspberry Pi OS 64-bit instalado

### Opción 1: Alpine/musl (Más liviano)

```bash
# 1. Compilar y extraer binarios
make all

# 2. Copiar archivos a Raspberry Pi
scp -r bin/ lib/ pi@raspberrypi.local:~/transcoder/

# 3. Conectar a Raspberry Pi
ssh pi@raspberrypi.local

# 4. Instalar binario y librerías
sudo mv ~/transcoder/bin/transcoding /usr/local/bin/transcoding
sudo chmod +x /usr/local/bin/transcoding
sudo mv ~/transcoder/lib/libbcg729.so* /usr/local/lib/
sudo ldconfig

# 5. Verificar instalación
transcoding --version
```

### Opción 2: Debian/glibc (Más compatible)

```bash
# 1. Compilar y extraer binarios (versión Debian)
make all-debian

# 2. Copiar archivos a Raspberry Pi
scp -r bin/ lib/ pi@raspberrypi.local:~/transcoder/

# 3. Conectar a Raspberry Pi
ssh pi@raspberrypi.local

# 4. Instalar binario y librerías
sudo mv ~/transcoder/bin/transcoding /usr/local/bin/transcoding
sudo chmod +x /usr/local/bin/transcoding
sudo mv ~/transcoder/lib/libbcg729.so* /usr/local/lib/
sudo ldconfig

# 5. Verificar instalación
transcoding --version
```

## 📦 Comandos del Makefile

El proyecto incluye un Makefile que automatiza todo el proceso de compilación:

| Comando | Descripción | Cuándo Usar |
|---------|-------------|-------------|
| `make help` | Muestra ayuda con todos los comandos | Primera vez |
| `make build` | Compila imagen ARM64 (Alpine) | Compilación manual |
| `make build-debian` | Compila imagen ARM64 (Debian) | Compilación manual |
| `make extract` | Extrae binarios de Alpine | Después de build |
| `make extract-debian` | Extrae binarios de Debian | Después de build-debian |
| `make all` | Compila + extrae (Alpine) | **Recomendado** |
| `make all-debian` | Compila + extrae (Debian) | Para Raspberry Pi OS |
| `make test` | Verifica el binario | Después de extract |
| `make clean` | Limpia todo | Empezar de cero |

### 🎯 Flujo de Trabajo Recomendado

```bash
# Ver ayuda
make help

# Compilar todo (Alpine - más pequeño)
make all

# O compilar todo (Debian - más compatible con Raspberry Pi OS)
make all-debian

# Verificar binario
make test

# Limpiar si necesitas recompilar
make clean
```

## 📁 Archivos Generados

Después de ejecutar `make all` o `make all-debian`:

```
.
├── bin/
│   └── transcoding          # Binario ARM64 (2.4MB)
└── lib/
    ├── libbcg729.so         # Librería principal (148KB)
    └── libbcg729.so.0       # Enlace simbólico (148KB)
```

### 🔍 Diferencias entre Alpine y Debian

| Aspecto | Alpine/musl | Debian/glibc |
|---------|-------------|--------------|
| **Tamaño** | Más pequeño | Ligeramente mayor |
| **Compatibilidad** | Requiere musl | Compatible con Raspberry Pi OS |
| **Velocidad** | Similar | Similar |
| **Recomendado para** | Sistemas custom | Raspberry Pi OS estándar |

## 🎯 Uso Básico en Raspberry Pi

Una vez instalado, puedes usar el transcoder directamente:

```bash
# Convertir a G.729 (alta compresión, 8 kbps)
transcoding input.wav output.g729 --format g729

# Convertir a μ-law (telefonía US, 64 kbps)
transcoding input.wav output.ulaw --format ulaw

# Convertir a A-law (telefonía europea, 64 kbps)
transcoding input.wav output.alaw --format alaw

# Convertir a SLIN (PCM crudo, 128 kbps)
transcoding input.wav output.slin --format slin
```

### 💡 Ejemplos Prácticos

```bash
# Convertir un solo archivo
transcoding audio.wav audio.g729 --format g729

# Procesamiento por lotes
for file in *.wav; do
    transcoding "$file" "${file%.wav}.g729" --format g729
done

# Conversión con verificación
transcoding input.wav output.g729 --format g729 && \
ffmpeg -f g729 -i output.g729 -ar 8000 -ac 1 -c:a pcm_s16le verify.wav
```

## 📊 Formatos Soportados

| Formato | Bitrate | Uso Principal | Calidad | Tamaño |
|---------|---------|---------------|---------|--------|
| **G.729** | 8 kbps | VoIP, máxima compresión | Buena para voz | Muy pequeño |
| **μ-law** | 64 kbps | Telefonía US | Buena para voz | Mediano |
| **A-law** | 64 kbps | Telefonía europea | Buena para voz | Mediano |
| **SLIN** | 128 kbps | Audio sin compresión | Perfecta | Grande |

## 🔧 Requisitos del Sistema

### Hardware
- **Raspberry Pi**: Modelo 5 (cualquier variante)
- **RAM**: Mínimo 1GB, recomendado 2GB+
- **Almacenamiento**: 100MB libres
- **Arquitectura**: ARM64 (aarch64)

### Software
- **OS**: Raspberry Pi OS 64-bit (Bookworm o superior)
- **Dependencias**: Se instalan automáticamente con ldconfig
- **Conexión**: No requiere internet después de la instalación

### Archivos WAV de Entrada
- **Formato**: PCM (AudioFormat = 1)
- **Canales**: Mono (1 canal)
- **Sample Rate**: 8000 Hz
- **Bits por muestra**: 16-bit

## 📊 Rendimiento en Raspberry Pi 5

Resultados de benchmark con diferentes formatos:

| Formato | Velocidad | Uso CPU | Uso RAM | Compresión |
|---------|-----------|---------|---------|------------|
| **G.729** | ~10x real-time | ~30% | ~50MB | 16:1 |
| **μ-law** | ~50x real-time | ~15% | ~30MB | 2:1 |
| **A-law** | ~50x real-time | ~15% | ~30MB | 2:1 |
| **SLIN** | ~100x real-time | ~5% | ~20MB | 1:1 |

*Medido en Raspberry Pi 5 (4GB RAM, Cortex-A76)*

## 🆘 Solución de Problemas

### Error: "command not found"

```bash
# Verificar instalación
which transcoding
ls -la /usr/local/bin/transcoding

# Reinstalar si es necesario
sudo cp bin/transcoding /usr/local/bin/
sudo chmod +x /usr/local/bin/transcoding
```

### Error: "libbcg729.so not found"

```bash
# Verificar librerías
ls -la /usr/local/lib/libbcg729.so*

# Actualizar cache de librerías
sudo ldconfig

# Agregar ruta temporalmente
export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

# O agregar permanentemente a ~/.bashrc
echo 'export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH' >> ~/.bashrc
source ~/.bashrc
```

### Error: "Architecture mismatch"

```bash
# Verificar arquitectura del binario
file /usr/local/bin/transcoding
# Debe mostrar: ARM aarch64, 64-bit LSB executable

# Verificar arquitectura del sistema
uname -m
# Debe mostrar: aarch64

# Si no coinciden, recompilar con el Makefile correcto
make clean
make all-debian  # Para Raspberry Pi OS
```

### Error: "WAV is not PCM" o formato incorrecto

```bash
# Convertir con FFmpeg a formato compatible
ffmpeg -i input.mp3 -ar 8000 -ac 1 -sample_fmt s16 -acodec pcm_s16le output.wav

# Verificar formato del WAV
ffmpeg -i input.wav
```

### Docker no disponible en el Makefile

```bash
# Verificar Docker en tu PC (no en Raspberry Pi)
docker --version

# Iniciar Docker si está detenido
sudo systemctl start docker

# Verificar que Docker está corriendo
docker info
```

## ✅ Verificación de Archivos

### Verificar conversión G.729
```bash
# Convertir de vuelta a WAV
ffmpeg -f g729 -i output.g729 -ar 8000 -ac 1 -c:a pcm_s16le verify.wav

# Reproducir para verificar calidad
aplay verify.wav
```

### Verificar conversión μ-law/A-law
```bash
# μ-law
ffmpeg -f mulaw -i output.ulaw -ar 8000 -ac 1 -c:a pcm_s16le verify-ulaw.wav

# A-law
ffmpeg -f alaw -i output.alaw -ar 8000 -ac 1 -c:a pcm_s16le verify-alaw.wav
```

### Verificar conversión SLIN
```bash
ffmpeg -f s16le -ar 8000 -ac 1 -i output.slin -c:a pcm_s16le verify-slin.wav
```

## 🔄 Actualización del Transcoder

```bash
# En tu PC: compilar nueva versión
cd /path/to/transcoding
git pull  # Si usas git
make clean
make all-debian

# Copiar a Raspberry Pi
scp -r bin/ lib/ pi@raspberrypi.local:~/transcoder/

# En Raspberry Pi: actualizar
ssh pi@raspberrypi.local
sudo mv ~/transcoder/bin/transcoding /usr/local/bin/
sudo mv ~/transcoder/lib/libbcg729.so* /usr/local/lib/
sudo ldconfig

# Verificar nueva versión
transcoding --version
```

## 🧹 Desinstalación

```bash
# En Raspberry Pi
sudo rm /usr/local/bin/transcoding
sudo rm /usr/local/lib/libbcg729.so*
sudo ldconfig

# Limpiar archivos temporales
rm -rf ~/transcoder/
```

## 🎓 Ejemplos Avanzados

### Script de conversión por lotes

```bash
#!/bin/bash
# convert_all.sh - Convierte todos los WAV a G.729

for file in *.wav; do
    if [ -f "$file" ]; then
        output="${file%.wav}.g729"
        echo "Procesando: $file -> $output"
        transcoding "$file" "$output" --format g729
        if [ $? -eq 0 ]; then
            echo "✅ Completado: $output"
        else
            echo "❌ Error: $file"
        fi
    fi
done

echo "Conversión por lotes completada!"
```

### Monitorear uso de recursos

```bash
# Mientras se ejecuta la conversión
top -p $(pgrep transcoding)

# O con htop
htop -p $(pgrep transcoding)
```

### Conversión con prioridad baja (para no afectar otros procesos)

```bash
nice -n 19 transcoding input.wav output.g729 --format g729
```

## 📚 Recursos Adicionales

### Documentación Completa
- [README.md](README.md) - Documentación en inglés
- [README.es.md](README.es.md) - Documentación en español
- [wav2multi-lib/README.md](wav2multi-lib/README.md) - Librería Go

### Referencias Técnicas
- [bcg729 Library](https://github.com/BelledonneCommunications/bcg729)
- [G.729 Specification](https://www.itu.int/rec/T-REC-G.729)
- [Raspberry Pi OS](https://www.raspberrypi.com/software/)

### Obtener Ayuda
```bash
# Ayuda del programa
transcoding --help

# Versión instalada
transcoding --version

# Ayuda del Makefile
make help
```

## 🔬 Información Técnica

### Dependencias del Binario

```bash
# Ver dependencias dinámicas
ldd /usr/local/bin/transcoding

# Resultado esperado:
# linux-vdso.so.1
# libbcg729.so.0 => /usr/local/lib/libbcg729.so.0
# libc.so.6 => /lib/aarch64-linux-gnu/libc.so.6
# /lib/ld-linux-aarch64.so.1
```

### Compilación del Binario

El binario se compila usando:
- **Go**: 1.23
- **CGO**: Habilitado
- **Target**: linux/arm64
- **Librería**: libbcg729 (compartida)

### Estructura de Directorios

```
/usr/local/
├── bin/
│   └── transcoding          # Binario principal
└── lib/
    ├── libbcg729.so         # Librería compartida
    └── libbcg729.so.0       # Versión específica
```

## 🎯 Casos de Uso

### 1. Sistema VoIP Casero
Convierte archivos de audio para usar en tu propio sistema VoIP con Asterisk.

```bash
# Preparar mensajes de voz para Asterisk
transcoding welcome.wav /var/lib/asterisk/sounds/custom/welcome.g729 --format g729
```

### 2. Optimización de Almacenamiento
Reduce el tamaño de grabaciones de audio largas.

```bash
# Conversión masiva para ahorrar espacio
for f in recordings/*.wav; do
    transcoding "$f" "compressed/$(basename ${f%.wav}.g729)" --format g729
done
```

### 3. Pruebas de Calidad de Codecs
Compara diferentes codecs en tu hardware.

```bash
# Generar todas las versiones
transcoding test.wav test.g729 --format g729
transcoding test.wav test.ulaw --format ulaw
transcoding test.wav test.alaw --format alaw
transcoding test.wav test.slin --format slin
```

## 💡 Consejos y Trucos

### 1. Automatización con Cron

```bash
# Editar crontab
crontab -e

# Agregar tarea diaria de conversión
0 2 * * * /home/pi/scripts/convert_daily.sh
```

### 2. Integración con Scripts Python

```python
import subprocess

def convert_to_g729(input_file, output_file):
    cmd = ['transcoding', input_file, output_file, '--format', 'g729']
    result = subprocess.run(cmd, capture_output=True, text=True)
    return result.returncode == 0
```

### 3. Monitoreo de Conversiones

```bash
# Crear log de conversiones
transcoding input.wav output.g729 --format g729 2>&1 | tee conversion.log
```

## 📞 Soporte

- **Autor**: Federico Pereira
- **Email**: fpereira@cnsoluciones.com
- **GitHub**: [lordbasex/wav2multi](https://github.com/lordbasex/wav2multi)
- **Empresa**: CNSoluciones - Soluciones de telecomunicaciones y VoIP

### Reportar Problemas

Si encuentras algún problema:
1. Verifica que tu Raspberry Pi sea ARM64 (64-bit)
2. Asegúrate de tener Raspberry Pi OS 64-bit
3. Ejecuta `make test` para verificar el binario
4. Revisa los logs con `transcoding --help`
5. Abre un issue en GitHub con los detalles

---

**¡Disfruta del transcoder de audio multi-formato en tu Raspberry Pi 5! 🍓🎵**

Made with ❤️ by CNSoluciones
