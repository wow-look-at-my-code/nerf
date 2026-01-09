#!/usr/bin/env bats

bats_load_library test_helper

@test "gh run watch is blocked" {
    trun gh run watch
    assert_failure 1
    assert_output --partial "gh wait-ci"
}

@test "gh run watch with args is blocked" {
    trun gh run watch 12345
    assert_failure 1
    assert_output --partial "gh wait-ci"
}

@test "gh other commands pass through" {
    trun gh --version
    assert_success
    assert_output --partial "gh version"
}
