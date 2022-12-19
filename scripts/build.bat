@echo off

REM Settings
SET WORKDIR="C:\Users\lucas\GolandProjects\TP_airport"
SET INSTALL_PATH="%homedrive%%homepath%\airport"
SET AIRPORT="NTE"

REM Change directory to the workdir
cd %WORKDIR%

REM Clean & Build
go clean .\...
rmdir /S /Q "%INSTALL_PATH%"
go install .\...

REM Making dir
mkdir "%INSTALL_PATH%"
mkdir "%INSTALL_PATH%\%AIRPORT%"
mkdir "%INSTALL_PATH%\%AIRPORT%\sensors"
mkdir "%INSTALL_PATH%\%AIRPORT%\sensors\pressure"
mkdir "%INSTALL_PATH%\%AIRPORT%\sensors\temperature"
mkdir "%INSTALL_PATH%\%AIRPORT%\sensors\wind"

REM Move bins
move "%GOPATH%\bin\pressure.exe" "%INSTALL_PATH%\%AIRPORT%\sensors\pressure\"
move "%GOPATH%\bin\temperature.exe" "%INSTALL_PATH%\%AIRPORT%\sensors\temperature\"
move "%GOPATH%\bin\wind.exe" "%INSTALL_PATH%\%AIRPORT%\sensors\wind\"

REM Move configs
copy "%WORKDIR%\configs\pubs\pressure\config.yml" "%INSTALL_PATH%\%AIRPORT%\sensors\pressure\"
copy "%WORKDIR%\configs\pubs\temperature\config.yml" "%INSTALL_PATH%\%AIRPORT%\sensors\temperature\"
copy "%WORKDIR%\configs\pubs\wind\config.yml" "%INSTALL_PATH%\%AIRPORT%\sensors\wind\"