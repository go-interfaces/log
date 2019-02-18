# Log interface <img align="left" width="60px" src="https://avatars0.githubusercontent.com/u/47711035?s=400&u=e8a2891cca67da66972ad478069588deb0299e4b&v=4">

Logger interface allows you to switch between loggers implementations at Go projects.



## Logger interface

This minimalistic Log interface is aimed to abstract golang loggers into a reduced set of functionality:
 - Write `Info` to log
 - Write `Error` to log
 - Obtain a leveled logger `V`
 - `SetLevel` for the logger

The only difference between `Info` and `Error` is that the later accepts the error as an explicit parameters. Some developers might consider Error not worth being part of the interface, but it's not uncommon to find loggers that will trigger hooks or use a different formatting for errors than the one used for regular logs.

The leveled logger function at the interface expects to return a copy of the current logger configured to write to the specific log level. Logger level configuration is managed through `SetLevel`

## Static functions

In order to use the package import as root for log related functions, and keep usage code clean, a set of static functions will allow you to:
 - Maintain loggers reference
 - Target default logger using `Info`, `Error`, `V` and `SetLevel` functions at the package level

## Design

- Initialization of logs is out of the scope of this interface, that should be managed among the logger implementation and the application using the logger
- Depending on the log implementation, leveled logs range from not supported, limited, to featurefull. In either case `V` should always return a valid `InfoWriter`.
- Loggers management from static functions keep all loggers in a map, including the default logger using the string `default` as a key.
- A `NoLog` implementation has been added to the library. It should be returned when a (leveled) logger is disabled, since if wont output any log.
- `SetLevel` is added as a complement to `V`. Some loggers will configure the level at initialization, out of this interface, and won't need to set again the logger level once the application is running. This suggest that `SetLevel` shouldn't be part of this interface, but since `V` is already part of the interface, a numeric leveling is assumed, hence a logger level management function is also included.
