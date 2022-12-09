# Logsnag.go

A go module for publishing events and insights to LogSnag.

## Installation

```sh
go get github.com/samuel-pratt/logsnag.go
```

## Usage

### Import Library

```go
import (
	"github.com/samuel-pratt/logsnag.go"
)
```

### Initialize Client

```go
logSnag := logsnag.NewLogSnag(
    "7f568d735724351757637b1dbf108e5",
    "my-saas"
)
```

### Publish Event

```go
logSnag.Publish(
    "waitlist",         // Channel
    "User Joined",      // Event
    "🛥️",               // Icon
    map[string]any{     // Tags
        "name": "john doe",
        "email": "john@example.com",
    },
    true,               // Notify
)
```

### Publish Insight

```go
logSnag.Insight(
    "User Count",   // Title
    "100",          // Value
    "🛥️",           // Icon
)
```
