# Adding some credentials

```
mkdir %USERPROFILE%\.aws\
notepad %USERPROFILE%\.aws\credentials
```

Format:

```
[default]
aws_access_key_id=--------------
aws_secret_access_key=-------------
aws_region=eu-west-1
```

## For developers

### Start the dev container

Note: windows users will need to link their .aws folder `ln -s /c/Users/User-pc/.aws ~/.aws`
`make up`

### Building

Run the command `make build` to build the main.exe

### Run

To run the examples. You call

- `go build; ./continues-terraforming plan eu-west-1 dev global`
  or once built
- `./continues-terraforming plan eu-west-1 dev global`
- `./continues-terraforming plan eu-west-1 dev setup`

Or you can build and plan using
TerraformCommand

- `make run arg="plan eu-west-1 dev global"`

# Resources

- https://nickjanetakis.com/blog/setting-up-docker-for-windows-and-wsl-to-work-flawlessly # for windows users
- https://blog.gruntwork.io/how-to-create-reusable-infrastructure-with-terraform-modules-25526d65f73d
