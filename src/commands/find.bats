#!/usr/bin/env bats

bats_load_library test_helper

@test "find -delete is blocked" {
    trun find /tmp -name "*.tmp" -delete
    assert_failure
    assert_output --partial "not allowed"
}

@test "find -exec is blocked" {
    trun find /tmp -exec ls {} \;
    assert_failure
    assert_output --partial "not allowed"
}

@test "find with -name uses fd" {
    trun find /etc -name "passwd" -maxdepth 1
    assert_success
    assert_output --partial "passwd"
}
