# spectrum_virtualize_exporter

![Go](https://github.com/bluecmd/spectrum_virtualize_exporter/workflows/Go/badge.svg)

Prometheus exporter for IBM Spectrum Virtualize (e.g. IBM Flashsystems).


![SpecV-Prometheus-Exporter](/images/specv-prometheus-exporter.jpg)




## Current Limitation
- Node Exporter won't currently work with SVC Controllers, as lsenclosurestats is not giving out information and lsnodecanister is different for some reason.. 
- No performance on volume and singel hosts yet. 
- Secure traffic, 
- Need to run 8.4.2 code or later. 

----

## Tested on:

 -  FlashSystem 8.4.2.x
 -  Won't work with SVC Controller, as lsenclosurestats is not working and ls

  - Won't work on code below: 8.4.0.x
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


## Future enhancements
- Invidual Volume Performance
- Invidual Host performance. 
- uptime, drive temp, 



## Usage for Node Exporter.

Example

Build the exporter with Go and start it with:

With certificate.

```yaml
./spectrum_virtualize_exporter \
  -auth-file ~/spectrum-monitor.yaml \
  -extra-ca-cert ~/namecheap.ca.crt
```

For insecure.

```yaml
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

The flag `-extra-ca-cert` is useful as it appears that at least V7000 on the
8.2 version is unable to attach an intermediate CA. 
from 8.5 you can attach chained root certificates.

**Users**
- Users on Spectrum Virtualize, minimum access is with Monitor Role. 



## Missing Metrics?

Please [file an issue](https://github.com/bluecmd/spectrum_virtualize_exporter/issues/new) describing what metrics you'd like to see.
Include as much details as possible please, e.g. how the perfect Prometheus metric would look for your use-case.


# Detail installation instructions

## Running the exporter for SpecV

## Option 1: Run on terminal with go installed.

1. Created a directory
2. Create a file `spectrum_virtualize_monitor.yml.`
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

4. Start the exporter pointing to the auth file and with insecure if you dont have the certificates. 

    ```yaml
    ./spectrum_virtualize_exporter \
      -auth-file ~/spectrum_virtualize_monitorr.yml \
      -insecure 
    ```






## Option 2: Running the Node Exporter in a Container



## Option 3: Running prebuilt Node Exporter in a Container




## Running the Prometus service/container

Create a folder to store the **prometheus.yml** file.

> mkdir /srv/prometheus


Default prometheus.yml
```  yaml
/etc/prometheus/prometheus.yml
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


cd into folder
> docker run -d -p 9090:9090 --restart always -v $PWD/prometheus.yml:/etc/prometheus/prometheus.yml --name prometheus prom/prometheus

Optional: to log into the container
>  docker exec -it prometheus sh



After you have started up the container, you should start up browser and point it to https://ip:9090 see something like this, with the spectrum_virtualize 

![Prometheus](/images/specv_exporter_1.png)




To Test that we can see data, we can do following: 

Press the endpoint http://ip:9797/probe and you will get information about probe metrics. 

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

https://grafana.com/api/dashboards/13753/images/9734/image



## Running the Grafana service/container

Create a folder to store data, /srv/opt

> docker run -d --volume "$PWD/data:/var/lib/grafana" -p 3000:3000 grafana/grafana-enterprise


Access the Grafana GUI with https://ip:3000

Then go to DataSources and add new Prometheus Source. 
Point the url to Prometheus instance with the IP/Hostname and the Port 9090
Running http for now..   
http method should be POST

Press save and test.

![Prometheus](/images/specv_exporter_9.png) 



Then Import the SpecV dashboard: 
orginally this was hosted here: Now i have updated it and the copy of the JSON is located in Github
https://grafana.com/grafana/dashboards/13753-v7000/


### ![Prometheus](/images/specv_exporter_5.png) 
<img src="/images/specv_exporter_5.png" width="700" height="500">










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
When testing on 8.2 code on a Storwize 7000
i got error: Probe request rejected; error is: Login code was 404, expected 200
Upgraded to 8.3.1.7 and then it worked, with the -insecure flag, this was with code that not called the version 1/v1 of the api but v0 (default if noy spesificed)
From 8.4.2.0 the restapi have been upgraded, so we need to change the url to contain rest/v1/*


### Some Error that i have experiance: 

**Error: 500**
```sh
Error: 500 - Internal error
```
This was trowned at me on newer Spectrum Virtualize code and SVC, could be that this was beacuse we where using v0 of the restapi interface, and now we are using v1 and it have better error handling.   or it beause of SVC

**Error 409**
``` sh
022/09/19 19:43:52 Error: Response code was 409, expected 200
```
This normally when you send and command or options in the restapi command that don't exist
Example i was trying to run this against a SVC and recived error 409, checked then the commands that would be run agains SpecV, example lsenclosurestats is not available SVC. 

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

