#!/bin/sh

# Settings
WORKDIR="$HOME/Documents/Projets/TP_airport/"
INSTALL_PATH="$HOME/airport"
AIRPORT="NTE"

# Change directory to the workdir
cd "$WORKDIR"

# Clean & Build
go clean ./...
rm -rd "$INSTALL_PATH"
go install ./...

# Making dir
mkdir "$INSTALL_PATH/"
mkdir "$INSTALL_PATH/$AIRPORT/"
mkdir "$INSTALL_PATH/$AIRPORT/sensors"
mkdir "$INSTALL_PATH/$AIRPORT/sensors/pressure"
mkdir "$INSTALL_PATH/$AIRPORT/sensors/temperature"
mkdir "$INSTALL_PATH/$AIRPORT/sensors/wind"
mkdir "$INSTALL_PATH/$AIRPORT/pubs"

# Move bins
mv "$GOPATH/bin/pressure" "$INSTALL_PATH/$AIRPORT/sensors/pressure/"
mv "$GOPATH/bin/temperature" "$INSTALL_PATH/$AIRPORT/sensors/temperature/"
mv "$GOPATH/bin/wind" "$INSTALL_PATH/$AIRPORT/sensors/wind/"