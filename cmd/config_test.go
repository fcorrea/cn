/*
 * Ceph Nano (C) 2018 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * Below main package has canonical imports for 'go get' and 'go build'
 * to work with all other clones of github.com/ceph/cn repository. For
 * more information refer https://golang.org/doc/go1.4#canonicalimports
 */

package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var configFile = "cn-test.toml"

func TestTitle(t *testing.T) {
	readConfigFile(configFile)
	assert.Equal(t, "Ceph Nano test configuration file", viper.Get("title"))
}

func TestMemorySize(t *testing.T) {
	// Testing the builtin configuration
	assert.Equal(t, "512MB", getMemorySize("default"))
	assert.Equal(t, "4GB", getMemorySize("huge"))

	// Testing with a configuration file
	readConfigFile(configFile)
	assert.Equal(t, "512MB", getMemorySize("test_nano_default"))
	assert.Equal(t, int64(536870912), getMemorySizeInBytes("test_nano_default"))
	assert.Equal(t, "1GB", getMemorySize("test_nano_no_default"))
	assert.Equal(t, int64(1073741824), getMemorySizeInBytes("test_nano_no_default"))
}

func TestUseDefault(t *testing.T) {
	readConfigFile(configFile)
	assert.Equal(t, false, useDefault(FLAVORS, "test_nano_no_default"))
	assert.Equal(t, true, useDefault(FLAVORS, "test_nano_default"))
}

func TestCephConf(t *testing.T) {
	readConfigFile(configFile)
	assert.Equal(t, map[string]interface{}{"osd_memory_target": int64(3841234556)}, getCephConf("test_nano_no_default"))
	expectedOutput := map[string]interface{}{
		"bluestore_cache_autotune_chunk_size": int64(8388608),
		"osd_max_pg_log_entries":              int64(10),
		"osd_memory_base":                     int64(268435456),
		"osd_memory_cache_min":                int64(33554432),
		"osd_memory_target":                   int64(3841234556),
		"osd_min_pg_log_entries":              int64(10),
		"osd_pg_log_dups_tracked":             int64(10),
		"osd_pg_log_trim_min":                 int64(10),
	}
	assert.Equal(t, expectedOutput, getCephConf("test_nano_default"))
}

func TestCPUCount(t *testing.T) {
	// Testing the builtin configuration
	assert.Equal(t, int64(1), getCPUCount("default"))
	assert.Equal(t, int64(2), getCPUCount("huge"))

	// Testing with a configuration file
	readConfigFile(configFile)
	assert.Equal(t, int64(1), getCPUCount("test_nano_default"))
	assert.Equal(t, int64(2), getCPUCount("test_nano_no_default"))
}

func TestIsEntryExist(t *testing.T) {
	assert.Equal(t, true, isEntryExists(FLAVORS, "default.use_default"))
	assert.Equal(t, false, isEntryExists(FLAVORS, "default.nawak"))
}

func TestImageName(t *testing.T) {
	// There is no configuration file
	configurationFile = ""

	// Without any configuration file, the default should be satisfied
	assert.Equal(t, DEFAULTIMAGE, getImageName())

	// Without any configuration file, any -i argument should be preserved
	imageName = "nawak"
	assert.Equal(t, "nawak", getImageName())

	// The default builtin should be kept too
	imageName = "mimic"
	assert.Equal(t, LATESTIMAGE+"mimic", getImageName())

	// Now, we have a configuration file
	configurationFile = readConfigFile(configFile)

	// Let's ensure the basic reading of the configuration file works
	assert.Equal(t, "ceph/daemon:latest-real1", getImageNameFromConfig("real1"))

	// If a -i is passed with a configuration file, let's report the image_name from the configuration file
	imageName = "complex"
	assert.Equal(t, "this.url.is.complex/cool/for-a-test", getImageName())
}
