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

// there is no output from lsenclosurestats on SVC

// func TestEnclosureStats_8_5_0_6_SVC(t *testing.T) {
// 	c := newFakeClient()
// 	c.prepare("rest/v1/lsenclosurestats", "testdata/specv_output_8_5_0_6_svc/lsenclosurestats.jsonnet")
// 	r := prometheus.NewPedanticRegistry()
// 	if !probeEnclosureStats(c, r, true) {
// 		t.Errorf("probeEnclosureStats() returned non-success")
// 	}

// 	em := `
// 	# HELP spectrum_power_watts Current power draw of enclosure in watts
// 	# TYPE spectrum_power_watts gauge
// 	spectrum_power_watts{enclosure="1"} 480
// 	# HELP spectrum_temperature Current enclosure temperature in celsius
// 	# TYPE spectrum_temperature gauge
// 	spectrum_temperature{enclosure="1"} 22
// 	`

//		if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
//			t.Fatalf("metric compare: err %v", err)
//		}
//	}

// there is no output from lsdrive on SVC
// func TestDrive_8_5_0_6_SVC(t *testing.T) {
// 	c := newFakeClient()
// 	c.prepare("rest/v1/lsdrive", "testdata/specv_output_8_5_0_6_svc/lsdrive.jsonnet")
// 	r := prometheus.NewPedanticRegistry()
// 	if !probeDrives(c, r) {
// 		t.Errorf("probeDrives() returned non-success")
// 	}

// 	em := `
// 	# HELP spectrum_drive_status Status of drive
// 	# TYPE spectrum_drive_status gauge
// 	`

// 	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
// 		t.Fatalf("metric compare: err %v", err)
// 	}
// }

// there is no output from TestEnclosurePSU on SVC
// func TestEnclosurePSU_8_5_0_6_SVC(t *testing.T) {
// 	c := newFakeClient()
// 	c.prepare("rest/v1/lsenclosurepsu", "testdata/specv_output_8_5_0_6_svc/lsenclosurepsu.jsonnet")
// 	r := prometheus.NewPedanticRegistry()
// 	if !probeEnclosurePSUs(c, r) {
// 		t.Errorf("probeEnclosurePSUs() returned non-success")
// 	}

// 	em := `
// 	# HELP spectrum_psu_status Status of PSU
// 	# TYPE spectrum_psu_status gauge
// 	`

// 	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
// 		t.Fatalf("metric compare: err %v", err)
// 	}
// }

func TestPool_8_5_0_6_SVC(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsmdiskgrp", "testdata/specv_output_8_5_0_6_svc/lsmdiskgrp.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probePool(c, r) {
		t.Errorf("probePool() returned non-success")
	}

	em := `
	# HELP spectrum_pool_capacity_bytes Capacity of pool in bytes
	# TYPE spectrum_pool_capacity_bytes gauge
	spectrum_pool_capacity_bytes{id="0",name="FS840-1"} 1.2303535114813e+13
	spectrum_pool_capacity_bytes{id="1",name="FS840-2"} 1.2303535114813e+13
	spectrum_pool_capacity_bytes{id="2",name="FS840_2_child_pool"} 5.49755813888e+11
	spectrum_pool_capacity_bytes{id="3",name="FS840_2_child_pool2"} 2.2011707392e+11
	spectrum_pool_capacity_bytes{id="4",name="storwize_power_pool"} 1.2567417905479e+13
	# HELP spectrum_pool_free_bytes Free bytes in pool
	# TYPE spectrum_pool_free_bytes gauge
	spectrum_pool_free_bytes{id="0",name="FS840-1"} 7.20480763904e+11
	spectrum_pool_free_bytes{id="1",name="FS840-2"} 1.803199069552e+12
	spectrum_pool_free_bytes{id="2",name="FS840_2_child_pool"} 5.49755813888e+11
	spectrum_pool_free_bytes{id="3",name="FS840_2_child_pool2"} 2.147483648e+11
	spectrum_pool_free_bytes{id="4",name="storwize_power_pool"} 1.792203953274e+12
	# HELP spectrum_pool_status Status of pool
	# TYPE spectrum_pool_status gauge
	spectrum_pool_status{id="0",name="FS840-1",status="offline"} 0
	spectrum_pool_status{id="0",name="FS840-1",status="online"} 1
	spectrum_pool_status{id="1",name="FS840-2",status="offline"} 0
	spectrum_pool_status{id="1",name="FS840-2",status="online"} 1
	spectrum_pool_status{id="2",name="FS840_2_child_pool",status="offline"} 0
	spectrum_pool_status{id="2",name="FS840_2_child_pool",status="online"} 1
	spectrum_pool_status{id="3",name="FS840_2_child_pool2",status="offline"} 0
	spectrum_pool_status{id="3",name="FS840_2_child_pool2",status="online"} 1
	spectrum_pool_status{id="4",name="storwize_power_pool",status="offline"} 0
	spectrum_pool_status{id="4",name="storwize_power_pool",status="online"} 1
	# HELP spectrum_pool_used_bytes Used bytes in pool
	# TYPE spectrum_pool_used_bytes gauge
	spectrum_pool_used_bytes{id="0",name="FS840-1"} 1.1127057673093e+13
	spectrum_pool_used_bytes{id="1",name="FS840-2"} 9.488785347706e+12
	spectrum_pool_used_bytes{id="2",name="FS840_2_child_pool"} 0
	spectrum_pool_used_bytes{id="3",name="FS840_2_child_pool2"} 786432
	spectrum_pool_used_bytes{id="4",name="storwize_power_pool"} 1.0511331161538e+13
	# HELP spectrum_pool_volume_count Number of volumes associated with pool
	# TYPE spectrum_pool_volume_count gauge
	spectrum_pool_volume_count{id="0",name="FS840-1"} 67
	spectrum_pool_volume_count{id="1",name="FS840-2"} 45
	spectrum_pool_volume_count{id="2",name="FS840_2_child_pool"} 0
	spectrum_pool_volume_count{id="3",name="FS840_2_child_pool2"} 1
	spectrum_pool_volume_count{id="4",name="storwize_power_pool"} 4
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestNodeStats_8_5_0_6_SVC(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsnodecanisterstats", "testdata/specv_output_8_5_0_6_svc/lsnodecanisterstats.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeNodeStats(c, r) {
		t.Errorf("probeNodeStats() returned non-success")
	}

	em := `
	# HELP spectrum_node_compression_usage_ratio Current ratio of allocated CPU for compresion
	# TYPE spectrum_node_compression_usage_ratio gauge
	spectrum_node_compression_usage_ratio{id="1"} 0.08
	spectrum_node_compression_usage_ratio{id="10"} 0
	spectrum_node_compression_usage_ratio{id="2"} 0.08
	spectrum_node_compression_usage_ratio{id="8"} 0
	# HELP spectrum_node_fc_bps Current bytes-per-second being transferred over Fibre Channel
	# TYPE spectrum_node_fc_bps gauge
    spectrum_node_fc_bps{id="1"} 1.9922944e+07
    spectrum_node_fc_bps{id="10"} 6.291456e+06
    spectrum_node_fc_bps{id="2"} 8.388608e+06
    spectrum_node_fc_bps{id="8"} 9.437184e+06
	# HELP spectrum_node_fc_iops Current I/O-per-second being transferred over Fibre Channel
	# TYPE spectrum_node_fc_iops gauge
	spectrum_node_fc_iops{id="1"} 8735
	spectrum_node_fc_iops{id="10"} 5640
	spectrum_node_fc_iops{id="2"} 7346
	# HELP spectrum_node_iscsi_bps Current bytes-per-second being transferred over iSCSI
	# TYPE spectrum_node_iscsi_bps gauge
	spectrum_node_iscsi_bps{id="1"} 0
	spectrum_node_iscsi_bps{id="10"} 0
	spectrum_node_iscsi_bps{id="2"} 0
	spectrum_node_iscsi_bps{id="8"} 0
	spectrum_node_fc_iops{id="8"} 6677
	# HELP spectrum_node_iscsi_iops Current I/O-per-second being transferred over iSCSI
	# TYPE spectrum_node_iscsi_iops gauge
	spectrum_node_iscsi_iops{id="1"} 0
	spectrum_node_iscsi_iops{id="10"} 0
	spectrum_node_iscsi_iops{id="2"} 0
	spectrum_node_iscsi_iops{id="8"} 0
	# HELP spectrum_node_mdisk_read_iops Current read I/O-per-second to mdisk
	# TYPE spectrum_node_mdisk_read_iops gauge
	spectrum_node_mdisk_read_iops{id="1"} 20
	spectrum_node_mdisk_read_iops{id="10"} 0
	spectrum_node_mdisk_read_iops{id="2"} 0
	spectrum_node_mdisk_read_iops{id="8"} 0
	# HELP spectrum_node_mdisk_read_mb Current Megabytes-per-second being read from mdisk
	# TYPE spectrum_node_mdisk_read_mb gauge
	spectrum_node_mdisk_read_mb{id="10"} 0
	spectrum_node_mdisk_read_mb{id="2"} 0
	spectrum_node_mdisk_read_mb{id="8"} 0
	spectrum_node_mdisk_read_mb{id="1"} 0
	# HELP spectrum_node_mdisk_read_ms Current milliseconds to read from mdisk
	# TYPE spectrum_node_mdisk_read_ms gauge
	spectrum_node_mdisk_read_ms{id="1"} 0.251
	spectrum_node_mdisk_read_ms{id="10"} 0
	spectrum_node_mdisk_read_ms{id="2"} 0.297
	spectrum_node_mdisk_read_ms{id="8"} 0.146
	# HELP spectrum_node_mdisk_write_iops Current write I/O-per-second to mdisk
	# TYPE spectrum_node_mdisk_write_iops gauge
	spectrum_node_mdisk_write_iops{id="1"} 5
	spectrum_node_mdisk_write_iops{id="10"} 0
	spectrum_node_mdisk_write_iops{id="2"} 0
	spectrum_node_mdisk_write_iops{id="8"} 40
	# HELP spectrum_node_mdisk_write_mb Current Megabytes-per-second being written to mdisk
	# TYPE spectrum_node_mdisk_write_mb gauge
	spectrum_node_mdisk_write_mb{id="1"} 0
	spectrum_node_mdisk_write_mb{id="10"} 0
	spectrum_node_mdisk_write_mb{id="2"} 0
	spectrum_node_mdisk_write_mb{id="8"} 2.097152e+06
	# HELP spectrum_node_mdisk_write_ms Current milliseconds to write to mdisk
	# TYPE spectrum_node_mdisk_write_ms gauge
	spectrum_node_mdisk_write_ms{id="1"} 0.176
	spectrum_node_mdisk_write_ms{id="10"} 0
	spectrum_node_mdisk_write_ms{id="2"} 0.101
	spectrum_node_mdisk_write_ms{id="8"} 0.428
	# HELP spectrum_node_sas_bps Current bytes-per-second being transferred over backend SAS
	# TYPE spectrum_node_sas_bps gauge
	spectrum_node_sas_bps{id="1"} 0
	spectrum_node_sas_bps{id="10"} 0
	spectrum_node_sas_bps{id="2"} 0
	spectrum_node_sas_bps{id="8"} 0
	# HELP spectrum_node_sas_iops Current I/O-per-second being transferred over backend SAS
	# TYPE spectrum_node_sas_iops gauge
	spectrum_node_sas_iops{id="1"} 0
	spectrum_node_sas_iops{id="10"} 0
	spectrum_node_sas_iops{id="2"} 0
	spectrum_node_sas_iops{id="8"} 0
	# HELP spectrum_node_system_usage_ratio Current ratio of allocated CPU for system
	# TYPE spectrum_node_system_usage_ratio gauge
	spectrum_node_system_usage_ratio{id="1"} 0.04
	spectrum_node_system_usage_ratio{id="10"} 0.06
	spectrum_node_system_usage_ratio{id="2"} 0.03
	spectrum_node_system_usage_ratio{id="8"} 0.06
	# HELP spectrum_node_total_cache_usage_ratio Total percentage for both the write and read cache usage for the node
	# TYPE spectrum_node_total_cache_usage_ratio gauge
	spectrum_node_total_cache_usage_ratio{id="1"} 0.79
	spectrum_node_total_cache_usage_ratio{id="10"} 0.79
	spectrum_node_total_cache_usage_ratio{id="2"} 0.79
	spectrum_node_total_cache_usage_ratio{id="8"} 0.77
	# HELP spectrum_node_vdisk_read_iops Current read I/O-per-second to vdisk
	# TYPE spectrum_node_vdisk_read_iops gauge
	spectrum_node_vdisk_read_iops{id="1"} 12
	spectrum_node_vdisk_read_iops{id="10"} 0
	spectrum_node_vdisk_read_iops{id="2"} 19
	spectrum_node_vdisk_read_iops{id="8"} 6
	# HELP spectrum_node_vdisk_read_mb Current Megabytes-per-second being read from vdisk
	# TYPE spectrum_node_vdisk_read_mb gauge
	spectrum_node_vdisk_read_mb{id="1"} 0
	spectrum_node_vdisk_read_mb{id="10"} 0
	spectrum_node_vdisk_read_mb{id="2"} 0
	spectrum_node_vdisk_read_mb{id="8"} 0
	# HELP spectrum_node_vdisk_read_ms Current milliseconds to read from vdisk
	# TYPE spectrum_node_vdisk_read_ms gauge
	spectrum_node_vdisk_read_ms{id="1"} 0.427
	spectrum_node_vdisk_read_ms{id="10"} 0.059
	spectrum_node_vdisk_read_ms{id="2"} 0.067
	spectrum_node_vdisk_read_ms{id="8"} 0.056
	# HELP spectrum_node_vdisk_write_iops Current write I/O-per-second to vdisk
	# TYPE spectrum_node_vdisk_write_iops gauge
	spectrum_node_vdisk_write_iops{id="1"} 410
	spectrum_node_vdisk_write_iops{id="10"} 0
	spectrum_node_vdisk_write_iops{id="2"} 65
	spectrum_node_vdisk_write_iops{id="8"} 34
	# HELP spectrum_node_vdisk_write_mb Current Megabytes-per-second being written to vdisk
	# TYPE spectrum_node_vdisk_write_mb gauge
	spectrum_node_vdisk_write_mb{id="1"} 4.194304e+06
	spectrum_node_vdisk_write_mb{id="10"} 0
	spectrum_node_vdisk_write_mb{id="2"} 0
	spectrum_node_vdisk_write_mb{id="8"} 0
	# HELP spectrum_node_vdisk_write_ms Current milliseconds to write to vdisk
	# TYPE spectrum_node_vdisk_write_ms gauge
	spectrum_node_vdisk_write_ms{id="1"} 0.556
	spectrum_node_vdisk_write_ms{id="10"} 0.211
	spectrum_node_vdisk_write_ms{id="2"} 0.381
	spectrum_node_vdisk_write_ms{id="8"} 0.203
	# HELP spectrum_node_write_cache_usage_ratio Ratio of the write cache usage for the node
	# TYPE spectrum_node_write_cache_usage_ratio gauge
	spectrum_node_write_cache_usage_ratio{id="1"} 0.22
	spectrum_node_write_cache_usage_ratio{id="10"} 0.33
	spectrum_node_write_cache_usage_ratio{id="2"} 0.22
	spectrum_node_write_cache_usage_ratio{id="8"} 0.33
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestFCPorts_8_5_0_6_SVC(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsportfc", "testdata/specv_output_8_5_0_6_svc/lsportfc.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeFCPorts(c, r) {
		t.Errorf("probeFCPorts() returned non-success")
	}

	em := `
	# HELP spectrum_fc_port_speed_bps Operational speed of port in bits per second
	# TYPE spectrum_fc_port_speed_bps gauge
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="1",node_id="1"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="1",node_id="10"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="1",node_id="2"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="1",node_id="8"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="2",node_id="1"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="2",node_id="10"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="2",node_id="2"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="2",node_id="8"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="3",node_id="1"} 0
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="3",node_id="10"} 0
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="3",node_id="2"} 0
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="3",node_id="8"} 0
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="4",node_id="1"} 0
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="4",node_id="10"} 0
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="4",node_id="2"} 0
	spectrum_fc_port_speed_bps{adapter_location="1",adapter_port_id="4",node_id="8"} 0
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="1",node_id="1"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="1",node_id="10"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="1",node_id="2"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="1",node_id="8"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="2",node_id="1"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="2",node_id="10"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="2",node_id="2"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="2",node_id="8"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="3",node_id="1"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="3",node_id="10"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="3",node_id="2"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="3",node_id="8"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="4",node_id="1"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="4",node_id="10"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="4",node_id="2"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="2",adapter_port_id="4",node_id="8"} 8e+09
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="1",node_id="1"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="1",node_id="10"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="1",node_id="2"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="1",node_id="8"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="2",node_id="1"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="2",node_id="10"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="2",node_id="2"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="2",node_id="8"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="3",node_id="1"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="3",node_id="10"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="3",node_id="2"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="3",node_id="8"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="4",node_id="1"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="4",node_id="10"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="4",node_id="2"} 0
	spectrum_fc_port_speed_bps{adapter_location="5",adapter_port_id="4",node_id="8"} 0
	# HELP spectrum_fc_port_status Status of Fibre Channel port
	# TYPE spectrum_fc_port_status gauge
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="1",status="active",wwpn="500507680C1111C8"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="1",status="inactive_configured",wwpn="500507680C1111C8"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="1",status="inactive_unconfigured",wwpn="500507680C1111C8"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="10",status="active",wwpn="500507680C1111C5"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="10",status="inactive_configured",wwpn="500507680C1111C5"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="10",status="inactive_unconfigured",wwpn="500507680C1111C5"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="2",status="active",wwpn="500507680C1111C6"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="2",status="inactive_configured",wwpn="500507680C1111C6"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="2",status="inactive_unconfigured",wwpn="500507680C1111C6"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="8",status="active",wwpn="500507680C1111C7"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="8",status="inactive_configured",wwpn="500507680C1111C7"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="1",node_id="8",status="inactive_unconfigured",wwpn="500507680C1111C7"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="1",status="active",wwpn="500507680C1211C8"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="1",status="inactive_configured",wwpn="500507680C1211C8"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="1",status="inactive_unconfigured",wwpn="500507680C1211C8"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="10",status="active",wwpn="500507680C1211C5"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="10",status="inactive_configured",wwpn="500507680C1211C5"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="10",status="inactive_unconfigured",wwpn="500507680C1211C5"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="2",status="active",wwpn="500507680C1211C6"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="2",status="inactive_configured",wwpn="500507680C1211C6"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="2",status="inactive_unconfigured",wwpn="500507680C1211C6"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="8",status="active",wwpn="500507680C1211C7"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="8",status="inactive_configured",wwpn="500507680C1211C7"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="2",node_id="8",status="inactive_unconfigured",wwpn="500507680C1211C7"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="3",node_id="1",status="active",wwpn="500507680C1311C8"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="3",node_id="1",status="inactive_configured",wwpn="500507680C1311C8"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="3",node_id="1",status="inactive_unconfigured",wwpn="500507680C1311C8"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="3",node_id="10",status="active",wwpn="500507680C1311C5"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="3",node_id="10",status="inactive_configured",wwpn="500507680C1311C5"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="3",node_id="10",status="inactive_unconfigured",wwpn="500507680C1311C5"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="3",node_id="2",status="active",wwpn="500507680C1311C6"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="3",node_id="2",status="inactive_configured",wwpn="500507680C1311C6"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="3",node_id="2",status="inactive_unconfigured",wwpn="500507680C1311C6"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="3",node_id="8",status="active",wwpn="500507680C1311C7"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="3",node_id="8",status="inactive_configured",wwpn="500507680C1311C7"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="3",node_id="8",status="inactive_unconfigured",wwpn="500507680C1311C7"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="4",node_id="1",status="active",wwpn="500507680C1411C8"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="4",node_id="1",status="inactive_configured",wwpn="500507680C1411C8"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="4",node_id="1",status="inactive_unconfigured",wwpn="500507680C1411C8"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="4",node_id="10",status="active",wwpn="500507680C1411C5"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="4",node_id="10",status="inactive_configured",wwpn="500507680C1411C5"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="4",node_id="10",status="inactive_unconfigured",wwpn="500507680C1411C5"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="4",node_id="2",status="active",wwpn="500507680C1411C6"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="4",node_id="2",status="inactive_configured",wwpn="500507680C1411C6"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="4",node_id="2",status="inactive_unconfigured",wwpn="500507680C1411C6"} 1
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="4",node_id="8",status="active",wwpn="500507680C1411C7"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="4",node_id="8",status="inactive_configured",wwpn="500507680C1411C7"} 0
	spectrum_fc_port_status{adapter_location="1",adapter_port_id="4",node_id="8",status="inactive_unconfigured",wwpn="500507680C1411C7"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="1",status="active",wwpn="500507680C2111C8"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="1",status="inactive_configured",wwpn="500507680C2111C8"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="1",status="inactive_unconfigured",wwpn="500507680C2111C8"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="10",status="active",wwpn="500507680C2111C5"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="10",status="inactive_configured",wwpn="500507680C2111C5"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="10",status="inactive_unconfigured",wwpn="500507680C2111C5"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="2",status="active",wwpn="500507680C2111C6"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="2",status="inactive_configured",wwpn="500507680C2111C6"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="2",status="inactive_unconfigured",wwpn="500507680C2111C6"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="8",status="active",wwpn="500507680C2111C7"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="8",status="inactive_configured",wwpn="500507680C2111C7"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="1",node_id="8",status="inactive_unconfigured",wwpn="500507680C2111C7"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="1",status="active",wwpn="500507680C2211C8"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="1",status="inactive_configured",wwpn="500507680C2211C8"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="1",status="inactive_unconfigured",wwpn="500507680C2211C8"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="10",status="active",wwpn="500507680C2211C5"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="10",status="inactive_configured",wwpn="500507680C2211C5"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="10",status="inactive_unconfigured",wwpn="500507680C2211C5"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="2",status="active",wwpn="500507680C2211C6"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="2",status="inactive_configured",wwpn="500507680C2211C6"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="2",status="inactive_unconfigured",wwpn="500507680C2211C6"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="8",status="active",wwpn="500507680C2211C7"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="8",status="inactive_configured",wwpn="500507680C2211C7"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="2",node_id="8",status="inactive_unconfigured",wwpn="500507680C2211C7"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="3",node_id="1",status="active",wwpn="500507680C2311C8"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="3",node_id="1",status="inactive_configured",wwpn="500507680C2311C8"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="3",node_id="1",status="inactive_unconfigured",wwpn="500507680C2311C8"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="3",node_id="10",status="active",wwpn="500507680C2311C5"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="3",node_id="10",status="inactive_configured",wwpn="500507680C2311C5"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="3",node_id="10",status="inactive_unconfigured",wwpn="500507680C2311C5"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="3",node_id="2",status="active",wwpn="500507680C2311C6"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="3",node_id="2",status="inactive_configured",wwpn="500507680C2311C6"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="3",node_id="2",status="inactive_unconfigured",wwpn="500507680C2311C6"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="3",node_id="8",status="active",wwpn="500507680C2311C7"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="3",node_id="8",status="inactive_configured",wwpn="500507680C2311C7"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="3",node_id="8",status="inactive_unconfigured",wwpn="500507680C2311C7"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="4",node_id="1",status="active",wwpn="500507680C2411C8"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="4",node_id="1",status="inactive_configured",wwpn="500507680C2411C8"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="4",node_id="1",status="inactive_unconfigured",wwpn="500507680C2411C8"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="4",node_id="10",status="active",wwpn="500507680C2411C5"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="4",node_id="10",status="inactive_configured",wwpn="500507680C2411C5"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="4",node_id="10",status="inactive_unconfigured",wwpn="500507680C2411C5"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="4",node_id="2",status="active",wwpn="500507680C2411C6"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="4",node_id="2",status="inactive_configured",wwpn="500507680C2411C6"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="4",node_id="2",status="inactive_unconfigured",wwpn="500507680C2411C6"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="4",node_id="8",status="active",wwpn="500507680C2411C7"} 1
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="4",node_id="8",status="inactive_configured",wwpn="500507680C2411C7"} 0
	spectrum_fc_port_status{adapter_location="2",adapter_port_id="4",node_id="8",status="inactive_unconfigured",wwpn="500507680C2411C7"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="1",node_id="1",status="active",wwpn="500507680C5111C8"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="1",node_id="1",status="inactive_configured",wwpn="500507680C5111C8"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="1",node_id="1",status="inactive_unconfigured",wwpn="500507680C5111C8"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="1",node_id="10",status="active",wwpn="500507680C5111C5"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="1",node_id="10",status="inactive_configured",wwpn="500507680C5111C5"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="1",node_id="10",status="inactive_unconfigured",wwpn="500507680C5111C5"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="1",node_id="2",status="active",wwpn="500507680C5111C6"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="1",node_id="2",status="inactive_configured",wwpn="500507680C5111C6"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="1",node_id="2",status="inactive_unconfigured",wwpn="500507680C5111C6"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="1",node_id="8",status="active",wwpn="500507680C5111C7"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="1",node_id="8",status="inactive_configured",wwpn="500507680C5111C7"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="1",node_id="8",status="inactive_unconfigured",wwpn="500507680C5111C7"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="2",node_id="1",status="active",wwpn="500507680C5211C8"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="2",node_id="1",status="inactive_configured",wwpn="500507680C5211C8"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="2",node_id="1",status="inactive_unconfigured",wwpn="500507680C5211C8"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="2",node_id="10",status="active",wwpn="500507680C5211C5"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="2",node_id="10",status="inactive_configured",wwpn="500507680C5211C5"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="2",node_id="10",status="inactive_unconfigured",wwpn="500507680C5211C5"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="2",node_id="2",status="active",wwpn="500507680C5211C6"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="2",node_id="2",status="inactive_configured",wwpn="500507680C5211C6"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="2",node_id="2",status="inactive_unconfigured",wwpn="500507680C5211C6"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="2",node_id="8",status="active",wwpn="500507680C5211C7"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="2",node_id="8",status="inactive_configured",wwpn="500507680C5211C7"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="2",node_id="8",status="inactive_unconfigured",wwpn="500507680C5211C7"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="3",node_id="1",status="active",wwpn="500507680C5311C8"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="3",node_id="1",status="inactive_configured",wwpn="500507680C5311C8"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="3",node_id="1",status="inactive_unconfigured",wwpn="500507680C5311C8"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="3",node_id="10",status="active",wwpn="500507680C5311C5"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="3",node_id="10",status="inactive_configured",wwpn="500507680C5311C5"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="3",node_id="10",status="inactive_unconfigured",wwpn="500507680C5311C5"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="3",node_id="2",status="active",wwpn="500507680C5311C6"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="3",node_id="2",status="inactive_configured",wwpn="500507680C5311C6"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="3",node_id="2",status="inactive_unconfigured",wwpn="500507680C5311C6"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="3",node_id="8",status="active",wwpn="500507680C5311C7"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="3",node_id="8",status="inactive_configured",wwpn="500507680C5311C7"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="3",node_id="8",status="inactive_unconfigured",wwpn="500507680C5311C7"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="4",node_id="1",status="active",wwpn="500507680C5411C8"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="4",node_id="1",status="inactive_configured",wwpn="500507680C5411C8"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="4",node_id="1",status="inactive_unconfigured",wwpn="500507680C5411C8"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="4",node_id="10",status="active",wwpn="500507680C5411C5"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="4",node_id="10",status="inactive_configured",wwpn="500507680C5411C5"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="4",node_id="10",status="inactive_unconfigured",wwpn="500507680C5411C5"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="4",node_id="2",status="active",wwpn="500507680C5411C6"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="4",node_id="2",status="inactive_configured",wwpn="500507680C5411C6"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="4",node_id="2",status="inactive_unconfigured",wwpn="500507680C5411C6"} 1
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="4",node_id="8",status="active",wwpn="500507680C5411C7"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="4",node_id="8",status="inactive_configured",wwpn="500507680C5411C7"} 0
	spectrum_fc_port_status{adapter_location="5",adapter_port_id="4",node_id="8",status="inactive_unconfigured",wwpn="500507680C5411C7"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestQuorumStatus_8_5_0_6_SVC(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lsquorum", "testdata/specv_output_8_5_0_6_svc/lsquorum.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeQuorum(c, r) {
		t.Errorf("probeQuorumStatus() returned non-success")
	}

	em := `
	# HELP spectrum_quorum_status Status of quorum
	# TYPE spectrum_quorum_status gauge
	spectrum_quorum_status{id="0",object_type="mdisk",status="offline"} 0
	spectrum_quorum_status{id="0",object_type="mdisk",status="online"} 1
	spectrum_quorum_status{id="1",object_type="mdisk",status="offline"} 0
	spectrum_quorum_status{id="1",object_type="mdisk",status="online"} 1
	spectrum_quorum_status{id="3",object_type="device",status="offline"} 0
	spectrum_quorum_status{id="3",object_type="device",status="online"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}

func TestHostStatus_8_5_0_6_SVC(t *testing.T) {
	c := newFakeClient()
	c.prepare("rest/v1/lshost", "testdata/specv_output_8_5_0_6_svc/lshost.jsonnet")
	r := prometheus.NewPedanticRegistry()
	if !probeHost(c, r) {
		t.Errorf("probeHostStatus() returned non-success")
	}

	em := `
	# HELP spectrum_host_status Status of hosts
	# TYPE spectrum_host_status gauge
	spectrum_host_status{hostname="Not_IN_Use_VMWARE01",id="0",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="Not_IN_Use_VMWARE01",id="0",port_count="2",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="Not_IN_Use_VMWARE01",id="0",port_count="2",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="Not_in_use_VMWARE02",id="1",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="Not_in_use_VMWARE02",id="1",port_count="2",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="Not_in_use_VMWARE02",id="1",port_count="2",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="Not_in_use_VMWARE03",id="2",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="Not_in_use_VMWARE03",id="2",port_count="2",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="Not_in_use_VMWARE03",id="2",port_count="2",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="Not_in_use_VMWARE04",id="3",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="Not_in_use_VMWARE04",id="3",port_count="2",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="Not_in_use_VMWARE04",id="3",port_count="2",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="P814_VIOS1",id="14",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="P814_VIOS1",id="14",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="P814_VIOS1",id="14",port_count="2",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="P814_VIOS2",id="13",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="P814_VIOS2",id="13",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="P814_VIOS2",id="13",port_count="2",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="VMWARE07",id="4",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="VMWARE07",id="4",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="VMWARE07",id="4",port_count="2",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="VMWARE08",id="5",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="VMWARE08",id="5",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="VMWARE08",id="5",port_count="2",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="VMWARE09",id="6",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="VMWARE09",id="6",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="VMWARE09",id="6",port_count="2",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="VMWARE10",id="7",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="VMWARE10",id="7",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="VMWARE10",id="7",port_count="2",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="infratsm02-09b178fd-00000073-87308356",id="15",port_count="4",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="infratsm02-09b178fd-00000073-87308356",id="15",port_count="4",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="infratsm02-09b178fd-00000073-87308356",id="15",port_count="4",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="p822_1_vios1",id="33",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="p822_1_vios1",id="33",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="p822_1_vios1",id="33",port_count="2",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="p822_1_vios2",id="34",port_count="2",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="p822_1_vios2",id="34",port_count="2",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="p822_1_vios2",id="34",port_count="2",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="powervc01",id="17",port_count="4",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="powervc01",id="17",port_count="4",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="powervc01",id="17",port_count="4",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="powervc01-6dbf61b7-0000009b-55993328",id="23",port_count="8",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="powervc01-6dbf61b7-0000009b-55993328",id="23",port_count="8",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="powervc01-6dbf61b7-0000009b-55993328",id="23",port_count="8",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="powervc02-eef93991-00000091-50278340",id="24",port_count="8",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="powervc02-eef93991-00000091-50278340",id="24",port_count="8",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="powervc02-eef93991-00000091-50278340",id="24",port_count="8",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="powervc03-dcd7f415-00000092-58388744",id="25",port_count="8",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="powervc03-dcd7f415-00000092-58388744",id="25",port_count="8",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="powervc03-dcd7f415-00000092-58388744",id="25",port_count="8",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="pvc_prep_sles-4ce9426a-00000097-84661545",id="20",port_count="8",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="pvc_prep_sles-4ce9426a-00000097-84661545",id="20",port_count="8",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="pvc_prep_sles-4ce9426a-00000097-84661545",id="20",port_count="8",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="pvc_testvm01-63540488",id="27",port_count="4",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="pvc_testvm01-63540488",id="27",port_count="4",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="pvc_testvm01-63540488",id="27",port_count="4",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="pvc_testvm01-89886324",id="21",port_count="4",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="pvc_testvm01-89886324",id="21",port_count="4",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="pvc_testvm01-89886324",id="21",port_count="4",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="pvc_testvm02-08436546",id="22",port_count="8",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="pvc_testvm02-08436546",id="22",port_count="8",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="pvc_testvm02-08436546",id="22",port_count="8",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="pvc_testvm03-41667268",id="26",port_count="4",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="pvc_testvm03-41667268",id="26",port_count="4",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="pvc_testvm03-41667268",id="26",port_count="4",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="pvc_testvm04-62314584",id="8",port_count="8",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="pvc_testvm04-62314584",id="8",port_count="8",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="pvc_testvm04-62314584",id="8",port_count="8",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="pvc_testvm20-88396381",id="28",port_count="8",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="pvc_testvm20-88396381",id="28",port_count="8",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="pvc_testvm20-88396381",id="28",port_count="8",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="rhel83-test-06429fc6-00000003-61056187",id="18",port_count="8",protocol="scsi",status="degraded"} 1
	spectrum_host_status{hostname="rhel83-test-06429fc6-00000003-61056187",id="18",port_count="8",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="rhel83-test-06429fc6-00000003-61056187",id="18",port_count="8",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="scale-test-st-be9c48c7-0000007c-98521743",id="19",port_count="8",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="scale-test-st-be9c48c7-0000007c-98521743",id="19",port_count="8",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="scale-test-st-be9c48c7-0000007c-98521743",id="19",port_count="8",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="sles15sp3-pre-58e27652-00000093-88592693",id="16",port_count="8",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="sles15sp3-pre-58e27652-00000093-88592693",id="16",port_count="8",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="sles15sp3-pre-58e27652-00000093-88592693",id="16",port_count="8",protocol="scsi",status="online"} 1
	spectrum_host_status{hostname="templa_RHEL77-8f099e81-00000072-15342990",id="10",port_count="4",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="templa_RHEL77-8f099e81-00000072-15342990",id="10",port_count="4",protocol="scsi",status="offline"} 1
	spectrum_host_status{hostname="templa_RHEL77-8f099e81-00000072-15342990",id="10",port_count="4",protocol="scsi",status="online"} 0
	spectrum_host_status{hostname="testvm-powerv-1ac93c35-0000000f-13266211",id="32",port_count="4",protocol="scsi",status="degraded"} 0
	spectrum_host_status{hostname="testvm-powerv-1ac93c35-0000000f-13266211",id="32",port_count="4",protocol="scsi",status="offline"} 0
	spectrum_host_status{hostname="testvm-powerv-1ac93c35-0000000f-13266211",id="32",port_count="4",protocol="scsi",status="online"} 1
	`

	if err := testutil.GatherAndCompare(r, strings.NewReader(em)); err != nil {
		t.Fatalf("metric compare: err %v", err)
	}
}
