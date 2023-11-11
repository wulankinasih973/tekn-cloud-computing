# Latihan Minggu 7
## Instalasi Docker
Jalankan file instaler Docker yang sudah diunduh. Centang opsi *Use WSL 2 instead of Hyper-V* dan *Add shortcut to desktop*.
![instal-docker](gambar-01.png)  
  
Tunggu hingga proses *unpacking files* selesai.   
![proses-unpacking](gambar-02.png)  
  
Instalasi telah selesai. Restart untuk menyelesaikan instalasi
![instalasi-selesai](gambar-03.png)
![setting-up](gambar-04.png)  
  
Jalankan Docker Desktop kemudian lanjutkan dengan *Sign in* apabila sudah memiliki akun docker atau pilih *Continue without signing in* jika belum.
![sign-in](gambar-05.png)  
  
## 2. Get Started - Docker
### 2.1 Containerize an application
Cloning repositori *getting-started-app*.  
```
git clone https://github.com/docker/getting-started-app.git 
```
![cloning](gambar-06.png)

Masuk ke dalam direktori `getting-started-app` kemudian buat file `Dockerfile`.  
```
cd getting-started-app
```
![getting-started-app](gambar-07.png)

Buatlah file Dockerfile di dalamnya, lalu tambahkan content berikut:
```
# syntax=docker/dockerfile:1

FROM node:18-alpine
WORKDIR /app
COPY . .
RUN yarn install --production
CMD ["node", "src/index.js"]
EXPOSE 3000
```
Build image dengan perintah:
```
docker build -t getting-started .
```
![build-image-1](gambar-08.png)  
![build-image-2](gambar-09.png)
Jalankan container dengan perintah:  
```
docker run -dp 127.0.0.1:3000:3000 getting-started
```
Buka browser kemudian akses ke halaman 127.0.0.1:3000.  
![tampilan-todo-app](gambar-10.png)  
  
Cek dengan perintah `docker ps` untuk melihat daftar container yang berjalan.  
![docker-ps](gambar-11.png)  
  
### 2.2 Update the application
Lakukan modifikasi pada file `src/static/js/app.js`.  
```
- <p className="text-center">No items yet! Add one above</p>
+ <p className="text-center">You have no todo items yet! Add one above!</p>
```
  
Hentikan container yang sedang berjalan.  
```
$ docker stop <ID container atau nama container>
```
Kemudian build ulang image dengan perintah:  
```
$ docker build -t getting-started .
```
![stop-and-build](gambar-12.png)

Jalankan container dengan perintah:  
```
$ docker run -dp 127.0.0.1:3000:3000 getting-started
```
Refresh halaman `127.0.0.1:3000`.  
![after-refresh](gambar-13.png)  
  
### 2.3 Share the application
Login ke [Docker Hub](https://hub.docker.com/), kemudian buatlah repositori baru.
  
Pada bagian *Repository Name* diisi dengan `getting-started` dan pastikan *Visibility* nya `Public`.  
![new-repo](gambar-14.png)  
  
Login ke docker hub melalui terminal dengan perintah berikut:  
```
$ docker login -u NAMA_USERNAME_ANDA
```
![login-docker-hub](gambar-15.png)  
  
Buat *tag* untuk image `getting-started` dengan perintah berikut:  
```
$ docker tag getting-started NAMA_USERNAME_ANDA/getting-started
```
![perintah-docker-tag](gambar-16.png)
![docker-images](gambar-17.png)  
  
Push image yang sudah diberi tag tadi ke docker hub dengan perintah:  
```
$ docker push NAMA_USERNAME_ANDA/getting-started
```
![push-image](gambar-18.png)  

Akses situs [Play with Docker](https://labs.play-with-docker.com/), kemudian login dengan akun docker yang digunakan.  
  
Klik tombol *__ADD NEW INSTANCE__*, kemudian jalankan container dari image yang sudah di-*push* ke docker hub tadi.  
```
$ docker run -dp 0.0.0.0:3000:3000 NAMA_USERNAME_ANDA/getting-started
```
![play-with-docker](gambar-19.png)  
  
Klik nomor port *__3000__* untuk mengakses halaman web.  
![tampilan-todo-app](gambar-20.png)  
  
### 2.4 Persist the DB
Buat sebuah *volume* dengan perintah:  
```
$ docker volume create NAMA_VOLUME
```
![membuat-volume](gambar-21.png)
![membuat-volume](gambar-22.png)  
  
Hapus container sebelumnya dengan perintah `docker rm -f <id>`  

Jalankan kembali container dari image `getting-started` dengan menambahkan parameter `--mount`:  
```
$ docker run -dp 127.0.0.1:3000:3000 --mount type=volume,src=todo-db,target=/etc/todos getting-started
```
![menjalankan-container](gambar-23.png)  
  
Akses Todo App melalui browser kemudian tambahkan beberapa data.  
![add-data](gambar-24.png)  
  
Hentikan dan hapus container yang sedang berjalan.  
![menghentikan-dan-menghapus-container](gambar-25.png)  
  
Jalankan container baru dengan menambahkan parameter `--mount` seperti pada langkah-langkah sebelumnya.  
  
Cek browser kembali dan pastikan data yang sebelumnya sudah pernah diinputkan masih ada.  
![tampilan-todo-app](gambar-26.png)  
  
Untuk mengetahui lokasi dimana volume berada, gunakan perintah:  
```
$ docker volume inspect NAMA_VOLUME
```
![hasil-inspect](gambar-27.png)  
  
### 2.5 Use bind mounts
Jalankan container baru dengan perintah:  
```
docker run -it --mount "type=bind,src=$pwd,target=/src" ubuntu bash
```
Perintah di atas akan melakukan *mount* direktori `src` yang ada di dalam container ke direktori yang saat ini sedang aktif pada host.  
![mount-bind](gambar-28.png)  

Setelah menjalankan perintah, Docker memulai sesi bash interaktif di direktori root sistem file container
![](gambar-29.png)

Dari dalam container, masuk ke direktori `src` kemudian cek isi direktorinya. Setelah itu buat sebuah file baru dengan nama `myfile.txt`.  
![membuat file baru di dalam direktori src]
![cd-src](gambar-30.png)
![create-new-file](gambar-31.png)    
  
Keluar dari container dengan perintah `exit`.
  
### 2.6 Multi-container apps
Buatlah sebuah network dengan perintah berikut:  
```
docker network create NAMA_NETWORK
```
![docker-network-ls](gambar-36.png)  
  
Jalankan sebuah container MySQL dan hubungkan container tersebut ke dalam network yang telah dibuat tadi. Berikut perintah lengkapnya:  
```
$ docker run -d \
    --network todo-app --network-alias mysql \
    -v todo-mysql-data:/var/lib/mysql \
    -e MYSQL_ROOT_PASSWORD=secret \
    -e MYSQL_DATABASE=todos \
    mysql:8.0
```
![menjalankan-container](gambar-32.png)  
  
Akses container MySQL menggunakan perintah 
```
docker exec -it <ID Container> mysql -u root -p
``` 
dan gunakan password `secure` untuk login. Kemudian cek isi database menggunakan perintah 
```
SHOW DATABASES;
```  
![mengakses-container](gambar-33.png)  
  
Selanjutnya jalankan container untuk todo app dan ketikkan perintah berikut:  
```
$ docker run -dp 127.0.0.1:3000:3000 \
  -w /app -v "$(pwd):/app" \
  --network todo-app \
  -e MYSQL_HOST=mysql \
  -e MYSQL_USER=root \
  -e MYSQL_PASSWORD=secret \
  -e MYSQL_DB=todos \
  node:18-alpine \
  sh -c "yarn install && yarn run dev"
```
![menjalankan-container-todo app](gambar-34.png)
![menjalankan-container-todo app](gambar-35.png)  