# hogor

Rings your door bell when someone opens the gate.

## Setup

Hogor asumes the following setup:

1. a raspberry pi or similar
2. gpio pin 2 is connected to a switch for activation/deactivation of the whole system
3. gpio pin 3 is connected to a switch on the gate which opens when the gate is open
4. gpio pin 4 is connected to a light/led for showing the status
5. gpio pin 5 is connected to the door bell

## What hogor does

When switched off (gpio pin 2 HIGH), hogor will do nothing (gpio pin 4/5 LOW).

When switched on (gpio pin 2 LOW), hogor will wait for 60s with a blinking status led (gpio pin 4 LOW/HIGH).
After 60s, hogor will ring the door bell (gpio pin 5 HIGH) when opening the gate (gpio pin 3 LOW).

In any state, hogor logs all events.
