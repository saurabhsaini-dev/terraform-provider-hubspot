# Terraform Hubspot Provider

This terraform provider allows to perform Create, Read, 
Update, Delete and Import Hubspot User(s).

## Requirements 

* [Go](https://golang.org/doc/install) 1.16 <br>
* [Terraform](https://www.terraform.io/downloads.html) 0.13.x <br/>
* [Hubspot](https://www.hubspot.com/) Any account(Free/Starter/Enterprise)

## Setup Hubspot Account
1. Create a Hubspot account with any subscription. (https://www.hubspot.com/)<br>
2. Sign in to the hubspot account.<br>
3. Go to your developer account.
4. Click on `Manage Apps`.
5. Click on `Create app`. Create app with required parameters. This app will provide us with Client Id, Client Secret and Scopes which will be needed to configure our provider and make request.<br>

### Generate Refresh Token
For generating Refresh Token, follow this page <br> (https://developers.hubspot.com/docs/api/oauth-quickstart-guide) <br>



## Initialise Hubspot Provider in local machine
1. Clone the respository to $GOPATH/src/github.com/hubspot/terraform-provider-hubspot <br>
2. Add the Client Id, Client Secret and Refresh Token to respective fields in `main.tf` <br>
3. Run the following command :
 ```golang
go mod init terraform-provider-hubspot
go mod tidy
```
4. Run `go mod vendor` to create a vendor directory that contains all the provider's dependencies. <br>

## Installation
1. Run the following command to create a vendor subdirectory which will comprise of all provider dependencies. <br>
```
~/.terraform.d/plugins/${host_name}/${namespace}/${type}/${version}/${target}
```
Command:
```bash
mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hubspot/0.1.0/[OS_ARCH]
```
For eg. `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/hubspot/0.1.0/windows_amd64` <br>

2. Run `go build -o terraform-provider-hubspot.exe`. This will save the binary (`.exe`) file in the main/root directory. <br>
3. Run this command to move this binary file to appropriate location.
 ```
 move terraform-provider-hubspot.exe %APPDATA%\terraform.d\plugins\hashicorp.com\edu\hubspot\0.1.0\[OS_ARCH]
 ```

 Otherwise you can manually move the file from current directory to destination directory.<br>


 [OR]

 1. Download required binaries <br>
 2. move binary `~/.terraform.d/plugins/[architecture name]/`


## Run the Terraform Provider

#### Create User
1. Add the user email and role id in the respective feild in `main.tf`
2. Initialize the terraform provider `terraform init`
3. Check the changes applicable using `terraform plan` and apply using `terraform apply`
4. You will see that a user has been successfully created and an account activation mail has been sent to the user.
5. Activate the account using the link provided in the mail.

#### Update the User
Update the data of the user in the `main.tf` file and apply using `terraform apply`

#### Read the User data
Add data and output blocks in the `main.tf` file and run `terraform plan` to read user data

#### Delete the User
Delete the resource block of the particular user from `main.tf`
file and run `terraform apply`.

#### Import a User
1. Write manually a resource configuration block for the User in `main.tf`, to which the imported object will be mapped.
2. Run the command `terraform import hubspot_user.user1 [EMAIL_ID]`
3. Check for the attributes in the `.tfstate` file and fill them accordingly in resource block.

### Testing the Provider
1. Navigate to the test file directory.
2. Run command `go test` . This command will give combined test result for the execution or errors if any failure occur.
3. If you want to see test result of each test function individually while running test in a single go, run command `go test -v`
4. To check test cover run `go test -cover`


## Example Usage 
```terraform
terraform {
    required_providers {
        hubspot = {
            version = "0.1.0"
            source  = "hashicorp.com/edu/hubspot"
        }
    }
}

provider "hubspot" {
    client_id     = "AAA"
    client_secret = "AAA"
    refresh_token = "AAA"
}

resource "hubspot_user" "user1" {
    email  = "[EMAIL_ID]"
    roleid = "[ROLE_ID]"
}

data "hubspot_user" "user2" {
    id = "[EMAIL_ID]"
}

output "user1" {
    value = data.hubspot_user.user2
}
```

## Argument Reference

* `client_id`     - The Hubspot Client Id
* `client_secret` - The Hubspot Client Secert
* `email`         - The email id associated with the user account.
* `roleid`        - The Role Id assigned to the user