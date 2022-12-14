@echo off

REM Settings
SET WORKDIR="%HOME%/Documents/Projets/TP_airport/"
SET INSTALL_PATH="%HOME%/airport"
SET AIRPORT="NTE"

REM Change directory to the workdir
cd %WORKDIR%

REM Clean & Build
go clean ./...
rmdir /S /Q "%INSTALL_PATH%"
go install ./...

REM Making dir
mkdir "%INSTALL_PATH%"
mkdir "%INSTALL_PATH%/%AIRPORT%"
mkdir "%INSTALL_PATH%/%AIRPORT%/sensors"
mkdir "%INSTALL_PATH%/%AIRPORT%/sensors/pressure"
mkdir "%INSTALL_PATH%/%AIRPORT%/sensors/temperature"
mkdir "%INSTALL_PATH%/%AIRPORT%/sensors/wind"
mkdir "%INSTALL_PATH%/%AIRPORT%/pubs"

REM Move bins
mv "%GOPATH%/bin/pressure" "%INSTALL_PATH%/%AIRPORT%/sensors/pressure/"
mv "%GOPATH%/bin/temperature" "%INSTALL_PATH%/%AIRPORT%/sensors/temperature/"
mv "%GOPATH%/bin/wind" "%INSTALL_PATH%/%AIRPORT%/sensors/wind/"

