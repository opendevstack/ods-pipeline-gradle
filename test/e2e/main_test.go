package e2e

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	ott "github.com/opendevstack/ods-pipeline/pkg/odstasktest"
	ttr "github.com/opendevstack/ods-pipeline/pkg/tektontaskrun"
)

var (
	namespaceConfig *ttr.NamespaceConfig
	rootPath        = "../.."
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	cc, err := ttr.StartKinDCluster(
		ttr.LoadImage(ttr.ImageBuildConfig{
			Dockerfile: "build/images/Dockerfile.gradle-toolset",
			ContextDir: rootPath,
		}),
	)
	if err != nil {
		log.Fatal("Could not start KinD cluster: ", err)
	}
	nc, cleanup, err := ttr.SetupTempNamespace(
		cc,
		ott.StartNexus(),
		ott.InstallODSPipeline(nil),
		ttr.InstallTaskFromPath(
			filepath.Join(rootPath, "build/tasks/build.yaml"),
			nil,
		),
		ttr.InstallTaskFromPath(
			filepath.Join(rootPath, "build/tasks/build.yaml"),
			map[string]string{
				"PostgresSidecar": "true",
			},
		),
	)
	if err != nil {
		log.Fatal("Could not setup temporary namespace: ", err)
	}
	defer cleanup()
	namespaceConfig = nc
	return m.Run()
}
