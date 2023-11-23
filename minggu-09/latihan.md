# Praktikum Minggu-09
## Docker for Beginners - Linux
### Task 0: Prerequisites
Clone repositori berikut:
```
git clone https://github.com/dockersamples/linux_tweet_app
```
![clone-repo](gambar-01.png)
### Task 1: Run some simple Docker containers
#### Run a single task in an Alpine Linux container
Jalankan perintah di bawah ini:
```
docker container run alpine hostname
```
![single-task](gambar-02.png)
Anda dapat melihat list semua container dengan perintah berikut:
```
docker container ls --all
```
![list-container](gambar-03.png)
#### Run an interactive Ubuntu container
Untuk menjalankan container secara interaktif, jalankan perintah berikut ini:
```
docker container run --interactive --tty --rm ubuntu bash
``` 
![run-interactive](gambar-04.png)
Di dalam shell, kita dapat mengeksekusi perintah-perintah seperti ls, ps, dan cat
![ls-ps-cat](gambar-05.png)
Eksekusi perintah `exit` untuk keluar dari shell.
#### Run a background MySQL container
Jalankan perintah di bawah ini:
```
docker container run \
 --detach \
 --name mydb \
 -e MYSQL_ROOT_PASSWORD=my-secret-pw \
 mysql:latest
```
![mysql-container](gambar-06.png)
![mysql-container](gambar-07.png)
Anda dapat mengecek container yang sedang berjalan dengan perintah berikut:
```
docker container ls
```
![checking-running-container](gambar-08.png)
Untuk menjalankan perintah di dalam container yang sedang berjalan, gunakan perintah berikut:
```
docker exec -it mydb \
 mysql --user=root --password=$MYSQL_ROOT_PASSWORD --version
```
![docker-exec](gambar-09.png)
### Task 2: Package and run a custom app using Docker
Masuk ke dalam direktori `linux_tweet_app`
![cd-linux-tweet-app](gambar-10.png)
Cek isi file Dockerfile dengan perintah:
```
cat Dockerfile
```
![cd-linux-tweet-app](gambar-11.png)
Export DOCKERID:
```
export DOCKERID=<NAMA DOCKER ID ANDA>
```
![echo-dockerid](gambar-12.png)
Gunakan perintah `docker image build` untuk membuat image Docker baru menggunakan instruksi di Dockerfile. `--tag` untuk memberikan nama khusus pada image.
```
docker image build --tag $DOCKERID/linux_tweet_app:1.0 .
```
![docker-image-buid](gambar-13.png)
![docker-image-buid](gambar-14.png)
Buat container baru menggunakan image yang sudah dibuat sebelumnya.
```
docker container run \
 --detach \
 --publish 80:80 \
 --name linux_tweet_app \
 $DOCKERID/linux_tweet_app:1.0
```
![docker-container-run](gambar-15.png)
Berikut adalah tampilan yang dihasilkan:
![hasil-tampilan](gambar-16.png)
Hapus container dengan perintah berikut:
```
docker container rm --force linux_tweet_app
```
![docker-container-rm](gambar-17.png)
### Task 3: Modify a running website
Jalankan perintah berikut:
```
docker container run \
 --detach \
 --publish 80:80 \
 --name linux_tweet_app \
 --mount type=bind,source="$(pwd)",target=/usr/share/nginx/html \
 $DOCKERID/linux_tweet_app:1.0
```
![docker-container-run](gambar-18.png)
Ganti `index.html` dengan `index-new.html`
```
cp index-new.html index.html
```
Berikut perubahan tampilan yang dihasilkan:
![hasil-tampilan](gambar-19.png)
Hapus container dengan perintah berikut ini:
```
docker rm --force linux_tweet_app
```
Jalankan container baru dengan perintah yang sama
```
docker container run \
 --detach \
 --publish 80:80 \
 --name linux_tweet_app \
 $DOCKERID/linux_tweet_app:1.0
```
![docker-container-run](gambar-20.png)
Tampilan halaman akan sama seperti sebelum dimodifikasi
![hasil-tampilan](gambar-21.png)
Hapus kembali container dengan perintah berikut:
```
docker rm --force linux_tweet_app
```
Lakukan update image dengan perintah berikut:
```
docker image build --tag $DOCKERID/linux_tweet_app:2.0 .
```
![docker-image-build](gambar-22.png)
Cek dengan perintah `ls`
```
docker image ls
```
![docker-image-ls](gambar-23.png)
Jalankan container menggunakan image yang telah diupdate:
```
docker container run \
 --detach \
 --publish 80:80 \
 --name linux_tweet_app \
 $DOCKERID/linux_tweet_app:2.0
```
![docker-container-run](gambar-24.png)
![docker-container-run](gambar-25.png)
Jalankan juga container menggunakan image yang belum diupdate:
```
docker container run \
 --detach \
 --publish 8080:80 \
 --name old_linux_tweet_app \
 $DOCKERID/linux_tweet_app:1.0
```
![docker-container-run](gambar-26.png)
![docker-container-run](gambar-27.png)
Push image ke DockerHub. Login dengan menggunakan perintah `docker login`
![docker-container-run](gambar-28.png)
Push image versi pertama:
![docker-image-push](gambar-29.png)
```
docker image push $DOCKERID/linux_tweet_app:1.0
```
Push image versi kedua:
```
docker image push $DOCKERID/linux_tweet_app:2.0
```
![docker-image-push](gambar-30.png)
Cek pada repo DockerHub anda:
![cek-dockerhub](gambar-31.png)
![cek-dockerhub](gambar-32.png)