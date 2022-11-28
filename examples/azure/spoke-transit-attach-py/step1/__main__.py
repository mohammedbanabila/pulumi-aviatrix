import pulumi
import pulumi_aviatrix as aviatrix
import os

#Create Aviatrix Transit VPC
vpc_tr3 = aviatrix.AviatrixVpc(
    "pulumi_transit1_vpc",
    cloud_type = 8,
    account_name = os.environ.get("AZURE_ACCT_NAME"),
    region = os.environ.get("AZURE_REGION"),
    name = "pulumi-tr3",
    cidr = "10.0.0.0/23",
    aviatrix_firenet_vpc = True
)

#Create Aviatrix Spoke VPC
vpc_spk3 = aviatrix.AviatrixVpc(
    "pulumi_spoke1_vpc",
    cloud_type = 8,
    account_name = os.environ.get("AZURE_ACCT_NAME"),
    region = os.environ.get("AZURE_REGION"),
    name = "pulumi-spk3",
    cidr = "10.1.0.0/24",
    aviatrix_firenet_vpc = False
)

#Create Aviatrix Spoke Gateway
spk3 = aviatrix.AviatrixSpokeGateway(
    "pulumi_spoke3_gw",
    cloud_type = 8,
    account_name = os.environ.get("AZURE_ACCT_NAME"),
    gw_name = "pulumi-spk3",
    vpc_id = vpc_spk3.vpc_id,
    vpc_reg = os.environ.get("AZURE_REGION"),
    gw_size= "Standard_B1ms",
    subnet = vpc_spk3.public_subnets[0].cidr
)

#Create Aviatrix Transit Gateway
tr3 = aviatrix.AviatrixTransitGateway(
    "pulumi_transit1_gw",
    cloud_type = 8,
    account_name = os.environ.get("AZURE_ACCT_NAME"),
    gw_name = "pulumi-tr3",
    vpc_id = vpc_tr3.vpc_id,
    vpc_reg = os.environ.get("AZURE_REGION"),
    gw_size= "Standard_B1ms",
    subnet = vpc_tr3.public_subnets[0].cidr,
    connected_transit = True
)

#Connect Spoke to Transit
aviatrix.AviatrixSpokeTransitAttachment(
     "pulumi_spk3_tr3_attach",
     spoke_gw_name = spk3.gw_name,
     transit_gw_name = tr3.gw_name
)
