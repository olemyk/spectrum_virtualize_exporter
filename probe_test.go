// Tests of spectrum_virtualize_exporter
//
// Copyright (C) 2020  Christian Svensson
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/google/go-jsonnet"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

type fakeClient struct {
	data map[string][]byte
}

func (c *fakeClient) prepare(path string, jfile string) {
	vm := jsonnet.MakeVM()
	b, err := os.ReadFile(jfile)
	if err != nil {
		log.Fatalf("Failed to read jsonnet %q: %v", jfile, err)
	}
	output, err := vm.EvaluateAnonymousSnippet(jfile, string(b))
	if err != nil {
		log.Fatalf("Failed to evaluate jsonnet %q: %v", jfile, err)
	}
	c.data[path] = []byte(output)
}

func (c *fakeClient) Get(path string, query string, obj interface{}) error {
	d, ok := c.data[path]
	if !ok {
		log.Fatalf("Tried to get unprepared URL %q", path)
	}
	return json.Unmarshal(d, obj)
}

func newFakeClient() *fakeClient {
	return &fakeClient{data: map[string][]byte{}}
}

func TestEnclosureStats(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsenclosurestats", "testdata/specv_output_8_5_2_2/lsenclosurestats.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeEnclosureStats(c, r) {
		t.Errorf("probeEnclosureStats() returned non-success")
	}

	em := `
	# HELP spectrum_power_watts Current power draw of enclosure in watts
	# TYPE spectrum_power_watts gauge
	spectrum_power_watts{enclosure="1"} 480
	# HELP spectrum_temperature Current enclosure temperature in celsius
	# TYPE spectrum_temperature gauge
	spectrum_temperature{enclosure="1"} 22
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
func TestDrive(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsdrive", "testdata/specv_output_8_5_2_2/lsdrive.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeDrives(c, r) {
		t.Errorf("probeDrives() returned non-success")
	}

	em := `
	# HELP spectrum_drive_status Status of drive
	# TYPE spectrum_drive_status gauge
	spectrum_drive_status{enclosure="1",id="0",slot_id="7",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="0",slot_id="7",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="0",slot_id="7",status="online"} 1
	spectrum_drive_status{enclosure="1",id="1",slot_id="6",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="1",slot_id="6",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="1",slot_id="6",status="online"} 1
	spectrum_drive_status{enclosure="1",id="10",slot_id="11",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="10",slot_id="11",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="10",slot_id="11",status="online"} 1
	spectrum_drive_status{enclosure="1",id="11",slot_id="12",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="11",slot_id="12",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="11",slot_id="12",status="online"} 1
	spectrum_drive_status{enclosure="1",id="2",slot_id="2",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="2",slot_id="2",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="2",slot_id="2",status="online"} 1
	spectrum_drive_status{enclosure="1",id="3",slot_id="8",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="3",slot_id="8",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="3",slot_id="8",status="online"} 1
	spectrum_drive_status{enclosure="1",id="4",slot_id="9",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="4",slot_id="9",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="4",slot_id="9",status="online"} 1
	spectrum_drive_status{enclosure="1",id="5",slot_id="3",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="5",slot_id="3",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="5",slot_id="3",status="online"} 1
	spectrum_drive_status{enclosure="1",id="6",slot_id="5",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="6",slot_id="5",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="6",slot_id="5",status="online"} 1
	spectrum_drive_status{enclosure="1",id="7",slot_id="4",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="7",slot_id="4",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="7",slot_id="4",status="online"} 1
	spectrum_drive_status{enclosure="1",id="8",slot_id="1",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="8",slot_id="1",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="8",slot_id="1",status="online"} 1
	spectrum_drive_status{enclosure="1",id="9",slot_id="10",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="9",slot_id="10",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="9",slot_id="10",status="online"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestEnclosurePSU(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsenclosurepsu", "testdata/specv_output_8_5_2_2/lsenclosurepsu.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeEnclosurePSUs(c, r) {
		t.Errorf("probeEnclosurePSUs() returned non-success")
	}

	em := `
	# HELP spectrum_psu_status Status of PSU
	# TYPE spectrum_psu_status gauge
	spectrum_psu_status{enclosure="1",id="1",status="degraded"} 0
	spectrum_psu_status{enclosure="1",id="1",status="offline"} 0
	spectrum_psu_status{enclosure="1",id="1",status="online"} 1
	spectrum_psu_status{enclosure="1",id="2",status="degraded"} 0
	spectrum_psu_status{enclosure="1",id="2",status="offline"} 0
	spectrum_psu_status{enclosure="1",id="2",status="online"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestPool(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsmdiskgrp", "testdata/specv_output_8_5_2_2/lsmdiskgrp.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probePool(c, r) {
		t.Errorf("probePool() returned non-success")
	}

	em := `
	# HELP spectrum_pool_capacity_bytes Capacity of pool in bytes
	# TYPE spectrum_pool_capacity_bytes gauge
	spectrum_pool_capacity_bytes{id="0",name="DRP_Pool0"} 1.96702630209126e+14
	spectrum_pool_capacity_bytes{id="1",name="Backup_0"} 1.96702630209126e+14
	spectrum_pool_capacity_bytes{id="2",name="Application"} 1.96702630209126e+14
	spectrum_pool_capacity_bytes{id="3",name="DRP_CPool_3"} 1.96702630209126e+14
	# HELP spectrum_pool_free_bytes Free bytes in pool
	# TYPE spectrum_pool_free_bytes gauge
	spectrum_pool_free_bytes{id="0",name="DRP_Pool0"} 1.92139656953856e+14
	spectrum_pool_free_bytes{id="1",name="Backup_0"} 1.92139656953856e+14
	spectrum_pool_free_bytes{id="2",name="Application"} 1.92139656953856e+14
	spectrum_pool_free_bytes{id="3",name="DRP_CPool_3"} 1.92139656953856e+14
	# HELP spectrum_pool_status Status of pool
	# TYPE spectrum_pool_status gauge
	spectrum_pool_status{id="0",name="DRP_Pool0",status="offline"} 0
	spectrum_pool_status{id="0",name="DRP_Pool0",status="online"} 1
	spectrum_pool_status{id="1",name="Backup_0",status="offline"} 0
	spectrum_pool_status{id="1",name="Backup_0",status="online"} 1
	spectrum_pool_status{id="2",name="Application",status="offline"} 0
	spectrum_pool_status{id="2",name="Application",status="online"} 1
	spectrum_pool_status{id="3",name="DRP_CPool_3",status="offline"} 0
	spectrum_pool_status{id="3",name="DRP_CPool_3",status="online"} 1
	# HELP spectrum_pool_used_bytes Used bytes in pool
	# TYPE spectrum_pool_used_bytes gauge
	spectrum_pool_used_bytes{id="0",name="DRP_Pool0"} 3.331520232161e+12
	spectrum_pool_used_bytes{id="1",name="Backup_0"} 0
	spectrum_pool_used_bytes{id="2",name="Application"} 0
	spectrum_pool_used_bytes{id="3",name="DRP_CPool_3"} 0
	# HELP spectrum_pool_volume_count Number of volumes associated with pool
	# TYPE spectrum_pool_volume_count gauge
	spectrum_pool_volume_count{id="0",name="DRP_Pool0"} 18
	spectrum_pool_volume_count{id="1",name="Backup_0"} 0
	spectrum_pool_volume_count{id="2",name="Application"} 17
	spectrum_pool_volume_count{id="3",name="DRP_CPool_3"} 8
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestNodeStats(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsnodecanisterstats", "testdata/specv_output_8_5_2_2/lsnodecanisterstats.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeNodeStats(c, r) {
		t.Errorf("probeNodeStats() returned non-success")
	}

	em := `
	# HELP spectrum_node_compression_usage_ratio Current ratio of allocated CPU for compresion
	# TYPE spectrum_node_compression_usage_ratio gauge
	spectrum_node_compression_usage_ratio{id="1"} 0.24
	spectrum_node_compression_usage_ratio{id="2"} 0
	# HELP spectrum_node_fc_bps Current bytes-per-second being transferred over Fibre Channel
	# TYPE spectrum_node_fc_bps gauge
	spectrum_node_fc_bps{id="1"} 1.048576e+06
	spectrum_node_fc_bps{id="2"} 0
	# HELP spectrum_node_fc_iops Current I/O-per-second being transferred over Fibre Channel
	# TYPE spectrum_node_fc_iops gauge
	spectrum_node_fc_iops{id="1"} 5
	spectrum_node_fc_iops{id="2"} 5
	# HELP spectrum_node_iscsi_bps Current bytes-per-second being transferred over iSCSI
	# TYPE spectrum_node_iscsi_bps gauge
	spectrum_node_iscsi_bps{id="1"} 0
	spectrum_node_iscsi_bps{id="2"} 0
	# HELP spectrum_node_iscsi_iops Current I/O-per-second being transferred over iSCSI
	# TYPE spectrum_node_iscsi_iops gauge
	spectrum_node_iscsi_iops{id="1"} 0
	spectrum_node_iscsi_iops{id="2"} 11
	# HELP spectrum_node_mdisk_read_iops Current read I/O-per-second to mdisk
	# TYPE spectrum_node_mdisk_read_iops gauge
	spectrum_node_mdisk_read_iops{id="1"} 0
	spectrum_node_mdisk_read_iops{id="2"} 0
	# HELP spectrum_node_mdisk_read_mb Current Megabytes-per-second being read from mdisk
	# TYPE spectrum_node_mdisk_read_mb gauge
	spectrum_node_mdisk_read_mb{id="1"} 0
	spectrum_node_mdisk_read_mb{id="2"} 0
	# HELP spectrum_node_mdisk_read_ms Current milliseconds to read from mdisk
	# TYPE spectrum_node_mdisk_read_ms gauge
	spectrum_node_mdisk_read_ms{id="1"} 0
	spectrum_node_mdisk_read_ms{id="2"} 0
	# HELP spectrum_node_mdisk_write_iops Current write I/O-per-second to mdisk
	# TYPE spectrum_node_mdisk_write_iops gauge
	spectrum_node_mdisk_write_iops{id="1"} 0
	spectrum_node_mdisk_write_iops{id="2"} 0
	# HELP spectrum_node_mdisk_write_mb Current Megabytes-per-second being written to mdisk
	# TYPE spectrum_node_mdisk_write_mb gauge
	spectrum_node_mdisk_write_mb{id="1"} 0
	spectrum_node_mdisk_write_mb{id="2"} 0
	# HELP spectrum_node_mdisk_write_ms Current milliseconds to write to mdisk
	# TYPE spectrum_node_mdisk_write_ms gauge
	spectrum_node_mdisk_write_ms{id="1"} 0
	spectrum_node_mdisk_write_ms{id="2"} 7
	# HELP spectrum_node_sas_bps Current bytes-per-second being transferred over backend SAS
	# TYPE spectrum_node_sas_bps gauge
	spectrum_node_sas_bps{id="1"} 0
	spectrum_node_sas_bps{id="2"} 0
	# HELP spectrum_node_sas_iops Current I/O-per-second being transferred over backend SAS
	# TYPE spectrum_node_sas_iops gauge
	spectrum_node_sas_iops{id="1"} 5
	spectrum_node_sas_iops{id="2"} 0
	# HELP spectrum_node_system_usage_ratio Current ratio of allocated CPU for system
	# TYPE spectrum_node_system_usage_ratio gauge
	spectrum_node_system_usage_ratio{id="1"} 0.01
	spectrum_node_system_usage_ratio{id="2"} 0.01
	# HELP spectrum_node_total_cache_usage_ratio Total percentage for both the write and read cache usage for the node
	# TYPE spectrum_node_total_cache_usage_ratio gauge
	spectrum_node_total_cache_usage_ratio{id="1"} 0.79
	spectrum_node_total_cache_usage_ratio{id="2"} 0.79
	# HELP spectrum_node_vdisk_read_iops Current read I/O-per-second to vdisk
	# TYPE spectrum_node_vdisk_read_iops gauge
	spectrum_node_vdisk_read_iops{id="1"} 0
	spectrum_node_vdisk_read_iops{id="2"} 4
	# HELP spectrum_node_vdisk_read_mb Current Megabytes-per-second being read from vdisk
	# TYPE spectrum_node_vdisk_read_mb gauge
	spectrum_node_vdisk_read_mb{id="1"} 0
	spectrum_node_vdisk_read_mb{id="2"} 0
	# HELP spectrum_node_vdisk_read_ms Current milliseconds to read from vdisk
	# TYPE spectrum_node_vdisk_read_ms gauge
	spectrum_node_vdisk_read_ms{id="1"} 1
	spectrum_node_vdisk_read_ms{id="2"} 1
	# HELP spectrum_node_vdisk_write_iops Current write I/O-per-second to vdisk
	# TYPE spectrum_node_vdisk_write_iops gauge
	spectrum_node_vdisk_write_iops{id="1"} 36
	spectrum_node_vdisk_write_iops{id="2"} 19
	# HELP spectrum_node_vdisk_write_mb Current Megabytes-per-second being written to vdisk
	# TYPE spectrum_node_vdisk_write_mb gauge
	spectrum_node_vdisk_write_mb{id="1"} 0
	spectrum_node_vdisk_write_mb{id="2"} 0
	# HELP spectrum_node_vdisk_write_ms Current milliseconds to write to vdisk
	# TYPE spectrum_node_vdisk_write_ms gauge
	spectrum_node_vdisk_write_ms{id="1"} 0
	spectrum_node_vdisk_write_ms{id="2"} 0
	# HELP spectrum_node_write_cache_usage_ratio Ratio of the write cache usage for the node
	# TYPE spectrum_node_write_cache_usage_ratio gauge
	spectrum_node_write_cache_usage_ratio{id="1"} 0.25
	spectrum_node_write_cache_usage_ratio{id="2"} 0.25
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestFCPorts(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsportfc", "testdata/specv_output_8_5_2_2/lsportfc.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeFCPorts(c, r) {
		t.Errorf("probeFCPorts() returned non-success")
	}

	em := `
	# HELP spectrum_fc_port_speed_bps Operational speed of port in bits per second
	# TYPE spectrum_fc_port_speed_bps gauge
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="1",node_id="1"} 3.2e+10
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="1",node_id="2"} 3.2e+10
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="2",node_id="1"} 3.2e+10
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="2",node_id="2"} 3.2e+10
	# HELP spectrum_fc_port_status Status of Fibre Channel port
	# TYPE spectrum_fc_port_status gauge
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="1",status="active",wwpn="500507680B11FFC8"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="1",status="inactive_configured",wwpn="500507680B11FFC8"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="1",status="inactive_unconfigured",wwpn="500507680B11FFC8"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="2",status="active",wwpn="500507680B11FFC9"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="2",status="inactive_configured",wwpn="500507680B11FFC9"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="2",status="inactive_unconfigured",wwpn="500507680B11FFC9"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="1",status="active",wwpn="500507680B12FFC8"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="1",status="inactive_configured",wwpn="500507680B12FFC8"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="1",status="inactive_unconfigured",wwpn="500507680B12FFC8"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="2",status="active",wwpn="500507680B12FFC9"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="2",status="inactive_configured",wwpn="500507680B12FFC9"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="2",status="inactive_unconfigured",wwpn="500507680B12FFC9"} 0
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestIPPorts(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsportip", "testdata/specv_output_8_5_2_2/lsportip.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeIPPorts(c, r) {
		t.Errorf("probeIPPorts() returned non-success")
	}

	em := `
	# HELP spectrum_ip_port_link_active Whether link is active
	# TYPE spectrum_ip_port_link_active gauge
	spectrum_ip_port_link_active{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="1"} 0
	spectrum_ip_port_link_active{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="2"} 0
	# HELP spectrum_ip_port_speed_bps Operational speed of port in bits per second
	# TYPE spectrum_ip_port_speed_bps gauge
	spectrum_ip_port_speed_bps{adapter_location="0",adapter_port_id="1",node_id="1"} 1e+09
	spectrum_ip_port_speed_bps{adapter_location="0",adapter_port_id="1",node_id="2"} 1e+09
	# HELP spectrum_ip_port_state Configuration state of Ethernet/IP port
	# TYPE spectrum_ip_port_state gauge
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="1",state="configured"} 0
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="1",state="management_only"} 0
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="1",state="unconfigured"} 1
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="2",state="configured"} 0
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="2",state="management_only"} 0
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="2",state="unconfigured"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestQuorumStatus(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsquorum", "testdata/specv_output_8_5_2_2/lsquorum.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeQuorum(c, r) {
		t.Errorf("probeQuorumStatus() returned non-success")
	}

	em := `
	# HELP spectrum_quorum_status Status of quorum
	# TYPE spectrum_quorum_status gauge
	spectrum_quorum_status{id="0",object_type="drive",status="offline"} 0
	spectrum_quorum_status{id="0",object_type="drive",status="online"} 1
	spectrum_quorum_status{id="1",object_type="drive",status="offline"} 0
	spectrum_quorum_status{id="1",object_type="drive",status="online"} 1
	spectrum_quorum_status{id="2",object_type="drive",status="offline"} 0
	spectrum_quorum_status{id="2",object_type="drive",status="online"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestHostStatus(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lshost", "testdata/specv_output_8_5_2_2/lshost.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeHost(c, r) {
		t.Errorf("probeHostStatus() returned non-success")
	}

	em := `
	# HELP spectrum_host_status Status of hosts
	# TYPE spectrum_host_status gauge
	spectrum_host_status{hostname="VMware_202",id="1",port_count="1",protocol="rdmanvme",status="degraded"} 0
	spectrum_host_status{hostname="VMware_202",id="1",port_count="1",protocol="rdmanvme",status="offline"} 0
	spectrum_host_status{hostname="VMware_202",id="1",port_count="1",protocol="rdmanvme",status="online"} 1
	spectrum_host_status{hostname="x3690-x5-01",id="0",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="x3690-x5-01",id="0",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="x3690-x5-01",id="0",port_count="2",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="x3690_x5_04",id="2",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="x3690_x5_04",id="2",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="x3690_x5_04",id="2",port_count="2",protocol="scsi",status="online"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
