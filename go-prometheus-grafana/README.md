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


   

