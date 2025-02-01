data "testlagger_lag" "test" {
  read_delay = 1000
  input      = "hello"
}