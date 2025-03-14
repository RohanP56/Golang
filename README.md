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