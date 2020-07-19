
# Manual: How to reuse Gravity data extractor

## Table Of Contents

1. Introduction
2. Initialization
3. First overview
4. Things to override
5. Extending the logic
 
## Introduction

This manual provides steps that will help you to build your own extractor tied to a specific endpoint. The example described in this manual uses a Binance price extractor. The data bridge used is HTTP.

## Initialization

First, clone the repo and build Swagger meta data:

>  bash ./bootstrap.sh

Enter info about your extractor. It will be visible in the Swagger UI spec.

> Enter extractor host: explorer.gravityhub.org
> Enter contact website: gravityhub.org
> Enter contact email: oracles@gravityhub.org

This command will add Swagger metadata to your extractor.

## First overview

The file directory is described below.

![Things to override](https://i.imgur.com/YnJuPpf.png)

1. Controller - implements data transition
2. Docs - a directory for documentation 
3. Model - model definitions, both internal and public
4. Router - entities responsible for routing

---

Let us take a look at the main.go file, specifically at the ***init*** and ***main*** functions:

Init function takes care of CLI parameters that can be crucial depending on the type of extractor that you are implementing.

![Init fn](https://i.imgur.com/ZClHaOE.png)

For example:

1. Almost *any* price extractor *will* process the pair name parameter from the CLI. Multiple extractors can be run for different price pairs, just by handling the params. (pair param)
2. Multiple extractors can be combined in one repo. It gives an ability to choose an implementation ***at runtime***, just by passing an extractor type. (type param)
3. It is also necessary to provide an appropriate extractor tag. (tag param)


---

The main.go file creates a concrete data transport controller instance (*ResponseController*).

The business logic of transported data is incapsulated in the *ResponseController*.

![Main fn](https://i.imgur.com/yRaJPgg.png)

"r" is a router singleton that provides reference orr routes, in this example for HTTP routes.

---

The extractor switching ***at runtime*** was mentioned above. In order to achieve such a goal, an enumerator is needed. [enum.go](models/enum.go)

![Extractor enumerator](https://i.imgur.com/Kl3BjS6.png)


Also, available extractor types must be declared internally. (*binance*, *metal*)


## Things to override

If the goal is to extend/mutate existing implementations, follow these steps:

1. Go into the *models directory*
2. Declare a new model or extend the existing one. You must implement all *IExtractor interface methods*  in order to declare a valid extractor.
![enter image description here](https://i.imgur.com/TD6FXmA.png)

Do not forget to mark structs you want to expose by specifying a Swagger model declaration.
>// swagger:model

You can find more details on goswagger [here]([https://goswagger.io/use/spec.html](https://goswagger.io/use/spec.html))

4. Provide an implementation with the name in *enum.go*
5. Open the *ResponseController* declaration file

![ResponseController internal extractor](https://i.imgur.com/oBJ5yc7.png)

Here you can provide logic for handling specific initialization data.

### Extending the logic

It is encouraged to extend the existing system by adding more layers of abstraction. However, it is necessary to keep in mind that violating the existing interface constraints mentioned in [scheme.md](docs/scheme.md) is considered as a system *anti-pattern*.
