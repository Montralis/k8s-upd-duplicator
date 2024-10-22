# UPD input stream duplicator

This simple Go program reads data from a given port and sends over UDP to specified addresses. It acts as an udp duplicator.


## Usage

The program can be used either locally or as a docker container.

Set the ENV Variables:

`LOG_LEVEL`: Can be `debug`, `info`, `warn`, `error`

`SOURCE_PORT`: Given as `int` for open container port where upd data is reveived
 
`DESTINATION_PORTS`: Using `net.UDPAddr` for destiantion adresses (comma seperated). See example:

``` env
DESTINATION_PORTS: 127.0.2.25:1234,receiver2:1235
```


```go
type UDPAddr struct {
	IP   IP
	Port int
	Zone string // IPv6 scoped addressing zone
}
```

## K8s service dns

`DESTINATION_PORTS` can also refer to Kubernetes services names.

Simply enter the name of the service that is to receive the duplicated data. The ip is resolved via Kubernetes internal dns and no hard-coded ip addresses of the pods need to be entered. This also works with docker compose network type bridge as given in the `docker-compose.yaml` files example.

Beware: In the kubernetes service and in the deyplozment of the pod, the container ports on which the data has been duplicated must be published.
