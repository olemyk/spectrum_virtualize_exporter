# spectrum_virtualize_exporter

![Go](https://github.com/bluecmd/spectrum_virtualize_exporter/workflows/Go/badge.svg)

Prometheus exporter for IBM Spectrum Virtualize (e.g. IBM FlashSystem).


![SpecV-Prometheus-Exporter](/images/specv-prometheus-exporter.jpg)

## Overview

- Prometheus is the main piece of this setup, it pulls metrics together from Node Exporer and collects them all in one place.

- Node exporter is a small application that queries the Spectrum Virtualize on RestAPI for a variety of metrics and exposes them over HTTP for other services to consume. Prometheus will query one or several Node Exporter instances to aggregate the metrics.

- Grafana is the cherry on top. It takes all the metrics Prometheus has aggregated and displays them as graphs and diagrams organized into dashboards.

![SpecV-Prometheus-Exporter-Grafana](/images/specv_grafana.png)

------

## Current Limitation
- Node Exporter currently dont't work with SVC Controllers, as lsenclosurestats is not giving out information and lsnodecanister is different for some reason.. 
- No performance on volume and singel hosts. (Missing function on SpecV)
- Secure traffic between 
- SpecV 8.4.2 code or later.  (New restapi interface from 8.4.2)

----

## Tested on:

 -  IBM FlashSystem 8.4.2.x

  - Won't work on SpecV code below: 8.4.0.x
    - 8.4.0.x - don't have the new rest api version rest/v1/x  - they you need change api call to be without v1/ in probe.go
    - 8.1.3.x - don't have the new rest api version rest/v1/x  - they you need change api call to be without v1/ in probe.go
    - 8.1.2.x - getting errors from api calls, 404, 500.  to old code to use time on. 

---

# Supported Metrics

 * `spectrum_power_watts`
 * `spectrum_temperature`
 * `spectrum_drive_status`
 * `spectrum_psu_status`
 * `spectrum_pool_capacity_bytes`
 * `spectrum_pool_free_bytes`
 * `spectrum_pool_status`
 * `spectrum_pool_used_bytes`
 * `spectrum_pool_volume_count`
 * `spectrum_node_compression_usage_ratio`
 * `spectrum_node_fc_bps`
 * `spectrum_node_fc_iops`
 * `spectrum_node_iscsi_bps`
 * `spectrum_node_iscsi_iops`
 * `spectrum_node_sas_bps`
 * `spectrum_node_sas_iops`
 * `spectrum_node_system_usage_ratio`
 * `spectrum_node_total_cache_usage_ratio`
 * `spectrum_node_write_cache_usage_ratio`

 * `spectrum_node_mdisk_read_iops`
 * `spectrum_node_mdisk_write_iops`
 * `spectrum_node_mdisk_read_ms`
 * `spectrum_node_mdisk_write_ms`
 * `spectrum_node_mdisk_read_mb`
 * `spectrum_node_mdisk_write_mb`
 
 * `spectrum_node_vdisk_read_iops`
 * `spectrum_node_vdisk_write_iops`
 * `spectrum_node_vdisk_read_ms`
 * `spectrum_node_vdisk_write_ms`
 * `spectrum_node_vdisk_read_mb`
 * `spectrum_node_vdisk_write_mb`

 * `spectrum_fc_port_speed_bps`
 * `spectrum_fc_port_status`
 * `spectrum_ip_port_link_active`
 * `spectrum_ip_port_speed_bps`
 * `spectrum_ip_port_state`

 * `spectrum_quorum_status`
 * `spectrum_host_status`

## Missing Metrics?

Please [file an issue](https://github.com/bluecmd/spectrum_virtualize_exporter/issues/new) describing what metrics you'd like to see.
Include as much details as possible please, e.g. how the perfect Prometheus metric would look for your use-case.

---


## Future enhancements
- IBM SVC support
- Test on IBM SpecV 8.5.x code. 
- Display microseconds instead of milliseconds
- Invidual Volume Performance (When API is available in restapi)
- Invidual Host performance.  (When API is available in restapi)
- Uptime, drive temp, 


-----

## Usage for SpecV Node Exporter.


### **Example**

Build the exporter with Go and start it with:

With certificate.

```sh
./spectrum_virtualize_exporter \
  -auth-file ~/spectrum-monitor.yaml \
  -extra-ca-cert ~/namecheap.ca.crt
```

Without the Certificate. use the flag -insecure

```sh
./spectrum_virtualize_exporter \
  -auth-file ~/spectrum-monitor.yaml \
  -insecure 
```

Where `~/spectrum-monitor.yaml` contains pairs of Spectrum targets
and login information in the following format:

```yaml
"https://my-v7000:7443":
  user: monitor
  password: passw0rd
"https://my-other-v7000:7443":
  user: monitor2
  password: passw0rd1
```

The flag `-extra-ca-cert` is useful as it appears that at least SpecV on the
8.2 version is unable to attach an intermediate CA. 
PS: from 8.5 you can attach chained root certificates.

### **Users**
- Users on IBM Spectrum Virtualize, minimum access is with Monitor Role. 


-----

# **Detailed installation instructions for Exporter, Prometeus and Grafana**

## Running the exporter for IBM Spectrum Virtualize


  ### **Option 1:** Running prebuilt Node Exporter in a Container
  > Image is buildt for `linux/amd64,linux/arm64,linux/ppc64le`
  ---

1. Created a directory `mkdir /srv/spectrum-virtualize-exporter`
2. Create a inventory file `spectrum-monitor.yaml`
3. Where `~/spectrum-monitor.yaml` contains pairs of Spectrum targets
and login information in the following format:

    ```yaml
      "https://my-v7000:7443":
        user: monitor
        password: passw0rd
      "https://my-other-v7000:7443":
        user: monitor2
        password: passw0rd1
    ```
4. `chmod +x spectrum-monitor.yaml`

- Run in background, from the folder you created in step 1
  > docker run --name specv_exporter -it -d --volume $PWD:/config:Z -p 9747:9747/tcp ghcr.io/olemyk/spectrum_virtualize_exporter:latest ./main -auth-file /config/spectrum-monitor.yaml -insecure
- Run interactiv
  > docker run --name specv_exporter -i --volume $PWD:/config:Z -p 9747:9747/tcp ghcr.io/olemyk/spectrum_virtualize_exporter:latest ./main -auth-file /config/spectrum-monitor.yaml -insecure
  
  ***PS: the :Z option in the volume is for selinux
    This will label the content inside the container with the exact MCS label that the container will run with, basically it runs chcon -Rt svirt_sandbox_file_t -l s0:c1,c2 /var/db where s0:c1,c2 differs for each container.***

### **Testing the probe:**
```sh
 % curl -G --data-urlencode "target=https://10.33.7.56:7443" http://localhost:9747/probe
# HELP probe_duration_seconds How many seconds the probe took to complete
# TYPE probe_duration_seconds gauge
probe_duration_seconds 0.785479
# HELP probe_success Whether or not the probe succeeded
# TYPE probe_success gauge
probe_success 1
# HELP spectrum_drive_status Status of drive
# TYPE spectrum_drive_status gauge
spectrum_drive_status{enclosure="1",id="0",slot_id="7",status="degraded"} 0
spectrum_drive_status{enclosure="1",id="0",slot_id="7",status="offline"} 0
spectrum_drive_status{enclosure="1",id="0",slot_id="7",status="online"} 1
spectrum_drive_status{enclosure="1",id="10",slot_id="12",status="degraded"} 0
spectrum_drive_status{enclosure="1",id="10",slot_id="12",status="offline"} 0
spectrum_drive_status{enclosure="1",id="10",slot_id="12",status="online"} 1
spectrum_drive_status{enclosure="1",id="11",slot_id="8",status="degraded"} 0
spectrum_drive_status{enclosure="1",id="11",slot_id="8",status="offline"} 0
spectrum_drive_status{enclosure="1",id="11",slot_id="8",status="online"} 1
```


### Option 2: Run on terminal with go installed.


> **Prerequisites: Go compiler**

1. Created a directory
2. Create a inventory file `spectrum-monitor.yaml`.yml.`
3. Where `~/spectrum-monitor.yaml` contains pairs of Spectrum targets
and login information in the following format:

    ```yaml
      "https://my-v7000:7443":
        user: monitor
        password: passw0rd
      "https://my-other-v7000:7443":
        user: monitor2
        password: passw0rd1
    ```
4. Build the Spectrum-Virtualize-Exporter. 

    ```sh
    export GOPATH=your_gopath
    cd your_gopath
    git clone git@github.com:olemyk/spectrum-virtualize-exporter.git
    cd spectrum-virtualize-exporter
    go build
    go install (Optional but recommended. This step will copy spectrum-virtualize-exporter binary package into $GOPATH/bin directory. It will be connvenient to copy the package to Monitoring docker image)
    ```

4. Start the exporter pointing to the auth file and with insecure if you dont have the certificates. 

    ```yaml
    ./spectrum_virtualize_exporter \
      -auth-file ~/spectrum_virtualize_monitor.yml 
      \ -insecure 
    ```


### Option 3: Build and run the Node Exporter in a Container

1. Build the Spectrum-Virtualize-Exporter. 
  ```sh
  export GOPATH=your_gopath
  cd your_gopath
  git clone git@github.com:olemyk/spectrum-virtualize-exporter.git
  cd spectrum-virtualize-exporter
  go build
  go install (Optional but recommended. This step will copy spectrum-virtualize-exporter binary package into $GOPATH/bin directory. It will be connvenient to copy the package to docker image)
  ```

> docker build -t spectrum_virtualize_exporter .

In the Docker files, there is allready a default command running:
  ``` sh
  `CMD "./main", "-auth-file", "/config/spectrum-monitor.yaml", "-extra-ca-cert", "~/tls.crt"` 
  ```

To override this we need to add the commands in the docker/podman command.

- To run it interactivly, with options. 
  > sudo docker run --name specv_exporter --rm -it --volume $PWD:/config -p 9747:9747/tcp spectrum-virtualize-exporter:latest ./main -auth-file /config/spectrum-monitor.yaml -insecure

- Run in background,
  > sudo docker run --name specv_exporter -it -d --volume $PWD:/config -p 9747:9747/tcp spectrum-virtualize-exporter:latest ./main -auth-file /config/spectrum-monitor.yaml -insecure


2. Created a directory `mkdir /srv/spectrum-virtualize-exporter`
3. Create a inventory file `spectrum-monitor.yaml`
4. Where `~/spectrum-monitor.yaml` contains pairs of Spectrum targets
and login information in the following format:

    ```yaml
      "https://my-v7000:7443":
        user: monitor
        password: passw0rd
      "https://my-other-v7000:7443":
        user: monitor2
        password: passw0rd1
    ```
5. `chmod +x spectrum-monitor.yaml`


  Run container in background, from the folder you created in step 2

  > docker run --name specv_exporter -it -d --volume $PWD:/config:Z -p 9747:9747/tcp spectrum_virtualize_exporter:latest ./main -auth-file /config/spectrum-monitor.yaml -insecure

 Run container interactive
  > docker run --name specv_exporter -i --volume $PWD:/config:Z -p 9747:9747/tcp spectrum_virtualize_exporter:latest ./main -auth-file /config/spectrum-monitor.yaml -insecure
  
  ***PS: the :Z option in the volume is for selinux
    This will label the content inside the container with the exact MCS label that the container will run with, basically it runs chcon -Rt svirt_sandbox_file_t -l s0:c1,c2 /var/db where s0:c1,c2 differs for each container.***



------

## Running the Prometus service/container

Create a folder to store the **prometheus.yml** file.

> mkdir /srv/prometheus


Default prometheus.yml > `/etc/prometheus/prometheus.yml`
```yaml
# my global config
global:
  scrape_interval: 15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.
  # scrape_timeout is set to the global default (10s).

# Alertmanager configuration
alerting:
  alertmanagers:
    - static_configs:
        - targets:
          # - alertmanager:9093

# Load rules once and periodically evaluate them according to the global 'evaluation_interval'.
rule_files:
  # - "first_rules.yml"
  # - "second_rules.yml"

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: "prometheus"

    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.

    static_configs:
      - targets: ["localhost:9090"]
```

Add the following lines, then change the **replacement** ip/hostname and **targets** ip/hostname

- **Replacement:** is Node Exporter ip and Port
- **Targets:** is SpecV Cluster IP.

```yaml
  - job_name: spectrum_virtualize
    metrics_path: /probe
    scheme: http
    relabel_configs:
    - source_labels: [__address__]
      separator: ;
      regex: (.*)
      target_label: __param_target
      replacement: $1
      action: replace
    - source_labels: [__address__]
      separator: ;
      regex: (?:.+)(?::\/\/)([^:]*).*
      target_label: instance
      replacement: $1
      action: replace
    - separator: ;
      regex: (.*)
      target_label: __address__
      replacement: '192.168.1.166:9747'
      action: replace
    static_configs:
    - targets:
      - https://10.33.7.56:7443

```

My Global config file
> /srv/prometheus/prometheus.yml

```yaml
# my global config
global:
  scrape_interval: 15s # Set the scrape interval to every 15 seconds. Default is every 1 minute.
  evaluation_interval: 15s # Evaluate rules every 15 seconds. The default is every 1 minute.
  # scrape_timeout is set to the global default (10s).

# Alertmanager configuration
alerting:
  alertmanagers:
    - static_configs:
        - targets:
          # - alertmanager:9093

# Load rules once and periodically evaluate them according to the global 'evaluation_interval'.
rule_files:
  # - "first_rules.yml"
  # - "second_rules.yml"

# A scrape configuration containing exactly one endpoint to scrape:
# Here it's Prometheus itself.
scrape_configs:
  # The job name is added as a label `job=<job_name>` to any timeseries scraped from this config.
  - job_name: "prometheus"

    # metrics_path defaults to '/metrics'
    # scheme defaults to 'http'.

    static_configs:
      - targets: ["localhost:9090"]

  - job_name: spectrum_virtualize
    metrics_path: /probe
    scheme: http
    relabel_configs:
    - source_labels: [__address__]
      separator: ;
      regex: (.*)
      target_label: __param_target
      replacement: $1
      action: replace
    - source_labels: [__address__]
      separator: ;
      regex: (?:.+)(?::\/\/)([^:]*).*
      target_label: instance
      replacement: $1
      action: replace
    - separator: ;
      regex: (.*)
      target_label: __address__
      replacement: 'localhost:9747'
      action: replace
    static_configs:
    - targets:
      - https://10.33.7.56:7443
```


 cd into folder you created and store the config and run: 
> docker run -d -p 9090:9090 --restart always -v $PWD/prometheus.yml:/etc/prometheus/prometheus.yml:Z --name prometheus docker.io/prom/prometheus
- Option Z is for SELINUX

Optional: to log into the container
> docker exec -it prometheus sh



After you have started up the container, you should start up browser and point it to https://ip:9090 see something like this, with the spectrum_virtualize 

![Prometheus](/images/specv_exporter_1.png)


To Test that we can see data, we can do following: 

Press the endpoint http://ip:9747/probe and you will get information about probe metrics. 

![Prometheus](/images/specv_exporter_2.png)


Then use one of the Probes strings to test,  press Graph and type in for example **spectrum_power_watts**

![Prometheus](/images/specv_exporter_7.png)

You can allso run curl to get values:

> curl -G --data-urlencode "target=https://specvip:7443" http://specv-exporter:9747/probe/


```sh
 % curl -G --data-urlencode "target=https://10.33.7.56:7443" http://localhost:9747/probe
# HELP probe_duration_seconds How many seconds the probe took to complete
# TYPE probe_duration_seconds gauge
probe_duration_seconds 0.785479
# HELP probe_success Whether or not the probe succeeded
# TYPE probe_success gauge
probe_success 1
# HELP spectrum_drive_status Status of drive
# TYPE spectrum_drive_status gauge
spectrum_drive_status{enclosure="1",id="0",slot_id="7",status="degraded"} 0
spectrum_drive_status{enclosure="1",id="0",slot_id="7",status="offline"} 0
spectrum_drive_status{enclosure="1",id="0",slot_id="7",status="online"} 1
spectrum_drive_status{enclosure="1",id="10",slot_id="12",status="degraded"} 0
spectrum_drive_status{enclosure="1",id="10",slot_id="12",status="offline"} 0
spectrum_drive_status{enclosure="1",id="10",slot_id="12",status="online"} 1
spectrum_drive_status{enclosure="1",id="11",slot_id="8",status="degraded"} 0
spectrum_drive_status{enclosure="1",id="11",slot_id="8",status="offline"} 0
spectrum_drive_status{enclosure="1",id="11",slot_id="8",status="online"} 1
```

## Running the Grafana service/container

Create a folder to store data, /srv/opt

> mkdir -p /srv/grafana/ 

> cd into /srv/grafana/

> docker run -d --name prometheus --volume "$PWD/data:/var/lib/grafana" -p 3000:3000 docker.io/grafana/grafana-enterprise

For ppc64le
I have not found any offical grafana for ppc64le (IBM Power) as a container. 

If you search for Grafana and IBM Power, and you will find it as RPM you can install.

https://www.power-devops.com/grafana

----
Access the Grafana GUI with https://ip:3000

- Then go to DataSources and add new Prometheus Source. 
  - The Name should be **prometheus_specv** or else you need to change the grafana   datasource in the json.
  - Point the url to Prometheus instance with the IP/Hostname and the Port 9090
    - Running http for now..   
- http method: POST
- Press Save & test and check that you get, Data source is working

![Prometheus](/images/specv_exporter_11.png) 



## Then Import the SpecV dashboard: 
- Orginally this was hosted here: https://grafana.com/grafana/dashboards/13753-v7000/

 - Now i have updated it and the copy of the JSON is located in the grafana_dekstop folder in the github repo. (will maybe create a Grafana id)
 - Access the Grafana GUI with https://ip:3000
    - Go to Dashboard and press import
    - Copy the content from `grafana_desktop_specv_rev_x.json` in the github repo. to the Import via panel json. 

    <img src="/images/specv_exporter_12.png" width="700" height="500">

  - Now Change the Name of the dashboard if you want, the UID should now be **promethes_specv** or the name you called the the Prometheus Datasource. 

    <img src="/images/specv_exporter_13.png" width="700" height="500">






------


## RestAPI Calls against Spectrum Virtualize - from 8.4.2.x

First we need to authenticate, so call /auth with your user. 
> % curl -k -X POST -H 'Content-Type:application/json' -H 'X-Auth-Username: superuser' -H 'X-Auth-Password: Passw0rd' https://10.10.10.182:7443/rest/v1/auth


Copy then the token to your next command:

```
{"token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpYXQiOjE2NjM5MzY0MjQsImV4cCI6MTY2Mzk0MDAyNCwianRpIjoiMjM0NTdhMGZlMTU1ZjA4NTRkOGQ4MzE0NTI5ZmQyOWUiLCJzdiI6eyJ1c2VyIjoic3VwZXJ1c2VyIn19.Fz1jqDdET-UDVazgyRFSvyV8zKRMu1Kmge0dbQxWmks-iGNOYCWNCrWcRNMGW5M-F5cnfWBq4vQ8QsteSv2WxA"}%
```

Replace the token and IP in the command. 
>curl -k -X POST -H 'Content-Type:application/json' -H 'X-Auth-Token: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpYXQiOjE2NjM5MzY0MjQsImV4cCI6MTY2Mzk0MDAyNCwianRpIjoiMjM0NTdhMGZlMTU1ZjA4NTRkOGQ4MzE0NTI5ZmQyOWUiLCJzdiI6eyJ1c2VyIjoic3VwZXJ1c2VyIn19.Fz1jqDdET-UDVazgyRFSvyV8zKRMu1Kmge0dbQxWmks-iGNOYCWNCrWcRNMGW5M-F5cnfWBq4vQ8QsteSv2WxA' https://10.32.64.182:7443/rest/v1/lsnodecanisterstats

> curl -k -X POST -H 'Content-Type:application/json' -H 'X-Auth-Token: ' https://10.33.7.56:7443/rest/v1/lsnodecanisterstats


To access the REST API Explorer, enter the following URL in a browser:

> https://<SpecV_ip_address | FQDN>:7443/rest/explorer
After login in you will be greeted with some example that is using CURL

------

## Troubleshooting 

If you are getting errors from API:

- 200: Ok
- 400: BadRequest
- 401: Unauthorized
- 403: Forbidden
- 404: NotFound
- 409: Conflict
- 429: TooManyRequests
- 500: InternalServerError
- 502: BadGateway


Errors:
When testing on 8.2 code on a SpecV. i got error: Probe request rejected; error is: Login code was 404, expected 200
Upgraded then to 8.3.1.7 and then it worked, with the -insecure flag, this was with code that's "called" the version /v0 of the API (default if not specified)
From 8.4.2.0 the SpecV restapi interface have been upgraded, so we need to change the url to contain rest/v1/*


### Some error codes that i have experiance: 

**Error: 500**
```sh
Error: 500 - Internal error
```
This was trowned at me on newer Spectrum Virtualize code and SVC, could be that this was beacuse we where using v0 of the restapi interface, and now we are using v1 and it have better error handling...

**Error 409**
```sh
022/09/19 19:43:52 Error: Response code was 409, expected 200
```
This normally when you send and command or options in the restapi command that don't exist
Example i was trying to run this against a SVC and received error 409, checked and then run then same command with CLI against SpecV, example lsenclosurestats is not available SVC. 

Se example below: 

```sh
% curl -k -X POST -H 'Content-Type:pplication/json' -H 'X-Auth-Token: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpYXQiOjE2NjM2MDk0NTAsImV4cCI6MTY2MzYxMzA1MCwianRpIjoiNjIwNmFmN2NlMjFhMzQ2MjZlYTUxNGVmODIwY2RkYzYiLCJzdiI6eyJ1c2VyIjoic3VwZXJ1c2VyIn19.y-K63TcAE8JlgAPuHqlER0ZMb8YkSrybYwmwSGRMLeXHl59XenV7NCdJ4mxhvPYXDwsw5jkky-fNzmfNIfgtUg' https://10.33.7.56:7443/rest/v1/lsenclosurestats -v
*
> POST /rest/v1/lsenclosurestats HTTP/2
> Host: 10.33.7.56:7443
> user-agent: curl/7.79.1
> accept: */*
> content-type:pplication/json
> x-auth-token: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzUxMiJ9.eyJpYXQiOjE2NjM2MDk0NTAsImV4cCI6MTY2MzYxMzA1MCwianRpIjoiNjIwNmFmN2NlMjFhMzQ2MjZlYTUxNGVmODIwY2RkYzYiLCJzdiI6eyJ1c2VyIjoic3VwZXJ1c2VyIn19.y-K63TcAE8JlgAPuHqlER0ZMb8YkSrybYwmwSGRMLeXHl59XenV7NCdJ4mxhvPYXDwsw5jkky-fNzmfNIfgtUg
>
* Connection state changed (MAX_CONCURRENT_STREAMS == 128)!
< HTTP/2 409
< server: nginx
< date: Mon, 19 Sep 2022 17:52:45 GMT
< content-type: application/json; charset=UTF-8
< content-length: 75
< strict-transport-security: max-age=31536000; includeSubDomains
<
* Connection #0 to host 10.33.7.56 left intact
"error code: 1, error text: CMMVC6051E An unsupported action was selected."%
```

