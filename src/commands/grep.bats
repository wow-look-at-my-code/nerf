#!/usr/bin/env bats

bats_load_library test_helper

@test "grep without args shows usage and exits" {
    trun grep
    assert_failure 2
    assert_output --partial "Usage:"
}

@test "grep with pattern works on piped input (fast)" {
    trun bash -c 'echo -e "foo\nbar\nbaz" | grep bar'
    assert_success
    assert_output "bar"
}

@test "grep with file argument works" {
    trun grep root /etc/passwd
    assert_success
    assert_output --partial "root"
}

@test "grep finds pattern in line with multiple matches" {
    trun bash -c 'echo "testestest" | grep test'
    assert_success
    assert_output "testestest"
}
