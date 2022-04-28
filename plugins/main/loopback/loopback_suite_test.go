// Copyright 2016 CNI authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package loopback_test

import (
	"os"
	"path/filepath"

	"github.com/onsi/gomega/gexec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var pathToLoPlugin string

func TestLoopback(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "plugins/main/loopback")
}

var _ = BeforeSuite(func() {
	var err error
	var pathToPlugins string
	pathToPlugins, err = gexec.Build("github.com/containernetworking/plugins/launcher/linux")
	pathToLoPlugin = filepath.Join(filepath.Dir(pathToPlugins), "loopback")
	os.Symlink(pathToPlugins, pathToLoPlugin)

	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
