
## Project Structure

This project is structured as follows:

- Back-end (BE) and front-end (FE) as distinct subprojects
- Common code for BE and FE exists in `shared` folder
    - this includes models, templates, and functions
- Front-end Go web app code is inside of `webui` folder


## Development

A convenient way for starting and reloading the app is to use `kick` tool.

`kick` watches for code changes in either BE or FE sides. If such a change is detected, it will
trigger the compilation on that side to capture the change and it will _restart_ the app. This improves considerably the development experience: you run this command once, instead of manually running `gopher build` command for getting the changes done on the FE side and running `go run orgone.go` command to restart the app (including the BE side).

