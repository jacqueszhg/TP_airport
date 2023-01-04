@echo off

REM Settings
SET WORKDIR="C:\Users\jacqu\OneDrive\Documents\1.IMT_Atlantique_cours\Architecture_distribuer\cours2_GO\Projet"
SET INSTALL_PATH="C:\Users\jacqu\Desktop\airport"
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
mkdir "%INSTALL_PATH%\%AIRPORT%\sub"
mkdir "%INSTALL_PATH%\%AIRPORT%\log"
mkdir "%INSTALL_PATH%\%AIRPORT%\log\temperature"
mkdir "%INSTALL_PATH%\%AIRPORT%\log\pressure"
mkdir "%INSTALL_PATH%\%AIRPORT%\log\wind"

REM Move bins
move "%GOPATH%\bin\pressure.exe" "%INSTALL_PATH%\%AIRPORT%\sensors\pressure\"
move "%GOPATH%\bin\temperature.exe" "%INSTALL_PATH%\%AIRPORT%\sensors\temperature\"
move "%GOPATH%\bin\wind.exe" "%INSTALL_PATH%\%AIRPORT%\sensors\wind\"
move "%GOPATH%\bin\database-sub.exe" "%INSTALL_PATH%\%AIRPORT%\sub"
move "%GOPATH%\bin\log-sub.exe" "%INSTALL_PATH%\%AIRPORT%\sub"

REM Move configs
copy "%WORKDIR%\configs\pubs\pressure\config.yml" "%INSTALL_PATH%\%AIRPORT%\sensors\pressure\"
copy "%WORKDIR%\configs\pubs\temperature\config.yml" "%INSTALL_PATH%\%AIRPORT%\sensors\temperature\"
copy "%WORKDIR%\configs\pubs\wind\config.yml" "%INSTALL_PATH%\%AIRPORT%\sensors\wind\"