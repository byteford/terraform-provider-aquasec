resource "aquasec_function_runtime_policy" "function_runtime_policy" {
  name = "function_runtime_policys"
  description = "function_runtime_policy"
  enabled = true
  enforce = false
  block_malicious_executables = true
  block_running_executables_in_tmp_folder = true
  block_malicious_executables_allowed_processes = [
    "proc1",
    "proc2"
  ]
  blocked_executables = [
    "exe1",
    "exe2",
  ]
}