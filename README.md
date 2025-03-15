# GOLANG Learning
**Go is a Strong and statically typed language.**

### Concurrency: 
refers to the ability of a system to handle multiple tasks or processes seemingly at the same time, even if they are not truly running simultaneously, improving responsiveness and resource utilization. 

### Parallelism:
It's technique where a problem is broken down into smaller, independent tasks that can be executed simultaneously by multiple processors or cores to achieve faster and more efficient computation. 

## Go Folder Structure: 
- **HOME/**
    - go/
        - bin/
            - (executable binaries)
        - pkg/
            - (compiled package files)
        - src/
            - myproject/
                - main.go
                - other files
#### HOME/go/src: This ensures that your project is a go project.

## Go Modules: 
If we want to create go projects outside the default "GOPATH", you can create using **Go Modules**, It allows us to create and manage projects outside of the GOPATH.

### How to initialize go module:
**go mod init**

### How to execute Go file?
**go run (file_name with extention)**

- In Go, package are used insted of classes, There are no concepts like OOPS in java.
- Each package in Go is essentially a directory in workspace. Each go file must belong to some package, and it should start with the keyword  **package** followed by package name.
- The main package is a special package in Go. An executable program must contain a main package.
- Goi used relative imports to bring package into current file. The relative path usually **"GOPATH/src"** since most package are stored in **pkg** directory.
- **main** function should be a part of package **main**.