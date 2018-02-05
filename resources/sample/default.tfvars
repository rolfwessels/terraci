aws_security_regions = ["back","middle","front","edge"]
allowed_secure_ips = [
    "155.93.248.0/22", # rolf 155.93.250.148/32 155.93.249.108
    "196.213.203.122/32", # gk
    "196.22.232.211/32" # infoslips
    ]

aws_amis_ubuntu = {
    eu-west-1 = "ami-674cbc1e"
    us-east-1 = "ami-1d4e7a66"
    us-west-1 = "ami-969ab1f6"
    us-west-2 = "ami-8803e0f0"
}
aws_amis_openswan = {
    eu-west-1 = "ami-ebd02392"
}
aws_beanstalk_application_viewer_name = "infoslips-viewer"
aws_beanstalk_application_viewer_description = "IIAB viewer deployement"

aws_beanstalk_application_api_name = "infoslips-api"
aws_beanstalk_application_api_description = "IIAB api deployement"


aws_beanstalk_application_application_name = "infoslips-application"
aws_beanstalk_application_application_description = "IIAB application deployement"

aws_beanstalk_application_evopdf_name = "infoslips-evopfd"
aws_beanstalk_application_evopdf_name_description = "IIAB viewer deployement"
