# Billion-Rows-Test

## Task

> [Original link](https://ecwid.to/GO)

You have a simple text file with IPv4 addresses. One line is one address, line by line:

```
145.67.23.4
8.34.5.23
89.54.3.124
89.54.3.124
3.45.71.5
...
```

The file is unlimited in size and can occupy tens and hundreds of gigabytes.

You should calculate the number of __unique addresses__ in this file using as little memory and time as possible.
There is a "naive" algorithm for solving this problem (read line by line, put lines into HashSet).
It's better if your implementation is more complicated and faster than this naive algorithm.

Before submitting an assignment, it will be nice to check how it handles this [file](https://ecwid-vgv-storage.s3.eu-central-1.amazonaws.com/ip_addresses.zip). Attention - the file weights about **20Gb**, and unzips to about **120Gb**.

## How to run

Just execute `make run` if you are on a Linux machine

It will generate "ip.txt" with `1.000.000.000` lines of random IPs (`14GB`), run the test and print time passed with the `time` command.

> I am low on free space, sorry 😢

## Test results

### My laptop specs:

```
> screenfetch

         #####
        #######
        ##O#O##
        #######
      ###########
     #############           OS: Pop 22.04 jammy
    ###############          Kernel: x86_64 Linux 6.9.3-76060903-generic
    ################         Shell: fish 3.3.1
   #################         Disk: 60G / 93G (68%)
 #####################       CPU: Intel Core i7-10870H @ 16x 5GHz
 #####################       GPU: NVIDIA GeForce RTX 3060 Laptop GPU
   #################         RAM: 8673MiB / 15862MiB

```

### Warning

I measure time via unix `time` command. Moreover, I continue using my laptop during test session, which means that **measurments are inaccurate**.

I know that I need to execute several time at least and take average. I do not care, however.

### Results

| Attempt Number | Approach | Execution Time | Diff | Commit(s) | Unique Ids | Comment |
|-----------------|---|---|---|--|--|--|
| 0 | Create a script that creates file `100.000.000` lines | 2.18.40 | - | [946d03a](https://github.com/vanyason/Billion-Rows-Test/commit/946d03a0fb515ed30058e1161cdd7bc76e14065a) and [81a7e7c](https://github.com/vanyason/Billion-Rows-Test/commit/81a7e7c9af8a20519c358d27f9a6a1552cad4a19) | - | File with `100.000.000` lines is `1.4GB` |
| 1 | Naive approach (open, read, save, print length)| 13.37.89 | - | [df039be](https://github.com/vanyason/Billion-Rows-Test/commit/df039bed06728c767020ff3d66abb323325f5ebc) | `98.843.672` | - |
| 2 | Still Naive approach (open, read, save, print length) but using `bufio.NewScanner` with buffer set to `100MB` | 0.33.36 | ~13 minutes faster | [ffd9ffe13](https://github.com/vanyason/Billion-Rows-Test/commit/ffd9ffe13d3ad6511dd281e762c26514ff9edef6) | `98.843.672` | I didn't expect such boost. At this point I am going to increase test file to 1.000.000.000 lines |
| 2.1 | Updated "File generator" and created `1.000.000.000` lines input | 05.06.33 | - | [99701fe4](https://github.com/vanyason/Billion-Rows-Test/commit/99701fe4feabd3b7310779b96371fd11c4bef5be) | - | File with `1.000.000.000` lines is `14GB`. What kind of file was originally provided in the task section ? 🤔 |
| 2.2 | Same as 2 but on `1 billion lines` | 31.57.81 | - | [ffd9ffe13](https://github.com/vanyason/Billion-Rows-Test/commit/ffd9ffe13d3ad6511dd281e762c26514ff9edef6) | `still unknown` | Total failure. Luckly I had `20GB` swap space. In a minute RAM was full, next 30 minutes script tried to survive but still crashed |
| 3 | Using bitset to store all the possible ips | 2.45.23 | - | [6751815](https://github.com/vanyason/Billion-Rows-Test/commit/675181544300c2d215f8bdf45a88d68af179f395) | `892117370` | As I discovered during step `2.2`, I can dramatically fail on lack of RAM. Because of that was decided to convert all the `IPs` to `int`. According to my calculations if all the IPs are `2^32` , and we assign each IP a bit, that is `2^32 bits = 536870912 bytes = 524288 Kb = 512 MB`. Bitset is a good datastructure for this case because uses only `512 MB` to store all the IPS possible + `concurrently safe` (Since it is append only and operation is Idempotent) |
| 4 | Added 32 (my CPU amount) goroutines to process file in equal sections | 0.39.67 | 6x faster | [bb859e0](https://github.com/vanyason/Billion-Rows-Test/commit/bb859e0cd86758be2fcee2638e1141001c986f6b) | `892117370` | Not impressed. I use bufio.Scanner which goes line by line. That does not allow goroutines to effectively jump to their chunk. Moreover, approach where goroutine reads and processes data at the same time feels inefficient |