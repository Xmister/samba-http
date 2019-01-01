# samba-http
An example usage of libdsm-go and libsmb2-go providing the FileSystem for a go http server.

## Building
There are helper scripts if you don't want to manually compile the required libs. These are using the `Dockerfile` and `Dockerfile-android` images for building the binaries.

### Binary for PC
`./build` will build a static binary to `$GOPATH/bin` that runs on your PC.

### JNI bindings for Android
`./build-android` will build `streamer.aar` package to `$GOPATH/bin` that you can import in Android studio and then the exported functions will be available from your java code: `StartServer()` and `StopServer()`
