# graphcommons-go

Go wrapper for [Graphcommons](https://graphcommons.github.io/api-v1/) API.

## Installation

```
go get github.com/erkanay/graphcommons-go
```

## Usage

#### Authentication
```go
import . "github.com/erkanay/graphcommons-go"

gc, _ := GraphCommons("API_KEY")
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
            Action: "edge_create",
            FromName: "Erkan",
            FromType: "Person",
            ToName: "Maximé",
            ToType: "Person",
            Name:"COLLABORATED",
            Weight: 1,
        },
    },
}
resp := gc.CreateGraph(body)
```

#### Get Graph
```go
resp := gc.Graphs("49beddc2-9616-409e-83aa-dd34335c69ee")
```

#### Update Graph
```go
type Signals struct{
	Signals  []Signal `json:"signals"`
}
body := Signals{
    Signals: []Signal{
        Signal{
            Action: "edge_create",
            FromName: "Aude",
            FromType: "Person",
            ToName: "Maximé",
            ToType: "Person",
            Name: "COLLABORATED",
            Weight: 1,
        },
        Signal{
            Action: "edge_create",
            FromName: "Kosta",
            FromType: "Person",
            ToName: "Bogdan",
            ToType: "Person",
            Name: "COLLABORATED",
            Weight: 1,
        },
    },
}
resp := gc.UpdateGraph("49beddc2-9616-409e-83aa-dd34335c69ee", body)
```

Constituted graph is [here](https://graphcommons.com/graphs/49beddc2-9616-409e-83aa-dd34335c69ee).

#### Delete Graph
```go
resp := gc.DeleteGraph("49beddc2-9616-409e-83aa-dd34335c69ee")
```
