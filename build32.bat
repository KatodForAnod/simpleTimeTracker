go env -w GOOS=windows
go env -w GOARCH=386

rsrc.exe -arch 386 -manifest simpleTimeTracker.exe.manifest -o rcrs.syso
windres -o res.syso simpleTimeTracker.rc -F pe-i386

go build -o simpleTimeTracker32.exe
