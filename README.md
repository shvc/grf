# grf
gen random files

### Build from source or download prebuild binary
```
git clone https://github.com/cchare/grf.git
cd grf
go build -o grf
```

### Usage
##### help page
```
./grf -h
Usage of ./grf:
  -n int
    	number of files (default 1)
  -o string
    	output dir (default ".")
  -p string
    	filename prefix (default "laod")
  -s string
    	size(K,M,G,T) of file (default "1M")
  -v	show version

```

##### Generate 5 random files with size 2.22M to outdir/
```
mkdir outdir
./grf -n 5 -s 2.2M -o outdir
```
