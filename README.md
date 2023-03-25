# Keep Raw Files

## The Purpose
Due to the camera's settings, a file will have both a JPG and a RAW file. My habit is to delete unwanted files from the JPG first, then go to the RAW files to delete the same ones.

```
.
├── DSCF6479.JPG
├── DSCF6480.JPG
├── DSCF6481.JPG
├── DSCF6482.JPG
├── DSCF6483.JPG
└── raw
    ├── DSCF6479.RAF
    ├── DSCF6480.RAF
    ├── DSCF6481.RAF
    ├── DSCF6482.RAF
    └── DSCF6483.RAF
```

After manually deleting, there are only 3 JPG files left.
```
.
├── DSCF6479.JPG
├── DSCF6481.JPG
├── DSCF6482.JPG
└── raw
    ├── DSCF6479.RAF
    ├── DSCF6480.RAF
    ├── DSCF6481.RAF
    ├── DSCF6482.RAF
    └── DSCF6483.RAF
```

After running this tool, I will delete the corresponding RAW files as well.
```
.
├── DSCF6479.JPG
├── DSCF6481.JPG
├── DSCF6482.JPG
└── raw
    ├── DSCF6479.RAF
    ├── DSCF6481.RAF
    └── DSCF6482.RAF
```

## How To Use?
```
go build main.go
./main {your picture path}
```