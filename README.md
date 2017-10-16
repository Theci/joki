# Joki
Joki is a lightweight latency tracker that makes use of fping. It has the ability to ping multiple targets through different QoS classes and links. The data is then exposed so Prometheus can pull it from the /metrics endpoint. The data can also be send to an Influxdb server.

## Configuration
Joki is configured through a config.toml that is located in either /etc/joki, $HOME/.joki or the working directory of the binary. An example config can be found in config.toml.example.

## Systemcd
Joki can be started as a service by copyin the .service file into any of systemd usefull folder (normally: /etc/systemd/system).
