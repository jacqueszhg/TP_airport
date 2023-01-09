# To run this script, build before and run it from the install folder

mosquitto -v &
./sub/log-sub.exe &
./sub/database-sub.exe &
./api/app.exe &