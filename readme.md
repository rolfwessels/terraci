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

### Building

Run the command `./build.com` to build the main.exe

### Run

To run the examples. You call

- `./main plan eu-west-1 dev global`
- `./main plan eu-west-1 dev region/env/setup`

Or you can build and plan using

- `./run plan eu-west-1 dev global`

# Resources

- https://blog.gruntwork.io/how-to-create-reusable-infrastructure-with-terraform-modules-25526d65f73d
