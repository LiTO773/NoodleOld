#!/bin/sh

APP=noodle
APPDIR=${APP}_1.0.0

mkdir -p $APPDIR/out/usr/bin
mkdir -p $APPDIR/out/usr/share/applications
mkdir -p $APPDIR/out/usr/share/icons/hicolor/1024x1024/apps
mkdir -p $APPDIR/out/usr/share/icons/hicolor/256x256/apps
mkdir -p $APPDIR/out/DEBIAN

go build -o $APPDIR/out/usr/bin/$APP

cp icons/icon.png $APPDIR/out/usr/share/icons/hicolor/1024x1024/apps/${APP}.png
cp icons/icon.png $APPDIR/out/usr/share/icons/hicolor/256x256/apps/${APP}.png

cat > $APPDIR/out/usr/share/applications/${APP}.desktop << EOF
[Desktop Entry]
Version=1.0
Type=Application
Name=$APP
Exec=$APP
Icon=$APP
Terminal=false
StartupWMClass=Lorca
EOF

cat > $APPDIR/out/DEBIAN/control << EOF
Package: ${APP}
Version: 1.0-0
Section: base
Priority: optional
Architecture: amd64
Maintainer: Serge Zaitsev <zaitsev.serge@gmail.com>
Description: Example for Lorca GUI toolkit
EOF

dpkg-deb --build $APPDIR/out
