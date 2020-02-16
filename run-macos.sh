#!/bin/sh

APP="Noodle.app"
mkdir -p $APP/out/Contents/{MacOS,Resources}
go build -o $APP/out/Contents/MacOS/noodle
cat > $APP/out/Contents/Info.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CFBundleExecutable</key>
	<string>noodle</string>
	<key>CFBundleIconFile</key>
	<string>icon.icns</string>
	<key>CFBundleIdentifier</key>
	<string>com.zserge.lorca.example</string>
</dict>
</plist>
EOF
cp icons/icon.icns $APP/out/Contents/Resources/icon.icns
find $APP/out
