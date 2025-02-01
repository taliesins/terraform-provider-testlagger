resource "testlagger_lag" "test" {
  create_delay = 1000
  read_delay   = 1000
  update_delay = 1000
  delete_delay = 1000
  input        = "hello"
}