#!/usr/bin/env bats

bats_load_library test_helper

@test "tail --help works" {
    trun tail --help
    assert_success
}

@test "tail with piped input works (fast)" {
    trun bash -c 'echo -e "1\n2\n3\n4\n5" | tail -2'
    assert_success
    assert_output $'4\n5'
}
