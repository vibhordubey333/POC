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