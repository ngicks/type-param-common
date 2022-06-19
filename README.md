# type-param-common

Type parameter primitives and commons.

Basically wrapper around go std library. If wrapping is inadequate, copy-and-paste and modify go codes.

## Non overlapping package name

Wrappers of Go std lib are suffixed with `-param`

## Package descriptions

### heap-param, list-param, ring-param

Wrappers around `container/heap`, `container/list` and `container/ring`. Getter/Setter and Unwrap methods are added to mutate/observe Value from outside wrapper.

Correct usage of these types involves direct mutation by assigning to Value. Since Go expose no way to trap property accesses, without additional getters/setters we have no way to access to those properties. Thus these methods should be justified.

### sync-param

Type-param aware wrappers around sync.Map and sync.Pool.

### slice

Deque, queue, stack and whatever that needs type-param. It eases pain of `write-deque-type-everywhere`.

- [x] Deque
- [x] Queue
- [x] Stack

### iterator

Iterator impl for go.

- [x] creating iterator
  - [x] from slice
  - [x] from list.List[T]
  - [x] from channel
- [x] Reverse, ForEach
- [x] Filter(Select and Exclude)
- [x] Map
- [x] Reduce
- [x] Skip and Take
