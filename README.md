# grf
gen random files

### Build
```
git clone https://github.com/cchare/grf.git
cd grf
go build
```

### Usage
```
./grf -h
Usage of ./grf:
  -n int
    	number of files (default 1)
  -o string
    	output dir (default ".")
  -s string
    	size(K,M,G,T) of file (default "1M")
  -v	show version

```

### Gen 5 2.2M random files to outdir/
```
mkdir outdir
./grf -n 5 -s 2.2M -o outdir
```
