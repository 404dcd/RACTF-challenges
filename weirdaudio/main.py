import math
import wave
import struct
import random

message = "Hello solver! I see that you are looking for a flag here. Well, it would be my honour to give it to you. The magic words are ractf{4udi0ph1le_p3Aks_0ct4l}. Good job on the solve, and have fun with the rest of the challenges."
audio = []

levels = [-12345, -8888, -6969, -1337, 1337, 6969, 8888, 12345]

for ch in message:
    for dig in oct(ord(ch)).split("o")[1].zfill(3):
        audio.append(levels[int(dig)])
    for i in range(64):
        audio.append(random.randint(-10, 10))

wav_file=wave.open("out.wav", "w")

nchannels = 1
sample_rate = 44100.0
sampwidth = 2
nframes = len(audio)
comptype = "NONE"
compname = "not compressed"
wav_file.setparams((nchannels, sampwidth, sample_rate, nframes, comptype, compname))

for sample in audio:
    wav_file.writeframes(struct.pack('h', sample))

wav_file.close()