## Rocket
Programm for monitoring nodes data.
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
### Creating binaries
```
fyne package -os *your os(windows/darwin/...)*
```

export ANDROID_NDK_HOME=/home/shpileuski/Android/Sdk/ndk/21.3.6528147
fyne package -os android -appID com.example.myapp -icon icon.png

fyne package -os ios -appID com.example.myapp -icon icon.png

fyne package -os darwin -icon icon.png
fyne package -os linux -icon icon.png
fyne package -os windows -icon icon.png


 adb install rocket.apk 
 adb logcat

Online
NodePerformance
Blocks
KESData
NodeState
