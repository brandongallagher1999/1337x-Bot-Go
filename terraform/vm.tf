


resource "azurerm_linux_virtual_machine" "main_vm" {
    name = "1337x Bot VM"
    resource_group_name = "default" #may need to change this depending on your resource group in azure
    location = "East US"
    size = "Standard B2s"
    admin_username = "username"

    admin_ssh_key {
        username   = "azureuser"
        public_key = file("~/.ssh/id_rsa.pub")
    }

    source_image_reference {
        publisher = "Canonical"
        offer     = "UbuntuServer"
        sku       = "20.04-LTS"
        version   = "latest"
    }
} 