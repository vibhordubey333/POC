### API Request

1. Get default  metrics exposed `curl localhost:8081/metrics`
2. Get connected devices `curl -i localhost:8081/devices`
```
[
  {
    "id": 0,
    "mac": "34-34-3422-2243-43",
    "firmware": "1.0"
  },
  {
    "id": 1,
    "mac": "34-34-3422-2243-44",
    "firmware": "2.0"
  }
]
```
3. Accessing Prometheus dashboard.<br/>
![image](https://github.com/vibhordubey333/POC/assets/22407855/253c82ea-5192-4024-86eb-fcfbc0753dd9)

4. Accessing `POC1_device_connected_list` metrics
   ![image](https://github.com/vibhordubey333/POC/assets/22407855/6cc8fa57-359b-4625-95c6-d2949993e40e)

5. Visualizing graph
   ![image](https://github.com/vibhordubey333/POC/assets/22407855/ff04cc80-b4af-469f-bfb4-c92b0ddbb89d)

6. Connecting with Grafana. `http://localhost:3000` Cred's `admin/devops123`
   - Click on "Add your first data source"
   ![image](https://github.com/vibhordubey333/POC/assets/22407855/1fd1f898-23d5-45b2-b9a7-bdd8e4f8b20d)

   - Configure <br/>
     If some error occurs while connecting use it as reference: https://stackoverflow.com/questions/74029504/spring-prometheus-grafana-err-reading-prometheus-post-http-localhost90 <br/>
     ![image](https://github.com/vibhordubey333/POC/assets/22407855/d5701d0b-a0d5-4b98-8116-0d1bd2a24867)

7. Configuring our Dashboard
   - Click on "Create your first dashboard"
   ![image](https://github.com/vibhordubey333/POC/assets/22407855/8b7fd1eb-c092-4590-974d-594ad6a0b142)
   -  Add new Panel
   ![image](https://github.com/vibhordubey333/POC/assets/22407855/d0de8f77-1631-4f5d-b803-1fa5cec358f4)
   - Select `Metric` and `Filters`
     ![image](https://github.com/vibhordubey333/POC/assets/22407855/cad3e1f7-78d6-405c-ad85-b341725c8ccc)
   - Click on "Run queries"
     ![image](https://github.com/vibhordubey333/POC/assets/22407855/2c6bd1b8-5f28-4f44-8f40-913236d42865)

  ---
  
  1. Counter:

The next metric type that we're going to implement is a counter. It is a cumulative metric that represents a single monotonically increasing counter. It can only go up and be reset to zero on restart. Typically you would use it with the rate function and measure the number of requests served, tasks completed, or errors. We're going to use it to count device upgrades.

Let's declare it as a CounterVec to add custom labels. We'll use a label to count upgrades of different device types. For example, the type can be a router, access point, modem, etc.
  ---
  References:
  - https://medium.com/@alcbotta/monitoring-you-golang-server-with-prometheus-and-grafana-97e64bb1d0e9
  - https://antonputra.com/monitoring/monitor-golang-with-prometheus/#gauge


 


   

