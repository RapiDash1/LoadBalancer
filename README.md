# LoadBalancer
This is a barebones load balancer implementation in Go.

What is a load balancer? 
> When multiple servers are joined to create a cluster. Clusters can use network load balancing whereby simultaneous cluster request are distributed between cluster servers. \
For more info, check out this wikipedia article -  
    >> https://en.wikipedia.org/wiki/Network_Load_Balancing#Server_load_balancing
---

## Getting Started

### Prerequisites
* Clone this repo to your local machine.
* Install Go if you havent.
* Install the following modules - \
    >```go get -u github.com/go-co-op/gocron```

### Running the app
* Create 5 python servers -
   > ```
    >python server.py "8000"
    >python server.py "8001"
    >python server.py "8002"
    >python server.py "8003"
    >python server.py "8004" 
    >```
* Run loadbalancer

    > ```go run .```
---
## Author

* RapiDash1