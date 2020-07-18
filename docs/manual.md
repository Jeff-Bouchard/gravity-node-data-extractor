
# Manual: How to reuse Gravity data extractor

## Table Of Contents

1. Introduction
2. Initialization
3. First overview
4. Things to override
5. Extending the logic
 
## Introduction

This manual provides steps that will help you to build your own extractor tied to specific endpoint. Our example shows the usage of Binance price extractor. The data bridge is HTTP.

## Initialization

Firstly, clone the repo and start building swagger meta data by calling:

>  bash ./bootstrap.sh

After that, enter info about your extractor, it will be visible in Swagger UI spec.

> Enter extractor host: explorer.gravityhub.org
> Enter contact website: gravityhub.org
> Enter contact email: oracles@gravityhub.org

This command will add swagger metadata to your extractor.

## First overview

Here is the file directory. 

![Things to override](https://i.imgur.com/YnJuPpf.png)

1. Controller - where data transition takes place
2. Docs - a place for docs 
3. Model - model definitions, both internal and public
4. Router - entities responsible for routing

---

Let's take a look at main.go file. We are interested in ***init*** and ***main*** functions:

Init function takes care of CLI parameters that can be crucial according to your extractor type.

![Init fn](https://i.imgur.com/ZClHaOE.png)

For example:

1. Almost *any* price extractor *will mostly* handle pair name parameter from CLI, it's convenient, we can run multiple extractors for different price pairs, just by handling the params. (pair param)
2. We can combine multiple extractors in one repo. It gives us an ability to choose implementation ***at runtime***, just by passing extractor type. (type param)
3. It goes without saying, providing appropriate extractor tag is crucial. (tag param)


---

Here we can see that main.go file creates concrete data transport controller instance (*ResponseController*).

Business logic of transported data is incapsulated in *ResponseController*.

![Main fn](https://i.imgur.com/yRaJPgg.png)

"r" is router singleton, it provides reference to routes. In this example - for HTTP routes.

---

We mentioned extractor switching ***at runtime*** just above. In order to achieve such a goal we need some kind of enumerator. [enum.go](models/enum.go)

![Extractor enumerator](https://i.imgur.com/Kl3BjS6.png)


Also, available extractor types must be declared internally. (*binance*, *metal*)


## Things to override

If your goal is to extend/mutate existing implementations, then follow the steps:

1. Take a look at *models directory*
2. Declare a new model or extend existing. You must implement all *IExtractor interface methods*  in order to declare a valid extractor.
![enter image description here](https://i.imgur.com/TD6FXmA.png)

Don't forget to mark structs you want to expose by specifying swagger model declaration.
>// swagger:model

You can find more details on goswagger [here]([https://goswagger.io/use/spec.html](https://goswagger.io/use/spec.html))

4. Provide implementation with the name in *enum.go* *(in order to change in at runtime by need)*
5. Open *ResponseController* declaration file

![ResponseController internal extractor](https://i.imgur.com/oBJ5yc7.png)

Here you can provide logic with handling specific initialization data.

### Extending the logic

Feel free to extend the existing system by adding more layers of abstraction. 

However, keep in mind that violating existing interface constraints mentioned in [scheme.md](docs/scheme.md) is considered as system *anti-pattern*.