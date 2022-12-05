Advent of Code 
==============

Fetch a day's data:
```
export DAY=3
go install github.com/GreenLightning/advent-of-code-downloader/aocdl@latest
export PATH=~/go/bin/:$PATH
export YEAR=2022
mkdir -p ${YEAR}/day${DAY}
aocdl -year ${YEAR} -day ${DAY} -output ${YEAR}/day${DAY}/day${DAY}.txt
```

Add a new day's BUILD.bazel (after you touch `${YEAR}/day${DAY}/day${DAY}.go`:

```
bazel run //:gazelle
```

and then be sure to add:
```
    data = [
        "dayX.txt",
        "dayX_ex.txt",
    ],
```

Fetch dependencies:

```
bazel run //:gazelle-update-repos
```

Run the latest day constantly with changes:
```bash
while :; do git ls-files -cmdo --exclude-standard | entr bash -c "bazel run \$(bazel query 'kind(go.*binary,  '//...')' | tail -1)"; done
```
