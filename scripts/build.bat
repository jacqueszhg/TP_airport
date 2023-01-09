@echo off

REM Settings
REM Get the current directory
SET FILEPATH=%~dp0
REM Get only the root project directory
SET WORKDIR="%FILEPATH:~0,-9%"
REM Set the path of your project installation
SET INSTALL_PATH="C:\airportInstall"
REM Set the airport IATA code
SET AIRPORT="AHO"

REM Change directory to the workdir
cd %WORKDIR%

REM Clean & Build
go clean .\...
go install .\...

REM Making dir
mkdir "%INSTALL_PATH%"
mkdir "%INSTALL_PATH%\%AIRPORT%"
mkdir "%INSTALL_PATH%\%AIRPORT%\sensors"
mkdir "%INSTALL_PATH%\%AIRPORT%\sensors\pressure"
mkdir "%INSTALL_PATH%\%AIRPORT%\sensors\temperature"
mkdir "%INSTALL_PATH%\%AIRPORT%\sensors\wind"
mkdir "%INSTALL_PATH%\sub"
mkdir "%INSTALL_PATH%\%AIRPORT%\log"
mkdir "%INSTALL_PATH%\%AIRPORT%\log\temperature"
mkdir "%INSTALL_PATH%\%AIRPORT%\log\pressure"
mkdir "%INSTALL_PATH%\%AIRPORT%\log\wind"
mkdir "%INSTALL_PATH%\api"

REM Move bins
move "%GOPATH%\bin\pressure.exe" "%INSTALL_PATH%\%AIRPORT%\sensors\pressure\"
move "%GOPATH%\bin\temperature.exe" "%INSTALL_PATH%\%AIRPORT%\sensors\temperature\"
move "%GOPATH%\bin\wind.exe" "%INSTALL_PATH%\%AIRPORT%\sensors\wind\"
move "%GOPATH%\bin\database-sub.exe" "%INSTALL_PATH%\sub"
move "%GOPATH%\bin\log-sub.exe" "%INSTALL_PATH%\sub"
move "%GOPATH%\bin\app.exe" "%INSTALL_PATH%\api"

REM Move configs
copy "%WORKDIR%\configs\pubs\pressure\config.yml" "%INSTALL_PATH%\%AIRPORT%\sensors\pressure\"
copy "%WORKDIR%\configs\pubs\temperature\config.yml" "%INSTALL_PATH%\%AIRPORT%\sensors\temperature\"
copy "%WORKDIR%\configs\pubs\wind\config.yml" "%INSTALL_PATH%\%AIRPORT%\sensors\wind\"
copy "%WORKDIR%\configs\subs\database\config.yml" "%INSTALL_PATH%\sub\databaseConfig.yml"
copy "%WORKDIR%\configs\subs\log\config.yml" "%INSTALL_PATH%\sub\logConfig.yml"
copy "%WORKDIR%\scripts\runservices.bat" "%INSTALL_PATH%\"
copy "%WORKDIR%\scripts\runsensors.bat" "%INSTALL_PATH%\%AIRPORT%\"