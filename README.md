

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
