# Log interface

## Logger interface

This minimalistic Log interface is aimed to abstract golang loggers into a reduced set of functionality:
 - Write `Info` to log
 - Write `Error` to log
 - Obtain a leveled logger `V`

The only difference between `Info` and `Error` is that the later accepts the error as an explicit parameters. Some developers might consider Error not worth being part of the interface, but it's not uncommon to find loggers that will trigger hooks or use a different formatting for errors than the one used for regular logs.

The leveled logger function at the interface expects to return a copy of the current logger configured to write to the specific log level

## Static functions

In order to use the package import as root for log related functions, and keep usage code clean, a set of static functions will allow you to:
 - Maintain loggers reference
 - Target default logger using `Info`, `Error` and `V` functions at the package level

## Design

- Initialization of logs is out of the scope of this interface, that should be managed among the logger implementation and the application using the logger
- Depending on the log implementation, leveled logs range from not supported, limited, to featurefull. In either case `V` should always return a valid `InfoWriter`.
- Loggers management from static functions keep all loggers in a map, including the default logger using the string `default` as a key.
- A `NoLog` implementation has been added to the library. It should be returned when a (leveled) logger is disabled, since if wont output any log.

