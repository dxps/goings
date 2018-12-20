## Notes

1. Added a small update to `kick` to consider also `.gohtml` extensions.

   Going to `$GOPATH/src/go.isomorphicgo.org/go/kick/`, in `kick.go` file I updated first line of `InitializeWatcher` function as such:
   ```go
   func initializeWatcher(shouldRestart chan bool, dirList []string) {

        supportedExtensions := map[string]int{".go": 1, ".html": 1, ".tmpl": 1, ".gohtml": 1}
   ```
   Then just ran `go install` to have it in `$GOPATH/bin`, which is part of my `PATH` env var, so I can use it anywhere.
   
   _Optional:_ In my (OCS tendency) case, I added a new line (`\n`) marker to the line that logs the reloading action, as such: `Instant KickStart Applied ....\n`

