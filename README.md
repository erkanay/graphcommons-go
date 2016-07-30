# graphcommons-go

Go wrapper for [Graphcommons](https://graphcommons.github.io/api-v1/) API.

## Installation

```
go get graphcommons
```

## Usage

#### Authentication
```go
gc, _ := graphcommons.GraphCommons("API_KEY")
resp := gc.Status()
fmt.Println(resp) // {"msg":"Working"}

```

#### New Graph
```go
body := Graph{
	Name: "Go graph",
	Description: "Go wrapper helps to create graph",
	Status: 1,
	Subtitle: "Graphcommons-go",
	Signals: []Signal{
	    Signal{
	        "edge_create",
	        "Erkan",
	        "Person",
	        "Maximé",
	        "Person",
	        "COLLABORATED",
	        1,
	    },
	},
}
resp := gc.CreateGraph(body)
```

#### Get Graph
```go
resp := gc.Graphs("49ef5458-ab17-40b2-b702-2ccad3ced756")
```

#### Update Graph
```go
type Signals struct{
	Signals  []Signal `json:"signals"`
}
body := Signals{
    Signals: []Signal{
        Signal{
            "edge_create",
            "Aude",
            "Person",
            "Maximé",
            "Person",
            "COLLABORATED",
            1,
        },
        Signal{
            "edge_create",
            "Kosta",
            "Person",
            "Bogdan",
            "Person",
            "COLLABORATED",
            1,
        },
    },
}
resp := gc.UpdateGraph("49ef5458-ab17-40b2-b702-2ccad3ced756", body)
```

#### Delete Graph
```go
resp := gc.DeleteGraph("49ef5458-ab17-40b2-b702-2ccad3ced756")
```
