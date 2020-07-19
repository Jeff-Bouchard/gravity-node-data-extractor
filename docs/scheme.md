
# Gravity built-in data extractor architecture

## Abstract

This document describes the UML architecture of the Gravity built-in data extractor. The architecture and implementation of extractors is developed by the **Gravity Core Team**. The implementation described in this document is the earliest initial version, so it is considered as *built-in*.

## Foreword

The architecture provided in this document is offered as ***best practice*** guidance for Gravity protocol node operators and contributors. 

Violating the specification and object relations can lead to unexpected behaviour of some system parts of the Gravity Node.

## Scheme Overview

![Gravity data extractor scheme](https://i.imgur.com/xkkFsrU.jpg)


## Main concepts

The main concepts behind the implementation of extractors are:
1. Providing a stateless system & avoiding data mutability
2. Conforming to ***any*** kind of data (reusability)
3. Manifesting available operations

### Stateless vs stateful

In designing modern systems, it is necessary to achieve a certain balance in how objects are manipulated.

The vast majority of applications combine both stateless and stateful system parts.

As regards to the *Gravity protocol*, it's not an exception. Such parts of the system that are responsible for data mutability and storage are considered stateful. 

By design, extractors are aimed to only perform data aggregation and mapping procedures, which is the reason for also introducing a stateless approach. In addition, *extractors are isolated*, meaning that they are not "aware" of any particular data consumers.
 

### Conformance to any data & Reusability

The current architecture gives an ability to ***transform*** and ***transport*** data in any way that is necessary for a particular use-case. There are two interfaces that implement this functionality:
1. Extractor<T, R> - This interface declares methods on how the data is ***transformed***,
2. IDataBridge<T, R> - This interface declares methods on how the data is ***transported***.

Generic types provided in declarations represent:
1. T - raw data type,
2. R - transformed response data type. 

Furthermore, such approach gives us an ability to represent any kind of data and deliver it in different ways.

The scheme below shows possible implementations:

![Gravity data extractor response controller examples](https://i.imgur.com/RnPi1Kw.png)

The ***reusability of every distinct system part*** translates into the ***reusability of the whole system***.

### Available operations

The design requires three accessible endpoints to be implemented. These are needed for:
1. Fetching raw data.
2. Fetching mapped/transformed data.
3. Fetching extractor info in JSON format, containing the data feed tag and description.
