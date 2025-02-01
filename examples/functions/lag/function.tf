output "lag_echo" {
  value = provider::testlagger::lag(1000, "hello")
}