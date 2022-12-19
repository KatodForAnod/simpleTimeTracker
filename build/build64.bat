go env set GOOS=windows
go env set GOARCH=amd64

rsrc.exe -manifest simpleTimeTracker.exe.manifest -o rcrs.syso
windres -o res.syso simpleTimeTracker.rc

go build -o simpleTimeTracker64.exe

del rcrs.syso res.syso
