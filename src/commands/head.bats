#!/usr/bin/env bats

bats_load_library test_helper

@test "head with -n works" {
    trun head -n 1 test/fixtures/sample.txt
    assert_success
    assert_output "foo"
}

@test "head with piped input works (fast)" {
    trun bash -c 'echo -e "1\n2\n3\n4\n5" | head -2'
    assert_success
    assert_output $'1\n2'
}
