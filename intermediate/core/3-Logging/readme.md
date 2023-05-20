# Logging

## What to log

- The timestamps for when an event occurred, or a log was generated
- Log levels such as debug, error, or info
- Contextual data

## log package

- The Go standard library has a built-in log package that provides most basic logging features.
- While it does not have log levels (such as debug, warning, or error), it still provides everything
  you need to get a basic strategy set up.
- Great for local development when getting fast feedback is more important than generating rich, structured logs.

## Logging frameworks

- Logging framework helps in standardizing the log data
- It's easier to read and understand the log data
- glog and logrus
- `logrus` is better maintained and used in popular projects like Docker.
