# Go base Alfred workflow generator

![GitHub](https://img.shields.io/github/license/cage1016/ak)

A generator for golang Alfred workflow that helps you create boilerplate code.

## Why ?

As a Gopher, i want to create a Golang base workflow quickly than sketching it out by hand. I also want to be able to share my way of doing things with others.

## Features

1. Create a new workflow with three type of patterns:
   1. Simple workflow with Alfred variables and arguments
   2. leverage [deanishe/awgo](https://github.com/deanishe/awgo) with Alfred items feedback
   3. leverage [deanishe/awgo](https://github.com/deanishe/awgo) and [spf13/cobra](https://github.com/spf13/cobra) with Alfred items feedback
2. Workflow development
   1. Build the workflow executable and output it into the `.workflow` subdirectory
   2. Display information about the workflow
   3. Link the `.workflow` subdirectory into Alfred's preferences directory, installing it.
   4. Package the workflow for distribution locally
   5. Unlink the `.workflow` subdirectory from Alfred's preferences directory, uninstalling it.
3. Additional patterns
   1. Add Github Action release to project. `-s` enable code sign and notarize
   1. Add license to project
   2. Add Makefile to project
4. Support `arm64` & `amd64`

## Installation

With Go 1.20 or higher:

```bash
go install github.com/cage1016/ak@latest
```

## Running the generator

```bash
A generator for awgo that helps you create boilerplate code

Usage:
  ak [flags]
  ak [command]

Available Commands:
  add         Used to add additional component to project
  alfred      Used to manage Go-based Alfred workflows
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initiates a workflow
  new         Used to create workflow package

Flags:
  -d, --debug           If you want to se the debug logs.
      --folder string   If you want to specify the base folder of the workflow.
  -f, --force           Force overwrite existing files without asking.
  -h, --help            help for ak
      --testing         If testing the generator.
  -v, --version         version for ak

Use "ak [command] --help" for more information about a command.
```

1. `ak init` to create a new workflow
2. revise `ak.json` as your workflow's information

    ```json
    {
        "go_mod_package": "github.com/xxx/ak-test",
        "workflow": {
            "folder": ".workflow",
            "name": "Ak Test",
            "description": "",
            "category_comment": "category: Tools, Internet, Productivity, Uncategorised",
            "category": "",
            "bundle_id": "com.xxx.aktest",
            "created_by": "",
            "web_address": "https://github.com/xxx/ak-test",
            "version": "0.1.0"
        },
        "update": {
            "github_repo": "https://github.com/xxx/ak-test"
        },
        "license": {
            "type_comment": "support license https://github.com/nishanths/license",
            "type": "mit",
            "year": "",
            "name": ""
        },
        "gon": {
            "application_identity": "Developer ID Application: KAI CHU CHUNG"
        }
    }    
    ```

4. Create one of three different workflow patterns 
   1. `ak new cmd` create cobra based workflow
5. Add additional components to the workflow
   1. `ak add githubAction` add Github Action release to project
      1. `-s` to enable code sign and notarize
   2. `ak add license` add license to project
   3. `ak add makefile` add Makefile to project
6. Workflow development
   1. `make build` to build the workflow executable and output it into the `.workflow` subdirectory
   2. `make info` to display information about the workflow
   3. `make link` to link the `.workflow` subdirectory into Alfred's preferences directory, installing it.
   4. `make package` to package the workflow for distribution locally
   5. `make unlink` to unlink the `.workflow` subdirectory from Alfred's preferences directory, uninstalling it.

## Examples

1. [ak/_example](https://github.com/cage1016/ak/tree/master/_example)

## License
This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.