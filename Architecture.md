- component basis structure

    when we need a networking module

    ``` go
        type Network struct {
            ...
        }

        func NewNetwork(config Config) (*Network, error) {

        }
    ```

    In the consumer side

    ``` go
        interface networkInterface {
            ReadMessage()
            SendMessage()
            ...
        }

        type LagrangeNode struct {
            network networkInterface
            ...
        }
    ```

- use config file and config parsing frameworks (like vyper)
