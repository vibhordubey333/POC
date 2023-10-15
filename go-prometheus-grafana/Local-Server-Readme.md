#### Installing Grafana

1. `docker run -d --name=grafana -p 3000:3000 grafana/grafana`
2. `docker run -d --name=prometheus -p 9090:9090 prom/prometheus`
2. `docker ps`
3. Verify by visiting `http://localhost:3000`. Default username/password `admin/admin`.
4. Hitting Metrics API `curl localhost:8080/metrics` <br/>
   
  ```
  # HELP API_Request_User_Count User Count
  # TYPE API_Request_User_Count counter
  API_Request_User_Count{status="501",user="Alpha"} 2
  API_Request_User_Count{status="501",user="Foxtrot"} 3
  # HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
  # TYPE go_gc_duration_seconds summary
  go_gc_duration_seconds{quantile="0"} 0
  go_gc_duration_seconds{quantile="0.25"} 0
  go_gc_duration_seconds{quantile="0.5"} 0
  go_gc_duration_seconds{quantile="0.75"} 0
  go_gc_duration_seconds{quantile="1"} 0
  go_gc_duration_seconds_sum 0
  go_gc_duration_seconds_count 0
  ```
