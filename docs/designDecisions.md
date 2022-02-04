# Design Decisions

While building the CAML application there were several design decisions made which are documented here, along with their rationale.

## Intended use of the Machine-Readable Metrics

Issue:  Is there both yaml and code implementing the policy or is there just yaml which is interpreted by a service which then builds the query to the metrics server?

Conversation held 2022/01/14, recorded in Issue [#31](https://github.com/ContiNube/CAML/issues/31).

The decision is that yaml file is read by an interpreter which builds a query message sent to the measures service.  By using this design we lessen the drift between the source of truth (the yaml file) and the implementation.  This also provides the abstraction of the policy in yaml from the implementation of query building.
