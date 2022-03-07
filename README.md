# lombok-toString

This is a package designed to parse the default format
of [Lombok's toString method](https://projectlombok.org/features/ToString) into a more common format: JSON.

This package is not designed for production use.

The Lombok's default toString implementation provides a quick and dirty string for application logs while preserving
data such as the class names. There have been many suggestions to improve the implementation that you can
follow [here](https://github.com/projectlombok/lombok/issues/1297).

This package is designed to convert such logs into something that is more conventional so that it may be more readable,
or reused for other purposes, such as building JSON requests.

# Usage

```bash
Usage:
  lombokString [string to be parsed] [flags]

Flags:
  -x, --exclude-null   exclude the fields with null value
  -h, --help           help for lombokString
  -m, --mini           minify output (i.e. remove all indents)
```

## Example

To give an example

```java
import lombok.ToString;

@ToString
public class User {
    String name;
    int age;
    String[] emails;
}
// user.toString() will return User(name=Bob Smith, age=42, emails=[bob@gmail.com, bob87@hotmail.com, smith_97@hotmail.com])
```

To generate a JSON version of the log, we can execute the following:

```bash
>>> lombok 'User(name=Bob Smith, age=42, emails=[bob@gmail.com, bob87@hotmail.com, smith_97@hotmail.com])'

{
    "age": 42,
    "emails": [
        "bob@gmail.com",
        "bob87@hotmail.com",
        "smith_97@hotmail.com"
    ],
    "name": "Bob Smith"
}
```

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

