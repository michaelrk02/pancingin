# PancingIN

## Prasyarat

1. Telah memasang Go (direkomendasikan versi terbaru)
2. Mengetahui cara menggunakan terminal / command line
3. Telah melakukan clone/download repository ini

## Prosedur pemasangan database

1. Jalankan GUI client apapun untuk Microsoft SQL Server
2. Buat database baru, misal: **DBMSProjectBKel7**
3. Jalankan **seluruh** perintah SQL yang tertera pada file **database.sql** dalam database yang telah dibuat tersebut

## Prosedur compile program

1. Buat folder **dist** pada folder ini
2. Buka terminal / command line yang mengarah ke folder ini
3. Jalankan perintah `go build -o dist/<program>` (misal: `go build -o dist/pancingin.exe` dalam Windows atau `go build -o dist/pancingin` dalam Linux)
4. Copy file **config.json** dan masukkan dalam folder **dist**
5. Edit file **config.json** yang terdapat dalam folder **dist** menyesuaikan konfigurasi database anda