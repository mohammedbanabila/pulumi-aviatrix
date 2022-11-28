import pulumi
import pulumi_aviatrix as aviatrix
import os

#Create Aviatrix  VPC
vpc_tr1 = aviatrix.AviatrixVpc(
    "pulumi_transit1_vpc",
    cloud_type = 8,
    account_name = os.environ.get("AZURE_ACCT_NAME"),
    region = os.environ.get("AZURE_REGION"),
    name = "pulumi-tr1",
    cidr = "10.0.0.0/23",
    aviatrix_firenet_vpc = True
)
