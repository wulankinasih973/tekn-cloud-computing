# Docker Orchestration Hands-on Lab
## Configure Swarm Mode
```
docker run -dt ubuntu sleep infinity
```
![](gambar-01.png)

Cek dengan perintah `docker ps`
```
docker ps
```
![](gambar-02.png)
### Create a Manager node
Pada node1, lakukan inisialisasi swarm
```
docker swarm init --advertise-addr $(hostname -i)
```
![](gambar-03.png)

Cek dengan perintah `docker info`
```
docker info
```
![](gambar-04.png)
![](gambar-05.png)
### Join Worker nodes to the Swarm
Tambahkan node2 dan node3 ke dalam swarm dengan menggunakan perintah `docker swarm join`
![](gambar-06.png)

Cek dengan perintah `docker node ls`
```
docker node ls
```
![](gambar-07.png)
## Deploy applications across multiple hosts
### Deploy the application components as Docker services
Buat service sleep-app pada node1
```
docker service create --name sleep-app ubuntu sleep infinity
```
![](gambar-08.png)

Cek dengan perintah `docker service ls`
```
docker service ls
```
![](gambar-09.png)
## Scale the application
```
docker service update --replicas 7 sleep-app
```
![](gambar-10.png)
```
docker service ps sleep-app
```
![](gambar-11.png)
## Drain a node and reschedule the containers
```
docker node ls
```
![](gambar-12.png)
![](gambar-13.png)
```
docker node update --availability drain <node ID>
```
![](gambar-14.png)
![](gambar-15.png)

Jalankan perintah `docker service ps sleep-app`.
![](gambar-16.png)
## Cleaning Up
```
docker service rm sleep-app
```
Hapus node1, node2, dan node3 dari swarm.
```
docker swarm leave --force
```
![](gambar-17.png)
