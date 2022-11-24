go env set GOOS=windows
go env set GOARCH=i386

rsrc.exe -manifest simpleTimeTracker.exe.manifest -o rcrs.syso
windres -o res.syso simpleTimeTracker.rc

go build -o simpleTimeTracker64.exe

del rcrs.syso res.syso