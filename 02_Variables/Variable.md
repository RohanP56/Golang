# Variable in Go
- To declare Variable in "Go" we have to use
    - **var(keyword) + variable name + variable type**
    - **var(keyword) + variable name**

        ``` Go
        var name string = "Rohan"
                or
        var name = "golang"
        ```

- **In Golang, if we have created any variable we must have to use it either we have to remove it.**

### Short hand syntax to declare variable
- **variable name + ":=" + assign_value**
```Go
name2 := "Python"
fmt.Println(name2)
```

- **If any variable we want to export in another file, we have to named that variable in Capital letter, same applicable for function.** This is the reason behind **"P"** is capital in **"Println"**,beacuse Println is imported from outsider package.