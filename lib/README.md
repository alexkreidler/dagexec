# docs

the dagexec API defines a simple but powerful API for executing DAG computational graphs on a variety of platforms

## Core concepts

Goals: 

Short-term:
- highly performant
- environment-agnostic
- Easy and convenient Go API using goroutines (possibly even easy for spawning accross machines: maybe use Nomad API to run `dagexec` worker on many machines)
- Easy serialization of graphs? - to graphviz, html/JS visualizations, protobuf? might not work for go functions


Longer-term:
- distributed (think about resource dependencies, e.g. an Output might not be  a tiny struct but could be a huge binary or ML model, which node to store on or orchestrate corresponding tasks onto)
- streaming/realtime data
- task/data syncronization? do we want to allow shared data between tasks? as of now: no - b/c simplicity

## About/Use Cases/Future work

`dagexec` is a single-threaded 1 machine, master task runner which can run tasks on a variety of platforms using the full capabilities of a DAG. You can create a graph programmatically, but then you have to `Compile` the graph. Then it can be serialized and run on any machine with the proper Executors and configuration, allowing you to run graphs on CI machines, etc. 

This project may also contain a dagexec style Executor which can be distributed on a network of machines (e.g. via ansible or etc) and then will create a cluster of executors which then will have its own orchestration logic for distributing tasks.

However, in most scenarios, dagexec will be run on one machine and will orchestrate/manage a bunch of resources somewhere else (e.g. in a Kubernetes cluster or a Docker Swarm). It can also be used for trivial examples like simple scripts or programs which need to create a bunch of objects in different orders, etc.

### Node
A node represents is one logical object in a graph. 

A node either can be used as a logical marker (e.g. as a marker to refer to a Section,  to group/package subsections of a graph), or it can correspond to exactly one `Task`

### Section

A section is a logical section of a graph, defined as all of the nodes along all of the edges from a `Start` Node to an `Finish` Node. A section can be reused an slotted in for any given Node as it implements the Node interface. That `SectionNode` will pass its `Input` to the `Input` of the `Start` Node and pass the `Output` of the `Finish` Node to its `Output`

### Task

A task represents one computational function to execute. A task references a `Function`

It takes in input and returns an output, and is a synchronous operation. The power of dagexec is that we can parallelize and represent dependencies between synchronous operations very effectively.

A task has a `State` which represents the state of this synchronous operations

### Trial
A Trial represents an attempt at running a Task. Ideally, only one Trial will be run per node as it will succeed immediately. However, nodes can create new Trials to retry if they fail.

A trial references a `Task` (its parent), which references a `Function`,  and also references an `Input`

A Trial also has a `State`

### Input and Output

Generic data types which represent data flowing through the graph.

### Function

A function is a type which contains all the information needed to produce an `Output` from an `Input`. It can vary depending on the `Executor` it is run with.

### Executor

An executor is an interface which can `Run(Trial)` aka `Run(Function, Action)` and return an `Output`

The `Run` function is expected to be a synchronous function, e.g. it waits for a long-running operation to continue.

It can modify the  

### Graph
A computational graph under which all nodes are created.

It has a `RootNode` and an `EndNode`.


#### To think about:

1. how to deal with streaming data
2. Better primitives for data (joining, mapping, reducing, etc)
3. Should we add branching etc?
4. Should we make the graph fully static, need to be entirely created at graph compile/build time, or should we allow Functions to manipulate the Graph
5. Enabling or disabling nodes?