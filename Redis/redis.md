## REDIS
in mempry data structure store used as a database, cache, and message broker.
- is an open-source key-value store
- hold its database entirely in the memory, using the disk only for persistence.
- rich set of data types: lists, sets, key-value, hashes, pub/sub

#### Key-Value:
set key value	O(1)
get key		O(1)
del key		O(1)
exists key	O(n)
expire key	O(1)
ttl key		O(1)
inc key		O(1)
decr key		O(1)
append key value		O(1)

#### Lists:
lpush key element 	O(1) => adds a new element to the head of a list
rpush key element 	O(1) => adds to the tail
lpop key [count]	O(n) => removes and returns an elements from the head of a list
rpop key [count]	O(n) => removes and return an elements from the tail of a list
llen key 	O(1) => returns the length of a list
larange key start strop I(S+N) 



#### Sets:
sadd set_name set_elements
smemebers set_name
scard set_name => numbers of elementis
sismember set_name set_element => check if set_element exists in set_name
sdiff set_name1 set_name2
sdiffstore new_set_name set_name1 set_name2
sinter set_name1 set_name2
sinterstore set_name1 set_name2
srem set_name element

sadd set_name set_elements

#### Hashes:
HSET key field value
HGET key field
HDEL key field
hgetall key
hkeys key
hvals key

#### Pub/Sub:
publish channel message
subscribe channel
unsubscribe channel
