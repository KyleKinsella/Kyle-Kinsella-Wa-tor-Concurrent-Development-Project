# Wa-tor Project, Made by Kyle Kinsella | C00273146

# What is wa-tor?
### Wa-tor is a population dynamics simulation. The planet of wa-tor is shaped like a torus and it is entirely covered with water. There are two types of animals in the world of wa-tor, Fish and Sharks. For more information see here: https://en.wikipedia.org/wiki/Wa-Tor

# Rules for fish
### 1. At each chronon, a fish moves randomly to one of the adjacent unoccupied squares. If there are no free squares, no movement takes place.<br>
### 2. Once a fish has survived a certain number of chronons it may reproduce. This is done as it moves to a neighbouring square, leaving behind a new fish in its old position. Its reproduction time is also reset to zero.<br>

# Rules for sharks
### 1. At each chronon, a shark moves randomly to an adjacent square occupied by a fish. If there is none, the shark moves to a random adjacent unoccupied square. If there are no free squares, no movement takes place.<br>
### 2. At each chronon, each shark is deprived of a unit of energy.<br>
### 3. Upon reaching zero energy, a shark dies.<br>
### 4. If a shark moves to a square occupied by a fish, it eats the fish and earns a certain amount of energy.<br>
### 5. Once a shark has survived a certain number of chronons it may reproduce in exactly the same way as the fish.

# I was not able to do the reproduction for the fish and sharks due to me not having enough time. 

# Below is an image of my wa-tor project:
<!-- ![Dining Philosophers](images/diningPhi.png) -->

# How to run my code
### To run my code you must do the following:
#### 1. Download my go file
#### 2. Type in: Go run filename, change filename to be the name of the file

# Licensing
### All of this work I have completed is licensed with my below license.
<p xmlns:cc="http://creativecommons.org/ns#" >This work by <span property="cc:attributionName">Kyle Kinsella</span> is licensed under <a href="https://creativecommons.org/licenses/by-nc-sa/4.0/?ref=chooser-v1" target="_blank" rel="license noopener noreferrer" style="display:inline-block;">Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International<img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/cc.svg?ref=chooser-v1" alt=""><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/by.svg?ref=chooser-v1" alt=""><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/nc.svg?ref=chooser-v1" alt=""><img style="height:22px!important;margin-left:3px;vertical-align:text-bottom;" src="https://mirrors.creativecommons.org/presskit/icons/sa.svg?ref=chooser-v1" alt=""></a></p> 