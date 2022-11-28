import pulumi
import pulumi_aviatrix as aviatrix
import os

#Create Aviatrix Spoke VPC
vpc_spk1 = aviatrix.AviatrixVpc(
    "pulumi_spoke1_vpc",
    cloud_type = 8,
    account_name = os.environ.get("AZURE_ACCT_NAME"),
    region = os.environ.get("AZURE_REGION"),
    name = "pulumi-spk1",
    cidr = "10.1.0.0/24",
    aviatrix_firenet_vpc = False
)

#Create Aviatrix Spoke Gateway
spk1 = aviatrix.AviatrixSpokeGateway(
    "pulumi_spoke1_gw",
    cloud_type = 8,
    account_name = os.environ.get("AZURE_ACCT_NAME"),
    gw_name = "pulumi-spk1",
    vpc_id = vpc_spk1.vpc_id,
    vpc_reg = os.environ.get("AZURE_REGION"),
    gw_size= "Standard_B1ms",
    subnet = vpc_spk1.public_subnets[0].cidr
)
