.PHONY : all dockerdiff

dockerdiff:
	go build -ldflags '-w -s' -o dockerdiff cmd/*.go

all: dockerdiff_linux_386 dockerdiff_linux_amd64 dockerdiff_linux_arm dockerdiff_linux_arm64 dockerdiff_win32.exe dockerdiff_win64.exe

dockerdiff_linux_386:
	GOOS="linux" GOARCH="386" go build -ldflags '-w -s' -o dockerdiff_linux_386 cmd/*.go

dockerdiff_linux_amd64:
	GOOS="linux" GOARCH="amd64" go build -ldflags '-w -s' -o dockerdiff_linux_amd64 cmd/*.go

dockerdiff_linux_arm:
	GOOS="linux" GOARCH="arm" go build -ldflags '-w -s' -o dockerdiff_linux_arm cmd/*.go

dockerdiff_linux_arm64:
	GOOS="linux" GOARCH="arm64" go build -ldflags '-w -s' -o dockerdiff_linux_arm64 cmd/*.go

#dockerdiff_win32.exe:
#	GOOS="windows" GOARCH="386" go build -ldflags '-w -s' -o dockerdiff_win32.exe cmd/*.go
#
#dockerdiff_win64.exe:
#	GOOS="windows" GOARCH="amd64" go build -ldflags '-w -s' -o dockerdiff_win64.exe cmd/*.go

install:
	install dockerdiff /usr/bin/dockerdiff

clean:
	rm -f dockerdiff dockerdiff_linux_386 dockerdiff_linux_amd64 dockerdiff_linux_arm dockerdiff_linux_arm64 dockerdiff_win32.exe dockerdiff_win64.exe