linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=386 go build github.com/getlantern/osversion/osversion

windows:
	CGO_ENABLED=1 GOOS=windows GOARCH=386 go build github.com/getlantern/osversion/osversion

android:
	GOOS=android GOARCH=arm gomobile build -target=android github.com/getlantern/osversion/osversion && \
  adb install -r osversion.apk
