# Settings
FILEPATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
WORKDIR="${FILEPATH:0:-9}"
# Set the path of your project installation
INSTALL_PATH="/home/user/project"
# Set the airport IATA code
AIRPORT="NTE"

# Change directory to the workdir
cd $WORKDIR

# Clean & Build
go clean ./...
# rm -r "$INSTALL_PATH"
go install ./...

# Making dir
mkdir "$INSTALL_PATH"
mkdir "$INSTALL_PATH/$AIRPORT"
mkdir "$INSTALL_PATH/$AIRPORT/sensors"
mkdir "$INSTALL_PATH/$AIRPORT/sensors/pressure"
mkdir "$INSTALL_PATH/$AIRPORT/sensors/temperature"
mkdir "$INSTALL_PATH/$AIRPORT/sensors/wind"
mkdir "$INSTALL_PATH/sub"
mkdir "$INSTALL_PATH/$AIRPORT/log"
mkdir "$INSTALL_PATH/$AIRPORT/log/temperature"
mkdir "$INSTALL_PATH/$AIRPORT/log/pressure"
mkdir "$INSTALL_PATH/$AIRPORT/log/wind"
mkdir "$INSTALL_PATH/api"

# Move bins
mv "$GOPATH/bin/pressure" "$INSTALL_PATH/$AIRPORT/sensors/pressure/"
mv "$GOPATH/bin/temperature" "$INSTALL_PATH/$AIRPORT/sensors/temperature/"
mv "$GOPATH/bin/wind" "$INSTALL_PATH/$AIRPORT/sensors/wind/"
mv "$GOPATH/bin/database-sub" "$INSTALL_PATH/sub"
mv "$GOPATH/bin/log-sub" "$INSTALL_PATH/sub"
mv "$GOPATH/bin/app" "$INSTALL_PATH/api"

# Move configs
cp "$WORKDIR/configs/pubs/pressure/config.yml" "$INSTALL_PATH/$AIRPORT/sensors/pressure/"
cp "$WORKDIR/configs/pubs/temperature/config.yml" "$INSTALL_PATH/$AIRPORT/sensors/temperature/"
cp "$WORKDIR/configs/pubs/wind/config.yml" "$INSTALL_PATH/$AIRPORT/sensors/wind/"
cp "$WORKDIR/scripts/runservices.sh" "$INSTALL_PATH/"
cp "$WORKDIR/scripts/runsensors.sh" "$INSTALL_PATH/$AIRPORT/"
