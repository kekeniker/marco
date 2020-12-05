# Marco

> Marco is a CLI tool for showing Spinnaker application, pipelines and pipeline templates

## Usecase

With `marco`, you can check these items organizantion-wide:

* All Spinnaker applications name 
* Pipeline and stages statistics

## Install

Install by this command.

```console
$ brew install kekeniker/tap/marco
```

## Configuration

`marco` uses the same configuration file with [spinnaker/spin](https://github.com/spinnaker/spin), the Spinnaker CLI.

See details with [Configure spin](https://www.spinnaker.io/setup/spin/#configure-spin). `marco` also shares the flags with `spin` for example `--gate-endpoint`, etc.
You don't have to learn anything.

## Usage

### Applications

```console
$ marco app list
+------------------------------+-------------------------------+-----------------+---------------+
|             NAME             |             EMAIL             | CLOUD PROVIDERS | INSTANCE PORT |
+------------------------------+-------------------------------+-----------------+---------------+
| app                          |                               | aws             |            80 |
+------------------------------+-------------------------------+-----------------+---------------+
| authority                    | marco@marco.com               | kubernetes      |            80 |
+------------------------------+-------------------------------+-----------------+---------------+
| marco                        | marco@marco.com               | kubernetes      |            80 |
+------------------------------+-------------------------------+-----------------+---------------+
```

You can check the name validation for each cloud provider with `--validate` option.

```console
$ marco app list  --validate
+------------------------------+-------------------------------+-----------------+---------------+-----------+-----+------+-----+---------------+---------------+-----------+--------------+-------+
|             NAME             |             EMAIL             | CLOUD PROVIDERS | INSTANCE PORT | APPENGINE | AWS | DCOS | GCE | KUBERNETES V1 | KUBERNETES V2 | OPENSTACK | TENCENTCLOUD | TITUS |
+------------------------------+-------------------------------+-----------------+---------------+-----------+-----+------+-----+---------------+---------------+-----------+--------------+-------+
| app                          |                               | aws             |            80 | ✅        | ✅  | ✅   | ✅  | ✅            | ✅            | ✅        | ✅           | ✅    |
+------------------------------+-------------------------------+-----------------+---------------+-----------+-----+------+-----+---------------+---------------+-----------+--------------+-------+
| authority                    | marco@marco.com               | kubernetes      |            80 | ✅        | ✅  | ✅   | ✅  | ✅            | ✅            | ✅        | ✅           | ✅    |
+------------------------------+-------------------------------+-----------------+---------------+-----------+-----+------+-----+---------------+---------------+-----------+--------------+-------+
| authority                    | masa@mercari.com              | kubernetes      |            80 | ✅        | ✅  | ✅   | ✅  | ✅            | ✅            | ✅        | ✅           | ✅    |
+------------------------------+-------------------------------+-----------------+---------------+-----------+-----+------+-----+---------------+---------------+-----------+--------------+-------+
```

### Pipelines

Specify the Spinnaker application by `-a` or `--application` .

```console
$ marco pipeline list -a "marco"
+----------+--------------------------------------+-----------------------------+----------------------------------+-------------+-------------------+----------------+
|   APP    |             PIPELINE ID              |            NAME             |             TEMPLATE             | STAGE REFID |       STAGE       |      TYPE      |
+----------+--------------------------------------+-----------------------------+----------------------------------+-------------+-------------------+----------------+
|  marco   | 1d733082-5368-475e-b8ud-cde5bcfe9eec | Trigger every minute        | (none)                           |           1 | Wait              | wait           |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           2 | Bake (Manifest)   | bakeManifest   |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           3 | Canary Analysis   | kayentaCanary  |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           4 | Disable Cluster   | disableCluster |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           5 | Deploy            | deploy         |
+          +--------------------------------------+-----------------------------+                                  +-------------+-------------------+----------------+
|          | 9b3c34c7-6f5a-5773-8a44-34f0e098f3e2 | Trigger every minite v2     |                                  |           1 | Wait              | wait           |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           2 | Bake (Manifest)   | bakeManifest   |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           3 | Canary Analysis   | kayentaCanary  |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           4 | Disable Cluster   | disableCluster |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           5 | Deploy            | deploy         |
+          +--------------------------------------+-----------------------------+----------------------------------+-------------+-------------------+----------------+
|          | ae74405e-942e-4p13-882f-3d783cbc7f10 | Test from Pipeline Template | deployWithManualJudgement:latest |           1 | Manual Judgment   | manualJudgment |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           2 | Deploy (Manifest) | deployManifest |
+----------+--------------------------------------+-----------------------------+----------------------------------+-------------+-------------------+----------------+
```

Use `--all` flag to see the pipelines for all applications. Also, if you don't want to expand the pipeline template use `--expand=false`.

```console
$ marco pipeline list -a "marco" --expand=false
+----------+--------------------------------------+-----------------------------+----------------------------------+-------------+-------------------+----------------+
|   APP    |             PIPELINE ID              |            NAME             |             TEMPLATE             | STAGE REFID |       STAGE       |      TYPE      |
+----------+--------------------------------------+-----------------------------+----------------------------------+-------------+-------------------+----------------+
|  marco   | 1d733082-5368-475e-b8ud-cde5bcfe9eec | Trigger every minute        | (none)                           |           1 | Wait              | wait           |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           2 | Bake (Manifest)   | bakeManifest   |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           3 | Canary Analysis   | kayentaCanary  |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           4 | Disable Cluster   | disableCluster |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           5 | Deploy            | deploy         |
+          +--------------------------------------+-----------------------------+                                  +-------------+-------------------+----------------+
|          | 9b3c34c7-6f5a-5773-8a44-34f0e098f3e2 | Trigger every minite v2     |                                  |           1 | Wait              | wait           |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           2 | Bake (Manifest)   | bakeManifest   |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           3 | Canary Analysis   | kayentaCanary  |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           4 | Disable Cluster   | disableCluster |
+          +                                      +                             +                                  +-------------+-------------------+----------------+
|          |                                      |                             |                                  |           5 | Deploy            | deploy         |
+          +--------------------------------------+-----------------------------+----------------------------------+-------------+-------------------+----------------+
|          | ae74405e-942e-4p13-882f-3d783cbc7f10 | Test from Pipeline Template | deployWithManualJudgement:latest |  (template) | (template)        | (template)     |
+----------+--------------------------------------+-----------------------------+----------------------------------+-------------+-------------------+----------------+
```

### Pipeline templates

```console
$ marco pipeline-template list
+--------------------------------------------+-----------+
|                     ID                     | PROTECTED |
+--------------------------------------------+-----------+
| deployWithManualJudgementWithDocker:latest | true      |
+--------------------------------------------+-----------+
| deployWithManualJudgementWithDocker:latest | true      |
+--------------------------------------------+-----------+
| deployWithManualJudgementWithDocker:latest | true      |
+--------------------------------------------+-----------+
| deployWithManualJudgement:latest           | true      |
+--------------------------------------------+-----------+
| deployWithManualJudgementWithDocker:latest | true      |
+--------------------------------------------+-----------+
| deployWithManualJudgementWithDocker:       | true      |
+--------------------------------------------+-----------+
| deployWithManualJudgement:latest           | true      |
+--------------------------------------------+-----------+
| deployWithManualJudgement:                 | true      |
+--------------------------------------------+-----------+
```

## Contribution

Please create a GitHub Issue or a pull request.
I welcome all contributions.

## Author

* [KeisukeYamashita](https://github.com/KeisukeYamashita)

## Licence

Copyright 2020 KeisukeYamashita. marco is released under the Apache License 2.0.
