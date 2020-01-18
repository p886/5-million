# 5 Million Steps in 2020

This little CLI app tells me if I'm on track for my goal of **5 Million Steps in 2020**.

## Build

```
go build -o app
```

## Usage

### CLI

```
./app [steps so far]
```

### Server

```
./app server
```

Then issue an HTTP request to `/?steps=[steps so far]`.

## Example

```
$ ./app 160000

Current target:	190000
Current steps:	160000
Current diff:	-30000
Relative diff:	-15.79%
```

HTTP JSON response

```
{
  "current_target": 245898,
  "current_steps": 200000,
  "diff": -45898,
  "relative_diff": "-18.67%"
}
```
