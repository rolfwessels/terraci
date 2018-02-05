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

```
terraform plan
terraform apply
terraform destroy
```


# Resources

 * https://blog.gruntwork.io/how-to-create-reusable-infrastructure-with-terraform-modules-25526d65f73d
