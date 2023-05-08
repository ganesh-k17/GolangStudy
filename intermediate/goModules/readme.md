# Module and package

## Library management

Library management is goes around below three concepts in Golang

- Repository (place)
- Modules
- Packages

## Modules

- A module is the root of a Gol library or application stored in a repository.
- A module consist of one or more package which give the module organization and structure.
- While we can store more than one module in a repository it is not encouraged.
- Maintaining two modules in one repository means tracking separate versions for two different projects in a single repository.
- We need to make sure that we have declared that the project is a module.
- Every module has globally unique identifier.
- Basically, if we publish a local module, this would be the path from which a module could be downloaded by Go tools.
- This could be a GitHub repository, or a private repository on a domain or anything but globally unique.

```text
- To Summarize, Go code is grouped int packages and packages are grouped into modules.
- A module specifies the dependencies needed to run our code, including the Go version and the set of other modules it requires in the go.mod file.
```

Initializing as module,

```bash
go mod init <module-path>

go mod  init example.com/learn  # example
```

## go.mod file

- Is a collectionof Go source code becomes a module when this is a valid go.mod file in its root directory.
- It consists of,
  1. the module declaration,
  2. the minimum compatible version for Go, and
  3. the dependencies for the imported third party packages.

## Creating and accessing package

_Import and export_:

If a function which is mentioned with caps letter then the function can be accessed outside the package. (We can assume it as public function.)

If a function name started with small letter then it cannot be accessed out side the package but only inside the package.

```bash
go mod tidy # used to download the packages which are used inside the code.
# so when a package is published and when we run <go mod tidy> then it would download the package.

go mod edit -replace github.com/arithmetic=../arithmetic
# if a package is unpublished the we have to mention the local folder path where the package should be referred from.  For this we use replace command
```

using replace key word we can mention to redirect to my local folder to get the package.

## golang commands

### go mod init

- this one initializes and writes a new go.mod file in the current directory.
- In effect creating a new module created at the current directory
- It accepts one optional argument which is the module path for the new module.

### go mod tidy

- The command ensures that the go.mod file matches the source code in the module.
- It adds any missing module requirements necessary to build the current module's packages and dependencies.
- and it removes requirements on modules that don't provide any relevent pacakges.

### go run <-file->

- The command compiles and run the program.
- Internally it,
  1. Compiles your code and builds an executable binary in a temporary location
  2. Launches the temp exe file and
  3. Finally cleans it when your app exists.
- This command is useful for testing small programs during the initial development stage of our application.

### go build

- This one build compiles and create the executable in current directory.

### go install

- Compile and
- move the executable to $GOPATH/bin
- run this executable from any path on the terminal

### go env GOPATH

- used to set the GO path

### go get

- Resolves its command line arguments to packages at specific module version
- updates go.mod to require those versions
- and download source code into the module cache
- add dependency for a package or update it to its latest version

  `go get example.com/pkg `

- To upgrade or downgrade a package to a specific version

  `go get example.com/pkg@v1.2.3`

- To find out gopath

  ```shell
  go env GOPATH
  # /users/ganesh/go
  ```

## Developing and publishing module

### Workflow

- Design and code the packages that the module will include.
- Commit code to your repository using convention that ensure its available to others via go tools.
- Publish the module to make it discoverable by developers.
- Over time, revise the module with versions that use a version number in convention that signals each version's stability and backward compatibility.

### Design and development

- When you are designing a module's public API, try to keep its functionality focused and discrete.
- Ensure backward compatibility to help users upgrade while minimizing churn to their own code.

### Decentralized publishing

- In go , your publish your module by tagging its code in your repository to make it available for other developers to use.
- After importing your package in their code developers use go tools (including the go get command) to download your module's code to compile with.

### Package discovery

- After we have published the module and someone has fetched it with Go tools, it will become visible on the Go package discovery site at pkg.go.dev
- There developers can search the site to find it and read its documentation

### Versioning

- As you revise and improve our module over time, you assign version numbers
- Designed to signal each version's stability and backward compatibility
- you indicate a module's version number by tagging the module's source in the repository with the number

### Module version Name confusion

- For each nuw releases, a module's release version number specifically reflects the nature of the module's changes since the preceading release.
- Versioning

  ```bash
  <major>.<minor>.<patch> - <pre-release-milestone>
  # eg: 1.4.0 - beta.2
  # major: does not guarantee backward compatibility
  # minor: backward compatible public changes
  # patch: a change does not affect any public  changes and dependencies.
  # pre-release mile stone:(alpha/beta): no stability guarantee
  ```

### Publishing a module

publishing it in GitHub

```shell
go init
git remote add origin <url>
git remote -v
git status
git add .
git commit -m "my changes"
git tag v1.0.0
git push origin v1.0.0
git pull --allow-unrelated-histories
```

Users can download the source code by the created tag,

```shell
go get github.com/ganesh/<module-name>@<tagged version>
```

### godoc

- In golang we used the go doc tool to create documentation
- Godoc parses go source code including comments and produces documentation as HTML or plain text
- The documentation created tightly coupled with the code its document
- The convention is simple to document a type, variable , constant, function or even for package.
- Write a proper comment directly preceding its declaration with no intervening blank line
- Use a blank comment to break your comment into multiple paragraphs
- If you have lengthy comments for the package, the convention to put the comment in a file to your package called doc.go

  ```shell
  go doc <PACKAGE-NAME>
  GO DOC <PACKAGE-NAME>.<IDENTIFIER-NAME>
  ```

### Package name collision

When we happen to import two package with same name we can use alias for a package.

```go
import (
    crand "crypto/rand"  // used alias
    "encoding/binary"
    "math/rand"
)
```

### Naming packages:

- Package name should be descriptive.
- Avoid repeating the name of the package in the names of functions and type within the package.
- Every Go file in a directory must have an identical package clause.
- As a rule, you should make the name of the package match the name of the directory that contains the package because it is hard to discover a package's name if it does not match the containing directory
