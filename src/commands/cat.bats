#!/usr/bin/env bats

bats_load_library test_helper

@test "cat truncates large files in terminal" {
    run_in_pty cat /etc/services
    assert_success
    assert_output --partial "truncated"
    assert_output --partial "Use the Read tool"
}

@test "cat passes through when piped" {
    trun bash -c 'cat /etc/passwd | head -1'
    assert_success
    assert_output --partial "root"
}
