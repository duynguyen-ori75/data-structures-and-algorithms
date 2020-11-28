# Consistent hashing

## Definition

Consistent hashing is a strategy for dividing up keys/data between multiple machines.

It works particularly well when the number of machines storing data may change. This makes it a useful trick for system design questions involving large, distributed databases, which have many machines and must account for machine failure.

## About the implementation

Very simple implementation, very easy to explain how consistent hashing works.

Besides, adding a new node is supported using gossip protocol instead of any consensus technique. This does not work in a production environment, as adding a new node to a cluster needs to be strongly consistent over the cluster.

### TODO (no plan yet)

- Replication factor
- Benchmarks (with death nodes, after replication factor is finished)

## References:
- https://en.wikipedia.org/wiki/Consistent_hashing
- https://github.com/papers-we-love/papers-we-love/blob/master/datastores/dynamo-amazons-highly-available-key-value-store.pdf
- https://medium.com/system-design-blog/consistent-hashing-b9134c8a9062
- https://www.interviewcake.com/concept/java/consistent-hashing