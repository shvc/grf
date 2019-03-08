# grf
gen random files
### Usage
```
./grf -h
Usage of grf:
  -n int
        number of files (default 1)
  -o string
        output dir (default ".")
  -s int
        size of file (default 1024)

```
### Gen 1024 random files
```
mkdir outdir
./grf -n 1024 -s 1048576 -o outdir
```
