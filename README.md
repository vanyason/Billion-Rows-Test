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

Before submitting an assignment, it will be nice to check how it handles this [file](https://ecwid-vgv-storage.s3.eu-central-1.amazonaws.com/ip_addresses.zip). Attention - the file weighs about **20Gb**, and unzips to about **120Gb**.

## How to run

Just execute `make` if you are on a Linux machine

It will generate "ip.txt" with `100.000.000` lines of random IPs (`1.4GB`), run the test and print time passed with the `time` command.

> I am low on free space, sorry 😢

## Test results

My laptop specs:

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
 #####################       CPU: Intel Core i7-10870H @ 16x 5GHz [63.0°C]
 #####################       GPU: NVIDIA GeForce RTX 3060 Laptop GPU
   #################         RAM: 8673MiB / 15862MiB

```

| Attempt Number | Approach | Execution Time | Diff | Commit |
|-----------------|---|---|---|--|
|0 | Create a script that creates file `100.000.000` lines | 2.18.40 | none| none |