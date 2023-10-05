# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables
# ---------------------------------------------------------------------------------------------------------------------

variable "COUNTER_API_URL" {
    sensitive = true
}
variable "DISCORD_BOT_TOKEN" {
    sensitive = true
}

variable "DOCKER_IMAGE_SHA" {}
