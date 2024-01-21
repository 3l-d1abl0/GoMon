# GoMon

GoMon is a monitoring service that exposes metrices of the System like network , CPU ,memory etc.

Can be used by external scrapper to gather System information.

## Installation

#### 1. Clone the repo
    git clone https://github.com/3l-d1abl0/GoMon.git

#### 2. Create the env file
    mkdir GoMon/envs
    touch GoMon/envs/.env

#### 3. Create .env files entries (example)
    SERVERPORT=3000
    DEBUG=false
    GIN_MODE=debug
#### 4. Create the binary
    cd GoMon
    go build -o GoMon
    ./GoMon

#### 5. Access the api at port specified in .env
    eg: curl http://localhost:3000/api/v1


The following end points are exposed :

 * GET /api/v1
 * GET /api/v1/resource/
 * GET /api/v1/resource/memory
 * GET /api/v1/resource/network
 * GET /api/v1/resource/cpu
 * GET /api/v1/resource/host
 * GET /api/v1/resource/load
 * GET /api/v1/resource/all


Built with: [Gin Gonic](https://github.com/gin-gonic "Gin-Gonic") , [gopsutil](https://github.com/shirou/gopsutil/ "gopsutil")
