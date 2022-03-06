# lombok-toString

This is a package designed to parse the default format of Lombok's toString method

# Build from Source

The following command builds from source directly into your local bin folder (which is likely to be within your PATH for
convenience). Alternatively, you can execute it without it being in the PATH as well.

```bash
go build -o /usr/local/bin/lombok
```

# Development

I would suggest to set up some form of live reload in order to rerun the unit tests while you edit the files.

[Reflex](https://github.com/cespare/reflex) is lightweight tool for enabling live reload.

```bash
go install github.com/cespare/reflex@latest
reflex -d none -r "\.go" -s -- sh -c "echo \"\n\" && go test -v ./...  
```

You can enable a tad of colour into the test logs by using the following

```bash
reflex -d none -r "\.go" -s -- sh -c "echo \"\n\" && go test -v ./... | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/'' | sed ''/RUN/s//$(printf "\033[36mRUN\033[0m")/''"
```

