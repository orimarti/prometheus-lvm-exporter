# LVM pool prometheus exporter
This docker, reads information about LVM thin pools and exports VG's size and VG's free space. Useful to know occupied space of VGs with thin pools and warn with prometheus if needed or metrics.

## Running it
The go binary will listen to port 9080 and serve metrics on the /metrics path, if we run it with docker, we should run the container as privileged.
```
docker run --name=lvm-exporter --privileged=true -p 9080:9080 orimarti/lvm-pool-exporter
```
