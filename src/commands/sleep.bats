#!/usr/bin/env bats

bats_load_library test_helper

@test "sleep returns quickly" {
    trun sleep 100
    assert_success
}
