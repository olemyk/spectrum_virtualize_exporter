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
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

// Section for 8.6.0.0 on IBM flashsystem with HyperSwap

func TestEnclosureStats_8600(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsenclosurestats", "testdata/specv_output_8_6_0_0/lsenclosurestats.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeEnclosureStats(c, r, false) {
		t.Errorf("probeEnclosureStats() returned non-success")
	}

	em := `
	# HELP spectrum_power_watts Current power draw of enclosure in watts
	# TYPE spectrum_power_watts gauge
	spectrum_power_watts{enclosure="1"} 401
	spectrum_power_watts{enclosure="2"} 337
	# HELP spectrum_temperature Current enclosure temperature in celsius
	# TYPE spectrum_temperature gauge
	spectrum_temperature{enclosure="1"} 28
	spectrum_temperature{enclosure="2"} 25
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
func TestDrive_8600(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsdrive", "testdata/specv_output_8_6_0_0/lsdrive.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeDrives(c, r) {
		t.Errorf("probeDrives() returned non-success")
	}

	em := `
	# HELP spectrum_drive_status Status of drive
	# TYPE spectrum_drive_status gauge
	spectrum_drive_status{enclosure="1",id="0",slot_id="5",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="0",slot_id="5",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="0",slot_id="5",status="online"} 1
	spectrum_drive_status{enclosure="1",id="1",slot_id="4",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="1",slot_id="4",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="1",slot_id="4",status="online"} 1
	spectrum_drive_status{enclosure="1",id="2",slot_id="12",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="2",slot_id="12",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="2",slot_id="12",status="online"} 1
	spectrum_drive_status{enclosure="1",id="3",slot_id="3",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="3",slot_id="3",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="3",slot_id="3",status="online"} 1
	spectrum_drive_status{enclosure="1",id="4",slot_id="6",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="4",slot_id="6",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="4",slot_id="6",status="online"} 1
	spectrum_drive_status{enclosure="1",id="5",slot_id="2",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="5",slot_id="2",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="5",slot_id="2",status="online"} 1
	spectrum_drive_status{enclosure="1",id="6",slot_id="1",status="degraded"} 0
	spectrum_drive_status{enclosure="1",id="6",slot_id="1",status="offline"} 0
	spectrum_drive_status{enclosure="1",id="6",slot_id="1",status="online"} 1
	spectrum_drive_status{enclosure="2",id="7",slot_id="8",status="degraded"} 0
	spectrum_drive_status{enclosure="2",id="7",slot_id="8",status="offline"} 0
	spectrum_drive_status{enclosure="2",id="7",slot_id="8",status="online"} 1
	spectrum_drive_status{enclosure="2",id="8",slot_id="2",status="degraded"} 0
	spectrum_drive_status{enclosure="2",id="8",slot_id="2",status="offline"} 0
	spectrum_drive_status{enclosure="2",id="8",slot_id="2",status="online"} 1
	spectrum_drive_status{enclosure="2",id="9",slot_id="6",status="degraded"} 0
	spectrum_drive_status{enclosure="2",id="9",slot_id="6",status="offline"} 0
	spectrum_drive_status{enclosure="2",id="9",slot_id="6",status="online"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestEnclosurePSU_8600(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsenclosurepsu", "testdata/specv_output_8_6_0_0/lsenclosurepsu.jsonnet")
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
	spectrum_psu_status{enclosure="2",id="1",status="degraded"} 0
	spectrum_psu_status{enclosure="2",id="1",status="offline"} 0
	spectrum_psu_status{enclosure="2",id="1",status="online"} 1
	spectrum_psu_status{enclosure="2",id="2",status="degraded"} 0
	spectrum_psu_status{enclosure="2",id="2",status="offline"} 0
	spectrum_psu_status{enclosure="2",id="2",status="online"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestPool_8600(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsmdiskgrp", "testdata/specv_output_8_6_0_0/lsmdiskgrp.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probePool(c, r) {
		t.Errorf("probePool() returned non-success")
	}

	em := `
	# HELP spectrum_pool_capacity_bytes Capacity of pool in bytes
	# TYPE spectrum_pool_capacity_bytes gauge
	spectrum_pool_capacity_bytes{id="0",name="DRP_Pool_S1"} 1.74096671142051e+14
	spectrum_pool_capacity_bytes{id="1",name="DRP_Pool_S2"} 2.1979237439242e+13
	spectrum_pool_capacity_bytes{id="2",name="Backup_Location"} 1.74096671142051e+14
	# HELP spectrum_pool_free_bytes Free bytes in pool
	# TYPE spectrum_pool_free_bytes gauge
	spectrum_pool_free_bytes{id="0",name="DRP_Pool_S1"} 1.65355553701232e+14
	spectrum_pool_free_bytes{id="1",name="DRP_Pool_S2"} 1.7515220230471e+13
	spectrum_pool_free_bytes{id="2",name="Backup_Location"} 1.65355553701232e+14
	# HELP spectrum_pool_status Status of pool
	# TYPE spectrum_pool_status gauge
	spectrum_pool_status{id="0",name="DRP_Pool_S1",status="offline"} 0
	spectrum_pool_status{id="0",name="DRP_Pool_S1",status="online"} 1
	spectrum_pool_status{id="1",name="DRP_Pool_S2",status="offline"} 0
	spectrum_pool_status{id="1",name="DRP_Pool_S2",status="online"} 1
	spectrum_pool_status{id="2",name="Backup_Location",status="offline"} 0
	spectrum_pool_status{id="2",name="Backup_Location",status="online"} 1
	# HELP spectrum_pool_used_bytes Used bytes in pool
	# TYPE spectrum_pool_used_bytes gauge
	spectrum_pool_used_bytes{id="0",name="DRP_Pool_S1"} 6.410152789934e+12
	spectrum_pool_used_bytes{id="1",name="DRP_Pool_S2"} 3.672368836771e+12
	spectrum_pool_used_bytes{id="2",name="Backup_Location"} 0
	# HELP spectrum_pool_volume_count Number of volumes associated with pool
	# TYPE spectrum_pool_volume_count gauge
	spectrum_pool_volume_count{id="0",name="DRP_Pool_S1"} 69
	spectrum_pool_volume_count{id="1",name="DRP_Pool_S2"} 63
	spectrum_pool_volume_count{id="2",name="Backup_Location"} 24
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestNodeStats_8600(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsnodecanisterstats", "testdata/specv_output_8_6_0_0/lsnodecanisterstats.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeNodeStats(c, r) {
		t.Errorf("probeNodeStats() returned non-success")
	}

	em := `
	# HELP spectrum_node_compression_usage_ratio Current ratio of allocated CPU for compresion
	# TYPE spectrum_node_compression_usage_ratio gauge
	spectrum_node_compression_usage_ratio{id="1"} 0
	spectrum_node_compression_usage_ratio{id="2"} 0
	spectrum_node_compression_usage_ratio{id="3"} 0
	spectrum_node_compression_usage_ratio{id="5"} 0
	# HELP spectrum_node_fc_bps Current bytes-per-second being transferred over Fibre Channel
	# TYPE spectrum_node_fc_bps gauge
	spectrum_node_fc_bps{id="1"} 0
	spectrum_node_fc_bps{id="2"} 0
	spectrum_node_fc_bps{id="3"} 0
	spectrum_node_fc_bps{id="5"} 0
	# HELP spectrum_node_fc_iops Current I/O-per-second being transferred over Fibre Channel
	# TYPE spectrum_node_fc_iops gauge
	spectrum_node_fc_iops{id="1"} 3282
	spectrum_node_fc_iops{id="2"} 3511
	spectrum_node_fc_iops{id="3"} 1445
	spectrum_node_fc_iops{id="5"} 1362
	# HELP spectrum_node_iscsi_bps Current bytes-per-second being transferred over iSCSI
	# TYPE spectrum_node_iscsi_bps gauge
	spectrum_node_iscsi_bps{id="1"} 0
	spectrum_node_iscsi_bps{id="2"} 0
	spectrum_node_iscsi_bps{id="3"} 0
	spectrum_node_iscsi_bps{id="5"} 0
	# HELP spectrum_node_iscsi_iops Current I/O-per-second being transferred over iSCSI
	# TYPE spectrum_node_iscsi_iops gauge
	spectrum_node_iscsi_iops{id="1"} 0
	spectrum_node_iscsi_iops{id="2"} 0
	spectrum_node_iscsi_iops{id="3"} 0
	spectrum_node_iscsi_iops{id="5"} 0
	# HELP spectrum_node_mdisk_read_iops Current read I/O-per-second to mdisk
	# TYPE spectrum_node_mdisk_read_iops gauge
	spectrum_node_mdisk_read_iops{id="1"} 0
	spectrum_node_mdisk_read_iops{id="2"} 0
	spectrum_node_mdisk_read_iops{id="3"} 0
	spectrum_node_mdisk_read_iops{id="5"} 0
	# HELP spectrum_node_mdisk_read_mb Current Megabytes-per-second being read from mdisk
	# TYPE spectrum_node_mdisk_read_mb gauge
	spectrum_node_mdisk_read_mb{id="1"} 0
	spectrum_node_mdisk_read_mb{id="2"} 0
	spectrum_node_mdisk_read_mb{id="3"} 0
	spectrum_node_mdisk_read_mb{id="5"} 0
	# HELP spectrum_node_mdisk_read_ms Current milliseconds to read from mdisk
	# TYPE spectrum_node_mdisk_read_ms gauge
	spectrum_node_mdisk_read_ms{id="1"} 0
	spectrum_node_mdisk_read_ms{id="2"} 0
	spectrum_node_mdisk_read_ms{id="3"} 0
	spectrum_node_mdisk_read_ms{id="5"} 0
	# HELP spectrum_node_mdisk_write_iops Current write I/O-per-second to mdisk
	# TYPE spectrum_node_mdisk_write_iops gauge
	spectrum_node_mdisk_write_iops{id="1"} 0
	spectrum_node_mdisk_write_iops{id="2"} 0
	spectrum_node_mdisk_write_iops{id="3"} 0
	spectrum_node_mdisk_write_iops{id="5"} 0
	# HELP spectrum_node_mdisk_write_mb Current Megabytes-per-second being written to mdisk
	# TYPE spectrum_node_mdisk_write_mb gauge
	spectrum_node_mdisk_write_mb{id="1"} 0
	spectrum_node_mdisk_write_mb{id="2"} 0
	spectrum_node_mdisk_write_mb{id="3"} 0
	spectrum_node_mdisk_write_mb{id="5"} 0
	# HELP spectrum_node_mdisk_write_ms Current milliseconds to write to mdisk
	# TYPE spectrum_node_mdisk_write_ms gauge
	spectrum_node_mdisk_write_ms{id="1"} 0
	spectrum_node_mdisk_write_ms{id="2"} 0
	spectrum_node_mdisk_write_ms{id="3"} 0
	spectrum_node_mdisk_write_ms{id="5"} 0
	# HELP spectrum_node_sas_bps Current bytes-per-second being transferred over backend SAS
	# TYPE spectrum_node_sas_bps gauge
	spectrum_node_sas_bps{id="1"} 0
	spectrum_node_sas_bps{id="2"} 0
	spectrum_node_sas_bps{id="3"} 0
	spectrum_node_sas_bps{id="5"} 0
	# HELP spectrum_node_sas_iops Current I/O-per-second being transferred over backend SAS
	# TYPE spectrum_node_sas_iops gauge
	spectrum_node_sas_iops{id="1"} 0
	spectrum_node_sas_iops{id="2"} 0
	spectrum_node_sas_iops{id="3"} 0
	spectrum_node_sas_iops{id="5"} 0
	# HELP spectrum_node_system_usage_ratio Current ratio of allocated CPU for system
	# TYPE spectrum_node_system_usage_ratio gauge
	spectrum_node_system_usage_ratio{id="1"} 0.02
	spectrum_node_system_usage_ratio{id="2"} 0.01
	spectrum_node_system_usage_ratio{id="3"} 0.01
	spectrum_node_system_usage_ratio{id="5"} 0.01
	# HELP spectrum_node_total_cache_usage_ratio Total percentage for both the write and read cache usage for the node
	# TYPE spectrum_node_total_cache_usage_ratio gauge
	spectrum_node_total_cache_usage_ratio{id="1"} 0.79
	spectrum_node_total_cache_usage_ratio{id="2"} 0.79
	spectrum_node_total_cache_usage_ratio{id="3"} 0.29
	spectrum_node_total_cache_usage_ratio{id="5"} 0.3
	# HELP spectrum_node_vdisk_read_iops Current read I/O-per-second to vdisk
	# TYPE spectrum_node_vdisk_read_iops gauge
	spectrum_node_vdisk_read_iops{id="1"} 5
	spectrum_node_vdisk_read_iops{id="2"} 1
	spectrum_node_vdisk_read_iops{id="3"} 6
	spectrum_node_vdisk_read_iops{id="5"} 2
	# HELP spectrum_node_vdisk_read_mb Current Megabytes-per-second being read from vdisk
	# TYPE spectrum_node_vdisk_read_mb gauge
	spectrum_node_vdisk_read_mb{id="1"} 0
	spectrum_node_vdisk_read_mb{id="2"} 0
	spectrum_node_vdisk_read_mb{id="3"} 0
	spectrum_node_vdisk_read_mb{id="5"} 0
	# HELP spectrum_node_vdisk_read_ms Current milliseconds to read from vdisk
	# TYPE spectrum_node_vdisk_read_ms gauge
	spectrum_node_vdisk_read_ms{id="1"} 0.672
	spectrum_node_vdisk_read_ms{id="2"} 0.083
	spectrum_node_vdisk_read_ms{id="3"} 0.057
	spectrum_node_vdisk_read_ms{id="5"} 0.077
	# HELP spectrum_node_vdisk_write_iops Current write I/O-per-second to vdisk
	# TYPE spectrum_node_vdisk_write_iops gauge
	spectrum_node_vdisk_write_iops{id="1"} 4
	spectrum_node_vdisk_write_iops{id="2"} 0
	spectrum_node_vdisk_write_iops{id="3"} 2
	spectrum_node_vdisk_write_iops{id="5"} 2
	# HELP spectrum_node_vdisk_write_mb Current Megabytes-per-second being written to vdisk
	# TYPE spectrum_node_vdisk_write_mb gauge
	spectrum_node_vdisk_write_mb{id="1"} 0
	spectrum_node_vdisk_write_mb{id="2"} 0
	spectrum_node_vdisk_write_mb{id="3"} 0
	spectrum_node_vdisk_write_mb{id="5"} 0
	# HELP spectrum_node_vdisk_write_ms Current milliseconds to write to vdisk
	# TYPE spectrum_node_vdisk_write_ms gauge
	spectrum_node_vdisk_write_ms{id="1"} 0.594
	spectrum_node_vdisk_write_ms{id="2"} 1.539
	spectrum_node_vdisk_write_ms{id="3"} 2.058
	spectrum_node_vdisk_write_ms{id="5"} 1.138
	# HELP spectrum_node_write_cache_usage_ratio Ratio of the write cache usage for the node
	# TYPE spectrum_node_write_cache_usage_ratio gauge
	spectrum_node_write_cache_usage_ratio{id="1"} 0.3
	spectrum_node_write_cache_usage_ratio{id="2"} 0.3
	spectrum_node_write_cache_usage_ratio{id="3"} 0.08
	spectrum_node_write_cache_usage_ratio{id="5"} 0.08
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestFCPorts_8600(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsportfc", "testdata/specv_output_8_6_0_0/lsportfc.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeFCPorts(c, r) {
		t.Errorf("probeFCPorts() returned non-success")
	}

	em := `
	# HELP spectrum_fc_port_speed_bps Operational speed of port in bits per second
	# TYPE spectrum_fc_port_speed_bps gauge
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="1",node_id="1"} 3.2e+10
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="1",node_id="2"} 3.2e+10
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="1",node_id="3"} 1.6e+10
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="1",node_id="5"} 1.6e+10
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="2",node_id="1"} 1.6e+10
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="2",node_id="2"} 1.6e+10
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="2",node_id="3"} 1.6e+10
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="2",node_id="5"} 1.6e+10
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="1",node_id="1"} 3.2e+10
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="1",node_id="2"} 3.2e+10
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="2",node_id="1"} 1.6e+10
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="2",node_id="2"} 1.6e+10
	# HELP spectrum_fc_port_status Status of Fibre Channel port
	# TYPE spectrum_fc_port_status gauge
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="1",status="active",wwpn="5005076810110E08"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="1",status="inactive_configured",wwpn="5005076810110E08"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="1",status="inactive_unconfigured",wwpn="5005076810110E08"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="2",status="active",wwpn="500507681011F110"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="2",status="inactive_configured",wwpn="500507681011F110"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="2",status="inactive_unconfigured",wwpn="500507681011F110"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="3",status="active",wwpn="500507680B11FFDC"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="3",status="inactive_configured",wwpn="500507680B11FFDC"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="3",status="inactive_unconfigured",wwpn="500507680B11FFDC"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="5",status="active",wwpn="500507680B11FFDD"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="5",status="inactive_configured",wwpn="500507680B11FFDD"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="5",status="inactive_unconfigured",wwpn="500507680B11FFDD"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="1",status="active",wwpn="5005076810120E08"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="1",status="inactive_configured",wwpn="5005076810120E08"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="1",status="inactive_unconfigured",wwpn="5005076810120E08"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="2",status="active",wwpn="500507681012F110"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="2",status="inactive_configured",wwpn="500507681012F110"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="2",status="inactive_unconfigured",wwpn="500507681012F110"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="3",status="active",wwpn="500507680B12FFDC"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="3",status="inactive_configured",wwpn="500507680B12FFDC"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="3",status="inactive_unconfigured",wwpn="500507680B12FFDC"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="5",status="active",wwpn="500507680B12FFDD"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="5",status="inactive_configured",wwpn="500507680B12FFDD"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="5",status="inactive_unconfigured",wwpn="500507680B12FFDD"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="1",status="active",wwpn="5005076810210E08"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="1",status="inactive_configured",wwpn="5005076810210E08"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="1",status="inactive_unconfigured",wwpn="5005076810210E08"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="2",status="active",wwpn="500507681021F110"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="2",status="inactive_configured",wwpn="500507681021F110"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="2",status="inactive_unconfigured",wwpn="500507681021F110"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="1",status="active",wwpn="5005076810220E08"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="1",status="inactive_configured",wwpn="5005076810220E08"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="1",status="inactive_unconfigured",wwpn="5005076810220E08"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="2",status="active",wwpn="500507681022F110"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="2",status="inactive_configured",wwpn="500507681022F110"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="2",status="inactive_unconfigured",wwpn="500507681022F110"} 0
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestIPPorts_8600(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsportip", "testdata/specv_output_8_6_0_0/lsportip.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeIPPorts(c, r) {
		t.Errorf("probeIPPorts() returned non-success")
	}

	em := `
	# HELP spectrum_ip_port_link_active Whether link is active
	# TYPE spectrum_ip_port_link_active gauge
	spectrum_ip_port_link_active{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="1"} 0
	spectrum_ip_port_link_active{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="2"} 0
	spectrum_ip_port_link_active{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="3"} 0
	spectrum_ip_port_link_active{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="5"} 0
	# HELP spectrum_ip_port_speed_bps Operational speed of port in bits per second
	# TYPE spectrum_ip_port_speed_bps gauge
	spectrum_ip_port_speed_bps{adapter_location="0",adapter_port_id="1",node_id="1"} 1e+09
	spectrum_ip_port_speed_bps{adapter_location="0",adapter_port_id="1",node_id="2"} 1e+09
	spectrum_ip_port_speed_bps{adapter_location="0",adapter_port_id="1",node_id="3"} 1e+09
	spectrum_ip_port_speed_bps{adapter_location="0",adapter_port_id="1",node_id="5"} 1e+09
	# HELP spectrum_ip_port_state Configuration state of Ethernet/IP port
	# TYPE spectrum_ip_port_state gauge
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="1",state="configured"} 0
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="1",state="management_only"} 0
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="1",state="unconfigured"} 1
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="2",state="configured"} 0
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="2",state="management_only"} 0
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="2",state="unconfigured"} 1
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="3",state="configured"} 0
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="3",state="management_only"} 0
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="3",state="unconfigured"} 1
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="5",state="configured"} 0
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="5",state="management_only"} 0
	spectrum_ip_port_state{adapter_location="0",adapter_port_id="1",mac="ff:ff:ff:ff:ff:ff",node_id="5",state="unconfigured"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestQuorumStatus_8600(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsquorum", "testdata/specv_output_8_6_0_0/lsquorum.jsonnet")
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
	spectrum_quorum_status{id="5",object_type="device",status="offline"} 0
	spectrum_quorum_status{id="5",object_type="device",status="online"} 1
	spectrum_quorum_status{id="6",object_type="device",status="offline"} 0
	spectrum_quorum_status{id="6",object_type="device",status="online"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestHostStatus_8600(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lshost", "testdata/specv_output_8_6_0_0/lshost.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeHost(c, r) {
		t.Errorf("probeHostStatus() returned non-success")
	}

	em := `
	# HELP spectrum_host_status Status of hosts
	# TYPE spectrum_host_status gauge
	spectrum_host_status{hostname="AIX_18_146_S1",id="7",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="AIX_18_146_S1",id="7",port_count="2",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="AIX_18_146_S1",id="7",port_count="2",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="AIX_18_147_S2",id="8",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="AIX_18_147_S2",id="8",port_count="2",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="AIX_18_147_S2",id="8",port_count="2",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="AIX_Site1",id="5",port_count="4",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="AIX_Site1",id="5",port_count="4",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="AIX_Site1",id="5",port_count="4",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="AIX_Site2",id="6",port_count="4",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="AIX_Site2",id="6",port_count="4",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="AIX_Site2",id="6",port_count="4",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="ESX_122",id="9",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="ESX_122",id="9",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="ESX_122",id="9",port_count="2",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="IBMi_1",id="11",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="IBMi_1",id="11",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="IBMi_1",id="11",port_count="2",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="IBMi_2",id="10",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="IBMi_2",id="10",port_count="2",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="IBMi_2",id="10",port_count="2",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="IBMi_App_Server",id="2",port_count="2",protocol="scsi",status="degraded"} 1
	spectrum_host_status{hostname="IBMi_App_Server",id="2",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="IBMi_App_Server",id="2",port_count="2",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="IBMi_Recov_Server",id="3",port_count="2",protocol="scsi",status="degraded"} 1
	spectrum_host_status{hostname="IBMi_Recov_Server",id="3",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="IBMi_Recov_Server",id="3",port_count="2",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="Site1_ESX",id="0",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="Site1_ESX",id="0",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="Site1_ESX",id="0",port_count="2",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="Site2_ESX",id="1",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="Site2_ESX",id="1",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="Site2_ESX",id="1",port_count="2",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="sshoeib_host",id="4",port_count="1",protocol="fcnvme",status="degraded"} 0
	spectrum_host_status{hostname="sshoeib_host",id="4",port_count="1",protocol="fcnvme",status="offline"} 1
	spectrum_host_status{hostname="sshoeib_host",id="4",port_count="1",protocol="fcnvme",status="online"} 0
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
