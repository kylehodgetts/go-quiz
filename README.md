# Go Quiz

## Build
### Binary
`go build`

### Docker
`docker build --rm -t go-quiz .`

## Run
### Binary
`./go-quiz [-csv path to problems csv | -timelimit int time limit]`

### Docker
`docker run -it go-quiz -csv problems.csv`
- Container has a default problems csv file in the root.
- Cannot pass a custom csv file with Docker yet. May add it later.