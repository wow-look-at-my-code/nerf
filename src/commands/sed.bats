#!/usr/bin/env bats

bats_load_library test_helper

@test "sed with piped input works" {
    trun bash -c 'echo "hello world" | sed "s/world/there/"'
    assert_success
    assert_output "hello there"
}

@test "sed without pipe fails with descriptive error" {
    trun sed 's/old/new/' /etc/passwd
    assert_failure
    assert_output --partial "not allowed for in-place file editing"
}
