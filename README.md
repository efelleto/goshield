# goshield: real-time log monitoring and intrusion detection engine

goshield is a high-performance security utility developed in go, designed to monitor web server access logs and identify malicious patterns through stateful analysis. it bridges the gap between passive logging and active threat detection.

## core architecture

the engine utilizes a non-blocking log-tailing strategy to monitor traffic with near-zero latency. it employs regular expressions to deconstruct nginx/apache log entries into actionable data points, enabling immediate response to reconnaissance activities.



## technical specifications

- **language:** go 1.21+
- **concurrency model:** event-driven log streaming via channel-based tailing.
- **state management:** in-memory threat scoring using thread-safe map structures.
- **pattern matching:** optimized regex engine for ipv4 and uri extraction.
- **terminal ui:** 24-bit color depth (truecolor) with information hierarchy (high-intensity status vs low-contrast metadata).

## detection logic

the system identifies threats based on directory traversal attempts and sensitive file discovery patterns. 

1. **ingestion:** parses raw log lines into ip and path variables.
2. **analysis:** checks for forbidden keywords (e.g., .env, admin, config).
3. **escalation:** increments an offense counter per source ip.
4. **mitigation:** triggers an alert at offense #1 and simulates a firewall block at offense #5.



## installation and usage

### environment setup
deploy the web server environment using docker:

```bash
docker-compose up -d
```

## build & run
compile the engine for your local environment:
```bash
go build -o goshield.exe main.go
./goshield.exe
```

## attack simulation
validate the detection engine by running the reconnaissance simulator:
```bash
go run attacker.go
```

## project structure
- `main.go`: the core detection engine and log parser.
- `attacker.go`: automated script for simulating malicious web traffic.
- `docker-compose.yml`: containerized web server environment.
- `logs/access.log`: shared volume for real-time log ingestion.

## roadmap
- [x] real-time log ingestion and regex parsing
- [x] stateful threat scoring system
- [x] interactive terminal ui with truecolor support
- [ ] implementation of native iptables/nftables integration
- [ ] persistent database layer for historical analysis
- [ ] webhook support for soc orchestration (siem/soar)

## demo

https://github.com/user-attachments/assets/0b405bd8-b8b3-4054-81ff-b23a10b7b51f



## license

distributed under the mit license. see `license` for more information.

## contact

efelleto - [github.com/efelleto](https://github.com/efelleto)

project link: [https://github.com/efelleto/goshield](https://github.com/efelleto/goshield)
