# weirdaudio
Flag: `ractf{4udi0ph1le_p3Aks_0ct4l}` (specified in source file message variable)

Just serve the users this file.

## To test
View in Audacity or something

## Solution

I take the message, encode it to a series of triple octals, and then for each octal digit I look up a funny number corresponding (there are 8).

They go low to high, so the lowest audio level corresponds to 0 and the highest to 7. There are 64 randomish samples inbetween each char. It's guessy until you look at it in a DAW and then it's not
