REM To run this script, build before and run it from the install folder

start mosquitto -v
start /d ".\sub\" log-sub.exe
start /d ".\sub\" database-sub.exe
start /d ".\api\" app.exe