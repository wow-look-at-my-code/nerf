#!/usr/bin/env bats

bats_load_library test_helper

@test "tail with -n works" {
    trun tail -n 1 test/fixtures/sample.txt
    assert_success
    assert_output "baz"
}

@test "tail with piped input works (fast)" {
    trun bash -c 'echo -e "1\n2\n3\n4\n5" | tail -2'
    assert_success
    assert_output $'4\n5'
}
