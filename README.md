# ods-pipeline-gradle

[![Tests](https://github.com/opendevstack/ods-pipeline-gradle/actions/workflows/main.yaml/badge.svg)](https://github.com/opendevstack/ods-pipeline-gradle/actions/workflows/main.yaml)

Tekton task for use with [ODS Pipeline](https://github.com/opendevstack/ods-pipeline) to build Gradle applications and libraries.

## Usage

```yaml
tasks:
- name: build
  taskRef:
    resolver: git
    params:
    - { name: url, value: https://github.com/opendevstack/ods-pipeline-gradle.git }
    - { name: revision, value: v0.1.0 }
    - { name: pathInRepo, value: tasks/build.yaml }
    workspaces:
    - { name: source, workspace: shared-workspace }
```

See the [documentation](https://github.com/opendevstack/ods-pipeline-gradle/blob/main/docs/build.adoc) for details and available parameters.

**TIP:** If you need a database to run alongside your tests, you can use the [`tasks/build-with-postgres.yaml`](https://github.com/opendevstack/ods-pipeline-gradle/blob/main/docs/build-with-postgres.adoc) variant of this task!

## About this repository

`docs` and `tasks` are generated directories from recipes located in `build`. See the `Makefile` target for how everything fits together.
