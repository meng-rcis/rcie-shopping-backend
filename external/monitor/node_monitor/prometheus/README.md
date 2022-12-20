# Setup Prometheus

1. Create a user for Prometheus on your system

   ```
    $ useradd -rs /bin/false prometheus
   ```

2. Create a new folder and a new configuration file for Prometheus
   You can see example prometheus.yml in this current directory

   ```
    $ mkdir /etc/prometheus
    $ touch /etc/prometheus/prometheus.yml
   ```

3. Create a data folder for Prometheus

   ```
    $ mkdir -p /data/prometheus
    $ chown prometheus:prometheus /data/prometheus /etc/prometheus/*
   ```

4. Edit the /etc/prometheus/prometheus.yml by adding its configuration (you can copy from ./prometheus.yml)

5. Get the userid for running Prometheus

   ```
    $ cat /etc/passwd | grep prometheus
   ```

   example return: `prometheus:x:997:997::/home/prometheus:/bin/false`

   you need to remember its userid (in this case above, 997:997)

6. Create the Prometheus container

   ```
   $ docker run --name myprom -d -p 9090:9090 --user {user-id}:{user-id}  --net=host -v /etc/prometheus:/etc/prometheus -v /data/prometheus:/data/prometheus prom/prometheus --config.file="/etc/prometheus/prometheus.yml" --storage.tsdb.path="/data/prometheus"
   ```

   example: docker run --name myprom -d -p 9090:9090 --user 997:997 --net=host -v /etc/prometheus:/etc/prometheus -v /data/prometheus:/data/prometheus prom/prometheus --config.file="/etc/prometheus/prometheus.yml" --storage.tsdb.path="/data/prometheus"

7. Check its logs

   ```
    $ docker logs myprom
   ```

8. Verify Prometheus by accessing `<HOST-IP>:9090/graph`

9. To restart Prometheus, get the PID & send SIGHUP signal

   ```
   $ ps aux | grep prometheus
   $ kill -HUP <PID>
   ```

### NOTE

cat prometheus configuration file inside container

`docker exec <container-id> cat /etc/prometheus/prometheus.yml`
