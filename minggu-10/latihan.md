# Docker Networking
## Section #1 - Networking Basics
### Step 1: The Docker Network Command
Perintah `docker network` adalah perintah utama untuk mengonfigurasi dan mengelola jaringan kontainer. Jalankan perintah `docker network` dari terminal pertama.
```
docker network
```
![docker-network](gambar-01.png)

### Step 2: List networks
Untuk melihat daftar yang ada pada Docker Host, jalankan perintah `docker network ls`.
```
docker network ls
```
![docker-network-ls](gambar-02.png)

### Step 3: Inspect a network
Untuk melihat konfigurasi network dengan nama bridge, jalankan perintah berikut:
```
docker network inspect bridge
```
![docker-network-inspect-bridge](gambar-03.png)

### Step 4: List network driver plugins
Jalankan perintah `docker info` untuk melihat informasi tentang Installasi Docker.
```
docker info
```
![docker-info](gambar-04.png)

## Section #2 - Bridge Networking
### Step 1: The Basics
Kita dapat melihat daftar network yang ada pada Docker dengan perintah `docker network ls`.
Install command `brctl`.
```
apk update
apk add bridge
```
![apk-update](gambar-05.png)
![apk-add-bridge](gambar-06.png)

List bridge pada Docker Host dengan menjalankan perintah berikut:
```
brctl-show
```
![brctl-show](gambar-07.png)

### Step 2: Connect a container
Jalankan container baru dengan perintah berikut:
```
docker run -dt ubuntu sleep infinity
```
![docker-run](gambar-08.png)

Jalankan kembali perintah `brctl show`.
![brctl-show](gambar-09.png)

Jalankan perintah `docker network inspect bridge` untuk melihat konfigurasi network.
![docker-network-inspect-bridge](gambar-10.png)

### Step 3: Test network connectivity
Lakukan ping dari Docker host ke container.
![ping](gambar-11.png)

Untuk melihat ID container, jalankan perintah berikut:
```
docker ps
```
![docker-ps](gambar-12.png)

Jalankan perintah `docker exec -it <CONTAINER ID> /bin/bash`
![docker-exec-it](gambar-13.png)

Install ping program:
```
apt-get update && apt-get install -y iputils-ping
```
![install-ping-program](gambar-14.png)

Lakukan ping www.github.com dengan running `ping -c5 www.github.com`
![ping](gambar-15.png)

### Step 4: Configure NAT for external connectivity
Jalankan container baru dengan image NGINX
```
docker run --name web1 -d -p 8080:80 nginx
```
![docker-run](gambar-16.png)
![](gambar-17.png)

## Section #3 - Overlay Networking
### Step 1: The Basics
Inisialisasi swarm, join a single worker node, dan verify the operations worked dengan perintah berikut:
```
docker swarm init --advertise-addr $(hostname -i)
```
![docker-swarm-init](gambar-18.png)

Lakukan join pada node ke-2:
![docker-swarm-join](gambar-19.png)
![docker-ls](gambar-20.png)

### Step 2: Create an overlay network
Jalankan perintah berikut pada node ke-1:
![docker-network-create](gambar-21.png)
![docker-ls](gambar-22.png)

### Step 3: Create a service
Jalankan perintah berikut pada terminal pertama untuk membuat service baru
```
docker service create --name myservice \
--network overnet \
--replicas 2 \
ubuntu sleep infinity
```
![docker-service-create](gambar-23.png)

Cek konfigurasi network
![docker-ls](gambar-24.png)

### Step 4: Test the network
Install ping command dengan perintah berikut:
```
apt-get update && apt-get install -y iputils-ping
```
![apt-get update](gambar-25.png)
![apt-get install iputils-ping](gambar-26.png)

Kemudian lakukan ping
![ping](gambar-27.png)

### Step 5: Test service discovery
Masih di dalam kontainer sebelumnya, jalankan perintah berikut:
![](gambar-28.png)

Lakukan ping
![ping](gambar-29.png)

Inspect konfigurasi myservice:
![](gambar-30.png)

## Cleaning Up
Jalankan perintah `docker service rm myservice` untuk menghapus service *myservice*
```
docker service rm myservice
```
![remove-myservice](gambar-31.png)

Menghapus node 1 dan node 2 dari Swarm. Jalankan pada node 1 dan node 2 perintah berikut ini:
![](gambar-32.png)
![](gambar-33.png)