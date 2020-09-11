# Data Dictionary

| Field Name | Data Type              | Data Format | Field Size                  | Description                                                  | Example                              |
| :--------- | :--------------------- | :---------- | :-------------------------- | :----------------------------------------------------------- | :----------------------------------- |
| Actors     | Actor                  | []Actor     | ?                           | Stores all actors                                            | [Actor{}]                            |
| CameraX    | int                    | (primitive) | 8 bytes                     | Camera X Pos                                                 | 0                                    |
| CameraY    | int                    | (primitive) | 8 bytes                     | Camera Y Pos                                                 | 0                                    |
| VelocityX  | float                  | (primitive) | 8 bytes                     | Camera/World/Actor X Velocity                                | 0.0                                  |
| VelocityY  | float                  | (primitive) | 8 bytes                     | Camera/World/Actor Y Velocity                                | 0.0                                  |
| TagTable   | map[string]int         | map         | 8 bytes per elem, excl. key | Maps actor tags to actor array pos in a LUT for quick access | map[string]int{“Player”:0}           |
| State      | map[string]interface{} | map         | ?                           | Provides a state (variables that survive over loops) for actors. The world is also stateful. | map[string]interface{}{“xp”:0}       |
| X, Y, Z    | int                    | (primitive) | 8 bytes per pos             | Actor world position                                         | 0, 0, 0                              |
| ActorLogic | pointer                | (primitive) | 8 bytes                     | A pointer to an actor’s logic function. Run once a frame.    | backgroundActorLogic or 0xc00002c008 |
| Image      | pointer                | (primitive) | 8 bytes                     | Pointer to the actor’s image to reduce garbage collection cycles and prevent unnecessary malloc. This field is optional as you can specify a renderhook for the engine to call. | 0xc00002c008                         |