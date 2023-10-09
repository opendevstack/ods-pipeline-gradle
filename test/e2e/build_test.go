package e2e

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	ott "github.com/opendevstack/ods-pipeline/pkg/odstasktest"
	"github.com/opendevstack/ods-pipeline/pkg/pipelinectxt"
	ttr "github.com/opendevstack/ods-pipeline/pkg/tektontaskrun"
	tekton "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

func TestBuildGradleTask(t *testing.T) {
	if err := runTask(
		ttr.WithStringParams(map[string]string{
			"cache-build": "false",
		}),
		ott.WithGitSourceWorkspace(t, "../testdata/workspaces/gradle-sample-app", namespaceConfig.Name),
		ttr.AfterRun(func(config *ttr.TaskRunConfig, run *tekton.TaskRun, logs bytes.Buffer) {
			wd := config.WorkspaceConfigs["source"].Dir

			ott.AssertFilesExist(t, wd,
				"docker/Dockerfile",
				"docker/app.jar",
				filepath.Join(pipelinectxt.XUnitReportsPath, "TEST-ods.java.gradle.sample.app.AppTest.xml"),
				filepath.Join(pipelinectxt.XUnitReportsPath, "TEST-ods.java.gradle.sample.app.AppTest2.xml"),
				filepath.Join(pipelinectxt.CodeCoveragesPath, "coverage.xml"),
			)

			logContains(t, logs,
				"ods-test-nexus",
				"Gradle 7.4.2",
				"Using GRADLE_OPTS=-Dorg.gradle.jvmargs=-Xmx512M",
				"Using GRADLE_USER_HOME=/workspace/source/.ods-cache/deps/gradle",
				"To honour the JVM settings for this build a single-use Daemon process will be forked.",
			)

		}),
	); err != nil {
		t.Fatal(err)
	}
}

func logContains(t *testing.T, logs bytes.Buffer, wantLogMsgs ...string) {
	t.Helper()
	for _, msg := range wantLogMsgs {
		if !strings.Contains(logs.String(), msg) {
			t.Fatalf("Want:\n%s\n\nGot:\n%s", msg, logs.String())
		}
	}
}
