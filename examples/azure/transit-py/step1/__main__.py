import pulumi
import pulumi_aviatrix as aviatrix
import os

#Create Aviatrix Transit VPC
vpc_tr2 = aviatrix.AviatrixVpc(
    "pulumi_transit2_vpc",
    cloud_type = 8,
    account_name = os.environ.get("AZURE_ACCT_NAME"),
    region = os.environ.get("AZURE_REGION"),
    name = "pulumi-tr2",
    cidr = "10.0.0.0/23",
    aviatrix_firenet_vpc = True
)

#Create Aviatrix Transit Gateway
tr2 = aviatrix.AviatrixTransitGateway(
    "pulumi_transit2_gw",
    cloud_type = 8,
    account_name = os.environ.get("AZURE_ACCT_NAME"),
    gw_name = "pulumi-tr2",
    vpc_id = vpc_tr2.vpc_id,
    vpc_reg = os.environ.get("AZURE_REGION"),
    gw_size= "Standard_B1ms",
    subnet = vpc_tr2.public_subnets[0].cidr,
    connected_transit = True
)
