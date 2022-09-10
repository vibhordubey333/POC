osboxes@osboxes:~/Documents/POC/go-redis-postgres$ curl http://localhost:8888/products
{
 "data": [
  {
   "product_id": 1,
   "product_name": "LEATHER BELT",
   "retail_price": 12.25
  },
  {
   "product_id": 2,
   "product_name": "WINTER JACKET",
   "retail_price": 89.65
  },
  {
   "product_id": 3,
   "product_name": "COTTON SOCKS",
   "retail_price": 2.85
  }
 ],
 "source": "Postgresql"
}
osboxes@osboxes:~/Documents/POC/go-redis-postgres$ curl http://localhost:8888/products
{
 "data": [
  {
   "product_id": 1,
   "product_name": "LEATHER BELT",
   "retail_price": 12.25
  },
  {
   "product_id": 2,
   "product_name": "WINTER JACKET",
   "retail_price": 89.65
  },
  {
   "product_id": 3,
   "product_name": "COTTON SOCKS",
   "retail_price": 2.85
  }
 ],
 "source": "RedisCache"
}

