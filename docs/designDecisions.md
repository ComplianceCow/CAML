# Design Decisions

While building the CAML application there were several design decisions made which are documented here, along with their rationale.

## Intended use of the Machine-Readable Metrics

Issue:  Is there both yaml and code implementing the policy or is there just yaml which is interpreted by a service which then builds the query to the metrics server?

Conversation held 2022/01/14, recorded in Issue [#31](https://github.com/ContiNube/CAML/issues/31).

The decision is that yaml file is read by an interpreter which builds a query message sent to the measures service.  By using this design we lessen the drift between the source of truth (the yaml file) and the implementation.  This also provides the abstraction of the policy in yaml from the implementation of query building.

## Output openmetric/prometheus metric data

Conversation held 2022/03/11, as related to Issue [#24]
(https://github.com/ContiNube/CAML/issues/24).

A working 'MVP' solution being critical to user engagement and ongoing improvements we need an (example) operational stack that goes from the .yaml file to a demonstratable metric output by a user. This isn't the only possible stack, but its a stack. 

To this end we provide an openmetric/prometheus formatted output from the CAML toolchain.  
An 'off the shelf' open source demonstration toolchain injesting, and then providing dashboard display is included & configured as part of the distribution. The stack includes node_exporter in the caml_dev container, and provisioned (pre-configured) prometheus & grafana containers. Each has a web interface that can be viewed to explore the metrics.

