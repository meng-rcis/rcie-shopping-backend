# Setup Grafana

1. Create a Grafana container

   ```
   $ docker run --name grafana -d -p 3000:3000 --net=host grafana/grafana

   or

   $ docker run -d -p 3000:3000 --name=grafana -v grafana-storage:/var/lib/grafana grafana/grafana
   ```

2. Verify Grafana by accessing `<HOST-IP>:3000/graph` with user is `admin` and password is `admin` by default
