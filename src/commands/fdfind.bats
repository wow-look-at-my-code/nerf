#!/usr/bin/env bats

bats_load_library test_helper

@test "fd -x is blocked" {
    trun fd -x echo {} /tmp
    assert_failure
    assert_output --partial "not allowed"
}

@test "fd --exec is blocked" {
    trun fd --exec echo {} /tmp
    assert_failure
    assert_output --partial "not allowed"
}

@test "fd -X is blocked" {
    trun fd -X echo {} /tmp
    assert_failure
    assert_output --partial "not allowed"
}

@test "fd --exec-batch is blocked" {
    trun fd --exec-batch echo {} /tmp
    assert_failure
    assert_output --partial "not allowed"
}

@test "fd without exec flags works" {
    trun fd passwd /etc --max-depth 1
    assert_success
    assert_output --partial "passwd"
}

@test "fdfind -x is blocked" {
    trun fdfind -x echo {} /tmp
    assert_failure
    assert_output --partial "not allowed"
}
