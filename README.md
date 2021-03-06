# types

Shared types for Sym integrations, delivered as protobufs.

Refer to the [API Docs](docs/index.md) for more details.

## Example Integration

Generate a project from one of our [starter templates](https://github.com/symopsio/serverless-templates) to integrate Sym Access workflows into your systems.

## Generated Libraries

### Typescript

TypeScript implementations are generated with [protobufjs](https://github.com/protobufjs/protobuf.js).

`import { sym } from "@symops/types"`

### Python

Python implementations are generated with [protoc](https://developers.google.com/protocol-buffers/docs/reference/python-generated).

```
from sym.enums.service_pb2 import Service
from sym.messages.approval_pb2 import Approval
from sym.models.user_pb2 import User
```

### Go

Golang implementations are generated with [protoc](https://developers.google.com/protocol-buffers/docs/reference/go-generated).

```
import "github.com/symopsio/types/go/enums"
import "github.com/symopsio/types/go/messages"
import "github.com/symopsio/types/go/models"
```

## Packages

| Package | Description |
| ------- | ----------- |
| messages | Types that are sent from system to system |
| models | Shared data types |
| enums | Shared enum types |

