# Workspaces Get

`workspaces/get` shows workspace lookup by ID.

~~~ shell
# Replace `id` in `main.tf` with the ID of the workspace that exists.
# To fetch the ID, consider leveraging the `workspaces` data source.
terraform apply -auto-approve
~~~

**Note: This Terraform provider is currently unpublished on the Terraform Registry and can only be executed in your local environment.**

**Note: `terraform init` does not work with `dev_overrides` for local development.**
