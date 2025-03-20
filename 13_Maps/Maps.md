# Maps
- Similar as HashMap in java

### Map Initialize:

```Go
m := make(map[key_type]value_type) //value type should be outside of the []
            or
m1 := map[string]int{
	"price": 100, 
	"phones": 3
}
```

- **Add Element in Map:** `map_name[key] = value`

```Go
m["India"] = "in"
```

- **Getting Element from Map:** `fmt.Println(map_name[key])`

```Go
fmt.Println(m["India"])
```
- ***If key does not exists in map returns zero value***
- **Printing full map:** `fmt.Println(map_name)`
- **Length of map:** `len(map_name)` same as array

```Go
fmt.Println("Length of Map:", len(m))
```

- **Loop Through map, this way we can also print map**
``` Go
for index, value := range m {
	fmt.Printf("Key is %s and value is %s \n", index, value)
}
```

- **Delete Element from Map:** `delete(map_name, key)`

```Go
delete(m, "India")
fmt.Println(m)
```

- **Empty full map:** `clear(map_name)`

```Go
clear(m)
fmt.Println(m)
```

- **Check map contains key or not:** `val1, val2 := map[key]`
	- **val1:** Value of the given key.
	- **val2:** boolean value that indicates the item is present in map or not.

```Go
val, ok := m["China"]
fmt.Println(val) // it will print it's key value
//_, ok := m["China"]  if we want we can skip it
if ok {
	fmt.Println("Exists")
} else {
	fmt.Println("Not Exists")
}
```
- **Map equality Check:** `maps.Equal(map1, map2)`
``` Go
m3 := map[string]int{"price": 10000, "Samsung": 3}
m4 := map[string]int{"price": 40000, "Apple": 3}
fmt.Println(maps.Equal(m3, m4))
```

