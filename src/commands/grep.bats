#!/usr/bin/env bats

bats_load_library test_helper

@test "grep without args shows usage and exits" {
    trun grep
    assert_failure 2
    assert_output --partial "usage"
}

@test "grep with pattern works on piped input (fast)" {
    trun bash -c 'echo -e "foo\nbar\nbaz" | grep bar'
    assert_success
    assert_output "bar"
}

@test "grep with file argument works" {
    trun grep hello test/fixtures/sample.txt
    assert_success
    assert_output "hello world"
}

@test "grep finds pattern in line with multiple matches" {
    trun bash -c 'echo "testestest" | grep test'
    assert_success
    assert_output "testestest"
}

@test "grep returns failure for pattern not in file" {
    trun grep notfound test/fixtures/sample.txt
    assert_failure 1
    assert_output ""
}
