#!/usr/bin/env bats

bats_load_library test_helper

@test "docker without CLAUDECODE passes through" {
    unset CLAUDECODE
    trun docker --version
    assert_success
    assert_output --partial "Docker"
}

@test "docker with CLAUDECODE passes through for non-compose commands" {
    trun docker --version
    assert_success
    assert_output --partial "Docker"
}
