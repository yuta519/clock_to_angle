# Clock To Angle

## Overview
Have you ever solved this math problem in school?

> It is 4 o’clock.  What is the measure of the angle formed between the hour hand and the minute hand?

This is a simple math problem. But I wonder how to solve this problem with programming.
So tried to implement with some languages.

## What does this repository do?
- This repo finds the angle of clock hands.
- When given current time or specified time like `4:00`, a program says `120°`.

## Thinking way
- Short hand angle
  - A short hand has 12 numbers. So when 1 hour moves, short hand angle moves 30°.
  - Also need to consider minutes.
  - In 1 hour(=30°), there is a 60 minutes. So when 1 minutes moves, short hand angle moves 0.5°.
- Long hand angle
  - A long hand has 60 numbers. So when 1 minutes moves, long hand moves 6°.
