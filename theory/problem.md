# Defining The Problem

### The problem

We are required to build a tile based Roleplaying Game.

It should include:

- The ability for the player to move
- Moving opponents
- A combat system
- Opponents and the player can die
- Weapons or Items or Special Abilities
- Characters that are able to leveled up
- Character events (eg: triggers)



### The steps that will be undertaken

I shall take the following steps.

I will code an engine that can facilitate the needs of the game. The engine will greatly assist in the rapid development of new features as it will be modular and object oriented in nature.

The engine will be comprised of actors who have their own states, we shall refer to these as stateful actors - states store information that can be written and read by logical functions and read by graphical rendering functions. Actors are stored in the world - a world is very similar to an actor albeit with it's own features and subroutines. Actors are able to communicate, so to speak, with other actors via modification of each others states. Actors essentially containerise features of the game and discourage monolithic programming practices. 

The true beauty of a stateful actor based system is that it trains you to think with or like an object, whereby each actor is essentially it's own person or agent performing an assigned task. Instead of thinking "how am I going to make this look like it's doing something", think instead "how am I going to make it do something".

I will implement the previously listed game features within the contraints of my engine's actor system.

I shall test, debug and remedy any errors within my program.

### Possible roadblocks or dilemmas

The engine should permit me the freedom to implement anything I please with great ease and eloquency.

Therefore, I cannot forsee any potential dilemmas that I may encounter during my programming of the game.

Though perfomance may be one thing I might have to look out for.

