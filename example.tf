provider "cloudkarafka" {}

resource "cloudkarafka_instance" "kafka_bat" {
  name   = "terraform-provider-test"
  plan   = "ducky"
  region = "amazon-web-services::us-east-1"
  //vpc_subnet = "10.201.0.0/24"
}

resource "cloudkarafka_instance_alarm" "test" {
  type = "cpu"
  apikey = "${cloudkarafka_instance.kafka_bat.apikey}"
  value_threshold = 85
  time_threshold = 300
}

output "kafka_brokers" {
  value = "${cloudkarafka_instance.kafka_bat.brokers}"
}
