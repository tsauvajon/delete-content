# Delete directories

I had a folder containing millions of files. On a Windows system, simply deleting
the folder didn't work. Deleting the content manually or with the command line
is extremely slow and inefficient.

This utility deletes all files in a repo leveraging go's concurrency, by using
a worker pool.

Build:

```
go build
```

Usage:

```
./delete-content [flags] [folder1] [folderX...]
```

-nb number of files to delete (default 0: delete every file)
-w number of concurrent workers (default 10)
-d only delete documents older than this number of days (default 0: all files)

Example:
```
# delete 10 files in the folders ./folder1, ./folder2 and ./folder3 that are at least 7 days old, with 5 concurrent workers
./delete-content -nb=10 -d=7 -w=5 folder1 folder2 folder3
```
