# glorydays
Flag: `ractf{Mus1c_M0rSe_MesS4g3}` (also in flag.txt)

Give users a link to the musescore. The source file is included just in case, but give them the link as musescore.com has an inbuilt player.

## To test
Click the link, press play on the web player.

## Solution

Top line of music (Alto Sax part) is morse code. Low pitch = dot, high pitch = dash. ALSO short = dot, long = dash. Spacing is barlines / audible space.

Decodes to Base 32 which is easy to figure out (had to for the morse as you can't encode a {), then decode that.
