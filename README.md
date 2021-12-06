## Rocket
Program for monitoring nodes data.
## How to run Rocket

### Installing Go
Rocket requires Go 1.13 to compile, please refer to the [official documentation](https://go.dev/doc/install) for how to install Go in your system.

### Installing Rocket:
```
go get github.com/adarocket/rocket 
```
### Installing Fyne:
```
go get fyne.io/fyne/v2/cmd/fyne
fyne install
```
### Creating binaries for PC
```
fyne package -os *your os(windows/darwin/...)* -icon icon.png
```

### Creating binaries for android
```
export ANDROID_NDK_HOME=/home/shpileuski/Android/Sdk/ndk/21.3.6528147
fyne package -os android -appID com.example.myapp -icon icon.png
```
