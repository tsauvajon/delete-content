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
delete-content [x] [total] [workers] [repo1] [repoX...]
```

x : displays an # every x deleted document  
total : total number of documents to delete  
workers : number of concurrent workers to run  
repo1, repo2, repo3 ... : folders to clean  

Example:
```
delete-content 1000 20000 10 done
```
