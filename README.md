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

## My investigation worklog

### 1 - I do not have 120 GB free space

Since I am low on free space, first thing to do is to **implement a script that generates described file**. I will test original provided file in the end