terraform {
  required_providers {
    azurerm = {
      source = "hashicorp/azurerm"
      version = "3.4.0"
    }
  }
}

provider "azurerm" {
    features {}
    subscription_id = "insert here" #Your subscription ID in Azure
}