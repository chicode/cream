# cream

The Chicode Remote Execution Action Mainframe.

## About

Cream works via a module system. The main Cream server is a simple Go websockets app that manages incoming connections and links them to modules. A module is a single-user docker container for a specific language environment; each module has the same API regardless of the language.

When a user connects to the server and requests a specific language, the main scheduler will start a module for the requested language environment, routing all of that users further communications to that module. Once the `disconnect` event is fired (manually or automatically), the module is then deleted.