## Rocket
Program for monitoring nodes data.
## How to run Rocket

### Installing Go
Rocket requires Go 1.13 to compile, please refer to the [official documentation](https://go.dev/doc/install) for how to install Go in your system.

### Cloning Rocket:
cd to your catalog
copy url of progect Rocket
```
git init
git clone *url* 
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
You need docker and go >= 1.13
```
go get github.com/fyne-io/fyne-cross
fyne-cross android -app-id rocket.android.app -icon icon.png
```
